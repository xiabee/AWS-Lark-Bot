package lib

import "encoding/json"

type UserIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principalId"`
	Arn         string `json:"arn"`
	AccountId   string `json:"accountId"`
	AccessKeyId string `json:"accessKeyId"`
	UserName    string `json:"userName"`
}

type Message struct {
	EventVersion    string       `json:"eventVersion"`
	UserIdentity    UserIdentity `json:"userIdentity"`
	EventTime       string       `json:"eventTime"`
	EventSource     string       `json:"eventSource"`
	EventName       string       `json:"eventName"`
	AwsRegion       string       `json:"awsRegion"`
	SourceIPAddress string       `json:"sourceIPAddress"`
	UserAgent       string       `json:"userAgent"`
	RequestID       string       `json:"requestID"`
	EventID         string       `json:"eventID"`
}

type LogEvent struct {
	ID        string  `json:"id"`
	Timestamp int64   `json:"timestamp"`
	Message   Message `json:"message"`
}

type Logs struct {
	MessageType         string     `json:"messageType"`
	Owner               string     `json:"owner"`
	LogGroup            string     `json:"logGroup"`
	LogStream           string     `json:"logStream"`
	SubscriptionFilters []string   `json:"subscriptionFilters"`
	LogEvents           []LogEvent `json:"logEvents"`
}

// UnmarshalJSON :Define a function to unmarshal the Message field of LogEvent
func (e *LogEvent) UnmarshalJSON(data []byte) error {
	type Alias LogEvent
	aux := &struct {
		Message string `json:"message"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var msg Message
	if err := json.Unmarshal([]byte(aux.Message), &msg); err != nil {
		return err
	}
	e.Message = msg
	return nil
}
