package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Make folder named after date
	mmddDate := time.Now().String()[5:10]
	//mmddDate := currentTime[5:10]
	path := "../" + mmddDate

	err := os.MkdirAll(path, 0777)
	if err != nil {
		log.Fatal(err)
	}

	inputsFile, err := os.Create(path + "/inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputsFile.Close()

	promptFile, err := os.Create(path + "/prompt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer promptFile.Close()

	mainGoFile, err := os.Create(path + "/main.go")
	if err != nil {
		log.Fatal(err)
	}
	mainGoFile.WriteString(
		`package main

func main() {}`)
	defer mainGoFile.Close()

	cmd := exec.Command("go", "mod", "init", mmddDate)
	cmd.Dir = "../" + mmddDate
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
