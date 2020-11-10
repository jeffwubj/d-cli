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
	"fmt"
	"github.com/jeffwubj/d-vm-operator/api/v1alpha1"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"encoding/json"

	"github.com/spf13/viper"
)

// listCmd represents the base command when called without any subcommands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list VMs",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		contents, err := d.DesktopVMClient.List(metav1.ListOptions{})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var specs = []v1alpha1.DesktopVMSpec{}
		for _, item := range contents.Items {
			specs = append(specs, item.Spec)
		}

		body, err := json.MarshalIndent(specs,"", "    ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(body))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func ExecuteListVMs() {
	if err := listCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	viper.BindPFlags(listCmd.PersistentFlags())
	listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(listCmd)
}
