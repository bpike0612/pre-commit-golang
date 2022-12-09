package cmd

import (
	"log"
	"os/exec"
)

func isInstalled(app string, pkg string) {
	_, err := exec.LookPath(app)
	if err != nil {
		log.Printf("%s is not installed, installing... %s\n", app, pkg)
		err = exec.Command("go", "install", pkg).Run()
		if err != nil {
			log.Printf("failed to install %s via go install %s\n", app, pkg)
			log.Fatal(err)
		}
	}
}
