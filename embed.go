package static

import "embed"

var (
	//go:embed assets
	Templates embed.FS
)
