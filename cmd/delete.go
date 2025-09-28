package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprimer un contact",
	Long:  "Supprimer un contact du CRM.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Erreur: ID du contact requis")
			return
		}

		id, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			fmt.Println("Erreur: ID invalide")
			return
		}

		contact, err := store.GetByID(uint(id))
		if err != nil {
			fmt.Printf("Contact non trouvé: %v\n", err)
			return
		}

		if err := store.Delete(uint(id)); err != nil {
			fmt.Printf("Erreur lors de la suppression: %v\n", err)
			return
		}

		fmt.Printf("Contact supprimé avec succès: %s (%s)\n", contact.Name, contact.Email)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
