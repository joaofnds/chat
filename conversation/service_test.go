package conversation_test

import (
	"app/adapters/logger"
	"app/adapters/postgres"
	"app/config"
	"app/conversation"
	"app/message"
	"app/test"
	. "app/test/matchers"
	"app/user"
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestConversation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "conversation suite")
}

var _ = Describe("conversation service", func() {
	var (
		app         *fxtest.App
		sut         *conversation.Service
		userService *user.Service
		ctx         context.Context
	)

	BeforeEach(func() {
		ctx = context.Background()
		app = fxtest.New(
			GinkgoT(),
			config.Module,
			logger.NopLoggerProvider,
			test.RandomAppConfigPort,
			test.Transaction,
			test.TestFiber,
			test.TestSocketIO,
			user.NopProbeProvider,
			user.Module,
			postgres.Module,
			conversation.Module,
			fx.Populate(&sut, &userService),
		)
		app.RequireStart()
	})

	AfterEach(func() { app.RequireStop() })

	Describe("conversation", func() {
		It("starts empty", func() {
			alice := Must2(userService.CreateUser("alice"))
			bob := Must2(userService.CreateUser("bob"))
			convo := Must2(sut.Create(ctx, alice, bob))

			found := Must2(sut.Find(ctx, convo.ID))
			Expect(found.Users).To(Equal([]user.User{alice, bob}))
			Expect(found.Messages).To(BeEmpty())
		})

		It("contains messages sent", func() {
			alice := Must2(userService.CreateUser("alice"))
			bob := Must2(userService.CreateUser("bob"))
			convo := Must2(sut.Create(ctx, alice, bob))

			hello := Must2(sut.SendMessage(ctx, convo, alice, "Hello"))
			world := Must2(sut.SendMessage(ctx, convo, bob, "World"))

			found := Must2(sut.Find(ctx, convo.ID))
			Expect(found.Messages).To(Equal([]message.Message{hello, world}))
		})

		It("finds for both users", func() {
			alice := Must2(userService.CreateUser("alice"))
			bob := Must2(userService.CreateUser("bob"))
			convo := Must2(sut.Create(ctx, alice, bob))

			aliceChats := Must2(sut.FindForUser(ctx, alice))
			Expect(aliceChats).To(Equal([]conversation.Conversation{convo}))

			bobChats := Must2(sut.FindForUser(ctx, bob))
			Expect(bobChats).To(Equal([]conversation.Conversation{convo}))
		})

		It("does not find when user is not in", func() {
			alice := Must2(userService.CreateUser("alice"))
			bob := Must2(userService.CreateUser("bob"))
			aliceAndBob := Must2(sut.Create(ctx, alice, bob))

			charlie := Must2(userService.CreateUser("charlie"))
			dave := Must2(userService.CreateUser("dave"))
			charlieAndDave := Must2(sut.Create(ctx, charlie, dave))

			aliceChats := Must2(sut.FindForUser(ctx, alice))
			Expect(aliceChats).NotTo(ContainElement(charlieAndDave))

			charlieChats := Must2(sut.FindForUser(ctx, charlie))
			Expect(charlieChats).NotTo(ContainElement(aliceAndBob))
		})
	})
})
