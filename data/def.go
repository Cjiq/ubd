package data

type Definition struct {
	Word    string `json:"word"`
	Text    string `json:"definition"`
	Example string `json:"example"`
	Rating  int    `json:"thumbs_up"`
}

type Result struct {
	Definitions []Definition `json:"list"`
}
