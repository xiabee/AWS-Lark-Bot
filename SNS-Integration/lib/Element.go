package lib

import (
	"AWS-Lark-Bot/resources"
	"fmt"
)

func blocked(flag bool) string {
	if flag {
		return "TRUE"
	} else {
		return "FALSE"
	}
}

func alertLevel(alertSeverity float64) string {
	if alertSeverity >= 7.0 {
		return "High"
	} else if alertSeverity >= 4.0 {
		return "Medium"
	} else {
		return "Low"
	}
}

// ProcElement get the servertiy of the alert
func ProcElement(event resources.Event, element *resources.Element) float64 {
	element.Tag = "div"
	element.Text.Tag = "lark_md"
	resourceType, ok := event.Detail.Resource["resourceType"].(string)
	if !ok {
		fmt.Println("Missing or invalid resourceType")
		return 0
	}

	fmt.Println("ResourceType: ", resourceType)
	switch resourceType {
	case "Instance":
		err := ProcEC2(event, element)
		if err != nil {
			fmt.Println("Failed to process EC2 alert: ", err)
			return 0
		}
	default:
		return 0
	}
	return 0
}
