package component

import "fmt"

type Graph interface {
	Draw()
}

type TextView struct{}

func (g *TextView) Draw() {
	fmt.Println("this is Draw() function of TextView")
}

type BorderView struct{}

func (g *BorderView) Draw() {
	fmt.Println("this is Draw() function of BorderView")
}

type ScrollView struct{}

func (g *ScrollView) Draw() {
	fmt.Println("this is Draw() function of ScrollView")
}

type GraphDecorator struct {
	Graph Graph
	Core  *GraphDecorator
}

func (g *GraphDecorator) Draw() {
	g.Graph.Draw()
	if g.Core != nil {
		g.Core.Draw()
	}
}

func (g *GraphDecorator) Decorate(i Graph) GraphDecorator {
	return GraphDecorator{
		Core:  g,
		Graph: i,
	}
}
