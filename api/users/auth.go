package users

import (
	"fmt"
	"time"

	"github.com/bk7987/timecards/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// JWTPair defines the return value for generating a JWT pair.
type JWTPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// GenerateSignedToken returns a signed JWT.
func GenerateSignedToken(user *User, expires time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = expires.Unix()

	signedToken, err := token.SignedString([]byte(config.GetConfig().JWTSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// generateJWTPair generates an access token and a refresh token for the given user.
func generateJWTPair(user *User) (*JWTPair, error) {
	accessToken, err := GenerateSignedToken(user, time.Now().Add(config.GetConfig().JWTExpiration))
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateSignedToken(user, time.Now().Add(config.GetConfig().JWTRefreshExpiration))
	if err != nil {
		return nil, err
	}

	return &JWTPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// ParseJWT parses a given JWT and verifies the signing method.
func ParseJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, err
}

// SetRefreshTokenCookie sets the refresh token cookie in the provided response context
func SetRefreshTokenCookie(ctx *fiber.Ctx, cookieValue string) {
	cookie := new(fiber.Cookie)
	cookie.Name = "refreshToken"
	cookie.Value = cookieValue
	cookie.Expires = time.Now().Add(config.GetConfig().JWTRefreshExpiration)
	cookie.HTTPOnly = true
	ctx.Cookie(cookie)
}

func (u *User) setPassword(password string) error {
	minLength := config.GetConfig().MinPasswordLength
	if len(password) < minLength {
		return fmt.Errorf("Password must be longer than %v characters", minLength)
	}

	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)
	return nil
}

// CheckPassword checks the provided password against the decrypted password belonging to the user.
func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
