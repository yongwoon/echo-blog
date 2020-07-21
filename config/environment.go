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

// IsPRD prod?
func IsPRD() bool {
	if Env() == PRD {
		return true
	}
	return false
}

// IsDEV dev?
func IsDEV() bool {
	if Env() == DEV {
		return true
	}
	return false
}

// IsLOCAL local?
func IsLOCAL() bool {
	if Env() == LOCAL {
		return true
	}
	return false
}

// IsTEST test?
func IsTEST() bool {
	if Env() == TEST {
		return true
	}
	return false
}

// Env get environment
func Env() string {
	if env := os.Getenv("GO_ENV"); env == TEST || env == DEV || env == PRD {
		return env
	}
	return LOCAL
}
