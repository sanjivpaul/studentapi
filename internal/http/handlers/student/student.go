package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sanjivpaul/studentapi/internal/types"
	"github.com/sanjivpaul/studentapi/internal/utils/response"
)



func New() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		// handle empty body error
		if errors.Is(err, io.EOF){
			// response.WriteJson(w, http.StatusBadRequest, err.Error()) // normal error show
			// response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err)) // with custom show
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body"))) // with custom and custom msg
			return
		}

		// handle error if another error is happen
		if err != nil{
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}


		// request validation
		if err := validator.New().Struct(student); err != nil{

			validateErrs := err.(validator.ValidationErrors) // typecast error format

			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			
			return
		}


		response.WriteJson(w, http.StatusCreated, map[string]string{"success":"OK"})
		// w.Write([]byte("create new handler is working"))
	}
}