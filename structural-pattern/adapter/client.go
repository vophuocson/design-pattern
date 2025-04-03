package adapter

import (
	externalproduct "design-pattern-practice/structural-pattern/adapter/external-product"
	"design-pattern-practice/structural-pattern/adapter/product"
)

func main() {
	externalCircle := &externalproduct.Circle{}
	var externalItem externalproduct.Drawing
	externalItem = externalCircle
	var item product.Item
	item = &product.Circle{Circle: externalItem}
	item.Sketch()
}
