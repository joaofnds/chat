package user

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewController(service *Service) *Controller {
	return &Controller{service}
}

type Controller struct {
	service *Service
}

func (c *Controller) Register(app *fiber.App) {
	app.Post("/users", c.Create)
	app.Get("/users", c.List)
	app.Get("/users/:id", c.Get)
	app.Delete("/users/:name", c.Delete)
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	var user User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	user, err = c.service.CreateUser(user.Name)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.Status(http.StatusCreated).JSON(user)
}

func (c *Controller) List(ctx *fiber.Ctx) error {
	users, err := c.service.List()
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.JSON(users)
}

func (c *Controller) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	user, err := c.service.Find(id)
	switch {
	case errors.Is(err, ErrNotFound):
		return ctx.SendStatus(http.StatusNotFound)
	case errors.Is(err, ErrRepository):
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.JSON(user)
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	if name == "" {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	user, err := c.service.FindByName(name)
	switch {
	case errors.Is(err, ErrNotFound):
		return ctx.SendStatus(http.StatusNotFound)
	case errors.Is(err, ErrRepository):
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	if err = c.service.Remove(user); err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.SendStatus(http.StatusOK)
}
