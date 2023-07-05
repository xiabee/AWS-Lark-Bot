package lib

import (
	"fmt"
	"time"
)

func ProcElement(log Logs, element *Element) {
	element.Tag = "div"
	element.Text.Tag = "lark_md"

	// Process the log data
	userLogs := make(map[string][]LogEvent)

	for _, logEvent := range log.LogEvents {
		username := logEvent.Message.UserIdentity.UserName
		userLogs[username] = append(userLogs[username], logEvent)
	}

	var sendStr string
	for username, logEvents := range userLogs {
		sendStr += fmt.Sprintf("[+] ** User Type** : %s\n", logEvents[0].Message.UserIdentity.Type)
		sendStr += fmt.Sprintf("[+] ** Account ID** : %s\n", logEvents[0].Message.UserIdentity.AccountId)
		sendStr += fmt.Sprintf("[+] ** AccessKey ID** : %s\n", logEvents[0].Message.UserIdentity.AccessKeyId)
		sendStr += fmt.Sprintf("[+] ** Event Region** : %s\n", logEvents[0].Message.AwsRegion)
		sendStr += fmt.Sprintf("[+] ** User Name** : %s\n", username)
		sendStr += fmt.Sprintf("[+] ** Event Time** : %s\n", ProcTime(logEvents[0].Timestamp))
		sendStr += fmt.Sprintf("[+] ** Source IP** : %s\n", logEvents[0].Message.SourceIPAddress)

		var eventNames []string
		for _, eveNow := range logEvents {
			if !Contains(eventNames, eveNow.Message.EventName) {
				eventNames = append(eventNames, eveNow.Message.EventName)
			}
		}

		sendStr += fmt.Sprintf("[+] **EventName** :\n")
		for _, eve := range eventNames {
			sendStr += fmt.Sprintf("    [-] %s\n", eve)
		}

	}
	element.Text.Content = sendStr
}

// ProcTime :Convert timestamp to time
func ProcTime(timeNum int64) string {
	timestamp := timeNum

	// Convert to seconds and nanoseconds
	sec := timestamp / 1000
	nsec := (timestamp % 1000) * int64(time.Millisecond)

	// Create a time.Time object
	t := time.Unix(sec, nsec)

	// Format and print the time
	return t.Format(time.RFC3339)
}

// Contains :Determine whether a string is in a string array
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
