package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// JwtHelper ...
type JwtHelper struct {
	jwt.StandardClaims
}

// CreateJwtToken ...
func (self *JwtHelper) CreateJwtToken(secret string, ID string) (string, error) {
	claims := JwtHelper{}
	claims.Id = ID
	claims.ExpiresAt = time.Now().Add(30 * time.Minute).Unix()
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, err
}

func (u *JwtHelper) GetJwtClaims(c echo.Context) jwt.MapClaims {
	user := c.Get("ID")
	if user == nil {
		return nil
	}

	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	return claims
}

func (u *JwtHelper) GetJwtClaim(c echo.Context, key string) interface{} {
	claims := u.GetJwtClaims(c)

	return claims[key]
}
