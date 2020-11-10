/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package vm

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"

	"github.com/spf13/viper"
)

var VMName = ""

// getCmd represents the base command when called without any subcommands
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get [Name]",
	Long: ``,
	Args: validateListArgs,
	Run: func(cmd *cobra.Command, args []string) {
		desktopVM, err := d.DesktopVMClient.Get(VMName,metav1.GetOptions{})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		body, err := json.MarshalIndent(desktopVM,"", "    ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(body))
	},
}

func validateListArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return fmt.Errorf("VM Name must be provided")
	}

	VMName = args[0]
	if VMName == "" {
		fmt.Errorf("VM name should not be empty")
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func ExecuteGetVM() {
	if err := getCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	viper.BindPFlags(getCmd.PersistentFlags())
	getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(getCmd)
}
