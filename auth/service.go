package auth

import "github.com/golang-jwt/jwt/v5"

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct {
	//
}

var SECRET_KEY = []byte("STARTUP_s3cr3T_k3Y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}