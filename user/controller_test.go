package user_test

import (
	apphttp "app/adapters/http"
	"app/adapters/logger"
	"app/adapters/postgres"
	"app/config"
	"app/test"
	"app/test/driver"
	. "app/test/matchers"
	"app/user"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

var _ = Describe("/users", Ordered, func() {
	var (
		app         *driver.Driver
		fxApp       *fxtest.App
		userService *user.Service
	)

	BeforeAll(func() {
		var httpConfig apphttp.Config

		fxApp = fxtest.New(
			GinkgoT(),
			logger.NopLoggerProvider,
			test.RandomAppConfigPort,
			config.Module,
			apphttp.Module,
			postgres.Module,
			user.Module,
			fx.Populate(&httpConfig, &userService),
		).RequireStart()

		app = driver.NewDriver(fmt.Sprintf("http://localhost:%d", httpConfig.Port))
	})

	AfterAll(func() {
		fxApp.RequireStop()
	})

	BeforeEach(func() {
		Must(userService.DeleteAll())
	})

	It("creates and lists users", func() {
		joao := Must2(app.CreateUser("joao"))
		fnds := Must2(app.CreateUser("fnds"))

		users := Must2(app.ListUsers())

		Expect(users).To(Equal([]user.User{joao, fnds}))
	})

	It("gets the user", func() {
		joao := Must2(app.CreateUser("joao"))
		found := Must2(app.GetUser(joao.ID))
		Expect(found).To(Equal(joao))
	})
})
