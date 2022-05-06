package HikvisionPassengerFlow

import (
	"encoding/xml"
	"github.com/gogf/gf/os/gtime"
)

type CountingStatisticsResult struct {
	XMLName            xml.Name `xml:"CountingStatisticsResult"`
	Version            string   `xml:"version,attr"`
	Xmlns              string   `xml:"xmlns,attr"`
	ResponseStatus     string   `xml:"responseStatus"`
	ResponseStatusStrg string   `xml:"responseStatusStrg"`
	MatchList          struct {
		MatchElement []struct {
			TimeSpan struct {
				StartTime string `xml:"startTime"`
				EndTime   string `xml:"endTime"`
			} `xml:"timeSpan"`
			EnterCount int `xml:"enterCount"`
			ExitCount  int `xml:"exitCount"`
		} `xml:"matchElement"`
	} `xml:"matchList"`
}

type Footfall struct {
	EnterCount   int        `json:"enterCount"`
	ExitCount    int        `json:"exitCount"`
	CurrentCount int        `json:"currentCount"`
	StartTime    gtime.Time `json:"startTime"`
	EndTime      gtime.Time `json:"endTime"`
}
