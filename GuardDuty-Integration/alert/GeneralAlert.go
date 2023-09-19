package alert

import (
	"AWS-Lark-Bot/resources"
	"errors"
	"strconv"
)

func ProcGerneral(event resources.Event, element *resources.Element) error {
	resourceType, ok := event.Detail.Resource["resourceType"].(string)
	if !ok {
		return errors.New("missing or invalid resourceType")
	}

	alertSeverity := event.Detail.Severity

	element.Text.Content = "[+] ** Type**:   " + event.DetailType + "\n" +
		"[+] ** Severity**:    " + Level(alertSeverity) + " " + strconv.FormatFloat(alertSeverity, 'f', -1, 64) + "\n" +
		"[+] ** Alert Time**:    " + event.Time + "\n" +
		"[+] ** Account**:    " + event.Account + "\n" +
		"[+] ** Region**:    " + event.Region + "\n" +
		"[+] ** Resource Type**:    " + resourceType + "\n" +
		"[+] ** Description**:    " + event.Detail.Description + "\n"
	return nil
}
