package prototype

import "design-pattern-practice/creational-pattern/prototype/product"

func main() {
	var pen product.Pen
	pencil := product.Pencil{}
	pen = &pencil
	p := pen.Clone()
	p.Draw()
}
