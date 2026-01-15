package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateJWTToken(secretKey string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
    	return "", err
    }
  	return tokenString, nil
}

func ValidateJWTToken(tokenString string, secretKey string) (jwt.MapClaims, error) {
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
      return []byte(secretKey), nil
   })
  
   if err != nil {
      return nil, err
   }
  
   if !token.Valid {
      return nil, errors.New("invalid token")
   }

   if claims, ok := token.Claims.(jwt.MapClaims); ok {
      return claims, nil
   }

   return nil, errors.New("invalid claims")
}