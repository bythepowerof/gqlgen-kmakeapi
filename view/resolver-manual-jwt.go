package gqlgen_kmakeapi

//go:generate go run github.com/99designs/gqlgen
import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func (r *mutationResolver) CreateJwt(ctx context.Context, input NewJwt) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":  input.User,
		"roles": input.Roles,

		"iss": Issuer,
		"sub": "gqlgen properties",
		"aud": "gqlgen",
		"exp": time.Now().Add(time.Minute * ExpiryMins).Unix(),
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		// iss	Issuer			Identifies principal that issued the JWT.
		// sub	Subject			Identifies the subject of the JWT.
		// aud	Audience		Identifies the recipients that the JWT is intended for. Each principal intended to process the JWT must identify itself with a value in the audience claim. If the principal processing the claim does not identify itself with a value in the aud claim when this claim is present, then the JWT must be rejected.
		// exp	Expiration Time	Identifies the expiration time on and after which the JWT must not be accepted for processing. The value must be a NumericDate:[9] either an integer or decimal, representing seconds past 1970-01-01 00:00:00Z.
		// nbf	Not Before		Identifies the time on which the JWT will start to be accepted for processing. The value must be a NumericDate.
		// iat	Issued at		Identifies the time at which the JWT was issued. The value must be a NumericDate.
		// jti	JWT ID			Case sensitive unique identifier of the token even among different issuers.
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(r.JwtSecret))

	return tokenString, err
}
