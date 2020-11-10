package vm

import (
	"fmt"
	"github.com/jeffwubj/d-cli/pkg/driver"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var d* driver.Driver

func init() {
	d = driver.InitDriver("", "default")
}


var rootCmd = &cobra.Command{
	Use:   "kubectl vm",
	Short: "",
	Long: ``,
}

func Execute()  {
	rootCmd.Execute()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".vm-rest" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".vm-rest")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}