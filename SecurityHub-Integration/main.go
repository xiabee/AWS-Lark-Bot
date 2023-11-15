package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
)

func main() {
	startTime := time.Now() // 记录程序开始时间

	// 加载配置
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		fmt.Println("配置加载错误:", err)
		return
	}

	// 创建 Security Hub 客户端
	client := securityhub.NewFromConfig(cfg)

	// 设置时间范围为过去一周
	oneWeekAgo := aws.String(time.Now().AddDate(0, 0, -7).Format(time.RFC3339))
	now := aws.String(time.Now().Format(time.RFC3339))

	// 创建查询输入
	input := &securityhub.GetFindingsInput{
		Filters: &types.AwsSecurityFindingFilters{
			RecordState: []types.StringFilter{
				{
					Comparison: types.StringFilterComparisonEquals,
					Value:      aws.String("ACTIVE"),
				},
			},
			UpdatedAt: []types.DateFilter{
				{
					Start: oneWeekAgo,
					End:   now,
				},
			},
		},
		MaxResults: 100,
	}

	// 用于统计 CRITICAL 和 HIGH 级别告警的数量
	countBySeverity := make(map[string]int)
	// 用于存储告警标题和出现次数
	titleCount := make(map[string]map[string]int)

	// 循环处理分页
	for {
		// 查询告警
		result, err := client.GetFindings(context.TODO(), input)
		if err != nil {
			fmt.Println("查询告警时发生错误:", err)
			return
		}

		// 统计 CRITICAL 和 HIGH 级别的结果
		for _, finding := range result.Findings {
			severityLabel := string(finding.Severity.Label)
			if severityLabel == "CRITICAL" || severityLabel == "HIGH" {
				countBySeverity[severityLabel]++
				title := aws.ToString(finding.Title)
				if titleCount[severityLabel] == nil {
					titleCount[severityLabel] = make(map[string]int)
				}
				titleCount[severityLabel][title]++
			}
		}

		// 检查是否还有更多的页面
		if result.NextToken == nil {
			break
		}

		// 更新输入，以获取下一个页面
		input.NextToken = result.NextToken
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
	fmt.Println("过去一周内 CRITICAL 和 HIGH 级别告警的数量统计:")
	for severity, count := range countBySeverity {
		fmt.Printf("级别 %s: %d 条告警\n", severity, count)
	}

	fmt.Printf("出现次数最多的 CRITICAL 级别告警: %s, 出现次数: %d\n", mostFrequentCritical, maxCountCritical)
	fmt.Printf("出现次数最多的 HIGH 级别告警: %s, 出现次数: %d\n", mostFrequentHigh, maxCountHigh)

	// 记录并打印程序运行时间
	endTime := time.Now() // 记录程序结束时间
	duration := endTime.Sub(startTime)
	fmt.Printf("程序运行时间: %s\n", duration)
}
