package registry

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

type ServiceName string

// Existing services are declared as constants
const (
	LogService = ServiceName("LogService")
)
