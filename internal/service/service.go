package service

import (
	"context"

	"github.com/dahaev/bottesting/pkg/models"
)

type Service struct {
	Repository Repo
}

func New(repo Repo) *Service {
	return &Service{
		Repository: repo,
	}
}

func (s *Service) CreateLadyAccount(ctx context.Context, account *models.Account) error {
	return s.Repository.CreateLadyAccount(ctx, account)
}

func (s *Service) GetAccountLady(ctx context.Context, userName string) (*models.Account, error) {
	account, err := s.Repository.GetAccountLady(ctx, userName)
	if err != nil {
		return nil, err
	}
	return account, err
}

func (s *Service) CreateDonAccount(ctx context.Context, username string) error {
	return s.Repository.CreateDonAccount(ctx, username)
}

func (s *Service) GetDonAccount(ctx context.Context, username string) (*models.DonAccount, error) {
	don, err := s.Repository.GetDonAccount(ctx, username)
	if err != nil {
		return nil, err
	}
	return don, err
}

func (s *Service) CreateReview(ctx context.Context, ladyName string, donName string, review string, rating int) error {
	return s.Repository.CreateReview(ctx, ladyName, donName, review, rating)
}

func (s *Service) GetLadyReviews(ctx context.Context, ladyName string) ([]models.Review, error) {
	res, err := s.Repository.GetLadyReviews(ctx, ladyName)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Repo interface {
	CreateLadyAccount(ctx context.Context, account *models.Account) error
	GetAccountLady(ctx context.Context, userName string) (*models.Account, error)
	CreateDonAccount(ctx context.Context, username string) error
	GetDonAccount(ctx context.Context, userName string) (*models.DonAccount, error)
	CreateReview(ctx context.Context, ladyName string, donName string, review string, rating int) error
	GetLadyReviews(ctx context.Context, ladyName string) ([]models.Review, error)
}
