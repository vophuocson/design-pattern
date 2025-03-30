package memento

type Editor struct {
	text string
}

func (t *Editor) GetText() string {
	return t.text
}

func (t *Editor) SetText(s string) {
	t.text = s
}

func (t *Editor) CreateMemento() *EditorMemento {
	return &EditorMemento{text: t.text}
}

func (t *Editor) RestoreState(m *EditorMemento) {
	t.text = m.text
}
