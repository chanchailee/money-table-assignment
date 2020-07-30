package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/chanchailee/money-table/assignment/pkg/model"
)

// HandleSuccessResp :
func HandleSuccessResp(w http.ResponseWriter, data []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// SendToCalculatedServer :
func SendToCalculatedServer(w http.ResponseWriter, req *http.Request, operation string) {
	url := os.Getenv("CAL_SERVER") + "/" + operation
	resp, err := http.Post(url, "application/json", req.Body)

	if err != nil {
		log.Panicf("%+v", err.Error())
		http.Error(w, "bad gateway", http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("%+v", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	HandleSuccessResp(w, body)
}

// MarshalResp :
func MarshalResp(method string, result float64) ([]byte, error) {
	return json.Marshal(model.Resp{
		Method: method,
		Result: result,
	})
}

// UnmarshalReq :
func UnmarshalReq(w http.ResponseWriter, req *http.Request) (*model.Req, error) {
	var num model.Req

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	err = json.Unmarshal(b, &num)
	if err != nil {
		return nil, err
	}

	return &num, nil
}

// Sum :
func Sum(num *model.Req) (float64, error) {
	return num.A + num.B, nil
}

// Mul :
func Mul(num *model.Req) (float64, error) {
	return num.A * num.B, nil
}

// Div :
func Div(num *model.Req) (float64, error) {
	if num.B == 0 {
		return 0, errors.New("invalid numbers : can't divide with 0")
	}

	return num.A / num.B, nil
}

// Sub :
func Sub(num *model.Req) (float64, error) {
	return num.A - num.B, nil
}
