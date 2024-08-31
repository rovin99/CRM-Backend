package services

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "crm-system/internal/models"
)

type TicketService struct {
    collection *mongo.Collection
}

func NewTicketService(db *mongo.Database) *TicketService {
    return &TicketService{
        collection: db.Collection("tickets"),
    }
}

func (s *TicketService) CreateTicket(ticket *models.Ticket) error {
    ticket.CreatedAt = time.Now()
    ticket.UpdatedAt = time.Now()
    ticket.Status = "open"

    _, err := s.collection.InsertOne(context.Background(), ticket)
    return err
}

func (s *TicketService) GetTicket(id string) (*models.Ticket, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var ticket models.Ticket
    err = s.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&ticket);

	return err
}

func (s *TicketService) UpdateTicket(id string, updates *models.Ticket) error {
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

func (s *TicketService) ResolveTicket(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		bson.M{"$set": bson.M{"status": "resolved", "resolved_at": time.Now()}},
	)

	return err
}

func (s *TicketService) DeleteTicket(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}

