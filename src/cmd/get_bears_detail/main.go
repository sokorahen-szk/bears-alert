package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sokorahen-szk/bears-alert/internal/bears_alert/repositories"
	"github.com/sokorahen-szk/bears-alert/internal/bears_alert/usecases"
	"github.com/sokorahen-szk/bears-alert/internal/fukui_bear_information"
)

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context) (*Response, error) {
	usecase := usecases.NewGetBearsDetailUsecase(
		repositories.NewBearsDetailRepository(),
		fukui_bear_information.NewFukuiBearInformationUsecase(),
	)

	err := usecase.Exec(ctx)
	if err != nil {
		return &Response{
			Message: "error",
		}, err
	}

	return &Response{
		Message: "hello world",
	}, nil
}

func main() {
	lambda.Start(handler)
}
