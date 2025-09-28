package cmd

import (
	"fmt"
	"loganalyzer/internal/config"
	"loganalyzer/internal/stores"
	"os"

	"github.com/spf13/cobra"
)

var (
	store stores.Storer
	cfg   *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "minicrm",
	Short: "Mini-CRM CLI - Gestionnaire de contacts",
	Long:  "Un gestionnaire de contacts simple et efficace en ligne de commande.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		cfg, err = config.LoadConfig()
		if err != nil {
			fmt.Printf("Erreur lors du chargement de la configuration: %v\n", err)
			os.Exit(1)
		}

		store, err = config.NewStore(cfg)
		if err != nil {
			fmt.Printf("Erreur lors de l'initialisation du store: %v\n", err)
			os.Exit(1)
		}
	},
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