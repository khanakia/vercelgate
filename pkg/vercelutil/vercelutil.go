package vercelutil

import (
	"app/pkg/jsonupdate"
	"app/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

var (
	authJsonFileName = "auth.json"
)

func SetAuthToken(token string) error {
	filePath, err := AuthJsonFile()
	if err != nil {
		return err
	}

	fileBytes, err := utils.OpenFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	jsonupd := jsonupdate.NewJsonUpdate(string(fileBytes))

	jsonupd.Set("token", token)

	err = os.WriteFile(filePath, []byte(jsonupd.String()), 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func SetCurrentTeam(teamID string) error {
	filePath, err := ConfigJsonFile()
	if err != nil {
		return err
	}

	fileBytes, err := utils.OpenFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	jsonupd := jsonupdate.NewJsonUpdate(string(fileBytes))

	jsonupd.Set("currentTeam", teamID)

	err = os.WriteFile(filePath, []byte(jsonupd.Pretty()), 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func DeleteCurrentTeam() error {
	filePath, err := ConfigJsonFile()
	if err != nil {
		return err
	}

	fileBytes, err := utils.OpenFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	jsonupd := jsonupdate.NewJsonUpdate(string(fileBytes))

	jsonupd.Deleete("currentTeam")

	err = os.WriteFile(filePath, []byte(jsonupd.String()), 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func ParseAuthFile(path string) (*AuthConfig, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var authConfig AuthConfig
	err = json.Unmarshal(fileBytes, &authConfig)
	if err != nil {
		return nil, err
	}
	return &authConfig, nil
}

type AuthConfig struct {
	Token string `json:"token"`
}

func AuthJsonFile() (string, error) {
	globalPath, err := GetGlobalPathConfig()
	if err != nil {
		return "", err
	}
	return filepath.Join(globalPath, authJsonFileName), nil
}

func ConfigJsonFile() (string, error) {
	globalPath, err := GetGlobalPathConfig()
	if err != nil {
		return "", err
	}
	return filepath.Join(globalPath, "config.json"), nil
}

func GetGlobalPathConfig() (string, error) {
	dirname := "com.vercel.cli"

	dirs := append(xdg.ConfigDirs, xdg.ConfigHome)

	for _, datadir := range dirs {

		dirPath := filepath.Join(datadir, dirname)
		// fmt.Println(dirPath)
		info, err := os.Stat(dirPath)
		if err != nil {
			continue
		}

		if info.IsDir() {
			return dirPath, nil
		}
	}

	return "", errors.New("global path config not found")
}
