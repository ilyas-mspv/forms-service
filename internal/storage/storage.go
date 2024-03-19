package storage

import "errors"

var (
	ErrDefault       = errors.New("db error")
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)
