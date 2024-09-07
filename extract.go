package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	var odin weapon
	err := json.Unmarshal(giveJson(), &odin)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the weapon name
	fmt.Printf("%v Skins:\n", odin.Name)

	// Loop over the skins in the weapon struct
	for i := 0; i < len(odin.Skins); {
		// Print skin name
		fmt.Printf("%v\n", odin.Skins[i].Name)
		// If the chromas != 1, loop through each variant
		if len(odin.Skins[i].Chromas) != 1 {
			for o := 0; o != len(odin.Skins[i].Chromas); {
				var variantName string

				// Make sure battle pass skins or 5-Tier skins cant escape the
				// removal of the skin title from the chroma names
				if strings.Contains(odin.Skins[i].Chromas[o].Name, "Level 4") {
					variantName = strings.ReplaceAll(odin.Skins[i].Chromas[o].Name, fmt.Sprintf("%v Level 4\n", odin.Skins[i].Name), "")
				} else if strings.Contains(odin.Skins[i].Chromas[o].Name, "Level 5") {
					variantName = strings.ReplaceAll(odin.Skins[i].Chromas[o].Name, fmt.Sprintf("%v Level 5\n", odin.Skins[i].Name), "")
				} else {
					variantName = strings.ReplaceAll(odin.Skins[i].Chromas[o].Name, fmt.Sprintf("%v\n", odin.Skins[i].Name), "")
				}

				if strings.ToLower(odin.Skins[i].Chromas[o].Name) == strings.ToLower(odin.Skins[i].Name) {
					fmt.Printf("  ╰─(Variant 0 Default): %v\n", odin.Skins[i].Chromas[o].Swatch)
				} else {
					fmt.Printf("  ╰─%v: %v\n", variantName, odin.Skins[i].Chromas[o].Swatch)
				}
				o += 1
			}
		}
		i += 1
	}

}

// Populate the skins array in weapon
func (w *weapon) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Name  string `json:"displayName"`
		Skins []skin `json:"skins"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	w.Name = aux.Name

	w.Skins = make([]skin, len(aux.Skins))
	for i, s := range aux.Skins {
		w.Skins[i] = s
	}

	return nil
}

// Populate the chromas in the skins arrays
func (s *skin) UnmarshalJSON(data []byte) error {
	aux := &struct {
		UUID    string   `json:"uuid"`
		Name    string   `json:"displayName"`
		Icon    string   `json:"displayIcon"`
		Chromas []chroma `json:"chromas"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	s.UUID = aux.UUID
	s.Name = aux.Name
	s.Icon = aux.Icon

	s.Chromas = make([]chroma, len(aux.Chromas))
	for i, c := range aux.Chromas {
		s.Chromas[i] = c
	}

	return nil
}

func getJson() string {
	resp, err := http.Get("https://valorant-api.com/v1/weapons")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)
	return sb
}
