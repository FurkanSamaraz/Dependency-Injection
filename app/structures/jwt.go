package structures

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID   float64 `json:"id"`
	Name string  `json:"name"`
	jwt.StandardClaims
}

func NewClaims() *Claims {
	return &Claims{}
}

func (m *Claims) Claims() {
	// Model1 için metodlar burada yer alır
}
