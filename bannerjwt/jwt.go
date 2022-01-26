package bannerjwt

import (
  "time"
  "crypto/rsa"
  "crypto/rand"
  "github.com/golang-jwt/jwt"
)

var verify_key *rsa.PublicKey
var sign_key *rsa.PrivateKey

func init(){
  sign_key, err := rsa.GenerateKey(rand.Reader , 1024)
  if err != nil{
    panic(err)
  }
  verify_key = &sign_key.PublicKey
}

type Claims struct {
	*jwt.StandardClaims
	Name string
}

func IsAuth(token_string string) (string, error){
  token, err := jwt.ParseWithClaims(token_string, &Claims{},
  func(token *jwt.Token) (interface{}, error) {
      return verify_key, nil
   },
  )
  if err != nil{
    return "", err
  }
  claims := token.Claims.(*Claims)
  return claims.Name , nil
}

func CreateToken(name string) (string, error){
  	token := jwt.New(jwt.GetSigningMethod("RS256"))
  	token.Claims = &Claims{
  		&jwt.StandardClaims{
  			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
  		},
  		name,
  	}
  	return token.SignedString(sign_key)
  }