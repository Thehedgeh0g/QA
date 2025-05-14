package utils

import (
	"github.com/tebeka/selenium"
	"testing"
)

type TestCase struct {
	Name string
	Test func(t *testing.T, driver selenium.WebDriver)
}

type TestGroup struct {
	Name  string
	Tests []TestCase
}
