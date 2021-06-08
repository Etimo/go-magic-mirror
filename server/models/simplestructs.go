package models

type Pomodoro struct {
	Simple string `json:"simple"`
}

type WelcomeMessage struct {
	Message string `json:"message"`
}

type LayoutMessage struct {
	Cols int `json:"cols"`
	Rows int `json:"rows"`
}
