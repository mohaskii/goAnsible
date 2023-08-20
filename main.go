package main

import (
	"fmt"
	"goAnsible/pkg/utils"
)

func main() {
	theP := utils.Playbook{}
	myPlayBook := utils.ConFig{}
	myPlayBook.Name = "lolo"
	myPlayBook.Hosts = "all"
	myPlayBook.RemoteUser = "bocal"
	theP.TheConfigs = append(theP.TheConfigs, myPlayBook)
	err := theP.ExecuteWithInventory("all")
	if err != nil {
		fmt.Println(err, "kdkdkdkkdkk")
	}
}
