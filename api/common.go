package api

// Response https://msx.benzac.de/wiki/index.php?title=Responses
type Response struct {
	Status  int         `json:"status"`
	Text    string      `json:"text,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

// Colors https://msx.benzac.de/wiki/index.php?title=Colors

// Icons https://msx.benzac.de/wiki/index.php?title=Icons

// Inline Expressions https://msx.benzac.de/wiki/index.php?title=Inline_Expressions

// Actions https://msx.benzac.de/wiki/index.php?title=Actions
