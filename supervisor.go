package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	api "wgen/mod/api"
	simulator "wgen/mod/simulator"
	workload "wgen/mod/workload"
)

func requiredFlags(flagName string) bool {
	res := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagName {
			res = true
		}
	})
	return res
}

func parseArgs() (string, string, uint) {
	w := flag.String("w", "", "workload/file/path")
	a := flag.String("a", "", "apispec/file/path")
	d := flag.Uint("d", 0, "simulated day length (in seconds)")

	flag.Parse()

	if !requiredFlags("a") {
		fmt.Println(errors.New("CALL ERROR: -w flag w/ args is required"))
		flag.PrintDefaults()
		os.Exit(0)
	}
	if !requiredFlags("w") {
		fmt.Println(errors.New("CALL ERROR: -w flag w/ args is required"))
		flag.PrintDefaults()
		os.Exit(0)
	}
	if !requiredFlags("d") {
		fmt.Println(errors.New("CALL ERROR: -w flag w/ args is required"))
		flag.PrintDefaults()
		os.Exit(0)
	}

	return *w, *a, *d
}

func main() {
	workload_file, apispec_file, sim_day_length := parseArgs()

	api_spec := api.ParseYamlAPISpecFile(apispec_file)
	workload := workload.ParseYamlWorkloadFile(workload_file, api_spec)

	simulator.Simulate(workload, uint32(sim_day_length))
}
