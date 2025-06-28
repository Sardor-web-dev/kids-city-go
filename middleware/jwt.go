// middleware/jwt.go или в init
package middleware

import (
	"os"
)

var JWTKey []byte

func InitJWTSecret() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET is not set")
	}
	JWTKey = []byte(secret)
}
