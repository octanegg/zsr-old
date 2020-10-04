package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/octanegg/zsr/internal/config"
)

// Response .
type Response struct {
	Message string `json:"message"`
}

// Jwks .
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys .
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// CustomClaims .
type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

const (
	audience = "zsr.octane.gg"
	domain   = "https://octanegg.us.auth0.com/"
	key      = "https://octanegg.us.auth0.com/.well-known/jwks.json"
	scope    = "modify:admin"
)

func jwtMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {

			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(audience, false)
			if !checkAud {
				return token, errors.New("invalid audience")
			}

			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(domain, false)
			if !checkIss {
				return token, errors.New("invalid issuer")
			}

			cert, err := getPemCert(token)
			if err != nil {
				panic(err.Error())
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	}).CheckJWT(w, r)

	if err == nil && next != nil {
		next(w, r)
	}
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(key)

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

func scopeMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, _ := jwt.ParseWithClaims(strings.Split(r.Header.Get(config.HeaderAuthorization), " ")[1], &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		cert, err := getPemCert(token)
		if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	})

	result := strings.Split(token.Claims.(*CustomClaims).Scope, " ")
	for i := range result {
		if result[i] == scope {
			next.ServeHTTP(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusForbidden)
}
