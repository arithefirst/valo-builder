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

func printSkins(response jsonData, index int) {
	// Print the weapon name
	fmt.Printf("%v Skins:\n", response.Data[index].Name)

	// Loop over the skins in the weapon struct
	for i := 0; i < len(response.Data[index].Skins); {
		// Print skin name
		fmt.Printf("%v\n", response.Data[index].Skins[i].Name)
		// If the chromas != 1, loop through each variant
		if len(response.Data[index].Skins[i].Chromas) != 1 {
			for o := 0; o != len(response.Data[index].Skins[i].Chromas); {
				var variantName string

				// Make sure battle pass skins or 5-Tier skins cant escape the
				// removal of the skin title from the chroma names
				if strings.Contains(response.Data[index].Skins[i].Chromas[o].Name, "Level 4") {
					variantName = strings.ReplaceAll(response.Data[index].Skins[i].Chromas[o].Name, fmt.Sprintf("%v Level 4\n", response.Data[index].Skins[i].Name), "")
				} else if strings.Contains(response.Data[index].Skins[i].Chromas[o].Name, "Level 5") {
					variantName = strings.ReplaceAll(response.Data[index].Skins[i].Chromas[o].Name, fmt.Sprintf("%v Level 5\n", response.Data[index].Skins[i].Name), "")
				} else {
					variantName = strings.ReplaceAll(response.Data[index].Skins[i].Chromas[o].Name, fmt.Sprintf("%v\n", response.Data[index].Skins[i].Name), "")
				}

				if strings.ToLower(response.Data[index].Skins[i].Chromas[o].Name) == strings.ToLower(response.Data[index].Skins[i].Name) {
					fmt.Printf("  ╰─(Variant 0 Default): %v\n", response.Data[index].Skins[i].Chromas[o].Swatch)
				} else {
					fmt.Printf("  ╰─%v: %v\n", variantName, response.Data[index].Skins[i].Chromas[o].Swatch)
				}
				o += 1
			}
		}
		i += 1
	}
}

func main() {
	var response jsonData
	bytes, err := readCache()

	// If the file is too old, request new data from the API
	if err != nil {
		if err.Error() == "cached file unusable: older than 48h" {
			fmt.Println("Cache file unusable. Requesting new file from https://valorant-api.com/v1/weapons")
			bytes = getJson()
			writeCache(bytes)
		} else {
			log.Fatalln(err)
		}
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for i := 0; i < 19; {
		printSkins(response, i)
		i += 1
	}

}
