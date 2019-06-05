package usuario

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var (
	SignKey   *rsa.PrivateKey
	VerifyKey *rsa.PublicKey
	once      sync.Once
)

func init() {
	once.Do(func() {
		loadSignFiles()
	})
}

// loadSignFiles Carga la información de los certificados de firma y confirmación
func loadSignFiles() {
	signBytes, err := ioutil.ReadFile("./certificados/app.rsa")
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v", err)
	}

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v", err)
	}

	verifyBytes, err := ioutil.ReadFile("./certificados/app.rsa.pub")
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v", err)
	}

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v", err)
	}
}

// generateJWT genera un token JWT nuevo
func generateJWT(u Model) (string, error) {
	var token string
	c := Claim{
		Usuario: u,
		StandardClaims: jwt.StandardClaims{
			// Tiempo de expiración del token: 1 semana
			ExpiresAt: time.Now().Add(time.Hour * 24 * 1).Unix(),
			Issuer:    "Cursos EDteam",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	token, err := t.SignedString(SignKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidateJWT Middleware para validar los JWT token
func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tokenString string
		tokenString, err := getTokenFromAuthorizationHeader(c.Request())
		if err != nil {
			tokenString, err = getTokenFromURLParams(c.Request())
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "no se encontró el token")
			}
		}

		verifyFunction := func(token *jwt.Token) (interface{}, error) {
			return VerifyKey, nil
		}

		texto := ""
		token, err := jwt.ParseWithClaims(tokenString, &Claim{}, verifyFunction)
		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError:
				vErr := err.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					texto = "Su token ha expirado, por favor vuelva a ingresar"
				default:
					texto = "Error de validación del token"
				}
			default:
				texto = "Error al procesar el token"
			}

			return c.JSON(http.StatusBadRequest, texto)
		}

		if !token.Valid {
			return c.JSON(http.StatusBadRequest, "token no valido")
		}

		email := token.Claims.(*Claim).Usuario.Email

		c.Set("email", email)

		return next(c)
	}
}
