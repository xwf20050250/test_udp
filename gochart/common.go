package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetAverage(datas []int64) float32 {
	var sum int64 = 0
	var index int = 0
	for _, v := range datas {
		sum += v
		index++
	}
	avg := float32(sum) / float32(len(datas))
	tmp := fmt.Sprintf("%.2f", avg)
	_ret, _ := strconv.ParseFloat(tmp, 64)
	return float32(_ret)
}

func GetVariance(datas []int64) float32 {
	avg := GetAverage(datas)
	var sum float64 = 0
	for _, v := range datas {
		sum += math.Pow(float64(float32(v)-avg), float64(2))
	}
	tmp := fmt.Sprintf("%.2f", float32(sum)/float32(len(datas)))
	_ret, _ := strconv.ParseFloat(tmp, 64)
	return float32(_ret)
}

func GetSquareRoot(datas []int64) float32 {
	variance := GetVariance(datas)
	sroot := math.Sqrt(float64(variance))
	tmp := fmt.Sprintf("%.2f", sroot)
	_ret, _ := strconv.ParseFloat(tmp, 64)
	return float32(_ret)
}

func GetSubtitle(details_avg map[string]float32, details_variance map[string]float32,
	details_root map[string]float32) string {
	ret_avg := "【平均值】"
	contents_avg := []string{}
	for k, v := range details_avg {
		tmp := k + ": " + strconv.FormatFloat(float64(v), 'f', -1, 32)
		contents_avg = append(contents_avg, tmp)
	}
	ret_avg += strings.Join(contents_avg, "  ")

	ret_variance := "【方差】"
	contents_variance := []string{}
	for k, v := range details_variance {
		tmp := k + ": " + strconv.FormatFloat(float64(v), 'f', -1, 32)
		contents_variance = append(contents_variance, tmp)
	}
	ret_variance += strings.Join(contents_variance, "  ")

	ret_root := "【标准差】"
	contents_root := []string{}
	for k, v := range details_root {
		tmp := k + ": " + strconv.FormatFloat(float64(v), 'f', -1, 32)
		contents_root = append(contents_root, tmp)
	}
	ret_root += strings.Join(contents_root, "  ")

	ret := ret_avg + "<br>" + ret_variance + "<br>" + ret_root
	return ret
}
