package factorymethod

import (
	combocreator "design-pattern-practice/factory-method/combo-creator"
)

func main() {
	combo := combocreator.CreateCombo("Protein Power Meal")
	combo.ShowCombo()
}
