package samosa

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// get root for the repo using the mod file location
func getRoot() ([]byte, error) {
	command := exec.Command("go", "env", "-json")
	data, err := command.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func decodeJSON(data []byte, buf map[string]interface{}) error {
	if err := json.Unmarshal(data, &buf); err != nil {
		return fmt.Errorf("can not decode json from input stream:%v ended with err:%v", string(data), err)
	}

	return nil
}

func getModDir(dataset map[string]interface{}) string {
	modDir := dataset["GOMOD"]
	v, ok := modDir.(string)
	if ok {
		return v
	}
	return ""
}

func getMod() (string, error) {
	dataSet := make(map[string]interface{})

	byteEnvData, err := getRoot()
	if err != nil {
		return "", err
	}

	err = decodeJSON(byteEnvData, dataSet)
	if err != nil {
		return "", err
	}

	return getModDir(dataSet), nil
}
