package ch04

import (
	"context"
	"ddd-in-go/ch02"
)

type accountKey = int

const AccountCtxKey = accountKey(1)

type BookingDomainService interface {
	CreateBooking(ctx context.Context, booking Booking) error
}

type BookingAppService struct {
	bookingRepo BookingRepository
	bookingDomainService BookingDomainService
}

func NewBookingAppService(bookingRepo BookingRepository, bookingDomainService BookingDomainService) *BookingAppService {
	return &BookingAppService{bookingRepo: bookingRepo, bookingDomainService: bookingDomainService}
}

func (b *BookingAppService) CreateBooking(ctx context.Context, booking Booking) error {
	_= ctx.Value(AccountCtxKey).(ch02.Customer)
	//TODO
	return nil
}
