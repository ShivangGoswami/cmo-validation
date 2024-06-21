package auth

import (
	"encoding/json"
	"net/http"
	neturl "net/url"
	"strings"
)

type Token struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	TokenType        string `json:"token_type"`
	IdToken          string `json:"id_token"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	Scope            string `json:"scope"`
}

const (
	scope = "openid api.iam.service_accounts"
)

func AuthToken(id, secret string) (Token, error) {
	url := "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token"

	data := neturl.Values{}
	data.Add("client_id", id)
	data.Add("client_secret", secret)
	data.Add("scope", scope)
	data.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return Token{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Token{}, err
	}
	defer resp.Body.Close()

	var auth Token
	err = json.NewDecoder(resp.Body).Decode(&auth)
	if err != nil {
		return Token{}, err
	}
	return auth, nil
}
