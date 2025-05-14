package page

import (
	"e2etests/pkg/utils"
	"fmt"
	"github.com/tebeka/selenium"
)

const (
	QuantityInput          = "quantity"
	ClearButton            = "//button[text()='Очистить корзину']"
	AddToCartButton        = "//a[@data-id='%s']"
	PopupItemNameValue     = "//table//td/a[contains(text(), '%s')]"
	PopupItemPriceValue    = "//table//td[contains(text(), '%s')]"
	PopupItemQuantityValue = "//table//td[contains(text(), '%s')]"
)

type Item struct {
	utils.SeleniumAdapter
}

func (p *Item) AddToCart(id string) error {
	elem, err := p.FindElement(selenium.ByXPATH, fmt.Sprintf(AddToCartButton, id))
	if err != nil {
		return err
	}

	return elem.Click()
}

func (p *Item) SetItemQuantity(quantity string) error {
	input, err := p.FindElement(selenium.ByName, QuantityInput)
	if err != nil {
		return err
	}

	if err := input.Clear(); err != nil {
		return fmt.Errorf("failed to clear input: %v", err)
	}

	return input.SendKeys(quantity)
}

func (p *Item) IsItemInCart(name, price, quantity string) error {
	_, err := p.FindElement(selenium.ByXPATH, fmt.Sprintf(PopupItemNameValue, name))
	if err != nil {
		return err
	}
	_, err = p.FindElement(selenium.ByXPATH, fmt.Sprintf(PopupItemPriceValue, price))
	if err != nil {
		return err
	}
	_, err = p.FindElement(selenium.ByXPATH, fmt.Sprintf(PopupItemQuantityValue, quantity))
	if err != nil {
		return err
	}

	buttonClear, err2 := p.FindElement(selenium.ByXPATH, ClearButton)
	if err2 != nil {
		return err2
	}
	err2 = buttonClear.Click()
	if err2 != nil {
		return err2
	}

	return nil
}
