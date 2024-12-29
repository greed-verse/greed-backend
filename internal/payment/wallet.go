package payment

import (
	"context"

	"github.com/google/uuid"
	"github.com/greed-verse/greed/internal/payment/repo"
	"github.com/jackc/pgx/v5/pgtype"
)

type walletService struct {
	repo *repo.Queries
}

func NewWalletService(module *Payment) *walletService {
	repo := module.GetRepo()
	return &walletService{
		repo: repo,
	}
}

func (ws *walletService) CreateWallet(userId uuid.UUID, balance pgtype.Numeric) (repo.Wallet, error) {
	params := repo.CreateWalletParams{
		UserID:  userId,
		Balance: balance,
	}
	return ws.repo.CreateWallet(context.Background(), params)
}
