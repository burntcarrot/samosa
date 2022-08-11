package samosa

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pterm/pterm"
)

type FuncInfoExport struct {
	FileName       string `json:"file"`
	PkgFileName    string `json:"pkg_file"`
	FunctionName   string `json:"function"`
	StartLine      int    `json:"start_line"`
	EndLine        int    `json:"end_line"`
	UncoveredLines int    `json:"uncovered_lines"`
}

func ExportJSON(filename string, fi []funcInfo) (err error) {
	funcInfosJSON := convertJSON(fi)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(funcInfosJSON, "", "	")
	if err != nil {
		return err
	}

	if filename == "" {
		pterm.Println(string(data))
		return nil
	}

	if err = ioutil.WriteFile(filename, data, 0644); err != nil {
		return err
	}

	pterm.Info.Printf("Saved results to %s!\n", filename)
	return nil
}

func convertJSON(funcInfos []funcInfo) []FuncInfoExport {
	var occurences []FuncInfoExport

	for _, fi := range funcInfos {
		functionInfo := FuncInfoExport(fi)

		occurences = append(occurences, functionInfo)
	}

	return occurences
}

func ExportCSV(filename string, fi []funcInfo) error {
	records := convertCSV(fi)

	if filename == "" {
		w := csv.NewWriter(os.Stdout)
		err := w.WriteAll(records)
		if err != nil {
			return err
		}
		return nil
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	w := csv.NewWriter(file)
	err = w.WriteAll(records)
	if err != nil {
		return err
	}

	pterm.Info.Printf("Saved results to %s!\n", filename)
	return nil
}

func convertCSV(funcInfos []funcInfo) [][]string {
	records := [][]string{{"file", "pkg_file", "function", "start_line", "end_line", "uncovered_lines"}}

	for _, fi := range funcInfos {
		functionInfo := []string{fi.FileName, fi.PkgFileName, fi.FunctionName, fmt.Sprint(fi.StartLine), fmt.Sprint(fi.EndLine), fmt.Sprint(fi.UncoveredLines)}

		records = append(records, functionInfo)
	}

	return records
}
