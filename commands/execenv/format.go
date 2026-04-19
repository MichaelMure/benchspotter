package execenv

import "github.com/thediveo/enumflag/v2"

type Format int

const (
	FormatText Format = iota
	FormatJSON
	FormatRaw
)

var FormatIds = map[Format][]string{
	FormatText: {"text"},
	FormatJSON: {"json"},
	FormatRaw:  {"raw"},
}

var FormatHelp = enumflag.Help[Format]{
	FormatText: "human-readable text",
	FormatJSON: "structured JSON for programmatic use",
	FormatRaw:  "native data format, pipeable to standard tools",
}
