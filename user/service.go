package user

import (
	"context"
)

type Service struct {
	repo  Repository
	probe Probe
}

func NewUserService(repo Repository, probe Probe) *Service {
	return &Service{repo, probe}
}

func (service *Service) CreateUser(name string) (User, error) {
	user := User{Name: name}

	err := service.repo.CreateUser(context.Background(), &user)
	if err != nil {
		service.probe.FailedToCreateUser(err)
	}
	service.probe.UserCreated()

	return user, err
}

func (service *Service) DeleteAll() error {
	err := service.repo.DeleteAll(context.Background())

	if err != nil {
		service.probe.FailedToDeleteAll(err)
	}

	return err
}

func (service *Service) List() ([]User, error) {
	return service.repo.All(context.Background())
}

func (service *Service) Find(id string) (User, error) {
	return service.repo.Find(context.Background(), id)
}

func (service *Service) FindByName(name string) (User, error) {
	user, err := service.repo.FindByName(context.Background(), name)
	if err != nil {
		service.probe.FailedToFindByName(err)
	}

	return user, err
}

func (service *Service) Remove(user User) error {
	err := service.repo.Delete(context.Background(), user)

	if err != nil {
		service.probe.FailedToRemoveUser(err, user)
	}

	return err
}
