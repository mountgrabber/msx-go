package api

// SelectionObject https://msx.benzac.de/wiki/index.php?title=Selection_Object
type SelectionObject struct {
	Important  bool        `json:"important,omitempty"`
	Headline   string      `json:"headline,omitempty"`
	Background string      `json:"background,omitempty"`
	Action     interface{} `json:"action,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
