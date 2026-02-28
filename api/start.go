package api

// StartObject https://msx.benzac.de/wiki/index.php?title=Start_Object
type StartObject struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Parameter string `json:"parameter"`
	Welcome   string `json:"welcome,omitempty"`
}
