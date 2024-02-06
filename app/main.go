package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
)

type ArgKeys struct {
	abort    string
	reboot   string
	poweroff string
	dryRun   string
}

var argKeys ArgKeys = ArgKeys{
	abort:    "-c",
	reboot:   "-r",
	poweroff: "-P",
	dryRun:   "-k",
}

var appName = "app"    // set by build.cmd
var appVer = "unknown" // set by build.cmd

var dryRun bool
var bind string

func main() {
	fmt.Printf("%s %s\n", appName, appVer)

	flag.StringVar(&bind, "bind", ":851", "Listen address")
	flag.BoolVar(&dryRun, "dry-run", false, "Dry run")

	flag.Parse()

	if runtime.GOOS == "windows" {
		argKeys = ArgKeys{
			abort:    "-a",
			reboot:   "-r",
			poweroff: "-s",
			dryRun:   "",
		}
	}

	fmt.Printf("Listen on %s\n", bind)

	if dryRun {
		if runtime.GOOS == "windows" {
			fmt.Println("Dry run not supported on Windows")
			return
		}
		fmt.Println("Works in dry mode")
	}

	serve()
}

func execute(args ...string) error {
	if dryRun {
		args = append(args, argKeys.dryRun)
	}
	cmd := exec.Command("shutdown", args...)
	fmt.Println(args)
	return cmd.Run()
}
