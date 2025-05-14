package tests

import (
	"e2etests/data/testdata"
	"e2etests/pkg/utils"
	"e2etests/pkg/utils/page"
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"testing"
)

func AuthTestGroup() utils.TestGroup {
	return utils.TestGroup{
		Name: "AuthTest",
		Tests: []utils.TestCase{
			{
				Name: "SuccessfulAuth",
				Test: SuccessfulAuth,
			},
			{
				Name: "FailedAuth",
				Test: FailedAuth,
			},
		},
	}
}

func SuccessfulAuth(t *testing.T, driver selenium.WebDriver) {
	testData := testdata.GetValidLoginData()

	authPage := page.Auth{}
	authPage.Init(driver)
	err := authPage.OpenPage(testdata.LoginUrl)
	assert.NoError(t, err, "Couldn't open auth page")

	err = authPage.Login(testData.Login, testData.Password)
	assert.NoError(t, err, "Writing auth data error")

	isLoginSuccessful, err := authPage.IsLoginSuccessful()
	assert.NoError(t, err, "Auth error")
	assert.True(t, isLoginSuccessful, "Auth should be successful")
}

func FailedAuth(t *testing.T, driver selenium.WebDriver) {
	testData := testdata.GetInvalidLoginData()

	authPage := page.Auth{}
	authPage.Init(driver)
	err := authPage.OpenPage(testdata.LoginUrl)
	assert.NoError(t, err, "Couldn't open auth page")

	err = authPage.Login(testData.Login, testData.Password)
	assert.NoError(t, err, "Writing auth data error")

	isLoginSuccessful, err := authPage.IsLoginSuccessful()
	assert.Error(t, err, "No auth error")
	assert.False(t, isLoginSuccessful, "Auth shouldn't be successful")
}
