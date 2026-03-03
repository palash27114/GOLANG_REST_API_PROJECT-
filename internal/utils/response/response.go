package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type Response struct {
	Status string `json:"status"`
	Err    string `json:"error"`
}

const (
	Statusok    = "OK"
	StatusError = "Empty Body"
)



func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralErr(err error) Response {
	return Response{
		Status: StatusError,
		Err:    err.Error(),
	}

}


func ValidationError(errs validator.ValidationErrors)Response{
	var errMsg [] string 
	for _,err:=range errs{
		switch err.ActualTag(){
		case "required":
			errMsg=append(errMsg, fmt.Sprintf("field %s is required field ",err.Field()))
		default:
			errMsg=append(errMsg, fmt.Sprintf("field %s is  field is invalid ",err.Field()))
		
		}

	}

	return Response{
		Status: StatusError,
		Err: strings.Join(errMsg,","),
	}

}