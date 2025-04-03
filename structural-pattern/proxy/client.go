package proxy

import "design-pattern-practice/structural-pattern/proxy/graphic"

func main() {
	var g graphic.Graphic
	g = graphic.NewProxyImage("./image1/png")
	g.Draw()
}
