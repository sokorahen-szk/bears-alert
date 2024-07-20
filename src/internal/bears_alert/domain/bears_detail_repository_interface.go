package domain

import "context"

type IBearsDetailRepository interface {
	Insert(context.Context, *BearsDetail) error
	List(context.Context) ([]*BearsDetail, error)
}
