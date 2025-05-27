// internal/usecase/message_usecase.go
package usecase

import (
	"github.com/Ulyana-kru00/forum-project/chat/internal/entity"
"github.com/Ulyana-kru00/forum-project/chat/internal/repository"
)

type MessageUseCase interface {
	SaveMessage(msg entity.Message) error
	GetMessages() ([]entity.Message, error)
}

type messageUseCase struct {
	repo repository.MessageRepository
}

func NewMessageUseCase(repo repository.MessageRepository) MessageUseCase {
	return &messageUseCase{repo: repo}
}

func (uc *messageUseCase) SaveMessage(msg entity.Message) error {
	return uc.repo.SaveMessage(msg)
}

func (uc *messageUseCase) GetMessages() ([]entity.Message, error) {
	return uc.repo.GetMessages()
}
