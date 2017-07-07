package main

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

type CompareData struct {
	InvokeNum int    // 请求
	ReqNum    int    //填充
	ViewNum   int    //曝光
	ClickNum  int    //点击
	Timepoint string //获取数据的时间点
}

type resultCompareData struct {
	InvokeDiff    int
	InvokePercent float64
	ReqDiff       int
	ReqPercent    float64
	ViewDiff      int
	ViewPercent   float64
	ClickDiff     int
	ClickPercent  float64
}
type reportData struct {
	DataPoint     string
	DataPointDesc string
	BeforeData    CompareData
	CurrentData   CompareData
	Result        resultCompareData
	Threshold     string
	Detail        string
}

var sendMessage = `<h2 style="text-align:center">{{.DataPointDesc}}</h2>
<h4 style="text-align:center">{{.Detail}}</h4>
<table border="1" cellpadding="10" width="100%">
    <tr>
	    <th>监控点\时间</th>
	    <th>{{.BeforeData.Timepoint}}</th>
	    <th>{{.CurrentData.Timepoint}}</th>
	    <th>差值</th>
	    <th>百分比</th>
	    <th>阈值</th>
    </tr>
    <tr>
	    <th>请求(invokeNum)</th>
	    <th>{{.BeforeData.InvokeNum}}</th>
	    <th>{{.CurrentData.InvokeNum}}</th>
	    <th>{{.Result.InvokeDiff}}</th></th>
	    <th>{{.Result.InvokePercent}}</th></th>
		<th>{{.Threshold}}</th>
    </tr>
    <tr>
	    <th>填充(reqNum)</th>
	    <th>{{.BeforeData.ReqNum}}</th>
	    <th>{{.CurrentData.ReqNum}}</th>
	    <th>{{.Result.ReqDiff}}</th></th>
	    <th>{{.Result.ReqPercent}}</th></th>
		<th>{{.Threshold}}</th>
    </tr>
	<tr>
		<th>曝光(view)</th>
	    <th>{{.BeforeData.ViewNum}}</th>
	    <th>{{.CurrentData.ViewNum}}</th>
	    <th>{{.Result.ViewDiff}}</th>
	    <th>{{.Result.ViewPercent}}</th>
		<th>{{.Threshold}}</th>
    </tr>
    <tr>
	    <th>点击(click)</th>
	    <th>{{.BeforeData.ClickNum}}</th>
	    <th>{{.CurrentData.ClickNum}}</th>
	    <th>{{.Result.ClickDiff}}</th>
	    <th>{{.Result.ClickPercent}}</th>
		<th>{{.Threshold}}</th>
    </tr>
 </table>`

const DATETIMEFORMAT = "2006-01-02 15:04:05"

func main() {

	var recover []string
	recover = append(recover, "312024054@qq.com")

	var data reportData
	data.DataPoint = "1_2_3_4"
	data.DataPointDesc = "某个应用的某个广告"
	data.BeforeData.ClickNum = 100
	data.BeforeData.InvokeNum = 200
	data.BeforeData.ReqNum = 300
	data.BeforeData.ViewNum = 400
	data.BeforeData.Timepoint = time.Now().Format(DATETIMEFORMAT)
	data.CurrentData.ClickNum = 101
	data.CurrentData.InvokeNum = 201
	data.CurrentData.ReqNum = 301
	data.CurrentData.ViewNum = 401
	data.CurrentData.Timepoint = time.Now().Format(DATETIMEFORMAT)
	data.Threshold = fmt.Sprintf("%.2f", 0.1)

	var mailFormat = template.Must(template.New("mailFormat").Parse(sendMessage))
	//var body bytes.Buffer
	body := bytes.NewBufferString("")

	calcResult(&data)

	if err := mailFormat.Execute(body, data); err != nil {
		fmt.Println(err)
	}

	fmt.Println(body)
}

func calcResult(data *reportData) {
	data.Result.InvokeDiff, data.Result.InvokePercent = calcDiffAndPercent(data.BeforeData.InvokeNum, data.CurrentData.InvokeNum)
	data.Result.ReqDiff, data.Result.ReqPercent = calcDiffAndPercent(data.BeforeData.ReqNum, data.CurrentData.ReqNum)
	data.Result.ViewDiff, data.Result.ViewPercent = calcDiffAndPercent(data.BeforeData.ViewNum, data.CurrentData.ViewNum)
	data.Result.ClickDiff, data.Result.ClickPercent = calcDiffAndPercent(data.BeforeData.ClickNum, data.CurrentData.ClickNum)
	data.Detail = "告警详情"
}

func calcDiffAndPercent(beforeData, currentData int) (int, float64) {
	differ := beforeData - currentData
	if differ < 0 {
		differ = differ * -1
	}

	var percent float64
	if beforeData != 0 && currentData != 0 {
		percent = float64(differ) / float64(beforeData)
	}

	return differ, percent
}
