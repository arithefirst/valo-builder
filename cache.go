package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func writeCache(data []byte) {
	filename := fmt.Sprintf("cache/%v-skins.json", time.Now().Unix())
	err := os.WriteFile(filename, data, 0777)
	if err != nil {
		// If folder is non-existent, re-create the folder, and re-try the file write
		if strings.Contains(err.Error(), "no such file or directory") {
			fmt.Println("\"cache\" folder not found. Creating empty dir called \"cache\"")
			cmd := exec.Command("mkdir", "cache")
			err = cmd.Start()
			if err != nil {
				log.Fatal(err)
				return
			}

			// Wait for mkdir to finish before re-trying
			err := cmd.Wait()
			if err != nil {
				log.Fatal(err)
				return
			}

			writeCache(data)
		} else {
			log.Fatal(err)
		}
	}
}

func readCache() ([]byte, error) {
	// Read the "cache" dir
	files, err := os.ReadDir("cache")

	// Prevent the program from killing itself if the cache dir is empty
	if len(files) == 0 {
		return nil, errors.New("open cache: no such file or directory")
	}

	if err != nil {
		return []byte{}, err
	}

	// Separate the file's timestamp from its name
	timestamp, err := strconv.Atoi(strings.Split(files[0].Name(), "-")[0])
	if err != nil {
		return []byte{}, err
	}

	// Dispose of the cache if it has been more than 48H
	if (timestamp + 172800) < int(time.Now().Unix()) {
		cmd := exec.Command("rm", "-rf", fmt.Sprintf("cache/%v-skins.json", timestamp))
		fserr := cmd.Start()
		if fserr != nil {
			return []byte{}, err
		}

		return []byte{}, errors.New("cached file unusable: older than 48h")
	}

	// Read the file as a byte array and then return it for the json parser
	byteArray, err := os.ReadFile(fmt.Sprintf("cache/%v-skins.json", timestamp))
	if err != nil {
		return []byte{}, err
	}

	return byteArray, nil

}
