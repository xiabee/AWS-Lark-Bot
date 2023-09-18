package lib

import (
	"AWS-Lark-Bot/alert"
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

// ProcElement get the servertiy of the alert
func ProcElement(event resources.Event, element *resources.Element) {
	element.Tag = "div"
	element.Text.Tag = "lark_md"
	resourceType, ok := event.Detail.Resource["resourceType"].(string)
	if !ok {
		fmt.Println("Missing or invalid resourceType")
		return
	}

	fmt.Println("ResourceType: ", resourceType)
	switch resourceType {
	case "Instance":
		err := alert.ProcEC2(event, element)
		if err != nil {
			fmt.Println("Failed to process EC2 alert: ", err)
			return
		}
	default:
		err := alert.ProcGerneral(event, element)
		if err != nil {
			fmt.Println("Failed to process EC2 alert: ", err)
			return
		}
	}
	return
}
