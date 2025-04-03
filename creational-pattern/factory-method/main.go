package factorymethod

import combocreator "design-pattern-practice/creational-pattern/factory-method/combo-creator"

func main() {
	combo := combocreator.CreateCombo("Protein Power Meal")
	combo.ShowCombo()
}
