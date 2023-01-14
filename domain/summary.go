package domain

type Summary struct {
	Greeting    string  `json:"greeting" validate:"required"`
	Temperature float32 `json:"temperature" validate:"required"`
	Heads_up    float32 `json:"heads_up" validate:"required"`
}
