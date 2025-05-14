package cmd

import (
	"e2etests/pkg/app/model"
	"e2etests/pkg/app/tests"
	"e2etests/pkg/infrastructure"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testGroups := getTestGroups()
	for _, group := range testGroups {
		fmt.Printf("Testing group %s\n", group.Name)
		for _, test := range group.Tests {
			groupName := group.Name
			test := test

			t.Run(groupName+"/chrome", func(t *testing.T) {
				t.Parallel()
				infrastructure.RunTestForBrowser(t, test.Name, "chrome", test.Test)
			})

			t.Run(groupName+"/firefox", func(t *testing.T) {
				t.Parallel()
				infrastructure.RunTestForBrowser(t, test.Name, "firefox", test.Test)
			})
		}
	}
}

func getTestGroups() []model.TestGroup {
	return []model.TestGroup{
		tests.AuthTestGroup(),
		tests.SearchTestGroup(),
		tests.AddToCartTestGroup(),
		tests.OrderTestGroup(),
	}
}
