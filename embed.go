package static

import "embed"

var (
	//go:embed views/templates
	Templates embed.FS
)
