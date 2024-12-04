package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Make folder named after date
	mmddDate := time.Now().String()[5:10]
	basePath := "../" + mmddDate

	err := os.MkdirAll(basePath, 0777)
	if err != nil {
		log.Fatal(err)
	}

	for _, suffix := range []string{"-p1", "-p2"} {
		path := basePath + "/" + mmddDate + suffix
		err := os.MkdirAll(path, 0777)
		if err != nil {
			log.Fatal(err)
		}

		files := map[string]string{
			"/inputs.txt": "",
			"/prompt.txt": "",
			"/main.go": `package main

func main() {}`,
		}

		for filename, content := range files {
			file, err := os.Create(path + filename)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(file, content)
			file.Close()
		}

		// Initialize Go module
		cmd := exec.Command("go", "mod", "init", mmddDate)
		cmd.Dir = path
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

	}
}
