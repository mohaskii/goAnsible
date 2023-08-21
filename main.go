package main

import (
	"fmt"
	"goAnsible/pkg/utils"
)

func main() {
	P := utils.InitPlaybook(1000)
	var myFirtConfig = make(map[string]interface{})
	myFirtConfig["name"] = "lolo"
	myFirtConfig["hosts"] = "all"
	myFirtConfig["remote_user"] = "bocal"
	P.Configs = append(P.Configs, myFirtConfig)
	P.HideOutput = true
	err := P.ExecuteWithInventory("all")
	if err != nil {
		fmt.Println(err)
	}
	// P.ConvertToYamlFile("lolo")

}
