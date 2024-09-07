package main

// Chromas struct
type chroma struct {
	Swatch string `json:"swatch"`
	Name   string `json:"displayName"`
}

// Skins struct
type skin struct {
	UUID    string   `json:"uuid"`
	Name    string   `json:"displayName"`
	Icon    string   `json:"displayIcon"`
	Chromas []chroma `json:"chromas"`
}

// Weapon struct
type weapon struct {
	Name  string `json:"displayName"`
	Skins []skin `json:"skins"`
}
