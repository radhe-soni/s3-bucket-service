package main

import (
	"fmt"

	"github.com/radhe-soni/ec2-service/config"
	"github.com/radhe-soni/ec2-service/service"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("EC2-WL: EC2 white listing api")
	config.InitConfig()
	for {
		if userInterface() {
			break
		}
	}

	fmt.Printf("Exiting app")
}

var services []string = []string{"describe", "update-ip", "exit"}

func userInterface() bool {
	ec2Service := viper.Get("ec2Service").(*service.EC2Service)
	fmt.Println("Enter you choice ...")
	for i, ser := range services {
		fmt.Printf("%d for %s \n", i+1, ser)
	}
	var choice int
	fmt.Scanln(&choice)
	exitApp := false
	switch choice {
	case 1:
		ec2Service.VerifySecurityGroups()
		break
	case 2:
		updateIPInterface(ec2Service.UpdateIPWith)
		break
	case 3:
		exitApp = true
		break
	default:
		fmt.Println("Enter a valid choice.")
	}
	return exitApp
}

func updateIPInterface(updateIPWith func(a, b string)) {
	newIP, err := service.FindMyPublicIP()
	if err != nil {
		fmt.Println("Enter New Ip to add -> ")
		fmt.Scanln(&newIP)
	}
	fmt.Println("Enter Rule description to update -> ")
	var ruleDescription string
	fmt.Scanln(&ruleDescription)
	fmt.Println()
	updateIPWith(newIP, ruleDescription)
}
