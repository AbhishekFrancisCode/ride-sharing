package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
)

type inmemRepository struct {
	trip      map[string]*domain.TripModel
	ridefares map[string]*domain.RideFareModel
}

func NewinmemRepository() *inmemRepository {
	return &inmemRepository{
		trip:      make(map[string]*domain.TripModel),
		ridefares: make(map[string]*domain.RideFareModel),
	}
}

func (r *inmemRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	r.trip[trip.ID.Hex()] = trip
	return trip, nil
}
