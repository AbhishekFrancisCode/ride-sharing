package main

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()
	repo := repository.NewinmemRepository()
	svc := service.NewService(repo)

	fare := &domain.RideFareModel{
		UserID:      "user123",
		PackageSlug: "standard",
		TotalPrice:  15.0,
	}
	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		panic(err)
	}
	println("Created trip with ID:", t.UserID)

	for {
		time.Sleep(time.Second)
	}
}
