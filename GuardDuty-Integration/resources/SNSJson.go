package resources

type Service struct {
	ServiceName string `json:"serviceName"`
	DetectorID  string `json:"detectorId"`
	Action      struct {
		ActionType string `json:"actionType"`
	}
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
		SchemaVersion string                 `json:"schemaVersion"`
		AccountID     string                 `json:"accountId"`
		Region        string                 `json:"region"`
		Partition     string                 `json:"partition"`
		ID            string                 `json:"id"`
		ARN           string                 `json:"arn"`
		Type          string                 `json:"type"`
		Resource      map[string]interface{} `json:"resource"`
		Service       Service                `json:"service"`
		Severity      float64                `json:"severity"`
		CreatedAt     string                 `json:"createdAt"`
		UpdatedAt     string                 `json:"updatedAt"`
		Title         string                 `json:"title"`
		Description   string                 `json:"description"`
	} `json:"detail"`
}
