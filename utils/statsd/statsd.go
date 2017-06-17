package statsd

import (
	"github.com/etsy/statsd/examples/go"
	"sweetcook-backend/config"
	"sweetcook-backend/utils"
	"sweetcook-backend/utils/logger"
	"time"
)

var client *statsd.StatsdClient

func init() {
	client = statsd.New("10.2.86.22", 8125)
	configJson := config.GetConfigJson()
	statsdInfo := configJson["statsd"]
	host,err := utils.GetStringFromInterfaceMap(statsdInfo, "host")
	if err == nil {
		portFloat,err := utils.GetFloat64FromInterfaceMap(statsdInfo, "port")
		port := int(portFloat)
		if err == nil {
			client = statsd.New(host, port)
			return
		}
		
	}
	logger.Info("statsd not init.")
}

func checkInitOk() bool {
	if client == nil {
		logger.Error("statsd config not found. request ignored.")
		return false
	}
	return true
}

func Timing(metric string, startTime time.Time, endTime time.Time)  {
	if !checkInitOk() {
		return
	}
	// Submit timing information
	duration := int64(endTime.Sub(startTime) / time.Millisecond)
	client.Timing(metric, duration)
}

func IncreMetric(metric string)  {
	if !checkInitOk() {
		return
	}
	// Increment a stat counter
	client.Increment(metric)
}

func DecreMetric(metric string)  {
	if !checkInitOk() {
		return
	}
	// Decrement a stat counter
	client.Decrement(metric)
}