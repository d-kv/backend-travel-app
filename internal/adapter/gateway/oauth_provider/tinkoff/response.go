package tinkoff

import "time"

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
