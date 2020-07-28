package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/chanchailee/money-table/assignment/pkg/helper"
)

func sendToCalculatedServer(w http.ResponseWriter, req *http.Request, operation string) {
	url := os.Getenv("CAL_SERVER") + "/" + operation
	resp, err := http.Post(url, "application/json", req.Body)

	if err != nil {
		http.Error(w, "bad gateway", http.StatusBadGateway)
		return
	}
	if err != nil {
		http.Error(w, "bad gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	helper.HandleSuccessResp(w, body)
}

func sum(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	sendToCalculatedServer(w, req, "sum")
}

func mul(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	sendToCalculatedServer(w, req, "mul")
}

func div(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	sendToCalculatedServer(w, req, "div")
}

func sub(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sendToCalculatedServer(w, req, "sub")
}

func main() {
	http.HandleFunc("/calculator.sum", sum)
	http.HandleFunc("/calculator.mul", mul)
	http.HandleFunc("/calculator.div", div)
	http.HandleFunc("/calculator.sub", sub)

	err := http.ListenAndServe(os.Getenv("ADDRESS"), nil)
	if err != nil {
		panic(err)
	}
}
