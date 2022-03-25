package test

import (
	"log"
	"regexp"

	"github.com/Mukunth-arya/golangapp/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

func Validate(Data3 models.Data) {

	err := validation.ValidateStruct(&Data3,
		validation.Field(&Data3.ProductName, validation.Required,
			validation.Length(5, 30).Error("please enter the value between a to z")),
		validation.Field(&Data3.Servicecomment, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`))),
		validation.Field(&Data3.Qualiycomment, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`))),
	)
	if err != nil {
		log.Fatal(err)
	}

}
