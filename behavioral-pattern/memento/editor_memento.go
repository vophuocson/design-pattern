package memento

type EditorMemento struct {
	text string
}

func (m *EditorMemento) getSavedState() string {
	return m.text
}
