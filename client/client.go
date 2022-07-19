package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	exec.LookPath("nix")
	cmd := exec.Command("nix", "show-derivation", "nixpkgs#coreutils", "-r")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	drv := ShowDerivation{}
	json.Unmarshal(out.Bytes(), &drv)
	outputPaths := drv.getStorePaths()
	fmt.Printf("%v", outputPaths)
}
