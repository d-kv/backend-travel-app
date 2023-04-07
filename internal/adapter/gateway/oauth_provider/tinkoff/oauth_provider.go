// TODO: use singletones for permanent request parameters
// TODO: add logging
package tinkoff

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/d-kv/backend-travel-app/pkg/adapter/igateway"
	"github.com/d-kv/backend-travel-app/pkg/infra/ilogger"
)

const introspectURL = "https://id.tinkoff.ru/auth/introspect"

type OAuthProvider struct {
	id     string
	secret string
	log    ilogger.LoggerI
	clt    *http.Client
}

func New(l ilogger.LoggerI, c *http.Client) *OAuthProvider {
	return &OAuthProvider{
		log: l,
		clt: c,
	}
}

var _ igateway.OAuthProviderI = (*OAuthProvider)(nil)

func (p *OAuthProvider) GetUserID(ctx context.Context, aToken string) (string, error) {
	bodyRdr := bytes.NewReader([]byte(fmt.Sprintf("token=%s", aToken)))
	req, err := http.NewRequestWithContext(
		ctx,
		"POST", //nolint:usestdlibvars // Using headers with the request
		introspectURL,
		bodyRdr,
	)
	if err != nil {
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

	resp, err := p.clt.Do(req) // TODO: add timeout
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var respJSON Response
	err = json.NewDecoder(resp.Body).Decode(&respJSON)
	if err != nil {
		return "", err
	}
	if !respJSON.Active { // HACK: can we rely only on this field?
		return "", igateway.ErrTokenIsExpired
	}
	return respJSON.CliendID, nil
}
