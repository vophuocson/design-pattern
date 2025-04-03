package singleton

import "design-pattern-practice/creational-pattern/singleton/product"

func main() {
	dbOne := product.GetConnection()
	dbOne.Show()
}
