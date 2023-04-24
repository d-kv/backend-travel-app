package ggonic

import "errors"

var (
	ErrAlreadyInContext = errors.New("already in the context")
	ErrNotInContext     = errors.New("not in the context")
)
