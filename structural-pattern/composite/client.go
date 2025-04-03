package composite

import "design-pattern-practice/structural-pattern/composite/item"

func main() {
	pdf := item.PDF{}
	doc := item.Doc{}
	childFolder := item.Folder{}
	childFolder.Add(&pdf)
	childFolder.Add(&doc)
	childFolder.Add(&pdf)
	image := item.Image{}
	parentFolder := item.Folder{}
	parentFolder.Add(&image)
	parentFolder.Add(&childFolder)
	parentFolder.ShowItem()
}
