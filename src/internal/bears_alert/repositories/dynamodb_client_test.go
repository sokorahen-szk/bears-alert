package repositories_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sokorahen-szk/bears-alert/internal/bears_alert/repositories"
	"github.com/stretchr/testify/assert"
)

func TestDynamodbClient_Insert(t *testing.T) {
	t.Skip("skip")
	ctx := context.Background()
	cfg, err := repositories.NewDynamodbConfig(ctx)
	if err != nil {
		t.Fatal(err)
	}

	tableName := "sample-table"

	tests := []struct {
		name    string
		items   []map[string]types.AttributeValue
		wantErr bool
	}{
		{
			name: "1件だけ登録",
			items: []map[string]types.AttributeValue{
				{
					"ID":   &types.AttributeValueMemberS{Value: "item1"},
					"Name": &types.AttributeValueMemberS{Value: "テスト1"},
				},
			},
			wantErr: false,
		},
		{
			name: "複数登録",
			items: []map[string]types.AttributeValue{
				{
					"ID":   &types.AttributeValueMemberS{Value: "item2"},
					"Name": &types.AttributeValueMemberS{Value: "テスト2"},
				},
				{
					"ID":   &types.AttributeValueMemberS{Value: "item3"},
					"Name": &types.AttributeValueMemberS{Value: "テスト3"},
				},
			},
			wantErr: false,
		},
		{
			name: "登録済みIDはエラーになる",
			items: []map[string]types.AttributeValue{
				{
					"ID":   &types.AttributeValueMemberS{Value: "item1"},
					"Name": &types.AttributeValueMemberS{Value: "テスト1"},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := repositories.NewDynamodbClient(cfg)

			err := client.Insert(ctx, tableName, tt.items)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
