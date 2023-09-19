package alert

import "AWS-Lark-Bot/resources"

func Level(alertSeverity float64) string {
	if alertSeverity >= 7.0 {
		return "High"
	} else if alertSeverity >= 4.0 {
		return "Medium"
	} else {
		return "Low"
	}
}

func GetAlertServerity(event resources.Event) float64 {
	return event.Detail.Severity
}
