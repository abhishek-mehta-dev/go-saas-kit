package embedfs

import "embed"

//go:embed all:../../web/templates
var TemplatesFS embed.FS
