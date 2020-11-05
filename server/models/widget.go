package models

type Widget struct {
	Id     string `json:"Id"`
	Type   string `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
type TextWidget struct {
	Widget
	Value string `json:"value"`
	Icon  string `json:"icon"`
}

func (widget *TextWidget) Init(id string, width int, height int, value string) {
	widget.Type = "Text"
	widget.Id = id
	widget.Width = width
	widget.Height = height
	widget.Value = value
}

type ListWidget struct {
	Widget
	Values   []string `json:"values"`
	ListIcon string   `json:"listIcon"`
}

func (widget *ListWidget) Init(id string, width int, height int, values []string) {
	widget.Type = "List"
	widget.Id = id
	widget.Width = width
	widget.Height = height
	widget.Values = values
}
