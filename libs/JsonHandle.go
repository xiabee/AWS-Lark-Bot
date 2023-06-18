package libs

import (
	"encoding/json"
)

func ProcessJSON(jsonStr string) (Event, error) {

	var event Event
	err := json.Unmarshal([]byte(jsonStr), &event)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}
