package test

import (
	"log"
	"regexp"

	"github.com/Mukunth-arya/golangapp/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

func Validate(Data3 models.Data) {

	err := validation.ValidateStruct(&Data3,
		validation.Field(&Data3.CakeName, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
		validation.Field(&Data3.Cakeflavour, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
		validation.Field(&Data3.TypeofCream, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
		validation.Field(&Data3.Toppings, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
		validation.Field(&Data3.Shape, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
		//validation.Field(&Data3.Satisfied, validation.Required,
		//validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
	)
	if err != nil {
		log.Fatal(err)
	}

}
