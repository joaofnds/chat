package conversation_test

import (
	"app/config"
	"app/conversation"
	"app/message"
	"app/test"
	"app/test/driver"
	"app/user"
	"fmt"
	"testing"

	"app/adapters/http"
	"app/adapters/logger"
	"app/adapters/postgres"
	. "app/test/matchers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestConversationController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "conversation controller")
}

var _ = Describe("/conversation", func() {
	var (
		app   *driver.Driver
		fxApp *fxtest.App
	)

	BeforeEach(func() {
		var httpConfig http.Config
		fxApp = fxtest.New(
			GinkgoT(),
			logger.NopLoggerProvider,
			test.RandomAppConfigPort,
			test.Transaction,
			config.Module,
			http.Module,
			postgres.Module,
			user.Module,
			conversation.Module,
			fx.Populate(&httpConfig),
		).RequireStart()
		url := fmt.Sprintf("http://localhost:%d", httpConfig.Port)
		app = driver.NewDriver(url)
	})

	AfterEach(func() { fxApp.RequireStop() })

	It("works", func() {
		alice := Must2(app.CreateUser("alice"))
		bob := Must2(app.CreateUser("bob"))
		convo := Must2(app.CreateConversation(alice, bob))
		msg := Must2(app.SendMessage(convo, alice, "Hello"))

		convo = Must2(app.GetConversation(convo.ID))
		Expect(convo.Users).To(Equal([]user.User{alice, bob}))
		Expect(convo.Messages).To(Equal([]message.Message{msg}))
	})
})
