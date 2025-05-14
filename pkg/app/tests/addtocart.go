package tests

import (
	"e2etests/pkg/app/page"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"e2etests/data/testdata"
	"e2etests/pkg/app/model"
)

func AddToCartTestGroup() model.TestGroup {
	return model.TestGroup{
		Name: "AddToCartTest",
		Tests: []model.TestCase{
			{
				Name: "TestAddToCartInMainPage",
				Test: TestAddToCartInMainPage,
			},
			{
				Name: "TestAddOneToCartInItemPage",
				Test: TestAddOneToCartInItemPage,
			},
			{
				Name: "TestAddSeveralToCartInItemPage",
				Test: TestAddSeveralToCartInItemPage,
			},
		},
	}
}

func TestAddToCartInMainPage(t *testing.T, driver selenium.WebDriver) {
	productCasio := testdata.ItemCasio
	productPage := page.Item{}
	productPage.Init(driver)
	err := productPage.OpenPage("")
	assert.NoError(t, err, "Не удалось открыть главную страницу")

	err = productPage.AddToCart(productCasio.ID)
	assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

	err = productPage.IsItemInCart(productCasio.Name, productCasio.Price, testdata.QuantityItemsOne)
	assert.NoError(t, err, "Товар отсутствует в корзине или неверные параметры")
}

func TestAddOneToCartInItemPage(t *testing.T, driver selenium.WebDriver) {
	productCasio := testdata.ItemCasio
	productPage := page.Item{}
	productPage.Init(driver)
	err := productPage.OpenPage(productCasio.URL)
	assert.NoError(t, err, "Не удалось открыть страницу товара")

	err = productPage.AddToCart(productCasio.ID)
	assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

	err = productPage.IsItemInCart(productCasio.Name, productCasio.Price, testdata.QuantityItemsOne)
	assert.NoError(t, err, "Товар отсутствует в корзине или неверные параметры")
}

func TestAddSeveralToCartInItemPage(t *testing.T, driver selenium.WebDriver) {
	productCasio := testdata.ItemCasio
	productPage := page.Item{}
	productPage.Init(driver)
	err := productPage.OpenPage(productCasio.URL)
	assert.NoError(t, err, "Не удалось открыть страницу товара")

	err = productPage.SetItemQuantity(testdata.QuantityItemsTen)
	assert.NoError(t, err, "Ошибка при установке количества товара")

	err = productPage.AddToCart(productCasio.ID)
	assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

	err = productPage.IsItemInCart(productCasio.Name, productCasio.Price, testdata.QuantityItemsTen)
	assert.NoError(t, err, "Товар отсутствует в корзине или неверное количество")
}
