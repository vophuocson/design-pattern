package singleton

import "design-pattern-practice/singleton/product"

func main() {
	dbOne := product.GetConnection()
	dbOne.Show()
}
