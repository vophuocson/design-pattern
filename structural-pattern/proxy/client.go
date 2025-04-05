package proxy

import "design-pattern-practice/structural-pattern/proxy/graphic"

func main() {
	var g graphic.Graphic
	// Initialize and hold a placeholder for the image
	g = graphic.NewProxyImage("./image1/png")
	// Suppose the image is only drawn when the user clicks to load it.
	g.Draw()
}
