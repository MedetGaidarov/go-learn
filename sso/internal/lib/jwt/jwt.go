// NewToken creates new JWT token for given user and app.

package jwt

import (
	"go-learn/sso/internal/domain/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Добавляем в токен всю необходимую информацию
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app_id"] = app.ID

	// Подписываем токен, используя секретный ключ приложения
	tokenString, err := token.SignedString([]byte(app.Secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}