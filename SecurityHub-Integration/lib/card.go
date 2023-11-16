package lib

func ProcCard(data *CardMessage, region string) {
	data.MsgType = "interactive"
	data.Card.Config.WideScreenMode = true
	data.Card.Header.Title.Tag = "markdown"
	title := "SecurityHub Alert of the Week in " + region
	data.Card.Header.Title.Content = title
	data.Card.Header.Template = "blue"

	var element Element
	ProcElement(&element, region)
	data.Card.Elements = append(data.Card.Elements, element)
}
