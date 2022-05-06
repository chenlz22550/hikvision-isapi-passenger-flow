package HikvisionPassengerFlow

import (
	"encoding/xml"
	"fmt"
	"github.com/gogf/gf/os/gtime"
	"github.com/loozhengyuan/hikvision-sdk/hikvision"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	SapiCountingPath = "/ISAPI/System/Video/inputs/channels/1/counting/search"
)

type HikvisionCamera struct {
	Username string
	Password string
	Host     string
	Port     string
	*hikvision.Client
}

func (h *HikvisionCamera) CheckOnline() (bool, error) {
	_, err := h.GetDeviceInfo()
	if err != nil {
		return false, err
	}
	return true, nil
}
func (h *HikvisionCamera) GetFootfall(startTime, endTime *gtime.Time) (ff Footfall, err error) {
	start := startTime.Format("Y-m-d\\TH:i:s")
	end := endTime.Format("Y-m-d\\TH:i:s")
	url := "http://" + h.Host + ":" + h.Port + SapiCountingPath
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf(`<CountingStatisticsDescription><reportType>daily</reportType><timeSpanList><timeSpan><startTime>%s</startTime><endTime>%s</endTime></timeSpan></timeSpanList><MinTimeInterval>hour</MinTimeInterval><child>false</child></CountingStatisticsDescription>`, start, end))
	client := h.Client.Client
	client.Timeout = time.Second * 10 //10秒超时
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	var r CountingStatisticsResult
	err = xml.Unmarshal(body, &r)
	if err != nil {
		return
	}
	for _, s := range r.MatchList.MatchElement {
		ff.EnterCount += s.EnterCount
		ff.ExitCount += s.ExitCount
		ff.CurrentCount += s.EnterCount - s.ExitCount
	}
	ff.StartTime = *startTime
	ff.EndTime = *endTime
	return ff, nil
}

func NewHikvisionCamera(host, port, username, password string) (*HikvisionCamera, error) {
	c, err := hikvision.NewClient(
		host+":"+port,
		username,
		password,
	)
	if err != nil {
		return nil, err
	}
	return &HikvisionCamera{Client: c, Username: username, Password: password, Port: port, Host: host}, nil
}
