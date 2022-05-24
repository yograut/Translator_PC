package main

import (
	"context"
	"time"

	"golang.org/x/text/language"
)

// Service is a Translator user.
type Service struct {
	translatorClient TranslatorAPI
}

func NewService() *Service {
	t := newTranslatorStub(
		100*time.Millisecond,
		500*time.Millisecond,
		0.1,
	)

	return &Service{
		translatorClient: t,
	}
}

func (s *Service) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {

	return s.translatorClient.Translate(ctx, from, to, data)
}
