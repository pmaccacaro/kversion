package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var options struct {
	help    bool
	version bool
}

func getHelp() {
	color.Green("\n== %s Configuration file ==\n", path.Base(os.Args[0]))
	color.Green(`
- Create a Configuration file in your HOME dir named .kversion.yaml: ex:  [~/.kversion.yaml]
- List all the binaries with full path in the "binaries" list. 
- Full config file example:
`)
	color.Yellow(`
	---
	binaries:
	 - /usr/local/bin/kubectl-1.11.5
	 - /usr/local/bin/kubectl-1.7.14
`)
	fmt.Printf("\n")
	os.Exit(0)
}

func main() {

	// configuration
	viper.AddConfigPath("$HOME/")
	viper.SetConfigName(".kversion")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	binaries := viper.GetStringSlice("binaries")

	flag.BoolVar(&options.help, "help", false, "Show how to configure the app")
	flag.Parse()

	if options.help {
		getHelp()
	}

	// Params parsing
	if len(os.Args) != 2 {
		color.Green("\n== Available Kubectls ==\n")
		color.Green("Usage: %s %s\n\n", os.Args[0], "number")
		for i, v := range binaries {
			color.Yellow("[%d] --> %s\n", i, v)
		}
		fmt.Print("\n")
		os.Exit(1)
	}

	selection := os.Args[1]
	selected, _ := strconv.Atoi(selection)

	if selected <= (len(binaries) - 1) {
		fmt.Printf("Selected: [%d] --> [%s]\n", selected, binaries[selected])
		target := binaries[selected]
		symlink := filepath.Join("/usr/local/bin/", "kubectl")
		fmt.Printf("Creating [%s] --> [%s]\n", target, symlink)
		err := os.Remove(symlink)
		if err != nil {
			fmt.Printf("Error removing [%s]!\n", symlink)
			os.Exit(1)
		}
		err = os.Symlink(target, symlink)
		if err != nil {
			fmt.Println("Error creating link!")
			os.Exit(1)
		}

	} else {
		fmt.Printf("Selection: [%d] not valid!\n", selected)
		os.Exit(1)
	}

}
