package payment

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
)

type item struct {
	Id     string
	Amount int64
}

func (p *Payment) Payment(c *fiber.Ctx) error {
	var req struct {
		Items []item `json:"items"`
	}

	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return err
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(req.Items)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	p.logger.Core().Info().Str("Payment Intent", pi.ClientSecret).Msg("Payment Intent Created")
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"client_secret": pi.ClientSecret,
	})
}

func calculateOrderAmount(items []item) int64 {
	var total int64
	for _, item := range items {
		total += item.Amount
	}
	return total
}
