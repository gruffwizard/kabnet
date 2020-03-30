package cmd

import (
	"fmt"
	"github.com/gruffwizard/kabnet/openshift"
	"github.com/spf13/cobra"
	"strings"
)

//FindAllStringSubmatch
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "info",
	Long:  "List information about available OpenShift versions and tools",
}

var infoDepsCmd = &cobra.Command{
	Use:   "deps",
	Short: "deps",
	Long:  "List information about Openshift Dependencies",

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println(strings.Join(openshift.OpenShiftMajorVersions(), ", "))
		} else {
			subversions := openshift.OpenShiftMinorVersion(args[0])
			if len(subversions) > 0 {
				fmt.Println(strings.Join(subversions, ", "))
			} else {
				files := openshift.Dependencies(args[0])
				if files == nil {
					fmt.Println("no information for version ", args[0])
				} else {
					fmt.Println("base url :", files.BaseURL)
					fmt.Println("Initramfs:", files.Initramfs)
					fmt.Println("Kernel   :", files.Kernel)
					fmt.Println("Metal    :", files.Metal)
				}
			}

		}
	},
}

var infoClientsCmd = &cobra.Command{
	Use:   "clients",
	Short: "clients",
	Long:  "List information about available OpenShift client tools",

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println(strings.Join(openshift.ClientVersions(), ", "))
		} else {
			files := openshift.Clients(args[0])
			fmt.Println("base url :", files.BaseURL)
			fmt.Println("Installer:", files.Installer)
			fmt.Println("Client   :", files.Client)

		}

	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.AddCommand(infoClientsCmd)
	infoCmd.AddCommand(infoDepsCmd)

}
