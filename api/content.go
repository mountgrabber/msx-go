package api

// ContentRootObject https://msx.benzac.de/wiki/index.php?title=Menu_Root_Object
type ContentRootObject struct {
	Name            string        `json:"name,omitempty"`
	Version         string        `json:"version,omitempty"`
	Reference       string        `json:"reference,omitempty"`
	Flag            string        `json:"flag,omitempty"`
	Reuse           bool          `json:"reuse,omitempty"`
	Cache           bool          `json:"cache,omitempty"`
	Restore         bool          `json:"restore,omitempty"`
	Important       bool          `json:"important,omitempty"`
	Wrap            bool          `json:"wrap,omitempty"`
	Compress        bool          `json:"compress,omitempty"`
	Preselect       bool          `json:"preselect,omitempty"`
	Refocus         int           `json:"refocus,omitempty"`
	Transparent     int           `json:"transparent,omitempty"`
	Type            string        `json:"type,omitempty"`
	Preload         string        `json:"preload,omitempty"`
	Headline        string        `json:"headline,omitempty"`
	Background      string        `json:"background,omitempty"`
	Extension       string        `json:"extension,omitempty"`
	Dictionary      string        `json:"dictionary,omitempty"`
	Template        interface{}   `json:"template,omitempty"` // Only for templated items
	Items           []interface{} `json:"items,omitempty"`    // Only for templated items
	Pages           []interface{} `json:"pages,omitempty"`    // Only for non-templated items
	Header          interface{}   `json:"header,omitempty"`
	Footer          interface{}   `json:"footer,omitempty"`
	Inserts         []interface{} `json:"inserts,omitempty"`
	Overlay         interface{}   `json:"overlay,omitempty"`
	Underlay        interface{}   `json:"underlay,omitempty"`
	Action          interface{}   `json:"action,omitempty"`
	Data            interface{}   `json:"data,omitempty"`
	Ready           interface{}   `json:"ready,omitempty"`
	Options         interface{}   `json:"options,omitempty"`
	Caption         string        `json:"caption,omitempty"`
	CaptionUnderlay int           `json:"captionUnderlay,omitempty"`
}

// ContentPageObject https://msx.benzac.de/wiki/index.php?title=Content_Page_Object
type ContentPageObject struct {
	Display         bool          `json:"display,omitempty"`
	Important       bool          `json:"important,omitempty"`
	Wrap            bool          `json:"wrap,omitempty"`
	Compress        bool          `json:"compress,omitempty"`
	Transparent     int           `json:"transparent,omitempty"`
	Headline        string        `json:"headline,omitempty"`
	Background      string        `json:"background,omitempty"`
	Area            string        `json:"area,omitempty"`
	Offset          string        `json:"offset,omitempty"`
	Position        string        `json:"position,omitempty"`
	Template        interface{}   `json:"template,omitempty"`
	Items           []interface{} `json:"items"`
	Action          interface{}   `json:"action,omitempty"`
	Data            interface{}   `json:"data,omitempty"`
	Options         interface{}   `json:"options,omitempty"`
	Caption         string        `json:"caption,omitempty"`
	CaptionUnderlay int           `json:"captionUnderlay,omitempty"`
}

// ContentItemObject https://msx.benzac.de/wiki/index.php?title=Content_Item_Object
type ContentItemObject struct {
	Id                string      `json:"id,omitempty"`
	Type              string      `json:"type,omitempty"`
	Key               string      `json:"key,omitempty"`
	Layout            string      `json:"layout,omitempty"` // Only for the template object and non-templated items
	Area              string      `json:"area,omitempty"`
	Offset            string      `json:"offset,omitempty"`
	Display           bool        `json:"display,omitempty"`
	Enable            bool        `json:"enable,omitempty"`
	Focus             bool        `json:"focus,omitempty"`
	Execute           bool        `json:"execute,omitempty"`
	Enumerate         bool        `json:"enumerate,omitempty"`
	Compress          bool        `json:"compress,omitempty"`
	Decompress        bool        `json:"decompress,omitempty"`
	Shortcut          bool        `json:"shortcut,omitempty"`
	Round             bool        `json:"round,omitempty"`
	Break             string      `json:"break,omitempty"`
	Group             string      `json:"group,omitempty"`
	Color             string      `json:"color,omitempty"`
	Title             string      `json:"title,omitempty"`
	TitleHeader       string      `json:"titleHeader,omitempty"`
	TitleFooter       string      `json:"titleFooter,omitempty"`
	Label             string      `json:"label,omitempty"`
	Icon              string      `json:"icon,omitempty"`
	IconSize          string      `json:"iconSize,omitempty"`
	Headline          string      `json:"headline,omitempty"`
	Text              string      `json:"text,omitempty"`
	Alignment         string      `json:"alignment,omitempty"`
	Truncation        string      `json:"truncation,omitempty"`
	Centration        string      `json:"centration,omitempty"`
	Separation        int         `json:"separation,omitempty"`
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
	ImageFiller       string      `json:"imageFiller,omitempty"`
	ImageWidth        int         `json:"imageWidth,omitempty"`
	ImageHeight       int         `json:"imageHeight,omitempty"`
	ImageOverlay      int         `json:"imageOverlay,omitempty"`
	ImagePreload      bool        `json:"imagePreload,omitempty"`
	ImageLabel        string      `json:"imageLabel,omitempty"`
	ImageColor        string      `json:"imageColor,omitempty"`
	ImageScreenFiller string      `json:"imageScreenFiller,omitempty"`
	ImageBoundary     bool        `json:"imageBoundary,omitempty"`
	PlayerLabel       string      `json:"playerLabel,omitempty"`
	Background        string      `json:"background,omitempty"`
	ExtensionIcon     string      `json:"extensionIcon,omitempty"`
	ExtensionLabel    string      `json:"extensionLabel,omitempty"`
	Action            string      `json:"action,omitempty"`
	Data              interface{} `json:"data,omitempty"`
	Properties        interface{} `json:"properties,omitempty"`
	Live              interface{} `json:"live,omitempty"`
	Selection         interface{} `json:"selection,omitempty"`
	Options           interface{} `json:"options,omitempty"`
}
