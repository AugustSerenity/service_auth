package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/AugustSerenity/service_auth/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//var jwtSecret = []byte("secret_key")

type Service struct {
	storage   Storage
	jwtSecret []byte
}

func New(st Storage, jwtSecret []byte) *Service {
	return &Service{
		storage:   st,
		jwtSecret: jwtSecret,
	}
}

func (s *Service) CreateToken(ctx context.Context, userID, ip string) (string, string, error) {
	accessToken, jti, err := s.generateAccessToken(userID, ip)
	if err != nil {
		return "", "", err
	}

	rawRefresh := generateRandomString(32)
	encodedRefresh := base64.StdEncoding.EncodeToString([]byte(rawRefresh))

	hashedRefresh, _ := bcrypt.GenerateFromPassword([]byte(encodedRefresh), bcrypt.DefaultCost)

	refresh := model.RefreshToken{
		UserID:      userID,
		AccessJTI:   jti,
		HashedToken: string(hashedRefresh),
		IP:          ip,
		Used:        false,
	}

	fmt.Print(refresh)
	s.storage.SaveToken(ctx, &refresh)

	return accessToken, encodedRefresh, nil
}

func (s *Service) generateAccessToken(userID, ip string) (string, string, error) {
	jti := generateRandomString(16)
	claims := model.Claims{
		UserID: userID,
		IP:     ip,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	return tokenString, jti, err
}

func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("failed to generate random string")
	}
	return hex.EncodeToString(bytes)[:length]
}

func (s *Service) RefreshToken(ctx context.Context, accessToken, refreshToken, currentIP string) (string, string, error) {
	claims, err := ParseJWT(accessToken)
	if err != nil {
		return "", "", fmt.Errorf("invalid access token: %w", err)
	}

	userID := claims.UserID
	accessJTI := claims.ID

	refreshFromDB, err := s.storage.FindRefreshToken(ctx, userID, accessJTI)
	if err != nil {
		return "", "", fmt.Errorf("refresh token not found: %w", err)
	}

	if refreshFromDB.Used {
		return "", "", errors.New("refresh token already used")
	}

	err = bcrypt.CompareHashAndPassword([]byte(refreshFromDB.HashedToken), []byte(refreshToken))
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	if refreshFromDB.IP != currentIP {
		go sendWarningEmailMock(userID, currentIP, refreshFromDB.IP)
	}

	if err = s.storage.MarkTokenUsed(ctx, refreshFromDB.ID); err != nil {
		return "", "", fmt.Errorf("could not mark token as used: %w", err)
	}

	newAccessToken, newJTI, err := s.generateAccessToken(userID, currentIP)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate new access token: %w", err)
	}

	rawRefresh := generateRandomString(32)

	encodedRefresh := base64.StdEncoding.EncodeToString([]byte(rawRefresh))

	hashedRefresh, err := bcrypt.GenerateFromPassword([]byte(encodedRefresh), bcrypt.DefaultCost)
	if err != nil {
		return "", "", fmt.Errorf("failed to hash refresh token: %w", err)
	}

	newRefresh := &model.RefreshToken{
		UserID:      userID,
		AccessJTI:   newJTI,
		HashedToken: string(hashedRefresh),
		IP:          currentIP,
		Used:        false,
	}

	if err = s.storage.SaveToken(ctx, newRefresh); err != nil {
		return "", "", fmt.Errorf("failed to save new refresh token: %w", err)
	}

	return newAccessToken, encodedRefresh, nil
}

func ParseJWT(tokenStr string) (*model.Claims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, &model.Claims{})
	if err != nil {
		return nil, errors.New("failed to parse token")
	}
	claims, ok := token.Claims.(*model.Claims)
	if !ok {
		return nil, errors.New("invalid claims in token")
	}
	return claims, nil
}

func sendWarningEmailMock(userID, oldIP, newIP string) error {
	message := fmt.Sprintf("Warning: IP address for user %s has changed!\nOld IP: %s\nNew IP: %s", userID, oldIP, newIP)

	log.Printf("Email Sent: %s", message)

	return nil
}

func generateJTI() string {
	randBytes := make([]byte, 16)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic("failed to generate random bytes for JTI")
	}
	return base64.StdEncoding.EncodeToString(randBytes)
}
