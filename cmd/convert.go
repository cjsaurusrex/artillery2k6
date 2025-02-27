package cmd

import (
	"fmt"
	artillery_parser "github.com/cjsaurusrex/arillery2k6/internal/artillery-parser"
	k6_builder "github.com/cjsaurusrex/arillery2k6/internal/k6-builder"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts Artillery scripts to k6 scripts",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := cmd.Flag("input").Value.String()
		fmt.Println("Trying to convert " + filePath + " to k6")
		data, _ := os.ReadFile(filePath)
		script := artillery_parser.Parse(data)
		k6Script := k6_builder.Build(script)
		tmpl, error := template.New("k6-script.tmpl").ParseFS(K6ScriptTemplate, "k6-script.tmpl")
		if error != nil {
			fmt.Println(error)
		}

		file, _ := os.Create("output.js")
		defer file.Close()
		tmpl.Execute(file, k6Script)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().String("input", "", "Input file")
}
