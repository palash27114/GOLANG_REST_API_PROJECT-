package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/palash27114/REST_API_PROJECT/internal/types"
	"github.com/palash27114/REST_API_PROJECT/internal/utils/response"
	"gopkg.in/go-playground/validator.v9"
)

func New() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){

		var student types.Student

		err:=json.NewDecoder(r.Body).Decode(&student)

	if errors.Is(err, io.EOF){
	response.WriteJson(w, http.StatusBadRequest,response.GeneralErr(fmt.Errorf("Empty every thing")))
	return
}


		slog.Info("creating a student")


		if err!=nil{
			response.WriteJson(w,http.StatusBadRequest,response.GeneralErr(err))
			return 

		}


		//request validate 
		//goland request validation 
		if err:=validator.New().Struct(student);err!=nil{
			validErrs:= err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest,response.ValidationError(validErrs))
			return

		}





		w.Write([]byte("welcome to student api"))

		response.WriteJson(w,http.StatusCreated,map[string]string{"success":"OK"})

}}