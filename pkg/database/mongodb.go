package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Name      string             `bson:"name"`
    Email     string             `bson:"email"`
    Password  string             `bson:"password"`
    Company   string             `bson:"company,omitempty"`
    Role      string             `bson:"role"`
    CreatedAt time.Time          `bson:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at"`
}

type Customer struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Name      string             `bson:"name"`
    Email     string             `bson:"email"`
    Phone     string             `bson:"phone"`
    Company   string             `bson:"company,omitempty"`
    Status    string             `bson:"status"`
    Notes     string             `bson:"notes,omitempty"`
    CreatedAt time.Time          `bson:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at"`
}

type Interaction struct {
    ID         primitive.ObjectID `bson:"_id,omitempty"`
    CustomerID primitive.ObjectID `bson:"customer_id"`
    UserID     primitive.ObjectID `bson:"user_id"`
    Type       string             `bson:"type"` // meeting, call, email, etc.
    Notes      string             `bson:"notes"`
    ScheduledAt time.Time         `bson:"scheduled_at"`
    CreatedAt  time.Time          `bson:"created_at"`
    UpdatedAt  time.Time          `bson:"updated_at"`
}

type Ticket struct {
    ID         primitive.ObjectID `bson:"_id,omitempty"`
    CustomerID primitive.ObjectID `bson:"customer_id"`
    Subject    string             `bson:"subject"`
    Description string            `bson:"description"`
    Status     string             `bson:"status"` // open, in-progress, resolved
    CreatedAt  time.Time          `bson:"created_at"`
    UpdatedAt  time.Time          `bson:"updated_at"`
    ResolvedAt time.Time          `bson:"resolved_at,omitempty"`
}