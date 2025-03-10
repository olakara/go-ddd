package chapter2

import "context"

type LeadRequest struct {
	Email string
}

type Lead struct {
	ID string
}

type LeadCreator interface {
	CreateLead(context context.Context, request LeadRequest) (Lead, error)
}

type Customer struct {
	leadID string
	userID string
}

func (c *Customer) UserID() string {
	return c.userID
}

func (c *Customer) SetUserID(id string) {
	c.userID = id
}

type LeadConverter interface {
	Covert(ctx context.Context, subSelection SubscriptionType) (Customer, error)
}

func (l Lead) Convert(ctx context.Context, subSelection SubscriptionType) (Customer, error) {
	// TODO implement conversion logic
	panic("Covert Lead to Customer needs implementation!")
}
