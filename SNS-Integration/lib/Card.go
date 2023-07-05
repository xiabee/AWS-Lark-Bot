package lib

func ProcCard(event Event, data *CardMessage) {
	data.MsgType = "interactive"
	data.Card.Config.WideScreenMode = true
	data.Card.Header.Title.Tag = "markdown"
	data.Card.Header.Title.Content = event.Detail.Title
	data.Card.Header.Template = "blue"

	var element Element
	ProcElement(event, &element)
	data.Card.Elements = append(data.Card.Elements, element)
}
