package testdata

const (
	QuantityItemsOne = "1"
	QuantityItemsTen = "10"
)

type ItemTestData struct {
	ID    string
	Name  string
	Price string
	URL   string
}

var ItemCasio = ItemTestData{
	ID:    "1",
	Name:  "Casio MRP-700-1AVEF",
	Price: "300",
	URL:   "/product/casio-mrp-700-1avef",
}
