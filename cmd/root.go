package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var basePath string

var rootCmd = &cobra.Command{
	Use:   "alea",
	Short: "Alea's cli productivity improver",
	Long: `Alea is a useful cli to help you improve productivity.
Do you want to improve it? Of course! Download the repo 
and start modifying!
			`,
}

func Execute() {

	viper.AddConfigPath(".")
	viper.AddConfigPath("~/.alea")
	viper.SetConfigName("config") 
	viper.ReadInConfig()

	defaultEnv := viper.GetString("default_env")

	viper.SetDefault("env", defaultEnv)
	basePath = viper.GetString("projects_folder")

	rootCmd.PersistentFlags().StringP("env", "e", "dev", "Environment to run the command against. By default 'dev'")
	viper.BindPFlag("env", rootCmd.PersistentFlags().Lookup("env"))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

