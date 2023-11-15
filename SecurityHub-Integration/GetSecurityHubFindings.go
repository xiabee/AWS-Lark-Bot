package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
)

// 查询 Security Hub 结果的通用函数
func querySecurityHubResults(region string, filters *types.AwsSecurityFindingFilters) ([]types.AwsSecurityFinding, error) {
	// 创建 AWS 客户端
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	client := securityhub.NewFromConfig(cfg)

	// 创建筛选条件，根据需要进行填充
	//filters := &types.AwsSecurityFindingFilters{}

	// 创建 GetFindingsInput 结构体
	input := &securityhub.GetFindingsInput{
		Filters:    filters, // 设置筛选条件
		MaxResults: 50,      // 设置每页最大结果数
		NextToken:  nil,     // 设置下一页查询的 NextToken
	}

	// 用于存储所有结果的切片
	var allFindings []types.AwsSecurityFinding

	for {
		// 查询 Findings
		resp, err := client.GetFindings(context.TODO(), input)
		if err != nil {
			return nil, err
		}

		// 将本页结果追加到 allFindings
		allFindings = append(allFindings, resp.Findings...)

		// 如果有下一页结果，继续查询
		if resp.NextToken == nil {
			break
		}

		// 设置下一页查询的 NextToken
		input.NextToken = resp.NextToken
	}

	return allFindings, nil
}
