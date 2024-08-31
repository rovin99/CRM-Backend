package services

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "crm-system/internal/models"
)

type CustomerService struct {
    collection *mongo.Collection
}

func NewCustomerService(db *mongo.Database) *CustomerService {
    return &CustomerService{
        collection: db.Collection("customers"),
    }
}

func (s *CustomerService) CreateCustomer(customer *models.Customer) error {
    customer.CreatedAt = time.Now()
    customer.UpdatedAt = time.Now()

    _, err := s.collection.InsertOne(context.Background(), customer)
    return err
}

func (s *CustomerService) GetCustomer(id string) (*models.Customer, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var customer models.Customer
    err = s.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&customer)
    if err != nil {
        return nil, err
    }

    return &customer, nil
}

func (s *CustomerService) UpdateCustomer(id string, updates *models.Customer) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    updates.UpdatedAt = time.Now()

    _, err = s.collection.UpdateOne(
        context.Background(),
        bson.M{"_id": objectID},
        bson.M{"$set": updates},
    )

    return err
}

func (s *CustomerService) DeleteCustomer(id string) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = s.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
    return err
}

func (s *CustomerService) ListCustomers(limit, offset int) ([]*models.Customer, error) {
    options := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
    cursor, err := s.collection.Find(context.Background(), bson.M{}, options)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var customers []*models.Customer
    for cursor.Next(context.Background()) {
        var customer models.Customer
        if err := cursor.Decode(&customer); err != nil {
            return nil, err
        }
        customers = append(customers, &customer)
    }

    return customers, nil
}