package uaa

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func RenewAccessToken(host string, refreshToken string) (string) {
	endpoint := fmt.Sprintf("https://%s:8443/oauth/token", host)
	params := url.Values{}
	params.Add("client_id", `pks_cli`)
	params.Add("client_secret", ``)
	params.Add("grant_type", `refresh_token`)
	params.Add("refresh_token", refreshToken)
	params.Add("response_type", `token`)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {

		payload := RefreshTokenResp{}
		if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
            log.Fatal(err)
        }

		return payload.AccessToken
	}
	return ""
}

type RefreshTokenResp struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	Jti          string `json:"jti"`
}

func GetTokenRemainingValidity(token string) int {
	
	parsedToken, _ := jwt.Parse(token, nil)
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {

		tm := time.Unix(int64(claims["exp"].(float64)), 0)
		fmt.Println("Time the token expires: ", tm)
		fmt.Println("Time now:               ", time.Now())

	    remainder := tm.Sub(time.Now())

		if remainder > 0 {
			return int(remainder.Seconds())
		}
	}
	return 0
}
