package item

import "fmt"

type Item interface {
	ShowItem()
}

type Image struct{}

func (i *Image) ShowItem() {
	fmt.Print("this is Image item")
}

type PDF struct{}

func (i *PDF) ShowItem() {
	fmt.Print("this is PDF item")
}

type Doc struct{}

func (i *Doc) ShowItem() {
	fmt.Print("this is Doc item")
}

type Folder struct {
	items []Item
}

func (i *Folder) ShowItem() {
	fmt.Print("this is Folder item")
	for _, i := range i.items {
		i.ShowItem()
	}
}

func (i *Folder) Add(item Item) {
	i.items = append(i.items, item)
}
