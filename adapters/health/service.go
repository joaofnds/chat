package health

import (
	"app/adapters/postgres"

	"context"
)

type Service struct {
	postgresHealth postgres.HealthChecker
}

func NewHealthService(postgresHealth postgres.HealthChecker) *Service {
	return &Service{postgresHealth}
}

func (s *Service) CheckHealth(ctx context.Context) Check {
	return Check{
		"postgres": NewStatus(s.postgresHealth.CheckHealth(ctx)),
	}
}
