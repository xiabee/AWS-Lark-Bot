package lib

import (
	"main/comp"
	"strconv"
)

func ProcElement(element *Element, region string) {
	element.Tag = "div"
	element.Text.Tag = "lark_md"

	// Process the log data
	countBySeverity := make(map[string]int)
	// 用于存储告警标题和出现次数
	titleCount := make(map[string]map[string]int)

	var sendStr string
	countBySeverity, titleCount, err := comp.GetAlert(region)
	if err != nil {
		sendStr = err.Error()
		return
	}
	// 找出高危和严重告警中出现次数最多的告警
	var mostFrequentCritical, mostFrequentHigh string
	var maxCountCritical, maxCountHigh int
	for title, count := range titleCount["CRITICAL"] {
		if count > maxCountCritical {
			maxCountCritical = count
			mostFrequentCritical = title
		}
	}
	for title, count := range titleCount["HIGH"] {
		if count > maxCountHigh {
			maxCountHigh = count
			mostFrequentHigh = title
		}
	}

	// 打印 CRITICAL 和 HIGH 级别告警的数量统计
	sendStr += "[+] SecurityHub alarm total statistics:\n"
	for severity, count := range countBySeverity {
		sendStr += "  [-] Level ** " + severity + " **: " + strconv.Itoa(count) + " alarms\n"
	}

	sendStr += "[+] The most frequent alarms:\n"
	if mostFrequentCritical != "" {
		sendStr += "  [-] ** CRITICAL **: " + mostFrequentCritical + "** Occurs **" + strconv.Itoa(maxCountCritical) + " times.\n"
	}
	if mostFrequentHigh != "" {
		sendStr += "  [-] ** HIGH **: " + mostFrequentHigh + "** Occurs **" + strconv.Itoa(maxCountHigh) + " times.\n"
	}
	element.Text.Content = sendStr
}
