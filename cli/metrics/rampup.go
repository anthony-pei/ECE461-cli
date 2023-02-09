package metrics

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/boyter/scc/v3/processor"
)

type RampUpMetric struct {
}

func (l RampUpMetric) CalculateScore(m Module) float64 {
	// Object l of type license matrix and m of type module with function get_url()

	dir := "temp"
	os.RemoveAll(dir)

	m.Clone(dir)
	processor.ProcessConstants() // Required to load the language information and need only be done once
	_, code, comments, _ := analyzeDir(dir)

	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Calculating ramp up metric for module:", m.GetGitHubUrl())
	fmt.Println(float64(comments) / float64(code) / 0.2)
	return float64(comments) / float64(code) / 0.2
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
