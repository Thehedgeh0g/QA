package tests

import (
	"e2etests/data/testdata"
	"e2etests/pkg/utils"
	"e2etests/pkg/utils/page"
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"testing"
)

func SearchTestGroup() utils.TestGroup {
	return utils.TestGroup{
		Name: "SearchTest",
		Tests: []utils.TestCase{
			{
				Name: "SearchInMainPage",
				Test: SearchInMainPage,
			},
			{
				Name: "TestSearchInItemPage",
				Test: TestSearchInItemPage,
			},
			{
				Name: "TestSearchInCategoryPage",
				Test: TestSearchInCategoryPage,
			},
			{
				Name: "TestSearchInSearchPage",
				Test: TestSearchInSearchPage,
			},
		},
	}
}

func SearchInMainPage(t *testing.T, driver selenium.WebDriver) {
	catalogPage := page.Catalog{}
	catalogPage.Init(driver)
	err := catalogPage.OpenPage("")
	assert.NoError(t, err, "Не удалось открыть главную страницу")

	err = catalogPage.SearchItem(testdata.ItemNameCasio)
	assert.NoError(t, err, "Ошибка при поиске товара Casio")

	err = catalogPage.FindItem(testdata.ItemNameCasio) // "Eldors Post"
	assert.NoError(t, err, "Товар Casio не найден на странице")
}

func TestSearchInItemPage(t *testing.T, driver selenium.WebDriver) {
	catalogPage := page.Catalog{}
	catalogPage.Init(driver)
	err := catalogPage.OpenPage(testdata.ItemPageUrl)
	assert.NoError(t, err, "Не удалось открыть страницу продукта")

	err = catalogPage.SearchItem(testdata.ItemNameRoyal)
	assert.NoError(t, err, "Ошибка при поиске товара Royal")

	err = catalogPage.FindItem(testdata.ItemNameRoyal) // "Eldors Post"
	assert.NoError(t, err, "Товар Royal не найден на странице продукта")
}

func TestSearchInCategoryPage(t *testing.T, driver selenium.WebDriver) {
	catalogPage := page.Catalog{}
	catalogPage.Init(driver)
	err := catalogPage.OpenPage(testdata.CategoryPageUrl)
	assert.NoError(t, err, "Не удалось открыть страницу категории")

	err = catalogPage.SearchItem(testdata.ItemNameRoyal)
	assert.NoError(t, err, "Ошибка при поиске товара Royal в категории")

	err = catalogPage.FindItem(testdata.ItemNameRoyal) // "Eldors Post"
	assert.NoError(t, err, "Товар Royal не найден в категории")
}

func TestSearchInSearchPage(t *testing.T, driver selenium.WebDriver) {
	catalogPage := page.Catalog{}
	catalogPage.Init(driver)
	err := catalogPage.OpenPage(testdata.SearchPageUrl)
	assert.NoError(t, err, "Не удалось открыть страницу поиска")

	err = catalogPage.SearchItem(testdata.ItemNameCitizen)
	assert.NoError(t, err, "Ошибка при поиске товара Citizen")

	err = catalogPage.FindItem(testdata.ItemNameCitizen) // "Eldors Post"
	assert.NoError(t, err, "Товар Citizen не найден на странице поиска")
}
