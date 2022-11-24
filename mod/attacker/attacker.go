package attacker

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	model "wgen/mod/model"
	printer "wgen/mod/printer"
	utils "wgen/mod/utils"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func methodPOSTTargeter(targetRequest model.Attack) vegeta.Targeter {
	return func(target *vegeta.Target) error {
		if target == nil {
			return vegeta.ErrNilTarget
		}

		target.Method = targetRequest.TargetApi.Method
		target.URL = targetRequest.TargetApi.BaseUrl + targetRequest.TargetApi.RelativeUrl

		if json.Valid([]byte(targetRequest.TargetApi.JSONParam)) {
			payload, err := json.Marshal(targetRequest.TargetApi.JSONParam)
			if err != nil {
				return err
			}
			target.Body = payload

			header := http.Header{}
			header.Add("Accept", "application/json")
			header.Add("Content-Type", "application/json")
			target.Header = header
		}

		return nil
	}
}

func exec(duration time.Duration, target model.Attack) {
	printer.PrintAttackPendingInfo(target)

	vegetaRate := vegeta.Rate{Freq: int(target.Rate), Per: target.RateType}
	var vegetaTargeter vegeta.Targeter

	if target.TargetApi.Method == "POST" {
		vegetaTargeter = methodPOSTTargeter(target)
	} else if target.TargetApi.Method == "GET" || target.TargetApi.Method == "DELETE" {
		vegetaTargeter = vegeta.NewStaticTargeter(vegeta.Target{
			Method: target.TargetApi.Method,
			URL:    target.TargetApi.BaseUrl + target.TargetApi.RelativeUrl,
		})
	} else {
		utils.Kill(errors.New("Illegal HTTP method"))
	}

	vegetaAttacker := vegeta.NewAttacker()

	start := time.Now()

	var metrics vegeta.Metrics
	for res := range vegetaAttacker.Attack(vegetaTargeter, vegetaRate, duration, "test_"+target.TargetApi.BaseUrl+target.TargetApi.RelativeUrl) {
		metrics.Add(res)
	}
	metrics.Close()

	printer.PrintAttackCompletedInfo(target, time.Since(start), metrics.Success)
}

func Attack(waitGroup *sync.WaitGroup, duration uint32, target model.Attack) {
	time, err := time.ParseDuration(fmt.Sprintf("%ds", duration))
	utils.Kill(err)

	if (1 * target.RateType) > time {
		printer.PrintAttackErrorInfo(target)
	} else {
		exec(time, target)
	}

	waitGroup.Done()
}
