package api

// MenuRootObject https://msx.benzac.de/wiki/index.php?title=Menu_Root_Object
type MenuRootObject struct {
	Name        string        `json:"name,omitempty"`
	Version     string        `json:"version,omitempty"`
	Reference   string        `json:"reference,omitempty"`
	Flag        string        `json:"flag,omitempty"`
	Reuse       bool          `json:"reuse,omitempty"`
	Cache       bool          `json:"cache,omitempty"`
	Restore     bool          `json:"restore,omitempty"`
	Refocus     int           `json:"refocus,omitempty"`
	Transparent int           `json:"transparent,omitempty"`
	Style       string        `json:"style,omitempty"`
	Logo        string        `json:"logo,omitempty"`
	LogoSize    string        `json:"logoSize,omitempty"`
	Headline    string        `json:"headline,omitempty"`
	Background  string        `json:"background,omitempty"`
	Extension   string        `json:"extension,omitempty"`
	Dictionary  string        `json:"dictionary,omitempty"`
	Menu        []interface{} `json:"menu"`
	Action      interface{}   `json:"action,omitempty"`
	Data        interface{}   `json:"data,omitempty"`
	Ready       interface{}   `json:"ready,omitempty"`
	Options     interface{}   `json:"options,omitempty"`
}

// MenuItemObject https://msx.benzac.de/wiki/index.php?title=Menu_Item_Object
type MenuItemObject struct {
	Id             string      `json:"id,omitempty"`
	Type           string      `json:"type,omitempty"`
	Display        bool        `json:"display,omitempty"`
	Enable         bool        `json:"enable,omitempty"`
	Focus          bool        `json:"focus,omitempty"`
	Execute        bool        `json:"execute,omitempty"`
	Transparent    int         `json:"transparent,omitempty"`
	Icon           string      `json:"icon,omitempty"`
	Image          string      `json:"image,omitempty"`
	Label          string      `json:"label,omitempty"`
	Background     string      `json:"background,omitempty"`
	ExtensionIcon  string      `json:"extensionIcon,omitempty"`
	ExtensionLabel string      `json:"extensionLabel,omitempty"`
	LineColor      string      `json:"lineColor,omitempty"`
	Data           interface{} `json:"data,omitempty"`
	Options        interface{} `json:"options,omitempty"`
}
