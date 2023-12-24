package orders

import "time"

type createOrderRequest struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func (c createOrderRequest) toOrderEntity() Order {
	return Order{
		Quantity: c.Quantity,
	}
}

type updateOrderRequest struct {
	Id                 string    `json:"id"`
	ExternalID         string    `json:"external_id"`
	UserID             string    `json:"user_id"`
	IsHigh             bool      `json:"is_high"`
	PaymentMethod      string    `json:"payment_method"`
	Status             string    `json:"status"`
	MerchantName       string    `json:"merchant_name"`
	Amount             int       `json:"amount"`
	PaidAmount         int       `json:"paid_amount"`
	BankCode           string    `json:"bank_code"`
	PaidAt             time.Time `json:"paid_at"`
	PayerEmail         string    `json:"payer_email"`
	Description        string    `json:"description"`
	Created            time.Time `json:"created"`
	Updated            time.Time `json:"updated"`
	Currency           string    `json:"currency"`
	PaymentChannel     string    `json:"payment_channel"`
	PaymentDestination string    `json:"payment_destination"`
}

func (u updateOrderRequest) toInvoiceEntity() Invoice {
	return Invoice{
		ExternalId: u.ExternalID,
		Status:     OrderStatus(u.Status),
		Amount:     float32(u.Amount),
	}
}
