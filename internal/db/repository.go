package db

import (
	"github.com/x-vanio/client-and-server-api/pkg/dto"
)

type Repository interface {
	Save(quote dto.Quote) error
}
