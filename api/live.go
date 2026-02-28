package api

// LiveObject https://msx.benzac.de/wiki/index.php?title=Live_Object
type LiveObject struct {
	Type              string      `json:"type,omitempty"`
	From              int         `json:"from,omitempty"`
	To                int         `json:"to,omitempty"`
	Duration          int         `json:"duration,omitempty"`
	Delay             int         `json:"delay,omitempty"`
	Source            string      `json:"source,omitempty"`
	Coming            interface{} `json:"coming,omitempty"`
	Running           interface{} `json:"running,omitempty"`
	Over              interface{} `json:"over,omitempty"`
	Execute           interface{} `json:"execute,omitempty"`
	Preload           bool        `json:"preload,omitempty"`
	Color             string      `json:"color,omitempty"`
	Title             string      `json:"title,omitempty"`
	TitleHeader       string      `json:"titleHeader,omitempty"`
	TitleFooter       string      `json:"titleFooter,omitempty"`
	Label             string      `json:"label,omitempty"`
	Icon              string      `json:"icon,omitempty"`
	Headline          string      `json:"headline,omitempty"`
	Text              string      `json:"text,omitempty"`
	Tag               string      `json:"tag,omitempty"`
	TagColor          string      `json:"tagColor,omitempty"`
	Badge             string      `json:"badge,omitempty"`
	BadgeColor        string      `json:"badgeColor,omitempty"`
	Stamp             string      `json:"stamp,omitempty"`
	StampColor        string      `json:"stampColor,omitempty"`
	Progress          int         `json:"progress,omitempty"`
	ProgressColor     string      `json:"progressColor,omitempty"`
	ProgressBackColor string      `json:"progressBackColor,omitempty"`
	WrapperColor      string      `json:"wrapperColor,omitempty"`
	Image             string      `json:"image,omitempty"`
	ExtensionIcon     string      `json:"extensionIcon,omitempty"`
	ExtensionLabel    string      `json:"extensionLabel,omitempty"`
	Action            string      `json:"action,omitempty"`
	Data              interface{} `json:"data,omitempty"`
}

// LiveInlineExpressions https://msx.benzac.de/wiki/index.php?title=Live_Inline_Expressions
