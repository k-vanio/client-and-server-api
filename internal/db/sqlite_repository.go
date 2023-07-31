package db

import (
	"context"
	"time"

	"github.com/x-vanio/client-and-server-api/internal/model"
	"github.com/x-vanio/client-and-server-api/pkg/dto"
	"gorm.io/gorm"
)

type SqliteRepository struct {
	DB *gorm.DB
}

func (sr *SqliteRepository) Save(quote dto.Quote) error {

	/*
		É importante ressaltar que o banco de dados está localmente integrado à aplicação,
		o que garante um tempo mínimo para a gravação e persistência dos dados.
		Para obter um erro, é necessário definir o tempo limite do contexto para um valor adequado,
		como: context.WithTimeout(context.Background(), time.Nanosecond*3)
	*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*20)
	defer cancel()

	insert := &model.Quote{
		Code:       quote.Currency.Code,
		Codein:     quote.Currency.Codein,
		Name:       quote.Currency.Name,
		High:       quote.Currency.High,
		Low:        quote.Currency.Low,
		VarBid:     quote.Currency.VarBid,
		PctChange:  quote.Currency.PctChange,
		Bid:        quote.Currency.Bid,
		Ask:        quote.Currency.Ask,
		Timestamp:  quote.Currency.Timestamp,
		CreateDate: quote.Currency.CreateDate,
	}

	if err := sr.DB.WithContext(ctx).Create(insert).Error; err != nil {
		return err
	}

	return nil
}
