package health

import "context"

func NewUnhealthyHealthService() UnhealthyHealthService {
	return UnhealthyHealthService{}
}

type UnhealthyHealthService struct{}

func (c UnhealthyHealthService) CheckHealth(_ context.Context) Check {
	return Check{
		"postgres": Status{Status: StatusDown},
	}
}
