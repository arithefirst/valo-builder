package main

// Items
// -------------------------------------

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

// Cards struct
type card struct {
	UUID string `json:"uuid"`
	Name string `json:"displayName"`
	Icon string `json:"largeArt"`
}

// Buddies struct
type buddy struct {
	UUID string `json:"uuid"`
	Name string `json:"displayName"`
	Icon string `json:"displayIcon"`
}

// Weapon struct
type weapon struct {
	Name  string `json:"displayName"`
	Skins []skin `json:"skins"`
}

// Responses
// -------------------------------------
type skinResp struct {
	Name string `json:"displayName"`
	Icon string `json:"fullRender"`
	UUID string `json:"uuid"`
}

type errResp struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

// Json outputs
// -------------------------------------

// Json output for skins
type skinJsonData struct {
	Status int      `json:"status"`
	Data   []weapon `json:"data"`
}

// Json output for cards
type cardJsonData struct {
	Status int    `json:"status"`
	Data   []card `json:"data"`
}

// Json output for buddies
type buddyJsonData struct {
	Status int     `json:"status"`
	Data   []buddy `json:"data"`
}
