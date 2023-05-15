package utils

import "fmt"

type Validator struct {
	// Validator için alanlar burada tanımlanır
	//...
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(data interface{}) error {
	// Veri doğrulama işlemleri burada gerçekleştirilir
	var err error
	s := NewLogger()
	if data != 1 {
		s.LogError(err)
		return err
	}
	fmt.Println(data)
	return err
}
