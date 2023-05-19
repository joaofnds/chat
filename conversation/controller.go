package conversation

import (
	"app/user"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewController(service *Service, userService *user.Service) *Controller {
	return &Controller{conversationService: service, userService: userService}
}

type Controller struct {
	conversationService *Service
	userService         *user.Service
}

func (c *Controller) Register(app *fiber.App) {
	convos := app.Group("/conversations")
	convos.Get("/", c.getConversations)
	convos.Post("/", c.createConversation)

	convo := convos.Group("/:id", c.middlewareFindConversation)
	convo.Get("/", c.getConversation)
	convo.Post("/", c.sendMessage)
}

func (c *Controller) getConversations(ctx *fiber.Ctx) error {
	userID := ctx.Query("userID")
	if userID == "" {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	u, err := c.userService.Find(userID)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return ctx.SendStatus(http.StatusNotFound)
		}
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	convos, err := c.conversationService.FindForUser(ctx.Context(), u)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}
	return ctx.JSON(convos)
}

func (c *Controller) createConversation(ctx *fiber.Ctx) error {
	var body CreateConversationBody
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	sender, err := c.userService.Find(body.SenderID)
	if err != nil {
		return ctx.SendStatus(http.StatusNotFound)
	}

	receiver, err := c.userService.Find(body.ReceiverID)
	if err != nil {
		return ctx.SendStatus(http.StatusNotFound)
	}

	convo, err := c.conversationService.Create(ctx.Context(), sender, receiver)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.Status(http.StatusCreated).JSON(convo)
}

func (c *Controller) sendMessage(ctx *fiber.Ctx) error {
	convo := ctx.Locals("conversation").(Conversation)

	var body SendMessageBody
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	author, err := c.userService.Find(body.AuthorID)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return ctx.SendStatus(http.StatusNotFound)
		}
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	msg, err := c.conversationService.SendMessage(ctx.Context(), convo, author, body.Text)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.Status(http.StatusCreated).JSON(msg)
}

func (c *Controller) getConversation(ctx *fiber.Ctx) error {
	return ctx.JSON(ctx.Locals("conversation"))
}

func (c *Controller) middlewareFindConversation(ctx *fiber.Ctx) error {
	convo, err := c.conversationService.Find(ctx.Context(), ctx.Params("id"))
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return ctx.SendStatus(http.StatusNotFound)
		}

		return ctx.SendStatus(http.StatusInternalServerError)
	}
	ctx.Locals("conversation", convo)
	return ctx.Next()
}

type CreateConversationBody struct {
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
}

type SendMessageBody struct {
	AuthorID string `json:"author_id"`
	Text     string `json:"text"`
}
