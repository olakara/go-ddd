package chapter2

import "context"

type UserType = int
type SubscriptionType = int

const (
	unknownUser UserType = iota
	lead
	customer
	churned
	lostLead
)

const (
	unknownSubscription SubscriptionType = iota
	basic
	premium
	exclusive
)

type UserAddRequest struct {
	Email          string
	SubType        SubscriptionType
	UserType       UserType
	PaymentDetails PaymentDetails
}

type UserModifyRequest struct {
	ID             string
	Email          string
	SubType        SubscriptionType
	UserType       UserType
	PaymentDetails PaymentDetails
}

type User struct {
	ID             string
	PaymentDetails PaymentDetails
}

type PaymentDetails struct {
	stripeTokenID string
}

type UserManager interface {
	AddUser(ctx context.Context, request UserAddRequest) (User, error)
	ModifyUser(ctx context.Context, request UserModifyRequest) (User, error)
}
