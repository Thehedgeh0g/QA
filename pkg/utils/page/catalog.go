package page

import (
	"fmt"

	"github.com/tebeka/selenium"

	"e2etests/pkg/utils"
)

const (
	ItemNameInCatalog  = "//input[@id='typeahead']"
	ItemNameInItemPage = "//h3[contains(text(), '%s')]"
)

type Catalog struct {
	utils.SeleniumAdapter
}

func (c *Catalog) SearchItem(text string) error {
	if err := c.typeInSearchInputSearchItem(text); err != nil {
		return err
	}

	return c.submitSearchItemWithEnter()
}

func (c *Catalog) FindItem(itemName string) error {
	_, err := c.FindElement(selenium.ByXPATH, fmt.Sprintf(ItemNameInItemPage, itemName))
	return err
}

func (c *Catalog) typeInSearchInputSearchItem(text string) error {
	input, err := c.FindElement(selenium.ByXPATH, ItemNameInCatalog)
	if err != nil {
		return err
	}

	if err := input.Clear(); err != nil {
		return err
	}

	return input.SendKeys(text)
}

func (c *Catalog) submitSearchItemWithEnter() error {
	input, err := c.FindElement(selenium.ByXPATH, ItemNameInCatalog)
	if err != nil {
		return err
	}

	if err := input.SendKeys(selenium.EnterKey); err != nil {
		return err
	}

	return nil
}
