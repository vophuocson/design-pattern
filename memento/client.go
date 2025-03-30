package memento

func main() {
	editor := Editor{}
	editor.SetText("1")
	editorMemento := editor.CreateMemento()
	stack := MementoStack{}
	stack.Push(editorMemento)
	editor.SetText("2")
	editorMemento = editor.CreateMemento()
	stack.Push(editorMemento)
	editor.SetText("3")
	editorMemento = editor.CreateMemento()
	stack.Push(editorMemento)
	editor.SetText("4")
	editorMemento = editor.CreateMemento()
	stack.Push(editorMemento)
	editor.RestoreState(stack.Pop())
	editor.RestoreState(stack.Pop())
}
