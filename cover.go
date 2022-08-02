package samosa

import (
	"fmt"
	"log"
)

func GetCoverageData(coverageFilePath string) ([]*funcInfo, int, int, error) {
	log.Default().Print("collecting cover profile for file:", coverageFilePath)
	profiles, err := getProfiles(coverageFilePath)
	if err != nil {
		return nil, 0, 0, err
	}
	fmt.Printf("acquired profiles successfully.....")

	fi, covered, total, err := getFunctionInfo(profiles)
	if err != nil {
		return nil, 0, 0, err
	}

	return fi, covered, total, nil
}
