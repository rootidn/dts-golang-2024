package controllers

import (
	"assignment-3/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	dataMu      sync.RWMutex
	currentData DataWrapper
	err         error
)

type DataWrapper struct {
	Script     template.HTML
	Data       models.Data
	Conditions Condition
}

type Condition struct {
	Water string
	Wind  string
}

// handler
func StatusMonitoring(w http.ResponseWriter, r *http.Request) {
	dataMu.RLock()
	defer dataMu.RUnlock()

	tmpl := template.Must(template.ParseFiles("assets/index.html"))

	currentData, err = getInitialData()
	if err != nil {
		http.Error(w, fmt.Sprint("failed to load data from json", err), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, currentData)
}

func UpdateData() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		newData, err := getUpdatedData()
		if err != nil {
			fmt.Printf("error updating data %v", err)
			break
		}
		dataMu.Lock()
		currentData = newData
		dataMu.Unlock()
	}
}

func getInitialData() (DataWrapper, error) {
	jsonFile, err := os.Open("./data/data.json")

	if err != nil {
		return DataWrapper{}, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return DataWrapper{}, err
	}

	var data models.Data

	err = json.Unmarshal([]byte(byteValue), &data)
	if err != nil {
		return DataWrapper{}, err
	}

	cond := Condition{Water: checkCondition(0, data.Status.Water), Wind: checkCondition(1, data.Status.Wind)}
	dw := DataWrapper{Script: template.HTML("<script>setTimeout(function(){window.location.reload(1);}, 15000);</script>"), Data: data, Conditions: cond}

	return dw, nil
}

func getUpdatedData() (DataWrapper, error) {
	data := models.Data{}
	var newWater = rand.Intn(100) + 1
	var newWind = rand.Intn(100) + 1
	data.Status.Water = newWater
	data.Status.Wind = newWind

	dataJson, err := json.Marshal(data)
	if err != nil {
		return DataWrapper{}, err
	}

	err = os.WriteFile("./data/data.json", dataJson, 0644)
	if err != nil {
		return DataWrapper{}, err
	}

	cond := Condition{Water: checkCondition(0, newWater), Wind: checkCondition(1, newWind)}

	return DataWrapper{Script: template.HTML("<script>setTimeout(function(){window.location.reload(1);}, 15000);</script>"), Data: data, Conditions: cond}, nil
}

func checkCondition(tipe int, status int) string {
	// 0 for water
	// 1 for wind

	if tipe == 0 {
		switch {
		case status < 5:
			return "Aman"
		case (status >= 6) && (status <= 8):
			return "Siaga"
		case status > 8:
			return "Bahaya"
		default:
			return "-"
		}
	} else if tipe == 1 {
		switch {
		case status <= 6:
			return "Aman"
		case (status >= 7) && (status <= 15):
			return "Siaga"
		case status > 15:
			return "Bahaya"
		default:
			return "-"
		}
	}
	return "-"
}
