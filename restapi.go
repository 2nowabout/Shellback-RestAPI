package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type notification struct {
	Type     string `json:"Type"`
	IpAdress string `json:"IP"`
	Value    string `json:"Value"`
}

type notifications []notification

func main() {
	requestHandler()
}

func requestHandler() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/getNotifications", getNotificationsHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/ hit")
}

func getNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	allNotifications := notifications{
		notification{Type: "testType", IpAdress: "192.168.0.1", Value: "CVE 9.0 Found"},
		notification{Type: "testType", IpAdress: "192.168.0.2", Value: "CVE 9.0 Found"},
	}
	fmt.Println(w, "loading all notifications")
	json.NewEncoder(w).Encode(allNotifications)
}
