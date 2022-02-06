package models

type InputText struct {
	Text string `json:"text"`
}

type OutputFrequency struct {
	Word string `json:"word"`
	Freq int    `json:"freq"`
}

type OutputJSON struct {
	Data    []OutputFrequency `json:"data"`
	Error   bool              `json:"error"`
	Message string            `json:"message"`
	Status  int               `json:"status"`
}
