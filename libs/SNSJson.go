package libs

type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type InstanceDetails struct {
	InstanceID   string `json:"instanceId"`
	InstanceType string `json:"instanceType"`
	LaunchTime   string `json:"launchTime"`
	Tags         []KV   `json:"tags"`
}

type Resource struct {
	ResourceType    string          `json:"resourceType"`
	InstanceDetails InstanceDetails `json:"instanceDetails"`
}

type Service struct {
	ServiceName string `json:"serviceName"`
	DetectorID  string `json:"detectorId"`
	Action      struct {
		ActionType      string `json:"actionType"`
		PortProbeAction struct {
			PortProbeDetails []struct {
				LocalPortDetails struct {
					Port     int    `json:"port"`
					PortName string `json:"portName"`
				} `json:"localPortDetails"`
				RemoteIpDetails struct {
					IpAddressV4  string `json:"ipAddressV4"`
					Organization struct {
						AsnOrg string `json:"asnOrg"`
						Org    string `json:"org"`
					} `json:"organization"`
					Country struct {
						CountryName string `json:"countryName"`
					} `json:"country"`
				} `json:"remoteIpDetails"`
			} `json:"portProbeDetails"`
			Blocked bool `json:"blocked"`
		} `json:"portProbeAction"`
	} `json:"action"`
	ResourceRole string `json:"resourceRole"`
}

type Event struct {
	Version    string   `json:"version"`
	ID         string   `json:"id"`
	DetailType string   `json:"detail-type"`
	Source     string   `json:"source"`
	Account    string   `json:"account"`
	Time       string   `json:"time"`
	Region     string   `json:"region"`
	Resources  []string `json:"resources"`
	Detail     struct {
		SchemaVersion string   `json:"schemaVersion"`
		AccountID     string   `json:"accountId"`
		Region        string   `json:"region"`
		Partition     string   `json:"partition"`
		ID            string   `json:"id"`
		ARN           string   `json:"arn"`
		Type          string   `json:"type"`
		Resource      Resource `json:"resource"`
		Service       Service  `json:"service"`
		Severity      int      `json:"severity"`
		CreatedAt     string   `json:"createdAt"`
		UpdatedAt     string   `json:"updatedAt"`
		Title         string   `json:"title"`
		Description   string   `json:"description"`
	} `json:"detail"`
}
