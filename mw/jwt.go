package mw

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

const (
	privKeyPath = "./sample_key"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "./sample_key.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
	serverPort int
)

type CustomerInfo struct {
	Id   int64
	Name string
}

type Claims struct {
	jwt.RegisteredClaims
	CustomerInfo
}

func JwtInit(privKeyPath string, pubKeyPath string) error {

	signBytes, err := os.ReadFile(privKeyPath)
	if err != nil {
		return err
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return err
	}

	verifyBytes, err := os.ReadFile(pubKeyPath)
	if err != nil {
		return err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return err
	}

	return nil

}

func CreateToken(userid int64, username string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims = &Claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 365)),
		},
		CustomerInfo{Id: userid, Name: username},
	}

	// Creat token string
	return t.SignedString(signKey)
}
func TokenStringGetUser(tokenString string) (int64, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	if err != nil {
		return int64(-1), "", err
	}

	claims := token.Claims.(*Claims)
	return claims.CustomerInfo.Id, claims.CustomerInfo.Name, nil
}
func Authen(tokenString string) bool {
	return true
}
