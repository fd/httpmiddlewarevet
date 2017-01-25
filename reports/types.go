package reports

type UnversionedPackage struct {
	Pkg      string     `json:"package"`
	Status   string     `json:"status"`
	Handlers []*Handler `json:"handlers"`
	Log      string     `json:"log"`
}

type Handler struct {
	Handler string  `json:"handler"`
	Status  string  `json:"status"`
	Tests   []*Test `json:"tests"`
}

type Test struct {
	Name   string   `json:"test"`
	Status string   `json:"status"`
	Protos []*Proto `json:"protos"`
}

type Proto struct {
	Name    string `json:"proto"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
