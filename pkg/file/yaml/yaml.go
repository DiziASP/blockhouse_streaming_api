package yaml

import (
	"gopkg.in/yaml.v3"
)

// YAML
var (
	Marshal    = yaml.Marshal
	Unmarshal  = yaml.Unmarshal
	NewDecoder = yaml.NewDecoder
	NewEncoder = yaml.NewEncoder
)
