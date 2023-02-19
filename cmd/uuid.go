package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

func init() {
	rootCmd.AddCommand(uuidCmd)
}

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a random UUID and copy it to clipboard (if possible)",
	Long:  `Generate a random UUID and copy it to clipboard (if possible)`,
	Run: func(cmd *cobra.Command, args []string) {

		uuid, _ := uuid.NewRandom()

		if err := clipboard.Init(); err == nil {
			clipboard.Write(clipboard.FmtText, []byte(uuid.String()))
		}
		
		fmt.Println(uuid)
	},
}


