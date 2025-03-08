package cmd

import (
	"fmt"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts Artillery scripts to k6 scripts",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		fmt.Println("Trying to convert " + filePath + " to k6")

		artilleryScript := artillery2k6.Parse(filePath)
		bc := helpers.NewBuilderConfig()
		if artilleryScript.Config.Environments != nil {
			bc.EnvironmentsInUse = true
			bc.RootVariableFormat = helpers.GlobalThis
		}

		k6Script := artillery2k6.BuildScript(bc, artilleryScript)
		k6ScriptContext := artillery2k6.BuildContext(k6Script, *bc)

		tmpl, error := template.New("k6-script.tmpl").ParseFS(K6ScriptTemplate, "k6-script.tmpl")
		if error != nil {
			fmt.Println(error)
		}

		file, _ := os.Create(cmd.Flag("output").Value.String())
		defer file.Close()
		tmpl.Execute(file, k6ScriptContext)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringP("output", "o", "output.js", "Output file")
}
