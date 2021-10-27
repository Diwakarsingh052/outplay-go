package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

const (
	RoleAdmin = "ADMIN"
	RoleUser  = "USER"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// Key is used to store/retrieve a Claims value from a context.Context.
const Key ctxKey = 1

type Claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

// HasRole returns true if the claims has at least one of the provided roles.
func (c Claims) HasRole(roles ...string) bool {
	for _, has := range c.Roles {
		for _, want := range roles {
			if has == want {
				return true
			}
		}
	}
	return false
}

// Valid is called during the parsing of a token.
func (c Claims) Valid() error {
	if err := c.RegisteredClaims.Valid(); err != nil {
		return fmt.Errorf("validating standard claims %w", err)
	}

	return nil
}

type Auth struct {
	privateKey *rsa.PrivateKey
	algorithm  string
	parser     *jwt.Parser
}

func NewAuth(privateKey *rsa.PrivateKey, algorithm string) (*Auth, error) {
	if privateKey == nil {
		return nil, errors.New("private key cannot be nil")
	}
	if jwt.GetSigningMethod(algorithm) == nil {
		return nil, fmt.Errorf("unknown algorithm %v", algorithm)
	}

	parser := jwt.Parser{
		ValidMethods: []string{algorithm}, // all valid algorithms thant we can use
	}

	a := Auth{
		privateKey: privateKey,
		algorithm:  algorithm,

		parser: &parser,
	}

	return &a, nil

}

// GenerateToken generates a signed JWT token string representing the user Claims.
func (a *Auth) GenerateToken(claims Claims) (string, error) {
	method := jwt.GetSigningMethod(a.algorithm)

	tkn := jwt.NewWithClaims(method, claims) // Declare the token with the algorithm used for signing, and the claims

	tokenStr, err := tkn.SignedString(a.privateKey)
	if err != nil {
		return "", fmt.Errorf("signing token %w", err)
	}

	return tokenStr, nil
}

// ValidateToken recreates the Claims that were used to generate a token. It
// verifies that the token was signed using our key.
func (a *Auth) ValidateToken(tokenStr string) (Claims, error) {

	var claims Claims
	token, err := a.parser.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return a.privateKey.Public(), nil
		//return []byte("hola"), nil
	})

	if err != nil {
		return Claims{}, fmt.Errorf("parsing token %w", err)
	}

	if !token.Valid {
		return Claims{}, errors.New("invalid token")
	}

	return claims, nil
}
