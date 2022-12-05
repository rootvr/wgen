package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	api "wgen/mod/api"
	printer "wgen/mod/printer"
	simulator "wgen/mod/simulator"
	utils "wgen/mod/utils"
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

func parseDayLength(dayLengthStr string) uint32 {
	unit := string(dayLengthStr[len(dayLengthStr)-1])
	time := string(dayLengthStr[:len(dayLengthStr)-1])

	if unit == "s" {
		res, err := strconv.ParseUint(time, 10, 0)
		utils.Kill(err)
		return uint32(res)
	} else if unit == "m" {
		res, err := strconv.ParseUint(time, 10, 0)
		utils.Kill(err)
		return uint32(res * 60)
	} else if unit == "h" {
		res, err := strconv.ParseUint(time, 10, 0)
		utils.Kill(err)
		return uint32(res * 3600)
	}

	panic("Invalid day length unit type! `" + unit + "` (must be `s`, `m` or `h`)")
}

func parseArgs() (string, string, uint32) {
	w := flag.String("w", "", "workload/file/path")
	a := flag.String("a", "", "apispec/file/path")
	d := flag.String("d", "", "simulated day length (examples: 30s, 10m, 1h)")
	var pd uint32

	flag.Parse()

	if !requiredFlags("a") {
		fmt.Println(errors.New("CALL ERROR: -a flag w/ args is required"))
		flag.PrintDefaults()
		os.Exit(0)
	}
	if !requiredFlags("w") {
		fmt.Println(errors.New("CALL ERROR: -w flag w/ args is required"))
		flag.PrintDefaults()
		os.Exit(0)
	}
	if !requiredFlags("d") {
		fmt.Println(errors.New("CALL ERROR: -d flag w/ args is required"))
		flag.PrintDefaults()
		os.Exit(0)
	}

	r, _ := regexp.Compile("[1-9]+[0-9]*[smh]{1}$")

	if r.MatchString(*d) {
		pd = parseDayLength(*d)
	} else {
		fmt.Println("CALL ERROR: -d flag w/ args must be `-d <UNSIGNED_INT>[smh]`, given", *d)
		flag.PrintDefaults()
		os.Exit(0)
	}

	return *w, *a, pd
}

func main() {
	workload_file, apispec_file, sim_day_length := parseArgs()

	api_spec := api.ParseYamlAPISpecFile(apispec_file)
	workload := workload.ParseYamlWorkloadFile(workload_file, api_spec)

	printer.PrintSimulationInfo(apispec_file, workload_file)
	simulator.Simulate(workload, uint32(sim_day_length))
}
