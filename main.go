package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Fetch the json from the weapons/skins api
func getJson() []byte {
	resp, err := http.Get("https://valorant-api.com/v1/weapons")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	return body
}

func main() {
	var response jsonData
	err := json.Unmarshal(getJson(), &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the weapon name
	fmt.Printf("%v Skins:\n", response.Data[0].Name)

	// Loop over the skins in the weapon struct
	for i := 0; i < len(response.Data[0].Skins); {
		// Print skin name
		fmt.Printf("%v\n", response.Data[0].Skins[i].Name)
		// If the chromas != 1, loop through each variant
		if len(response.Data[0].Skins[i].Chromas) != 1 {
			for o := 0; o != len(response.Data[0].Skins[i].Chromas); {
				var variantName string

				// Make sure battle pass skins or 5-Tier skins cant escape the
				// removal of the skin title from the chroma names
				if strings.Contains(response.Data[0].Skins[i].Chromas[o].Name, "Level 4") {
					variantName = strings.ReplaceAll(response.Data[0].Skins[i].Chromas[o].Name, fmt.Sprintf("%v Level 4\n", response.Data[0].Skins[i].Name), "")
				} else if strings.Contains(response.Data[0].Skins[i].Chromas[o].Name, "Level 5") {
					variantName = strings.ReplaceAll(response.Data[0].Skins[i].Chromas[o].Name, fmt.Sprintf("%v Level 5\n", response.Data[0].Skins[i].Name), "")
				} else {
					variantName = strings.ReplaceAll(response.Data[0].Skins[i].Chromas[o].Name, fmt.Sprintf("%v\n", response.Data[0].Skins[i].Name), "")
				}

				if strings.ToLower(response.Data[0].Skins[i].Chromas[o].Name) == strings.ToLower(response.Data[0].Skins[i].Name) {
					fmt.Printf("  ╰─(Variant 0 Default): %v\n", response.Data[0].Skins[i].Chromas[o].Swatch)
				} else {
					fmt.Printf("  ╰─%v: %v\n", variantName, response.Data[0].Skins[i].Chromas[o].Swatch)
				}
				o += 1
			}
		}
		i += 1
	}

	var odin weapon = response.Data[0]
	var ares weapon = response.Data[1]

}
