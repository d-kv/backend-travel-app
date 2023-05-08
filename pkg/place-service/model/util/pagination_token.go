package util

const (
	DefaultSkipN = 0
	DefaultResN  = 0
)

type PaginationToken struct {
	SkipN int64
	ResN  int64
}

type PaginationTokenOptions func(*PaginationToken)

func NewPaginationToken(skipN, resN int64) *PaginationToken {
	return &PaginationToken{
		SkipN: skipN,
		ResN:  resN,
	}
}
