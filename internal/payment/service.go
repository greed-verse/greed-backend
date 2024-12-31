package payment

import (
	"context"
	"math/big"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/greed-verse/greed/internal/payment/repo"
	"github.com/jackc/pgx/v5/pgtype"
)

func (p *Payment) CreateWallet(userId uuid.UUID, balance pgtype.Numeric) (repo.Wallet, error) {
	params := repo.CreateWalletParams{
		UserID:  userId,
		Balance: balance,
	}
	return p.repo.CreateWallet(context.Background(), params)
}

func (p *Payment) walletHandler(msg *message.Message) error {
	userid, err := uuid.FromBytes(msg.Payload)
	if err != nil {
		return err
	}
	var balance pgtype.Numeric
	balance.Int.Set(big.NewInt(int64(0.0)))

	_, err = p.CreateWallet(userid, balance)
	if err != nil {
		return err
	}

	return nil
}
