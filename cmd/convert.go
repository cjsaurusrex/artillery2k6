package cmd

import (
	"fmt"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
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
		SetupBuilderConfig(bc, artilleryScript)

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

func SetupBuilderConfig(bc *helpers.BuilderConfig, script models.ArtilleryScript) {
	if script.Config.Environments != nil {
		bc.EnvironmentsInUse = true
		bc.RootVariableFormat = helpers.GlobalThis
	}

	bc.PayloadsInUse = len(script.Config.Payloads.Payloads) > 0
	if !bc.PayloadsInUse && bc.EnvironmentsInUse {
		envsUsingPayloads := false
		for _, env := range script.Config.Environments {
			if len(env.Payload.Payloads) > 0 {
				envsUsingPayloads = true
				break
			}
		}
		bc.PayloadsInUse = envsUsingPayloads
	}
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringP("output", "o", "output.js", "Output file")
}
