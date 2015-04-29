package main

import (
	"fmt"
	"log"

	"github.com/samalba/dockerclient"
)

func colorPrint(names string, data string) {
	fmt.Printf("\x1b[31;1m%v\x1b[0m :  %v\n", names, data)
}

func main() {
	// Init the client
	docker, _ := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	containers, err := docker.ListContainers(true, false, "")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-------------------------------------------------")
	for _, c := range containers {
		//fmt.Println("\x1b[31;1mContainer ID\x1b[0m : ", c.Id[:20])
		colorPrint("Container ID", c.Id[:20])
		colorPrint("Names       ", c.Names[0][1:])
		colorPrint("Image       ", c.Image)
		colorPrint("Command     ", c.Command)
		colorPrint("Status      ", c.Status)

		for _, port := range c.Ports {
			fmt.Printf("\x1b[31;1mPort\x1b[0m         :  %v:%v -> %v/%v \n",
				port.IP, port.PrivatePort, port.PublicPort, port.Type)
		}
		fmt.Println("---------------------------------------------------")
	}

}
