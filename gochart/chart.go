package main

import (
    "sync"

    "github.com/fananchong/gochart"
)

type Chart struct {
    gochart.ChartTime
    raknet []int64
    kcp    []int64
    tcp    []int64
    // kcp_fast []int64
    // kcp_fec  []int64
    m sync.Mutex
}

func NewChart() *Chart {
    this := &Chart{raknet: make([]int64, 0), kcp: make([]int64, 0)}
    this.TickUnit = 100
    this.RefreshTime = DEFAULT_REFRESH_TIME
    this.SampleNum = DEFAULT_SAMPLE_NUM
    this.ChartType = "line"
    this.Title = "网络丢包测试"
    this.SubTitle = *showtext1
    this.YAxisText = "收包延时"
    this.YMax = "2000"
    this.ValueSuffix = "ms"
    this.TickLabelStep = "100"
    this.PlotLinesY = "{ color:'red', dashStyle:'longdashdot', value:100, width:1, label:{ text:'100ms', align:'left' } }"
    this.PlotLinesY += ",{ color:'red', dashStyle:'longdashdot', value:200, width:1, label:{ text:'200ms', align:'left' } }"
    this.PlotLinesY += ",{ color:'red', dashStyle:'longdashdot', value:300, width:1, label:{ text:'300ms', align:'left' } }"
    this.PlotLinesY += ",{ color:'red', dashStyle:'longdashdot', value:400, width:1, label:{ text:'400ms', align:'left' } }"
    return this
}

func (this *Chart) Update(now int64) (map[string][]interface{}, string) {
    datas := make(map[string][]interface{})
    details_avg := make(map[string]float32)      // 均值
    details_variance := make(map[string]float32) // 方差
    details_root := make(map[string]float32)     // 标准差
    this.m.Lock()
    datas["raknet"] = make([]interface{}, 0)
    for _, v := range this.raknet {
        datas["raknet"] = append(datas["raknet"], v)
    }
    details_avg["raknet_avg"] = GetAverage(this.raknet)
    details_variance["raknet_∑"] = GetVariance(this.raknet)
    details_root["raknet_σ"] = GetSquareRoot(this.raknet)

    datas["kcp"] = make([]interface{}, 0)
    for _, v := range this.kcp {
        datas["kcp"] = append(datas["kcp"], v)
    }
    details_avg["kcp_avg"] = GetAverage(this.kcp)
    details_variance["kcp_∑"] = GetVariance(this.kcp)
    details_root["kcp_σ"] = GetSquareRoot(this.kcp)

    datas["tcp"] = make([]interface{}, 0)
    for _, v := range this.tcp {
        datas["tcp"] = append(datas["tcp"], v)
    }
    details_avg["tcp_avg"] = GetAverage(this.tcp)
    details_variance["tcp_∑"] = GetVariance(this.tcp)
    details_root["tcp_σ"] = GetSquareRoot(this.tcp)

    subTitleNew := GetSubtitle(details_avg, details_variance, details_root)

    // datas["kcp_fast"] = make([]interface{}, 0)
    // for _, v := range this.kcp_fast {
    //     datas["kcp_fast"] = append(datas["kcp_fast"], v)
    // }
    // datas["kcp_fec"] = make([]interface{}, 0)
    // for _, v := range this.kcp_fec {
    //     datas["kcp_fec"] = append(datas["kcp_fec"], v)
    // }

    this.raknet = this.raknet[:0]
    this.kcp = this.kcp[:0]
    this.tcp = this.tcp[:0]
    // this.kcp_fast = this.kcp_fast[:0]
    // this.kcp_fec = this.kcp_fec[:0]
    this.m.Unlock()
    return datas, subTitleNew
}

func (this *Chart) AddRakNetData(v int64) {
    this.m.Lock()
    this.raknet = append(this.raknet, v)
    this.m.Unlock()
}

func (this *Chart) AddKcpData(v int64) {
    this.m.Lock()
    this.kcp = append(this.kcp, v)
    this.m.Unlock()
}

func (this *Chart) AddTcpData(v int64) {
    this.m.Lock()
    this.tcp = append(this.tcp, v)
    this.m.Unlock()
}

func (this *Chart) AddK1Data(v int64) {
    // this.m.Lock()
    // this.kcp_fast = append(this.kcp_fast, v)
    // this.m.Unlock()
}

func (this *Chart) AddK2Data(v int64) {
    // this.m.Lock()
    // this.kcp_fec = append(this.kcp_fec, v)
    // this.m.Unlock()
}
