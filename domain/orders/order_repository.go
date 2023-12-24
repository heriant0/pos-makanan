package orders

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	dbPg    *sqlx.DB
	dbMongo *mongo.Client
}

func newRespository(db *sqlx.DB, dbMongo *mongo.Client) repository {
	return repository{db, dbMongo}
}

// insertOrder implements orderRepository.
func (r repository) insertOrder(ctx context.Context, order Order) (err error) {
	collection := r.dbMongo.Database(mongoDatabase).Collection(mongoOrderCollection)

	order.Id = primitive.NewObjectID()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.DeletedAt = nil

	_, err = collection.InsertOne(ctx, order)
	if err != nil {
		return
	}

	return
}

// findLastestOrder implements orderRepository.
func (r repository) findLastestOrderInvoiceId(ctx context.Context) (invoiceId string, err error) {
	collection := r.dbMongo.Database(mongoDatabase).Collection(mongoOrderCollection)
	options := options.FindOne().SetSort(bson.M{"invoice_id": -1})

	lastOrder := Order{}
	err = collection.FindOne(ctx, bson.M{}, options).Decode(&lastOrder)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("oyeah")
			return
		}
		err = nil
	}

	invoiceId = lastOrder.InvoiceId
	return
}

// updateOrderStatus implements orderRepository.
func (r repository) updateOrderStatus(ctx context.Context, invoice Invoice) (err error) {
	collection := r.dbMongo.Database(mongoDatabase).Collection(mongoOrderCollection)

	update := bson.M{
		"status": invoice.Status.value(),
	}
	_, err = collection.
		UpdateOne(
			ctx,
			bson.M{"invoice_id": invoice.ExternalId},
			bson.M{"$set": update},
		)

	return
}
