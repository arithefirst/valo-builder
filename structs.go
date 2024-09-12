package main

// Chromas struct
type chroma struct {
	Swatch string `json:"swatch"`
	Icon   string `json:"fullRender"`
	Name   string `json:"displayName"`
}

// Skins struct
type skin struct {
	UUID    string   `json:"uuid"`
	Name    string   `json:"displayName"`
	Icon    string   `json:"fullRender"`
	Chromas []chroma `json:"chromas"`
}

// Weapon struct
type weapon struct {
	Name  string `json:"displayName"`
	Skins []skin `json:"skins"`
}

type jsonData struct {
	Status int      `json:"status"`
	Data   []weapon `json:"data"`
}

type skinResp struct {
	Name string `json:"displayName"`
	Icon string `json:"fullRender"`
	UUID string `json:"uuid"`
}
