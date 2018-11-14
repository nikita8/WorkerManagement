package siam

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt"
	"github.com/graniticio/granitic/iam"
)

type JwtConstruct struct {
	Algo   string
	NValue string
}

func DefaultJwtConstruct() JwtConstruct {
	jwtConstruct := JwtConstruct{}
	jwtConstruct.Algo = "RS256"
	jwtConstruct.NValue = "wHHNfEYo9R256uCwO5l0SzxDlo8KpNbI6JSINdfJg6kwcgQwTxs-nRT-GH35vPFrx7NPeXzJUPoidBnGptG9TbWRVahE0s7dWDBpHDrsAGX9s6DyWTM9K0W3ToFW2YCY7FifgYKrM2StVrXfi6vNTBQY1RknrVoVZxTYJuj-GbOCo0NVdeOLB304OnzZ8-GEfhpmxYgdnQHTztEAS2ecVuNjV7PFJ0ycVVzf17JIImI3ai6i87Y58__3IGnhklG7aK1OYq12LsYepGe88vcT5d5oWW1ldtlaiFAfcKsTdJ_jki8oPHuSzxQjdS2YGxOa_asQ0y-PeZ84kbDgmq6YdQ"
	return jwtConstruct
}

type UserIdentifier struct {
}

func (u *UserIdentifier) Identify(ctx context.Context, req *http.Request) (iam.ClientIdentity, context.Context) {
	//(iam.ClientIdentity, context.Context)
	anonym := iam.NewAnonymousIdentity()
	authHeader := req.Header.Get("Authorization")

	if authHeader == "" {
		return anonym, ctx
	}

	accessAndIdToken := strings.Split(authHeader, " ")[1]
	accessToken := strings.Split(accessAndIdToken, " : ")[0]

	publicPem := jwkToPem(DefaultJwtConstruct().NValue)

	validToken, errMsg := parseAndValidate(accessToken, publicPem)

	var userIdentity iam.ClientIdentity
	if validToken {
		userIdentity = iam.NewAuthenticatedIdentity("email")
	} else {
		userIdentity = anonym
		fmt.Println(errMsg)
	}

	return userIdentity, ctx

	//return iam.ClientIdentity, ctx
}

type Token struct {
	*jwt.JWT
	IsLoggedIn  bool   `json:"isLoggedIn"`
	CustomField string `json:"customField,omitempty"`
}

// state: 0 => Error, 1 => Success
func parseAndValidate(accessToken string, publicPem rsa.PublicKey) (valid bool, msg string) {

	now := time.Now()
	rs256 := jwt.NewRS256(nil, &publicPem)

	// First, extract the payload and signature.
	// This enables unmarshaling the JWT first and
	// verifying it later or vice versa.
	payload, sig, err := jwt.Parse(accessToken)
	if err != nil {
		return false, err.Error()
	}
	if err = rs256.Verify(payload, sig); err != nil {
		return false, err.Error()
	}
	var jot Token
	if err = jwt.Unmarshal(payload, &jot); err != nil {
		return false, err.Error()
	}

	// Validate fields.
	iatValidator := jwt.IssuedAtValidator(now)
	expValidator := jwt.ExpirationTimeValidator(now)
	//audValidator := jwt.AudienceValidator("admin")

	//if err = jot.Validate(iatValidator, expValidator, audValidator); err != nil {
	if err = jot.Validate(iatValidator, expValidator); err != nil {
		switch err {
		case jwt.ErrIatValidation:
			// handle "iat" validation error
			return false, err.Error()
		case jwt.ErrExpValidation:
			// handle "exp" validation error
			return false, err.Error()
		case jwt.ErrAudValidation:
			// handle "aud" validation error
			return false, err.Error()
		}
	}

	return true, ""
}

func jwkToPem(nStr string) rsa.PublicKey {

	var e uint64
	mockRSA := rsa.PublicKey{N: big.NewInt(0), E: int(e)}

	decN, err := base64.RawURLEncoding.DecodeString(nStr)
	if err != nil {
		fmt.Println(err)
		return mockRSA
	}
	n := big.NewInt(0)
	n.SetBytes(decN)

	eStr := "AQAB"
	decE, err := base64.RawURLEncoding.DecodeString(eStr)
	if err != nil {
		fmt.Println(err)
		return mockRSA
	}
	var eBytes []byte
	if len(decE) < 8 {
		eBytes = make([]byte, 8-len(decE), 8)
		eBytes = append(eBytes, decE...)
	} else {
		eBytes = decE
	}
	eReader := bytes.NewReader(eBytes)

	err = binary.Read(eReader, binary.BigEndian, &e)
	if err != nil {
		fmt.Println(err)
		return mockRSA
	}
	return rsa.PublicKey{N: n, E: int(e)}
}
