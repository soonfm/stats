//
// SFM 2.0 Stats Service
//

package main

import (
	"github.com/Sirupsen/logrus"

	"github.com/soonfm/stats/cmd"
)

// Version Number - Set at build time
var VERSION = "unset"

func main() {
	if err := cmd.Execute(VERSION); err != nil {
		logrus.Fatalf("Fatal Error: %v", err)
	}
}
