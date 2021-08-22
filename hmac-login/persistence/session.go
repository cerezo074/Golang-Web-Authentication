package persistence

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	. "example/models"
	"fmt"
	"net/http"
	"strings"
)

type session map[string]string

const (
	privateKey = "Save this key in your cloud infrastructure instead of hardcoding it"
)

var (
	tokens *session = &session{}
)

func RegisterToken(user User) (string, error) {
	if !user.IsValidID() {
		return "", fmt.Errorf("Invalid user to create session, %v", user.Email)
	}

	h := hmac.New(sha256.New, []byte(privateKey))
	h.Write([]byte(user.Username))
	token := fmt.Sprintf("%x", h.Sum(nil))
	(*tokens)[user.Email] = token

	return token, nil
}

func DeleteToken(user User) error {
	if !user.IsValidID() {
		return fmt.Errorf("Invalid user to delete session, %v", user.Email)
	}

	delete(*tokens, user.Email)

	return nil
}

func UserForSession(request *http.Request) (*User, error) {
	authHeader := request.Header.Get("Authorization")
	return userForToken(authHeader)
}

func userForToken(token string) (*User, error) {
	if strings.Contains(token, "Bearer") {
		tokenWithoutPrefix := strings.ReplaceAll(token, "Bearer", "")
		rawToken := strings.ReplaceAll(tokenWithoutPrefix, " ", "")
		email := (*tokens)[rawToken]
		return GetUser(email)
	}

	return nil, errors.New("Invalid token type")
}
