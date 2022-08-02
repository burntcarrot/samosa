package samosa

func GetCoverageData(coverageFilePath string) ([]*funcInfo, int, int, error) {
	profiles, err := getProfiles(coverageFilePath)
	if err != nil {
		return nil, 0, 0, err
	}

	fi, covered, total, err := getFunctionInfo(profiles)
	if err != nil {
		return nil, 0, 0, err
	}

	return fi, covered, total, nil
}
