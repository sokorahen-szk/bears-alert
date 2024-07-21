package repositories

import (
	"context"
	"os"
	"slices"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type IDynamodbClient interface {
}

type DynamodbClient struct {
	client *dynamodb.Client
}

var _ IDynamodbClient = (*DynamodbClient)(nil)

func NewDynamodbConfig(ctx context.Context) (*aws.Config, error) {
	if slices.Contains([]string{"local", "testing"}, os.Getenv("APP_ENV")) {
		if cfg, err := config.LoadDefaultConfig(
			ctx,
			config.WithRegion(os.Getenv("AWS_REGION")),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				os.Getenv("AWS_ACCESS_KEY_ID"),
				os.Getenv("AWS_SECRET_ACCESS_KEY"),
				os.Getenv("AWS_SESSION_KEY")),
			),
		); err == nil {
			return &cfg, nil
		}
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(os.Getenv("AWS_REGION")))
	return &cfg, err
}

func NewDynamodbClient(cfg *aws.Config) *DynamodbClient {
	return &DynamodbClient{
		client: dynamodb.NewFromConfig(*cfg),
	}
}

func (dc DynamodbClient) Insert(ctx context.Context, tableName string, items []map[string]types.AttributeValue) error {
	transactionItems := make([]types.TransactWriteItem, len(items))
	for i, item := range items {
		transactionItems[i] = types.TransactWriteItem{
			Put: &types.Put{
				TableName:           aws.String(tableName),
				Item:                item,
				ConditionExpression: aws.String("attribute_not_exists(ID)"),
			},
		}
	}

	_, err := dc.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: transactionItems,
	})
	return err
}
