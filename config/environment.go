package config

import (
	"os"
)

const (
	// LOCAL local environment
	LOCAL = "local"
	// TEST test environment
	TEST = "test"
	// DEV development environment
	DEV = "development"
	// PRD production environment
	PRD = "production"
)

// IsPRD 本番かどうか
func IsPRD() bool {
	if Env() == PRD {
		return true
	}
	return false
}

// IsDEV 開発かどうか
func IsDEV() bool {
	if Env() == DEV {
		return true
	}
	return false
}

// IsLOCAL ローカル環境かどうか
func IsLOCAL() bool {
	if Env() == LOCAL {
		return true
	}
	return false
}

// IsTEST テスト環境かどうか
func IsTEST() bool {
	if Env() == TEST {
		return true
	}
	return false
}

// Env ステージを取得
func Env() string {
	if env := os.Getenv("GO_ENV"); env == TEST || env == DEV || env == PRD {
		return env
	}
	return LOCAL
}
