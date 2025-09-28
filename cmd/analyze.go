package cmd

import (
	"fmt"
	"loganalyzer/internal/analyzer"
	"loganalyzer/internal/config"
	"loganalyzer/internal/reporter"

	"github.com/spf13/cobra"
)

var configFile string
var outputFile string

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyser les logs selon la configuration",
	Long:  "Lit un fichier de configuration JSON et analyse les logs en parallèle.",
	Run: func(cmd *cobra.Command, args []string) {
		configs, err := config.LoadConfig(configFile)
		if err != nil {
			fmt.Printf("Erreur lors du chargement de la config: %v\n", err)
			return
		}

		fmt.Printf("Analyse de %d logs en cours...\n", len(configs))
		
		results := analyzer.AnalyzeLogsConcurrently(configs)
		
		reporter.PrintResults(results)
		
		if outputFile != "" {
			err = reporter.ExportResults(results, outputFile)
			if err != nil {
				fmt.Printf("Erreur lors de l'export: %v\n", err)
			} else {
				fmt.Printf("Résultats exportés vers: %s\n", outputFile)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringVarP(&configFile, "config", "c", "", "Fichier de configuration JSON (requis)")
	analyzeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Fichier de sortie JSON")
	analyzeCmd.MarkFlagRequired("config")
}
