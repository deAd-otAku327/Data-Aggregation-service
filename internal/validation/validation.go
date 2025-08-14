package validation

import (
	"data-aggregation-service/internal/types/dto"
	"fmt"
	"time"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validation struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func NewValidation() *Validation {
	validate := validator.New()

	_ = validate.RegisterValidation("afterdate", func(fl validator.FieldLevel) bool {
		firstDateFieldName := fl.Param()
		firstDateFieldValue := fl.Parent().FieldByName(firstDateFieldName)
		if !firstDateFieldValue.IsValid() {
			fmt.Println(firstDateFieldValue)
			return false
		}

		firstDate, err := time.Parse(dto.DateTimeLayout, firstDateFieldValue.String())
		if err != nil {
			return false
		}

		secondDate, err := time.Parse(dto.DateTimeLayout, fl.Field().String())
		if err != nil {
			return false
		}

		return secondDate.After(firstDate)
	})

	uni := ut.New(en.New(), ru.New())
	translate, _ := uni.GetTranslator("en")

	translations := map[string]string{
		"required":         "{0} value required",
		"required_without": "required at least one of values: {0}, {1}",
		"gte":              "{0} value must be > {1}",
		"uuid4":            "{0} value must meet uuid4 format",
		"datetime":         "{0} datetime value must be in {1} format",
		"afterdate":        "{0} datetime value must be after {1}",
	}

	for tag, translation := range translations {
		_ = validate.RegisterTranslation(tag, translate,
			func(ut ut.Translator) error {
				return ut.Add(tag, translation, true)
			},
			func(ut ut.Translator, fe validator.FieldError) string {
				param := fe.Param()
				if param == "" {
					t, _ := ut.T(tag, fe.Field())
					return t
				}
				t, _ := ut.T(tag, fe.Field(), param)
				return t
			},
		)
	}

	return &Validation{
		Validator:  validate,
		Translator: translate,
	}
}

func CollectValidationErrors(err error, translator ut.Translator) []string {
	msgs := make([]string, 0)
	for _, e := range err.(validator.ValidationErrors) {
		msgs = append(msgs, e.Translate(translator))
	}

	return msgs
}
