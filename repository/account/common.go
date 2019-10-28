package account

import (
	"math/rand"
	"time"

	"github.com/daoraimi/dagger/api"
	"github.com/daoraimi/dagger/config"
	"github.com/daoraimi/dagger/share"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func GenerateRandomPassword(length int) []byte {
	var pwd []byte
	peLength := len(share.PasswordElements)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		pwd = append(pwd, share.PasswordElements[rand.Intn(peLength)])
	}
	return pwd
}

func EncryptUserPassword(plainPassword []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(plainPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(hashedPassword), nil
}

func ValidateCredential(plainPassword, hashedPassword []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, plainPassword); err != nil {
		return false
	}
	return true
}

func GenerateUserToken(userID uint64) (string, error) {
	tokenSecret := config.GetString("token.secret")
	durationString := config.GetString("token.expireDuration")
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return "", errors.WithStack(err)
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&api.TokenClaim{
			userID,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(duration).Unix(),
			},
		},
	)

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", errors.WithStack(err)
	}

	return tokenString, nil
}
