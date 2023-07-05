package lib

func ProcCard(logs Logs, data *CardMessage) {
	data.MsgType = "interactive"
	data.Card.Config.WideScreenMode = true
	data.Card.Header.Title.Tag = "markdown"
	title := logs.LogEvents[0].Message.UserIdentity.AccessKeyId + " DETECTED IN USE"
	data.Card.Header.Title.Content = title
	data.Card.Header.Template = "red"

	var element Element
	ProcElement(logs, &element)
	data.Card.Elements = append(data.Card.Elements, element)
}
