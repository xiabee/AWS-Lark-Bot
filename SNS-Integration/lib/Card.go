package lib

import "AWS-Lark-Bot/resources"

func ProcCard(event resources.Event, data *resources.CardMessage) float64 {
	data.MsgType = "interactive"
	data.Card.Config.WideScreenMode = true
	data.Card.Header.Title.Tag = "markdown"
	data.Card.Header.Title.Content = event.Detail.Title
	data.Card.Header.Template = "blue"

	var element resources.Element
	serverity := ProcElement(event, &element)
	data.Card.Elements = append(data.Card.Elements, element)

	return serverity
}
