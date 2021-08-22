package persistence

import (
	"errors"
	. "example/models"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type session map[string]string

type jwtPayload struct {
	jwt.StandardClaims
	Id string
}

const (
	privateKey = "Save this key in your cloud infrastructure instead of hardcoding it"
)

var (
	uuids *session = &session{}
)

func RegisterToken(user User) (string, error) {
	if !user.IsValidID() {
		return "", errors.New("invalid user to create session")
	}

	oldToken := currentToken(user)
	if oldToken != "" {
		delete(*uuids, oldToken)
	}

	rawUuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	UUID := rawUuid.String()
	token, err := createJWTToken(UUID)
	if err != nil {
		return "", err
	}

	(*uuids)[UUID] = user.Email
	return token, nil
}

func DeleteToken(request *http.Request) error {
	token, err := extractToken(request)
	if err != nil {
		return err
	}

	uuid, err := getUUIDFromSession(token)
	if err != nil {
		return err
	}

	if (*uuids)[uuid] == "" {
		return errors.New("invalid user to delete session")
	}

	delete(*uuids, uuid)

	return nil
}

func UserForSession(request *http.Request) (*User, error) {
	token, err := extractToken(request)
	if err != nil {
		return nil, err
	}

	uuid, err := getUUIDFromSession(token)
	if err != nil {
		return nil, err
	}

	email := (*uuids)[uuid]
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
	for key, value := range *uuids {
		if value == user.Email {
			return key
		}
	}

	return ""
}

func createJWTToken(id string) (string, error) {
	payload := jwtPayload{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
		Id: id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &payload)
	ss, err := token.SignedString([]byte(privateKey))

	if err != nil {
		return "", fmt.Errorf("couldn't SignedString %w", err)
	}

	return ss, nil
}

func getUUIDFromSession(token string) (string, error) {
	if token == "" {
		return "", errors.New("token if empty")
	}

	afterVerificationToken, err := jwt.ParseWithClaims(token, &jwtPayload{}, func(beforeVerificationToken *jwt.Token) (interface{}, error) {
		if beforeVerificationToken.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", beforeVerificationToken.Header["alg"])
		}

		return []byte(privateKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := afterVerificationToken.Claims.(*jwtPayload); ok && afterVerificationToken.Valid {
		return claims.Id, nil
	} else {
		return "", errors.New("can't extract session ID from current token")
	}
}
