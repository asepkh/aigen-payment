package snap

import (
	"github.com/midtrans/midtrans-go/snap"

	"github.com/asepkh/aigen-go-payment/invoice"
)

func NewOtherBankVA(inv *invoice.Invoice) (*snap.Request, error) {
	return newBuilder(inv).
		AddPaymentMethods(snap.PaymentTypeOtherVA).
		Build()
}
