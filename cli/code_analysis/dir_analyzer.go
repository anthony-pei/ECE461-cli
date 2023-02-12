package code_analysis

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/boyter/scc/processor"
	log "github.com/sirupsen/logrus"
)

func AnalyzeCodeBase(dir string) (int64, int64, int64, int64) {
	processor.ProcessConstants() // Required to load the language information and need only be done once
	return analyzeDir(dir)

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

func CleanDir(dir string, failOnError bool) {
	err := os.RemoveAll(dir)

	if failOnError && (err != nil) {
		log.Debug("Error removing directory.")
	}
}
