package usecase

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"github.com/wesmota/go-jobsity-chat-server/models"
	"github.com/wesmota/go-jobsity-chat-server/usecase/storage"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo storage.AuthorizationRepo
}

func NewAuthService(repo storage.AuthorizationRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (models.Login, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil || user.ID == 0 {
		return models.Login{}, errors.New("invalid credentials")
	}
	jwtTTL, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	if err != nil {
		return models.Login{}, err
	}
	expiresAt := time.Now().Add(time.Hour * time.Duration(jwtTTL)).Unix()
	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		return models.Login{}, errors.New("invalid credentials")
	}
	tkObj := &models.Token{
		UserID:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tkObj)
	jwtSecret := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return models.Login{}, err
	}
	return models.Login{
		Token: tokenString,
	}, nil

}

func (s *AuthService) SignUp(ctx context.Context, user models.User) (models.Login, error) {
	userDB, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err == nil && userDB.ID != 0 {
		return models.Login{}, errors.New("email already exists")
	}

	jwtTTL, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	if err != nil {
		return models.Login{}, err
	}
	expiresAt := time.Now().Add(time.Hour * time.Duration(jwtTTL)).Unix()
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.Login{}, err
	}

	hashStr := string(hashPwd)
	user.Password = hashStr
	log.Info().Interface("user", user).Msg("SignUp")

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return models.Login{}, err
	}
	tk := &models.Token{
		UserID:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	jwtSecret := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return models.Login{}, err
	}

	return models.Login{
		Token: tokenString,
	}, nil

}
