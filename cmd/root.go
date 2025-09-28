package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Outil d'analyse de logs distribuée",
	Long:  "Un outil CLI pour analyser des fichiers de logs en parallèle et générer des rapports JSON.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
