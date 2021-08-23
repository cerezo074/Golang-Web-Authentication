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

type userSession struct {
	email    string
	jwtToken string
}

type session map[string]userSession

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

	oldUUID := currentUUIDByUser(user)
	if oldUUID != "" {
		delete(*uuids, oldUUID)
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	UUID := newUUID.String()
	jwtToken, err := createJWTToken(UUID)
	if err != nil {
		return "", err
	}

	(*uuids)[UUID] = userSession{email: user.Email, jwtToken: jwtToken}
	return jwtToken, nil
}

func UserForSession(request *http.Request) (*User, error) {
	token, err := extractToken(request)
	if err != nil {
		return nil, err
	}

	uuid, err := getUUIDFromSession(token)
	if err != nil {
		removeInvalidToken(err, token)
		return nil, err
	}

	email := (*uuids)[uuid].email
	return GetUser(email)
}

func DeleteToken(request *http.Request) error {
	token, err := extractToken(request)
	if err != nil {
		return err
	}

	uuid, err := getUUIDFromSession(token)
	if err != nil {
		removeInvalidToken(err, token)
		return err
	}

	if (*uuids)[uuid].email == "" {
		return errors.New("invalid user to delete session")
	}

	delete(*uuids, uuid)

	return nil
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

func removeInvalidToken(err error, token string) {
	if validationError, ok := err.(*jwt.ValidationError); ok {
		validationError.Inner = errors.New("session was cancelled and removed")
		uuid := currentUUIDByJWTToken(token)
		delete(*uuids, uuid)
	}
}

func currentUUIDByUser(user User) string {
	for key, value := range *uuids {
		if value.email == user.Email {
			return key
		}
	}

	return ""
}

func currentUUIDByJWTToken(token string) string {
	for key, value := range *uuids {
		if value.jwtToken == token {
			return key
		}
	}

	return ""
}

func createJWTToken(id string) (string, error) {
	payload := jwtPayload{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Second).Unix(),
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

func getUUIDFromSession(jwtToken string) (string, error) {
	if jwtToken == "" {
		return "", errors.New("token if empty")
	}

	afterVerificationToken, err := jwt.ParseWithClaims(jwtToken, &jwtPayload{}, func(beforeVerificationToken *jwt.Token) (interface{}, error) {
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
