package main

import "time"

func main() {
	nowDate := time.Now()
	firstDate := time.Date(nowDate.Year(), nowDate.Month(), 1, 0, 0, 0, 0, time.Local)
	lastDate := firstDate.AddDate(0, 1, -1)
	firstWeekday := firstDate.Weekday()
	lastDay := lastDate.Day()
	for i := 0; i <= lastDay; i++ {
		for j := 0; j <= 7; j++ {
			// TODO 初日記載位置計算から
			if j == 0 {
				if j == int(firstWeekday) {
					
				}
			}
		}
	}
}
