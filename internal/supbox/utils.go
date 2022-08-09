package supbox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ReadRekordboxConfig(path string) RekordboxConfig {
	// read file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	// json data
	var obj RekordboxConfig

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetEncryptedPassword(asarPath string) string {
	f, err := os.Open(asarPath)
	if err != nil {
		panic("Cannot open asar file: " + asarPath)
	}

	// Search for the password
	data, err := ioutil.ReadAll(f)
	re := regexp.MustCompile(`pass: \".*\"`)
	result := re.FindAllString(string(data), 10)[0]

	password := strings.Split(result, ": ")[1]

	// Remove the quotation marks
	password = strings.Replace(password, `"`, "", -1)

	return password

}

func GetEncryptedPasswordDataFromConfig(config RekordboxConfig) string {
	if len(config.Options) < 2 {
		panic("Unable to read Rekordbox Config file to get password.")
	}

	if len(config.Options[1]) < 2 {
		panic("Unable to read Rekordbox Config file to get password.")
	}

	return config.Options[1][1]
}

func GetDatabaseFilePath(root string) string {
	return filepath.Join(root, "master.db")
}

func GetImagePath(dataPath string, imagePath string) string {
	return filepath.Join(dataPath, "share", imagePath)
}
