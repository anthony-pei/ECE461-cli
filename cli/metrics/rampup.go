package metrics

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/boyter/scc/v3/processor"
)

type RampUpMetric struct {
}

var analyzeDirFunction = analyzeDir

func (l RampUpMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()
	log.Println("Calculating ramp up metric for module:", m.GetGitHubUrl())
	dir := "temp"
	cleanDir(dir, false)
	m.Clone(dir)
	processor.ProcessConstants()                    // Required to load the language information and need only be done once
	_, code, comments, _ := analyzeDirFunction(dir) // Returns total lines, lines of code, comments, and blank lines. Not using first and last at the moment
	log.Printf("Code Lines: %v, Comment Lines: %v", code, comments)
	cleanDir(dir, true)

	if code == 0 {
		return 0.0
	}
	score := float64(comments) / float64(code) / 0.2
	score = math.Min(score, 1)
	score = math.Max(score, 0)
	return score
}

func cleanDir(dir string, failOnError bool) {
	err := os.RemoveAll(dir)

	if failOnError && (err != nil) {
		log.Fatal(err)
	}
}

func analyzeDir(dir string) (int64, int64, int64, int64) {
	dir_lines, dir_code, dir_comments, dir_blank := int64(0), int64(0), int64(0), int64(0)
	items, _ := ioutil.ReadDir(dir)
	for _, item := range items {
		new_path := filepath.Join(dir, item.Name())
		total_lines, code_lines, comment_lines, blank_lines := int64(0), int64(0), int64(0), int64(0)
		if item.IsDir() {
			total_lines, code_lines, comment_lines, blank_lines = analyzeDir(new_path)
		} else {
			total_lines, code_lines, comment_lines, blank_lines = analyze(new_path)
			//log.Printf("%v: Total Lines: %v, Code Lines: %v, Comment Lines: %v, Blank Lines: %v\n", item.Name(), total_lines, code_lines, comment_lines, blank_lines)
		}
		dir_lines += total_lines
		dir_code += code_lines
		dir_comments += comment_lines
		dir_blank += blank_lines
	}

	return dir_lines, dir_code, dir_comments, dir_blank
}

func analyze(filename string) (int64, int64, int64, int64) {
	bts, _ := ioutil.ReadFile(filename)
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
