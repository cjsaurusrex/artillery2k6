/*
Copyright © 2024 Christopher Orchard <EMAIL ADDRESS>
*/
//package main
//
//import (
//	"ArtilleryToK6/cmd"
//	"embed"
//)
//
////go:embed k6-script.tmpl
//var K6ScriptTemplate embed.FS
//
//func main() {
//	cmd.K6ScriptTemplate = K6ScriptTemplate
//	cmd.Execute()
//}

package main

import (
	"github.com/cjsaurusrex/arillery2k6/internal/artillery2k6"
	"os"
	"text/template"
)

func main() {
	//data, _ := os.ReadFile("input.yml")
	script := artillery2k6.Parse("input.yml")
	k6Script := artillery2k6.Build(script)

	tmpl, _ := template.New("k6-script.tmpl").ParseFiles("k6-script.tmpl")

	file, _ := os.Create("output.js")
	defer file.Close()
	tmpl.Execute(file, k6Script)
}
