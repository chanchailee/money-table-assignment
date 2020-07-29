package main

import (
	"net/http"
	"os"

	"github.com/chanchailee/money-table/assignment/pkg/service"
)

func sum(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	service.SendToCalculatedServer(w, req, "sum")
}

func mul(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	service.SendToCalculatedServer(w, req, "mul")
}

func div(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	service.SendToCalculatedServer(w, req, "div")
}

func sub(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	service.SendToCalculatedServer(w, req, "sub")
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
