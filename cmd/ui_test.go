package cmd

import (
	"e2etests/pkg/tests"
	"fmt"
	"testing"

	"e2etests/pkg/utils"
)

func Test(t *testing.T) {
	testGroups := getTestGroups()
	for _, group := range testGroups {
		fmt.Printf("Testing group %s\n", group.Name)
		for _, test := range group.Tests {
			test := test

			t.Run(test.Name+"/chrome", func(t *testing.T) {
				t.Parallel()
				utils.RunTestForBrowser(t, test.Name, "chrome", test.Test)
			})

			t.Run(test.Name+"/firefox", func(t *testing.T) {
				t.Parallel()
				utils.RunTestForBrowser(t, test.Name, "firefox", test.Test)
			})
		}
	}
}

func getTestGroups() []utils.TestGroup {
	return []utils.TestGroup{
		tests.AuthTestGroup(),
		tests.SearchTestGroup(),
		tests.AddToCartTestGroup(),
		tests.OrderTestGroup(),
	}
}
