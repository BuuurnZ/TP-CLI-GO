package analyzer

import (
	"loganalyzer/internal/config"
	"math/rand"
	"os"
	"time"
)

type AnalysisResult struct {
	LogID       string `json:"log_id"`
	FilePath    string `json:"file_path"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

func AnalyzeLog(logConfig config.LogConfig) AnalysisResult {
	result := AnalysisResult{
		LogID:    logConfig.ID,
		FilePath: logConfig.Path,
	}

	if _, err := os.Stat(logConfig.Path); os.IsNotExist(err) {
		result.Status = "FAILED"
		result.Message = "Fichier introuvable."
		result.ErrorDetails = err.Error()
		return result
	}

	file, err := os.Open(logConfig.Path)
	if err != nil {
		result.Status = "FAILED"
		result.Message = "Fichier non accessible."
		result.ErrorDetails = err.Error()
		return result
	}
	defer file.Close()

	time.Sleep(time.Duration(rand.Intn(151)+50) * time.Millisecond)

	result.Status = "OK"
	result.Message = "Analyse terminée avec succès."
	result.ErrorDetails = ""

	return result
}

func AnalyzeLogsConcurrently(configs []config.LogConfig) []AnalysisResult {
	results := make([]AnalysisResult, len(configs))
	resultsChan := make(chan AnalysisResult, len(configs))

	for _, logConfig := range configs {
		go func(cfg config.LogConfig) {
			result := AnalyzeLog(cfg)
			resultsChan <- result
		}(logConfig)
	}

	for i := 0; i < len(configs); i++ {
		results[i] = <-resultsChan
	}

	return results
}
