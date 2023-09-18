package lib

import "AWS-Lark-Bot/resources"

func ProcCard(event resources.Event, data *resources.CardMessage, serverity float64) {
	data.MsgType = "interactive"
	data.Card.Config.WideScreenMode = true
	data.Card.Header.Title.Tag = "markdown"
	data.Card.Header.Title.Content = event.Detail.Type
	data.Card.Header.Template = "blue"

	if serverity >= 7.0 {
		data.Card.Header.Template = "red"
	} else if serverity >= 4.0 {
		data.Card.Header.Template = "yellow"
	} else {
		data.Card.Header.Template = "blue"
	}
	// make different color for different severity

	var element resources.Element
	ProcElement(event, &element)
	data.Card.Elements = append(data.Card.Elements, element)
}
