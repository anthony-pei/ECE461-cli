package metrics

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/boyter/scc/v3/processor"
)

type RampUpMetric struct {
}

func (l RampUpMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()

	dir := "temp"
	os.RemoveAll(dir)
	//m.Clone(dir)
	// Analyze code

	total_lines, code_lines, comment_lines, blank_lines := analyze("go.sum")
	log.Printf("Total Lines: %v, Code Lines: %v, Comment Lines: %v, Blank Lines: %v\n", total_lines, code_lines, comment_lines, blank_lines)
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Calculating ramp up metric for module:", m.GetGitHubUrl())
	return 0.0
}

func analyze(filename string) (int64, int64, int64, int64) {
	bts, _ := ioutil.ReadFile(filename)
	processor.ProcessConstants() // Required to load the language information and need only be done once
	language, _ := processor.DetectLanguage(filename)
	if len(language) == 0 {
		language = []string{""}
	}
	filejob := &processor.FileJob{
		Filename: filename,
		Language: language[0],
		Content:  bts,
		Bytes:    int64(len(bts)),
	}

	processor.CountStats(filejob)
	return filejob.Lines, filejob.Code, filejob.Comment, filejob.Blank

}
