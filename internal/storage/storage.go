package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrUrlExist    = errors.New("db url exist")
)
