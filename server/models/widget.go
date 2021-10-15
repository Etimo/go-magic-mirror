package models

type Widget struct {
	Id   string `json:"Id"`
	Type string `json:"type"`
}

type TextIcon struct {
	Value string `json:"value"`
	Icon  string `json:"icon"`
}
type TextWidget struct {
	Widget
	TextIcon
}

func (widget *TextWidget) Init(id string, value string) {
	widget.Type = "Text"
	widget.Id = id
	widget.Value = value
}

type ListWidget struct {
	Widget
	Values []TextIcon `json:"values"`
}

func (widget *ListWidget) Init(id string, values []TextIcon) {
	widget.Type = "List"
	widget.Values = values
}
