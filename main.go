package main

import (
	"fmt"

	"github.com/sshaunn/train-go/ms1/avro"
	"github.com/sshaunn/train-go/ms1/routes"
)

func main() {
	fmt.Println("schema regi like what??: ", avro.SchemaRegistry())
	routes.Routers()
}
