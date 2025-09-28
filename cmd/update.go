package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mettre à jour un contact",
	Long:  "Mettre à jour les informations d'un contact existant.",
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

		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		company, _ := cmd.Flags().GetString("company")

		if name != "" {
			contact.Name = name
		}
		if email != "" {
			contact.Email = email
		}
		if phone != "" {
			contact.Phone = phone
		}
		if company != "" {
			contact.Company = company
		}

		if err := store.Update(contact); err != nil {
			fmt.Printf("Erreur lors de la mise à jour: %v\n", err)
			return
		}

		fmt.Printf("Contact mis à jour avec succès (ID: %d)\n", contact.ID)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("name", "n", "", "Nouveau nom")
	updateCmd.Flags().StringP("email", "e", "", "Nouvel email")
	updateCmd.Flags().StringP("phone", "p", "", "Nouveau téléphone")
	updateCmd.Flags().StringP("company", "c", "", "Nouvelle entreprise")
}
