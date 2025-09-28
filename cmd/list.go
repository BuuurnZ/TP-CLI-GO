package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister tous les contacts",
	Long:  "Afficher la liste de tous les contacts du CRM.",
	Run: func(cmd *cobra.Command, args []string) {
		contacts, err := store.GetAll()
		if err != nil {
			fmt.Printf("Erreur lors de la récupération des contacts: %v\n", err)
			return
		}

		if len(contacts) == 0 {
			fmt.Println("Aucun contact trouvé.")
			return
		}

		fmt.Printf("Contacts (%d):\n", len(contacts))
		fmt.Println("ID\tNom\t\tEmail\t\t\tTéléphone\tEntreprise")
		fmt.Println("--\t---\t\t-----\t\t\t---------\t---------")

		for _, contact := range contacts {
			fmt.Printf("%d\t%s\t\t%s\t\t%s\t\t%s\n",
				contact.ID,
				contact.Name,
				contact.Email,
				contact.Phone,
				contact.Company,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
