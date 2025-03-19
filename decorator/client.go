package decorator

import "design-pattern-practice/decorator/component"

func Main() {
	var g component.Graph
	g = &component.TextView{}
	decorator := component.GraphDecorator{Graph: g}
	g = &component.BorderView{}
	decorator.Decorate(g)
	g = &component.ScrollView{}
	decorator.Decorate(g)
	decorator.Draw()
}
