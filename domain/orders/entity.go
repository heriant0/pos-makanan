package orders

import (
	"time"

	paymentgateway "github.com/heriant0/pos-makanan/external/payment-gateway"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id          int       `db:"id" json:"id"`
	CategoryId  int       `db:"category_id" json:"categoryId"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Price       float64   `db:"price" json:"price"`
	Stock       int       `db:"stock" json:"stock"`
	ImageUrl    string    `db:"image_url" json:"image_url"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}

type ProductRecord struct {
	Id    int     `bson:"id"`
	Price float32 `bson:"price"`
	Name  string  `bson:"name"`
}

type Order struct {
	Id         primitive.ObjectID `bson:"_id"`
	InvoiceId  string             `bson:"invoice_id"`
	Product    ProductRecord      `bson:"product"`
	Quantity   int                `bson:"quantity"`
	Amount     float32            `bson:"amount"`
	InvoiceUrl string             `bson:"invoice_url"`
	Status     string             `bson:"status"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	DeletedAt  *time.Time         `bson:"deleted_at"`
}

type OrderStatus string

func (o OrderStatus) value() string {
	return string(o)
}

const (
	Paid    OrderStatus = "Paid"
	Pending OrderStatus = "Pending"
	Expired OrderStatus = "Expired"
)

type Invoice struct {
	ExternalId string  `json:"external_id"`
	Amount     float32 `json:"amount"`
	Status     OrderStatus
	InvoiceUrl string
}

func (i Invoice) toInvoiceRequest() paymentgateway.InvoiceRequest {
	return paymentgateway.InvoiceRequest{
		Amount:     i.Amount,
		ExternalId: i.ExternalId,
	}
}
