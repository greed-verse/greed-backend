package payment

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/greed-verse/greed/internal/payment/repo"
	"github.com/jackc/pgx/v5/pgtype"
)

func (p *Payment) CreateWallet(userId string, balance pgtype.Numeric) (repo.Wallet, error) {
	params := repo.CreateWalletParams{
		UserID:  userId,
		Balance: balance,
	}
	return p.repo.CreateWallet(context.Background(), params)
}

func (p *Payment) walletHandler(msg *message.Message) error {
	f := 0.0

	num := &pgtype.Numeric{}
	err := num.Scan(fmt.Sprintf("%.2f", f)) // .2 specifies precision
	if err != nil {
		return err
	}
	fmt.Println(num)

	_, err = p.CreateWallet(string(msg.Payload), *num)
	if err != nil {
		return err
	}
	return nil
}
