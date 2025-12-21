package service

import (
	"context"
	"fmt"
	"time"

	"ride-sharing/services/payment-service/internal/domain"
	"ride-sharing/services/payment-service/pkg/types"

	"github.com/google/uuid"
)

type paymentService struct {
	paymentProcessor domain.PaymentProcessor
}

// NewPaymentService creates a new instance of the payment service
func NewPaymentService(paymentProcessor domain.PaymentProcessor) domain.Service {
	return &paymentService{
		paymentProcessor: paymentProcessor,
	}
}
