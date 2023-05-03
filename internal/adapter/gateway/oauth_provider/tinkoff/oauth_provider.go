// TODO: use singletones for permanent request parameters
package tinkoff

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/d-kv/backend-travel-app/pkg/adapter/gateway"
)

type Response struct {
	Active      bool      `json:"active"`
	Scope       []string  `json:"scope"`
	CliendAppID string    `json:"cliend_id"`
	TokenType   string    `json:"token_type"`
	ExpireAt    time.Time `json:"exp"`
	IssuedAt    time.Time `json:"iat"`
	UserID      string    `json:"sub"`
	Audience    []string  `json:"aud"`
	Issuer      string    `json:"iss"`
}

const introspectURL = "https://id.tinkoff.ru/auth/introspect"

type OAuthProvider struct {
	id     string
	secret string
	clt    *http.Client
}

func New(id, secret string, cl *http.Client) *OAuthProvider {
	return &OAuthProvider{
		id:     id,
		secret: secret,
		clt:    cl,
	}
}

var _ gateway.OAuthProviderI = (*OAuthProvider)(nil)

func (p *OAuthProvider) UserID(ctx context.Context, aToken string) (string, error) {
	bodyRdr := bytes.NewReader([]byte(fmt.Sprintf("token=%s", aToken)))
	req, err := http.NewRequestWithContext(
		ctx,
		"POST", //nolint:usestdlibvars // Using headers with the request
		introspectURL,
		bodyRdr,
	)
	if err != nil {
		log.Error().
			Err(err)
		return "", err
	}
	opaqueCreds := base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%s:%s",
			p.id,
			p.secret,
		),
	))

	req.Header.Add(
		"Authorization",
		fmt.Sprintf("Basic %s", opaqueCreds),
	)
	req.Header.Add(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	resp, err := p.clt.Do(req)
	if err != nil {
		log.Error().
			Err(err)
		return "", err
	}
	defer resp.Body.Close()

	var respJSON Response
	err = json.NewDecoder(resp.Body).Decode(&respJSON)
	if err != nil {
		log.Error().
			Err(err)
		return "", err
	}
	if !respJSON.Active || respJSON.CliendAppID != p.id {
		log.Info().
			Str("user_id", respJSON.UserID).
			Str("expired_at", respJSON.ExpireAt.String()).
			Msg("bad auth attempt")
		return "", gateway.ErrTokenIsExpired
	}
	return respJSON.UserID, nil
}
