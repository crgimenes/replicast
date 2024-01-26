package assets

import (
	"embed"
	"net/http"
)

//go:embed *.lua
var assets embed.FS

var FS = http.FS(assets)
