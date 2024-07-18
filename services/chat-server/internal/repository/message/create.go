package message

import (
	"context"
	"fmt"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/message/converter"
)

func (r *repository) CreateMessage(ctx context.Context, messageConverted model.MessageCreate) error {
	message, err := converter.MessageCreateModelToData(messageConverted)
	if err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrInvalidID, err)
	}

	_, err = r.collectionMessages.InsertOne(ctx, message)
	if err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrDatabase, err)
	}

	return nil
}
