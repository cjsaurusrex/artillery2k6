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

import "github.com/cjsaurusrex/arillery2k6/cmd"

func main() {
	cmd.Execute()
}
