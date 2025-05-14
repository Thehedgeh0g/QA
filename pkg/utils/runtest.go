package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
)

func RunTestForBrowser(t *testing.T, testName, browserName string, testFunc func(*testing.T, selenium.WebDriver)) {
	t.Helper()
	t.Run(fmt.Sprintf("%s-%s", testName, browserName), func(t *testing.T) {
		caps := selenium.Capabilities{
			"browserName": browserName,
		}
		driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
		if !assert.NoError(t, err, "Failed to start "+browserName+" session") {
			return
		}
		defer driver.Quit()
		testFunc(t, driver)
	})
}
