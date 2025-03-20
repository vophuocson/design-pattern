package graphic

import "fmt"

type Graphic interface {
	Draw()
}

type Image struct {
	Path string
}

func (g *Image) Load() {
	fmt.Println("load image from: " + g.Path)
}

func (g *Image) Draw() {
	fmt.Println("Draw image")
}

type ProxyImage struct {
	image *Image
	Path  string
}

func NewProxyImage(path string) *ProxyImage {
	return &ProxyImage{Path: path}
}

func (g *ProxyImage) Draw() {
	if g.image == nil {
		// load image
		g.image = &Image{Path: g.Path}
		g.image.Load()
	}
	g.image.Draw()
}
