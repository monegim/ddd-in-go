package main

import "context"

type LeadRequest struct {
	email string
}

type Lead struct {
	id string
}

type LeadCreator interface {
	createLead(ctx context.Context, request LeadRequest) (Lead, error)
}

type Customer struct {
	leadID string
	userID string
}

func (c *Customer) UserID() string {
	return c.userID
}

func (c *Customer) SetUserID(userID string) {
	c.userID = userID
}

type LeadConvertor interface {
	Convert(ctx context.Context, subSelection SubscriptionType)(Customer, error)
}

func (l Lead)Convert(ctx context.Context, subSelection SubscriptionType)(Customer, error) {
	//TODO implement me
	panic("implement me")
}
