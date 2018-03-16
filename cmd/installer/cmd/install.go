// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/helm/cmd/helm/installer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"fmt"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install tiller service in cluster",
	Long: `Install tiller service in kubernetes cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace:= viper.GetString("namespace")
		image:= viper.GetString("image")
		serviceaccount := viper.GetString("serviceaccount")
		// creates the in-cluster config
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		// creates the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			fmt.Printf("Failed to create kubernetes client")
		}
		opts:= &installer.Options{Namespace:namespace,ImageSpec:image,ServiceAccount:serviceaccount}
		err =installer.Install(clientset,opts)
		if err != nil {
			fmt.Printf("Failed to install tiller service %v",err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
