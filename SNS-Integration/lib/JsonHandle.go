package lib

import (
	"AWS-Lark-Bot/resources"
	"encoding/json"
)

func ProcessJSON(jsonStr string) (resources.Event, error) {

	var event resources.Event
	err := json.Unmarshal([]byte(jsonStr), &event)
	if err != nil {
		return resources.Event{}, err
	}
	return event, nil
}
