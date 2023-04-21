package v1

type EnvironmentList map[string]Environment

type Environment struct {
	Name         string `json:"name"`
	ClusterAPI   string `json:"clusterapi"`
	AppendSuffix bool   `json:"appendsuffix"`
}
