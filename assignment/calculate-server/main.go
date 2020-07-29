package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chanchailee/money-table/assignment/pkg/model"
	"github.com/chanchailee/money-table/assignment/pkg/service"
)

func calculate(w http.ResponseWriter, req *http.Request, operation func(c *model.Req) (float64, error), operationName string) {
	num, err := service.UnmarshalReq(w, req)
	if err != nil {
		log.Panicf("%+v", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	result, err := operation(num)
	if err != nil {
		log.Panicf("%+v", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	data, err := service.MarshalResp(operationName, result)
	if err != nil {
		log.Panicf("%+v", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	service.HandleSuccessResp(w, data)
}

func handleSum(w http.ResponseWriter, req *http.Request) {
	calculate(w, req, service.Sum, "sum")
}

func handleMul(w http.ResponseWriter, req *http.Request) {
	calculate(w, req, service.Mul, "mul")
}

func handleDiv(w http.ResponseWriter, req *http.Request) {
	calculate(w, req, service.Div, "div")
}

func handleSub(w http.ResponseWriter, req *http.Request) {
	calculate(w, req, service.Sub, "sub")
}

func main() {
	http.HandleFunc("/sum", handleSum)
	http.HandleFunc("/mul", handleMul)
	http.HandleFunc("/div", handleDiv)
	http.HandleFunc("/sub", handleSub)

	err := http.ListenAndServe(os.Getenv("ADDRESS"), nil)
	if err != nil {
		panic(err)
	}
}
