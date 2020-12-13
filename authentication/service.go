package authentication

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/brother14th/locationmapping/db"

)

//Service provides authentication
type Service interface {
	Authenticate(string, string) (string, error)
}

type service struct {
	userRepository db.UserRepository
	key            []byte
}

func (s *service) Authenticate(username string, password string) (string, error) {
	if username == "" {
		return "", errors.New("empty username")
	}
	if password == "" {
		return "", errors.New("empty password")
	}

	isValid, err := s.userRepository.ValidateUser(username, password)
	if !isValid {
		return "", errors.New("invalid username or password")
	}

	token, err := createToken(s.key, username)
	if err != nil {
		return "", errors.New("error creating token")
	}
	return token, nil
}

//NewService creates a new service
func NewService(userRepository db.UserRepository, key []byte) Service {
	return &service{
		userRepository: userRepository,
		key:            key,
	}
}

func createToken(key []byte, username string) (string, error) {
	expirationTime := time.Now().Add(20 * time.Minute)

	//create JWT claims
	claims := jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: expirationTime.Unix(),
	}

	// create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
