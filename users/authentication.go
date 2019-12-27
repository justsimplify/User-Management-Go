package users

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	st "himanshu.com/sample/structures"
)

const (
	signingKey = "s3cre3t"
)

// CreateToken JWT Token
func CreateToken(userName string, password string, id int) string {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(15 * time.Second)

	// Create the Claims
	claims := st.UserAuthClaim{
		UserName: userName,
		Password: password,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "hikumar",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(signingKey))
	if err != nil {
		panic(err)
	}
	return ss
}

// VerifyToken Verify Token
func VerifyToken(tokenString string) (*st.UserAuthClaim, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &st.UserAuthClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if claims, ok := token.Claims.(*st.UserAuthClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid Token")
}
