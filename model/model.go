package model

type Movie struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}


