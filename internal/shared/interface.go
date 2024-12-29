package shared

import (
	"github.com/google/uuid"
	"github.com/greed-verse/greed/internal/payment/repo"
	"github.com/jackc/pgx/v5/pgtype"
)

type WalletService interface {
	CreateWallet(userId uuid.UUID, balance pgtype.Numeric) (repo.Wallet, error)
}
