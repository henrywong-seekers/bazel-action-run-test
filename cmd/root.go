package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var (
	output  string
	rootCmd = &cobra.Command{
		Use:   "hello",
		Short: "Say hello",
		Long:  `Say hello to you in a text file`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			hello := fmt.Sprintf("Hello %s!", args[0])
			err := ioutil.WriteFile(output, []byte(hello), 0644)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output file (required)")

	rootCmd.MarkPersistentFlagRequired("output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
