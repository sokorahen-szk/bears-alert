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

type Handler struct {
	usecase *usecases.GetBearsDetailUsecase
}

func NewHandler(usecase *usecases.GetBearsDetailUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h Handler) Handle(ctx context.Context) (*Response, error) {
	err := h.usecase.Exec(ctx)
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
	handler := NewHandler(
		usecases.NewGetBearsDetailUsecase(
			repositories.NewBearsDetailRepository(),
			fukui_bear_information.NewFukuiBearInformationRepository(),
		),
	)
	lambda.Start(handler.Handle)
}
