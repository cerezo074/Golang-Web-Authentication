package persistence

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	. "example/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
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
		return "", errors.New("invalid user to create session")
	}

	oldToken := currentToken(user)
	if oldToken != "" {
		delete(*tokens, oldToken)
	}

	h := hmac.New(sha256.New, []byte(privateKey))
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	h.Write([]byte(uuid.String()))
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

func currentToken(user User) string {
	for key, value := range *tokens {
		if value == user.Email {
			return key
		}
	}

	return ""
}
