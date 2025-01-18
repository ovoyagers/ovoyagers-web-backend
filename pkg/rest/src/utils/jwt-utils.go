package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/petmeds24/backend/config"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/tailscale/golang-x-crypto/bcrypt"
)

type JWTUtil struct {
	cfg         *config.Config
	redisClient *redis.Client
}

var (
	ACCESS_TOKEN_EXPIRY  = 15 * time.Minute    // 15 minutes
	REFRESH_TOKEN_EXPIRY = 15 * 24 * time.Hour // 15 days
)

func NewJWTUtil() *JWTUtil {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
		panic(err)
	}
	redisClient := config.NewRedisConfig().GetRedisClient()
	return &JWTUtil{cfg: cfg, redisClient: redisClient}
}

type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

func (ju *JWTUtil) CreateToken(userID string, email string) (*TokenDetails, error) {
	acToken := ju.cfg.ACCESS_TOKEN_SECRET
	rfToken := ju.cfg.REFRESH_TOKEN_SECRET

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(ACCESS_TOKEN_EXPIRY).Unix()
	td.RtExpires = time.Now().Add(REFRESH_TOKEN_EXPIRY).Unix()

	accessTokenClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"email":      email,
		"exp":        td.AtExpires,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(acToken))
	if err != nil {
		return nil, err
	}
	td.AccessToken = accessTokenString

	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     td.RtExpires,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(rfToken))
	if err != nil {
		return nil, err
	}
	td.RefreshToken = refreshTokenString

	return td, nil
}
func (ju *JWTUtil) ValidateAccessToken(tokenString string) (map[string]interface{}, error) {
	acToken := ju.cfg.ACCESS_TOKEN_SECRET
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(acToken), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, err
	}
	userInfo := map[string]interface{}{
		"id":    claims["user_id"],
		"email": claims["email"],
	}
	return userInfo, nil
}

func (ju *JWTUtil) ValidateRefreshToken(tokenString string) (map[string]interface{}, error) {
	rfToken := ju.cfg.REFRESH_TOKEN_SECRET
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(rfToken), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, err
	}
	userInfo := map[string]interface{}{
		"id":    claims["user_id"],
		"email": claims["email"],
	}
	return userInfo, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateOTP generates a random n-digit OTP using a cryptographically secure random number generator.
// The OTP is a string of 'n' random digits.
// If n is less than 1, GenerateOTP returns an error.
// The generated OTP is a string of digits, not a number.
func GenerateOTP(n int) (otp string, err error) {
	if n < 1 {
		return "", fmt.Errorf("n should be greater than 0")
	}

	otpBytes := make([]byte, n)
	for i := range otpBytes {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", fmt.Errorf("failed to generate secure random number: %v", err)
		}
		otpBytes[i] = '0' + byte(num.Int64())
	}

	return string(otpBytes), nil
}
