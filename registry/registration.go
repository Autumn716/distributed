package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName
	ServiceUpdateURL string
}

type ServiceName string

// Existing services are declared as constants
const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)

// 每一条更新
type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
