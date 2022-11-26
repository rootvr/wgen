package printer

import (
	"time"

	logger "wgen/mod/logger"
	model "wgen/mod/model"
)

func PrintDayInfo(day uint32, duration uint32) {
	logger.InfoMessage("day=%d,length=%d,unit=s\n", day, duration)
}

func PrintSimulationInfo(apispec string, workload string) {
	logger.InfoMessage("apispec=%s,workload=%s\n", apispec, workload)
}

func PrintTargetInfo(target model.Attack, dayDuration uint32) {
	if target.RateType == time.Second {
		logger.InfoMessage("url=%s,rate=%d,unit=s,requests=%d\n", target.TargetApi.BaseUrl+target.TargetApi.RelativeUrl, target.Rate, target.Rate*dayDuration)
	} else if target.RateType == time.Minute {
		logger.InfoMessage("url=%s,rate=%d,unit=m,requests=%d\n", target.TargetApi.BaseUrl+target.TargetApi.RelativeUrl, target.Rate, target.Rate*(dayDuration/60))
	} else if target.RateType == time.Hour {
		logger.InfoMessage("url=%s,rate=%d,unit=h,requests=%d\n", target.TargetApi.BaseUrl+target.TargetApi.RelativeUrl, target.Rate, target.Rate*(dayDuration/(60*60)))
	}

}

func PrintAttackPendingInfo(target model.Attack) {
	if target.RateType == time.Second {
		logger.PendingMessage("attack=%s,rate=%d,unit=s\n", target.TargetApi.BaseUrl+target.TargetApi.RelativeUrl, target.Rate)
	} else if target.RateType == time.Minute {
		logger.PendingMessage("attack=%s,rate=%d,unit=m\n", target.TargetApi.BaseUrl+target.TargetApi.RelativeUrl, target.Rate)
	} else if target.RateType == time.Hour {
		logger.PendingMessage("attack=%s,rate=%d,unit=h\n", target.TargetApi.BaseUrl+target.TargetApi.RelativeUrl, target.Rate)
	}
}

func PrintAttackErrorInfo(target model.Attack) {
	logger.ErrorMessage("skip=%s,why=attack time rate > day length\n", target.TargetApi.BaseUrl+target.TargetApi.RelativeUrl)
}

func PrintAttackCompletedInfo(target model.Attack, duration time.Duration, success float64) {
	logger.CompletedMessage("finish=%s,time=%f,unit=s,success=%.2f%%\n", target.TargetApi.BaseUrl+target.TargetApi.RelativeUrl, duration.Seconds(), success*100)
}
