package reporter

import (
	"encoding/json"
	"loganalyzer/internal/analyzer"
	"os"
	"path/filepath"
	"time"
)

func ExportResults(results []analyzer.AnalysisResult, outputPath string) error {
	if outputPath == "" {
		return nil
	}

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	now := time.Now()
	dateStr := now.Format("060102")
	baseName := filepath.Base(outputPath)
	ext := filepath.Ext(outputPath)
	nameWithoutExt := baseName[:len(baseName)-len(ext)]
	
	finalPath := filepath.Join(dir, dateStr+"_"+nameWithoutExt+ext)

	file, err := os.Create(finalPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(results)
}

func PrintResults(results []analyzer.AnalysisResult) {
	for _, result := range results {
		println("=== " + result.LogID + " ===")
		println("Chemin: " + result.FilePath)
		println("Statut: " + result.Status)
		println("Message: " + result.Message)
		if result.ErrorDetails != "" {
			println("Erreur: " + result.ErrorDetails)
		}
		println()
	}
}
