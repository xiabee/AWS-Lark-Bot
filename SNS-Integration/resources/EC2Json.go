package resources

type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type InstanceDetails struct {
	InstanceID   string      `json:"instanceId"`
	InstanceType string      `json:"instanceType"`
	LaunchTime   interface{} `json:"launchTime"`
}
