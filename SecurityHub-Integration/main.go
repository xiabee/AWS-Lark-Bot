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
	}

	// 用于统计各个级别告警的数量
	countBySeverity := make(map[string]int)

	// 循环处理分页
	for {
		// 查询告警
		result, err := client.GetFindings(context.TODO(), input)
		if err != nil {
			fmt.Println("查询告警时发生错误:", err)
			return
		}

		// 处理并统计结果
		for _, finding := range result.Findings {
			fmt.Println("告警标题:", aws.ToString(finding.Title))
			fmt.Println("告警描述:", aws.ToString(finding.Description))
			fmt.Println("最后更新时间:", *finding.UpdatedAt)
			fmt.Println("告警级别:", finding.Severity.Label)
			fmt.Println("------------------------------------------------")

			// 更新告警级别计数
			severityLabel := string(finding.Severity.Label)
			countBySeverity[severityLabel]++
		}

		// 检查是否还有更多的页面
		if result.NextToken == nil {
			break
		}

		// 更新输入，以获取下一个页面
		input.NextToken = result.NextToken
	}

	// 打印告警级别统计
	fmt.Println("告警级别统计:")
	for severity, count := range countBySeverity {
		fmt.Printf("级别 %s: %d 条告警\n", severity, count)
	}
}
