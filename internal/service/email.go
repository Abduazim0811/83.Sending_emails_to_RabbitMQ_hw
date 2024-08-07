package service

import (
	"log"
	"mailer-rabbitmq/internal/models"
	"mailer-rabbitmq/internal/rabbitmq"
)

type EmailService struct {
	Publisher *rabbitmq.Publisher
}

func NewEmailService(publisher *rabbitmq.Publisher) *EmailService {
	return &EmailService{
		Publisher: publisher,
	}
}

func (s *EmailService) SendEmail(email models.Email) error {
	log.Printf("Preparing to send email to %v", email.To)

	// Publish email to RabbitMQ
	if err := s.Publisher.Publish(email); err != nil {
		return err
	}

	return nil
}
