package cmd

import (
	"fmt"
	"loganalyzer/internal/models"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un nouveau contact",
	Long:  "Ajouter un nouveau contact au CRM.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		company, _ := cmd.Flags().GetString("company")

		if name == "" || email == "" {
			fmt.Println("Erreur: Le nom et l'email sont obligatoires")
			return
		}

		contact := &models.Contact{
			Name:    name,
			Email:   email,
			Phone:   phone,
			Company: company,
		}

		if err := store.Create(contact); err != nil {
			fmt.Printf("Erreur lors de la création du contact: %v\n", err)
			return
		}

		fmt.Printf("Contact créé avec succès (ID: %d)\n", contact.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("name", "n", "", "Nom du contact (obligatoire)")
	addCmd.Flags().StringP("email", "e", "", "Email du contact (obligatoire)")
	addCmd.Flags().StringP("phone", "p", "", "Téléphone du contact")
	addCmd.Flags().StringP("company", "c", "", "Entreprise du contact")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
}
