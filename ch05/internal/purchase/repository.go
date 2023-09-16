package purchase

import (
	coffeeco "coffeeco/internal"
	"coffeeco/internal/payment"
	"coffeeco/internal/store"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Repository interface {
	Store(ctx context.Context, purchase Purchase) error
}

type MongoRepository struct {
	purchases *mongo.Collection
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}
	purchases := client.Database("coffeeco").Collection("purchases")
	return &MongoRepository{
		purchases: purchases,
	}, nil
}

func (mr *MongoRepository) Store(ctx context.Context, purchase Purchase) error {
	mongP := toMongoPurchase(purchase)
	_, err := mr.purchases.InsertOne(ctx, mongP)
	if err != nil {
		return fmt.Errorf("failed to persist purchase: %w", err)
	}
	return nil
}

type mongoPurchase struct {
	ID                 uuid.UUID          `bson:"ID"`
	Store              store.Store        `bson:"Store"`
	ProductsToPurchase []coffeeco.Product `bson:"products_purchased"`
	Total              int64              `bson:"purchase_total"`
	PaymentMeans       payment.Means      `bson:"payment_means"`
	TimeOfPurchase     time.Time          `bson:"created_at"`
	CartToken          *string            `bson:"cart_token"`
}

func toMongoPurchase(p Purchase) mongoPurchase {
	return mongoPurchase{
		ID:                 p.id,
		Store:              p.Store,
		ProductsToPurchase: p.ProductsToPurchase,
		Total:              p.total.Amount(),
		PaymentMeans:       p.PaymentMeans,
		TimeOfPurchase:     p.timeOfPurchase,
		CartToken:          p.CardToken,
	}
}
