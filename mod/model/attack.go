package model

import (
	"time"
)

type Attack struct {
	TargetApi API
	RateType  time.Duration
	Rate      uint32
}
