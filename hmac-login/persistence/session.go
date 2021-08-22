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
	(*tokens)[token] = user.Email

	return token, nil
}

func DeleteToken(request *http.Request) error {
	token, err := extractToken(request)
	if err != nil {
		return err
	}

	if (*tokens)[token] == "" {
		return errors.New("invalid user to delete session")
	}

	delete(*tokens, token)

	return nil
}

func UserForSession(request *http.Request) (*User, error) {
	token, err := extractToken(request)
	if err != nil {
		return nil, err
	}

	email := (*tokens)[token]
	return GetUser(email)
}

func extractToken(request *http.Request) (string, error) {
	authHeader := request.Header.Get("Authorization")

	if strings.Contains(authHeader, "Bearer") {
		tokenWithoutPrefix := strings.ReplaceAll(authHeader, "Bearer", "")
		rawToken := strings.ReplaceAll(tokenWithoutPrefix, " ", "")
		return rawToken, nil
	}

	return "", nil
}
