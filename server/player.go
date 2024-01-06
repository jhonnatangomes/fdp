package server

type Player struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Lives int    `json:"lives"`
}
