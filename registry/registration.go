package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredService  []ServiceName // 添加服务依赖项
	ServiceUpdateURL string        // 服务注册和依赖服务沟通的URL
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
