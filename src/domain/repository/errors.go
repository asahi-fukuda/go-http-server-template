package repository

import "errors"

var (
	// ErrModelNotFound は Model が見つからない場合に発生するエラー
	ErrModelNotFound = errors.New("model is not found")
)
