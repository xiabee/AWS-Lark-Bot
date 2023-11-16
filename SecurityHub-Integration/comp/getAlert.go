package comp

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"time"
)

func GetAlert(region string) (map[string]int, map[string]map[string]int, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Println("Configuration loading error:", err)
		return nil, nil, err
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
		MaxResults: aws.Int32(100),
	}
	countBySeverity := make(map[string]int)
	//用于存储告警标题和出现次数
	titleCount := make(map[string]map[string]int)
	// 循环处理分页
	for {
		// 查询告警
		result, err := client.GetFindings(context.TODO(), input)
		if err != nil {
			fmt.Println("An error occurred while querying alarms:", err)
			return nil, nil, err
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

	fmt.Println("Get alert in region: ", region)
	return countBySeverity, titleCount, nil
}
