package models

type Widget struct {
	Id     string `json:"Id"`
	Type   string `json:"type"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type TextIcon struct {
	Value string `json:"value"`
	Icon  string `json:"icon"`
}
type TextWidget struct {
	Widget
	TextIcon
}

func (widget *TextWidget) Init(id string, x int, y int, width int, height int, value string) {
	widget.Type = "Text"
	widget.Id = id
	widget.X = x
	widget.Y = y
	widget.Width = width
	widget.Height = height
	widget.Value = value
}

type ListWidget struct {
	Widget
	Values []TextIcon `json:"values"`
}

func (widget *ListWidget) Init(id string, x int, y int, width int, height int, values []TextIcon) {
	widget.Type = "List"
	widget.Id = id
	widget.X = x
	widget.Y = y
	widget.Width = width
	widget.Height = height
	widget.Values = values
}
