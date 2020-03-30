package cmd

import (
	"fmt"
	"github.com/gruffwizard/kabnet/util"
	"os"
	"strconv"
)

func err(msg string) {
	fmt.Println(msg)
	os.Exit(-1)
}

func validateAsNumber(v string, msg string) int {

	if v == "" {
		return 0
	}
	r, error := strconv.Atoi(v)
	if error == nil {
		return r
	}
	err(msg)
	return 0

}

func validateAsExistingDirectory(d string, msg string) {

	if d == "" {
		err(msg)
	}
	if !util.FileExists(d) {
		err(msg)
	}

}
