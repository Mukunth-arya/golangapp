package Customerr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

//This Datastruct is used to customize the error
type Myerror struct {
	when    time.Time
	Kind    string `json:"kind"`
	code    string `json:"code"`
	Message string `json:"message"`
}

type Serviceerror struct {
	Error Myerror `json: "error"`
}

func UnknownErrorResponse(rw http.ResponseWriter, lgr zerolog.Logger, err error) {

	data1 := Serviceerror{Error: Myerror{
		when:    time.Now(),
		Kind:    "unanticipated",
		code:    "unanticipated",
		Message: "Please enter the valid content",
	},
	}
	lgr.Error().Err(err).Msg("unknown error")
	data2, err := json.Marshal(data1)
	if err != nil {
		log.Fatal(err)
	}

	data3 := string(data2)
	rw.Header().Set("Content-type", "application/json")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	rw.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(rw, data3)
}
