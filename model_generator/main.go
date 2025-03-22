package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"mivar_robot_api/generator"
)

func main() {
	file, err := os.Create("model.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	_, err = file.Write([]byte(xml.Header))
	if err != nil {
		fmt.Println("Error writing XML header:", err)
		return
	}

	gen := generator.NewGenerator()

	modelOutput, err := gen.GenerateModel()
	if err != nil {
		fmt.Println("Error generating model:", err)
		return
	}

	_, err = file.Write(modelOutput)
	if err != nil {
		fmt.Println("Error writing XML:", err)
		return
	}

	fmt.Println("XML generated successfully")
}
