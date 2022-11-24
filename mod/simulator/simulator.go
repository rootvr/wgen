package simulator

import (
	"sort"
	"sync"

	attacker "wgen/mod/attacker"
	model "wgen/mod/model"
	printer "wgen/mod/printer"
)

func attack(duration uint32, targets []model.Attack) {
	var waitGroup sync.WaitGroup

	for _, target := range targets {
		waitGroup.Add(1)
		go attacker.Attack(&waitGroup, duration, target)
	}

	waitGroup.Wait()
}

func Simulate(workload map[uint32][]model.Attack, duration uint32) {
	var days []uint32

	for day := range workload {
		days = append(days, day)
	}

	sort.Slice(days, func(i, j int) bool { return days[i] < days[j] })

	for _, day := range days {
		targets := workload[day]

		printer.PrintDayInfo(day, duration)

		for _, target := range targets {
			printer.PrintTargetInfo(target, duration)
		}

		attack(duration, targets)
	}
}
