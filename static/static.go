package static

import "embed"

//go:embed *
var Css embed.FS

//go:embed cat.png
var Cat embed.FS
