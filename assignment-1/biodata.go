package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
	Function to show selected friend biodata based on program arguments
	run
		go run main.go [friends_number int]
*/

type Friend struct {
	nama                      string
	alamat                    string
	pekerjaan                 string
	alasan_pilih_kelas_golang string
}

func main() {
	// define friends data
	friends := []Friend{
		{nama: "Joko", alamat: "Pancoran", pekerjaan: "Sales", alasan_pilih_kelas_golang: "Asik"},
		{nama: "Budi", alamat: "Kebon Jeruk", pekerjaan: "Petani", alasan_pilih_kelas_golang: "Keren"},
		{nama: "Fina", alamat: "Kuningan", pekerjaan: "Dokter", alasan_pilih_kelas_golang: "Seru"},
		{nama: "Eko", alamat: "Tanah Abang", pekerjaan: "Masinis", alasan_pilih_kelas_golang: "Kebutuhan"},
		{nama: "Ayu", alamat: "Mangga Dua", pekerjaan: "Perawat", alasan_pilih_kelas_golang: "Penasaran"},
	}
	// get and validate friend's number from args
	friendsNum := getFriendsNum(len(friends))
	// show data on selected friends
	showFriendsData(friends, friendsNum)
}

func getFriendsNum(friendsLength int) int {
	argsRaw := os.Args
	// check args length must be 2 : file-location and number
	strNum, err := checkArgs(argsRaw)
	if err != nil {
		fmt.Println("Error occured: ", err)
		os.Exit(0)
	}
	// check if input number is int
	num, err := checkInputNum(strNum, friendsLength)
	if err != nil {
		fmt.Println("Error occured: ", err)
		os.Exit(0)
	}
	return num
}

func checkArgs(args []string) (string, error) {
	if len(args) < 2 || len(args) > 2 {
		return "", fmt.Errorf("insert exactly 1 integer argument to choose friend's data you want to show")
	}
	return args[1], nil
}

func checkInputNum(strNum string, friendsLength int) (int, error) {
	// check can be converted to int
	num, err := strconv.Atoi(strNum)
	if err != nil {
		return 0, fmt.Errorf("friend's number input argument must be on integer format")
	}
	// check if out of range
	if num < 0 || num >= friendsLength {
		return 0, fmt.Errorf("friend's number input argument is out of range (choose 0 to %d)", friendsLength-1)
	}
	return num, nil
}

func showFriendsData(friends []Friend, num int) {
	friend := friends[num]
	fmt.Printf("----- Showing friends's data (No.%d) -----\n", num)
	fmt.Println("nama :", friend.nama)
	fmt.Println("alamat :", friend.alamat)
	fmt.Println("pekerjaan :", friend.pekerjaan)
	fmt.Println("alasan milih kelas golang :", friend.alasan_pilih_kelas_golang)
}
