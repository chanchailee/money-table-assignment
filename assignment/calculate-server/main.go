package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/chanchailee/money-table/assignment/pkg/helper"
	"github.com/chanchailee/money-table/assignment/pkg/object"
)

func marshalResp(method string, result float64) ([]byte, error) {
	return json.Marshal(object.Resp{
		Method: method,
		Result: result,
	})
}

func unmarshalReq(w http.ResponseWriter, req *http.Request) (*object.Req, error) {
	var num object.Req
	b, err := ioutil.ReadAll(req.Body)

	defer req.Body.Close()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &num)
	if err != nil {
		return nil, err
	}

	return &num, nil
}

func sum(num *object.Req) (float64, error) {
	return num.A + num.B, nil
}

func mul(num *object.Req) (float64, error) {
	return num.A * num.B, nil
}

func div(num *object.Req) (float64, error) {
	if num.B == 0 {
		return 0, errors.New("invalid numbers : can't divide 0 with 0")
	}

	return num.A / num.B, nil
}

func sub(num *object.Req) (float64, error) {
	return num.A - num.B, nil
}

func calculate(w http.ResponseWriter, req *http.Request, operation func(c *object.Req) (float64, error), operationName string) {
	num, err := unmarshalReq(w, req)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	result, err := operation(num)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	data, err := marshalResp(operationName, result)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	helper.HandleSuccessResp(w, data)
}

func handleSum(w http.ResponseWriter, req *http.Request) {
	calculate(w, req, sum, "sum")
}

func handleMul(w http.ResponseWriter, req *http.Request) {
	calculate(w, req, mul, "mul")
}

func handleDiv(w http.ResponseWriter, req *http.Request) {
	calculate(w, req, div, "div")
}

func handleSub(w http.ResponseWriter, req *http.Request) {
	calculate(w, req, sub, "sub")
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
