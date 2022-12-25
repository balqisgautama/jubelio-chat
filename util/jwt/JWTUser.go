package jwt

import (
	"encoding/json"
	"github.com/balqisgautama/jubelio-chat/config"
	"github.com/balqisgautama/jubelio-chat/constanta"
	res "github.com/balqisgautama/jubelio-chat/dto/response"
	"github.com/balqisgautama/jubelio-chat/model"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type PayloadJWTUser struct {
	ClientID string `json:"cid"`
	jwt.StandardClaims
}

func GenerateJWTUser(userID int64, clientID string, issuer string) (token string, output res.APIResponse) {
	tokenCode := PayloadJWTUser{
		ClientID: clientID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constanta.Time8Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
			Subject:   strconv.Itoa(int(userID)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenCode)
	token, err := jwtToken.SignedString([]byte(config.ApplicationConfiguration.GetJWTKey()))
	if err != nil {
		output = model.GenerateJWTError("JWTUser.go", "GenerateJWTUser", err)
		return
	}
	return
}

func ValidatorJWTUser(jwtToken string) (tokenData *PayloadJWTUser, output res.APIResponse) {

	claims := &PayloadJWTUser{}
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ApplicationConfiguration.GetJWTKey()), nil
	})

	if err != nil {
		output = model.GenerateJWTError("JWTUser.go", "ValidatorJWTUser", err)
		return
	}

	tokenData = convertTokenData(token.Claims.(*PayloadJWTUser))
	return
}

func convertTokenData(input interface{}) *PayloadJWTUser {
	bolB, _ := json.Marshal(input)
	tokenData := PayloadJWTUser{}
	json.Unmarshal(bolB, &tokenData)
	return &tokenData
}
