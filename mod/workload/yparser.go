package workload

import (
	"fmt"
	"os"
	"time"

	api "wgen/mod/api"
	model "wgen/mod/model"
	utils "wgen/mod/utils"

	yaml "gopkg.in/yaml.v3"
)

type YAttackData struct {
	Rate uint   `yaml:"rate"`
	Unit string `yaml:"unit"`
}

type YDay struct {
	Attacks map[string]YAttackData `yaml:"day"`
}

type YWorkload struct {
	Workload []YDay `yaml:"workload"`
}

func getUnitType(unit string) time.Duration {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}()
	if unit == "s" {
		return time.Second
	}
	if unit == "m" {
		return time.Minute
	}
	if unit == "h" {
		return time.Hour
	}
	panic("Invalid rate type! `" + unit + "` (must be `s`, `m` or `h`)")
}

func ParseYamlWorkloadFile(filename string, APISpec api.YAPISpec) map[uint32][]model.Attack {
	file, err := os.ReadFile(filename)
	if err != nil {
		utils.Kill(err)
	}

	var yWorkload YWorkload
	err = yaml.Unmarshal(file, &yWorkload)
	if err != nil {
		utils.Kill(err)
	}

	workload := make(map[uint32][]model.Attack)

	for day, day_data := range yWorkload.Workload {
		for api_name, api_data := range day_data.Attacks {
			APIReq := model.API{}
			APIReq.BaseUrl = APISpec.BaseUrl
			APIReq.RelativeUrl = APISpec.APIs[api_name].RelativeUrl
			APIReq.Method = APISpec.APIs[api_name].Method
			SimReq := model.Attack{TargetApi: APIReq}
			SimReq.Rate = uint32(api_data.Rate)
			SimReq.RateType = getUnitType(api_data.Unit)

			workload[uint32(day)] = append(workload[uint32(day)], SimReq)
		}
	}

	return workload
}
