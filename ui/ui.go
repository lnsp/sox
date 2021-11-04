package ui

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var distFS embed.FS
var FS, _ = fs.Sub(distFS, "dist")
