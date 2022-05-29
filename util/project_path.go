package util

import (
	"path/filepath"
	"runtime"
)

// Package utils contains some common paths used in configuration and tests
var (
	_, b, _, _ = runtime.Caller(0)
	// ProjectRoot Root folder of this project
	ProjectRoot = filepath.Join(filepath.Dir(b), "/..")
)
