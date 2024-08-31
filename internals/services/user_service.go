package services

import (
    "context"
    "errors"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "crm-system/internal/models"
    "crm-system/pkg/utils"
)

type UserService struct {
    collection *mongo.Collection
}

func NewUserService(db *mongo.Database) *UserService {
    return &UserService{
        collection: db.Collection("users"),
    }
}

func (s *UserService) CreateUser(user *models.User) error {
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword

    _, err = s.collection.InsertOne(context.Background(), user)
    return err
}

func (s *UserService) GetUser(id string) (*models.User, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var user models.User
    err = s.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (s *UserService) UpdateUser(id string, updates *models.User) error {
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

func (s *UserService) DeleteUser(id string) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = s.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
    return err
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    err := s.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}