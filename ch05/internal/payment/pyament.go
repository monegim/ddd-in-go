package payment

import (
	"github.com/stripe/stripe-go/v73/client"
)
type StripeService struct {
	stripeClient *client.API
}