package services

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "crm-system/internal/models"
)

type InteractionService struct {
    collection *mongo.Collection
}

func NewInteractionService(db *mongo.Database) *InteractionService {
    return &InteractionService{
        collection: db.Collection("interactions"),
    }
}

func (s *InteractionService) CreateInteraction(interaction *models.Interaction) error {
    interaction.CreatedAt = time.Now()
    interaction.UpdatedAt = time.Now()

    _, err := s.collection.InsertOne(context.Background(), interaction)
    return err
}

func (s *InteractionService) GetInteraction(id string) (*models.Interaction, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var interaction models.Interaction
    err = s.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&interaction)
    if err != nil {
        return nil, err
    }

    return &interaction, nil
}

func (s *InteractionService) UpdateInteraction(id string, updates *models.Interaction) error {
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

func (s *InteractionService) DeleteInteraction(id string) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = s.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
    return err
}

func (s *InteractionService) ListInteractionsByCustomer(customerID string) ([]*models.Interaction, error) {
    customerObjectID, err := primitive.ObjectIDFromHex(customerID)
    if err != nil {
        return nil, err
    }

    cursor, err := s.collection.Find(context.Background(), bson.M{"customer_id": customerObjectID})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var interactions []*models.Interaction
    for cursor.Next(context.Background()) {
        var interaction models.Interaction
        if err := cursor.Decode(&interaction); err != nil {
            return nil, err
        }
        interactions = append(interactions, &interaction)
    }

    return interactions, nil
}

func (s *InteractionService) ScheduleInteraction(interaction *models.Interaction) error {
    interaction.CreatedAt = time.Now()
    interaction.UpdatedAt = time.Now()

    _, err := s.collection.InsertOne(context.Background(), interaction)
    return err
}