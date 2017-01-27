package reports

// UnversionedPackage is used to structure all report data for a pkg.
type UnversionedPackage struct {
	Pkg      string     `json:"package"`
	Status   string     `json:"status"`
	Handlers []*Handler `json:"handlers"`
	Log      string     `json:"log"`
}

// Handler is used to structure report data for a single handler.
type Handler struct {
	Handler string  `json:"handler"`
	Status  string  `json:"status"`
	Tests   []*Test `json:"tests"`
}

// Test is used to structure report data for a single test.
type Test struct {
	Name   string   `json:"test"`
	Status string   `json:"status"`
	Protos []*Proto `json:"protos"`
}

// Proto is used to structure report data for a Proto version.
type Proto struct {
	Name    string `json:"proto"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

const (
	// StatusPassed indicates a succesfull test
	StatusPassed = "passed"
	// StatusFailed indicates a failed test
	StatusFailed = "failed"
	// StatusSkipped indicates a skipped test
	StatusSkipped = "skipped"
)
