package lib

func ProcCard(logs Logs, data *CardMessage) {
	data.MsgType = "interactive"
	data.Card.Config.WideScreenMode = true
	data.Card.Header.Title.Tag = "markdown"
	data.Card.Header.Title.Content = "STATIC ACCESS KEY " + logs.LogEvents[0].Message.UserIdentity.AccessKeyId + " DETECTED IN USE"
	data.Card.Header.Template = "blue"

	var element Element
	ProcElement(logs, &element)
	data.Card.Elements = append(data.Card.Elements, element)
}
