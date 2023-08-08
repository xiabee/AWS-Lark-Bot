package lib

import (
	"strconv"
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
func ProcElement(event Event, element *Element) float64 {
	element.Tag = "div"
	element.Text.Tag = "lark_md"

	var kv []KV = nil
	for _, tags := range event.Detail.Resource.InstanceDetails.Tags {
		if tags.Key == "servicetype" || tags.Key == "eks:cluster-name" || tags.Key == "component" || tags.Key == "tenant" {
			kv = append(kv, tags)
		}
	}
	eksInfo := ""
	for _, tags := range kv {
		eksInfo = eksInfo + "  [-] **" + tags.Key + "**:    " + tags.Value + "\n"
	}
	// Load EKS information

	alertSeverity := event.Detail.Severity

	element.Text.Content = "[+] ** Type**:   " + event.DetailType + "\n" +
		"[+] ** Severity**:    " + alertLevel(alertSeverity) + " " + strconv.FormatFloat(alertSeverity, 'f', -1, 64) + "\n" +
		"[+] ** Time**:    " + event.Time + "\n" +
		"[+] ** Account**:    " + event.Account + "\n" +
		"[+] ** Region**:    " + event.Region + "\n" +
		"[+] ** Resource Type**:    " + event.Detail.Resource.ResourceType + "\n" +
		"[+] ** ID**:    " + event.Detail.Resource.InstanceDetails.InstanceID + "\n" +
		"[+] ** Launch Time**:    " + event.Detail.Resource.InstanceDetails.LaunchTime + "\n" +
		"[+] ** EKS info**:\n" + eksInfo +
		"[+] ** Action Type**:    " + event.Detail.Service.Action.ActionType + "\n" +
		"[+] ** Description**:    " + event.Detail.Description + "\n" +
		"[+] ** Local Port**:    " + strconv.Itoa(event.Detail.Service.Action.PortProbeAction.PortProbeDetails[0].LocalPortDetails.Port) + "\n" +
		"[+] ** Remote IP**:    " + event.Detail.Service.Action.PortProbeAction.PortProbeDetails[0].RemoteIpDetails.IpAddressV4 + "\n" +
		"[+] ** AsnOrg**:    " + event.Detail.Service.Action.PortProbeAction.PortProbeDetails[0].RemoteIpDetails.Organization.AsnOrg + "\n" +
		"[+] ** Country**:    " + event.Detail.Service.Action.PortProbeAction.PortProbeDetails[0].RemoteIpDetails.Country.CountryName + "\n" +
		"[+] ** Blocked**:    " + blocked(event.Detail.Service.Action.PortProbeAction.Blocked) + "\n"

	return alertSeverity
}
