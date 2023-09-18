package alert

import (
	"AWS-Lark-Bot/resources"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

func ProcEC2(event resources.Event, element *resources.Element) error {
	var instanceDetails resources.InstanceDetails
	resourceType, ok := event.Detail.Resource["resourceType"].(string)
	if !ok {
		return errors.New("missing or invalid resourceType")
	}
	instanceDetailsJSON, ok := event.Detail.Resource["instanceDetails"].(map[string]interface{})
	if ok != true {
		return errors.New("missing or invalid instanceDetails")
	}

	marshaledDetailsJSON, err := json.Marshal(instanceDetailsJSON)

	err = json.Unmarshal(marshaledDetailsJSON, &instanceDetails)
	if err != nil {
		return err
	}

	alertSeverity := event.Detail.Severity
	var launchTime string

	switch tim := instanceDetails.LaunchTime.(type) {
	case string:
		launchTime = tim
	case float64:
		timestampMillis := int64(instanceDetails.LaunchTime.(float64))
		timestampSeconds := timestampMillis / 1000
		t := time.Unix(timestampSeconds, 0)
		launchTime = t.Format("2006-01-02 15:04:05 MST")
	default:
		launchTime = ""
	}
	element.Text.Content = "[+] ** Type**:   " + event.DetailType + "\n" +
		"[+] ** Severity**:    " + Level(alertSeverity) + " " + strconv.FormatFloat(alertSeverity, 'f', -1, 64) + "\n" +
		"[+] ** Alert Time**:    " + event.Time + "\n" +
		"[+] ** Account**:    " + event.Account + "\n" +
		"[+] ** Region**:    " + event.Region + "\n" +
		"[+] ** Resource Type**:    " + resourceType + "\n" +
		"[+] ** Instance ID**:    " + instanceDetails.InstanceID + "\n" +
		"[+] ** Launch Time**:    " + launchTime + "\n" +
		"[+] ** Action Type**:    " + event.Detail.Service.Action.ActionType + "\n" +
		"[+] ** Description**:    " + event.Detail.Description + "\n"
	return nil
}
