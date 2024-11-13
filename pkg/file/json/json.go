package json

import (
	"github.com/goccy/go-json"
)

// JSON metadata
var (
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
)
