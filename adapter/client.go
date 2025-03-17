package adapter

import (
	externalproduct "design-pattern-practice/adapter/external-product"
	"design-pattern-practice/adapter/product"
)

func main() {
	externalCircle := &externalproduct.Circle{}
	var externalItem externalproduct.Drawing
	externalItem = externalCircle
	var item product.Item
	item = &product.Circle{Circle: externalItem}
	item.Sketch()
}
