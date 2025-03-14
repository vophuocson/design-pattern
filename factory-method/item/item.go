package item

type Item interface {
	GetName() string
}
type GrilledChickenBreast struct{}

func (g GrilledChickenBreast) GetName() string { return "Grilled Chicken Breast" }

type SteamedBroccoli struct{}

func (s SteamedBroccoli) GetName() string { return "Steamed Broccoli with Garlic" }

type ScrambledEggs struct{}

func (s ScrambledEggs) GetName() string { return "Scrambled Eggs with Cheese" }

type BrownRice struct{}

func (b BrownRice) GetName() string { return "Brown Rice with Sesame Seeds" }

type GreekYogurt struct{}

func (g GreekYogurt) GetName() string { return "Greek Yogurt with Honey and Nuts" }

type BeefSteak struct{}

func (b BeefSteak) GetName() string { return "Beef Steak with Mashed Potatoes" }

type CreamyMushroomSoup struct{}

func (c CreamyMushroomSoup) GetName() string { return "Creamy Mushroom Soup" }

type CaesarSalad struct{}

func (c CaesarSalad) GetName() string { return "Caesar Salad with Grilled Shrimp" }

type GarlicButterBread struct{}

func (g GarlicButterBread) GetName() string { return "Garlic Butter Bread" }

type ChocolateLavaCake struct{}

func (c ChocolateLavaCake) GetName() string { return "Chocolate Lava Cake" }
