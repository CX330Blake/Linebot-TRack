package data

import (
	"fmt"
)

func parseInput() {

}

func fetchAPI(from string, to string, date string) {
	// 使用 fmt.Sprintf 格式化 URL
	url := fmt.Sprintf(
		"https://tdx.transportdata.tw/api/basic/v3/Rail/TRA/DailyTrainTimetable/OD/Inclusive/%d/to/%d/%s?$format=json",
		from,
		to,
		date,
	)
}
