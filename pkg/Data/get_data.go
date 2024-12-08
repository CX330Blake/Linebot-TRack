/*
API:
https://tdx.transportdata.tw/api/basic/v3/Rail/TRA/DailyTrainTimetable/OD/3360/to/4220/2024-12-08?$format=json
*/

package data

import (
	"fmt"
	"net/http"
	"regexp"
)

type TrainData struct {
	TrainNo string `json:"TrainNo"`
	From    string `json:""`
}

func GetData(query string) {
	date, from, to, err := parseInput(query)
	if err == nil {
		fetchAPI(from, to, date)
	}
}

func parseInput(query string) (date string, from string, to string, err error) {
	regex := regexp.MustCompile(`(\d{4}-\d{2}-\d{2})\s+(\S+?)到(\S+)`)

	// 匹配輸入字串
	matches := regex.FindStringSubmatch(query)
	if len(matches) != 4 {
		return "", "", "", fmt.Errorf("Error: 無法解析輸入 %s", query)
	}

	// 提取日期和地點
	return matches[1], matches[2], matches[3], nil
}

func fetchAPI(from string, to string, date string) {
	// 使用 fmt.Sprintf 格式化 URL
	url := fmt.Sprintf(
		"https://tdx.transportdata.tw/api/basic/v3/Rail/TRA/DailyTrainTimetable/OD/%d/to/%d/%s?$format=json",
		from,
		to,
		date,
	)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP 請求失敗: %w", err)
	}
}

func getStationID(station string) (id string) {
	trainStations := map[string]string{
		"基隆":   "0900",
		"三坑":   "0910",
		"八堵":   "0920",
		"七堵":   "0930",
		"百福":   "0940",
		"五堵":   "0950",
		"汐止":   "0960",
		"汐科":   "0970",
		"南港":   "0980",
		"松山":   "0990",
		"臺北":   "1000",
		"萬華":   "1010",
		"板橋":   "1020",
		"浮洲":   "1030",
		"樹林":   "1040",
		"南樹林":  "1050",
		"山佳":   "1060",
		"鶯歌":   "1070",
		"桃園":   "1080",
		"內壢":   "1090",
		"中壢":   "1100",
		"埔心":   "1110",
		"楊梅":   "1120",
		"富岡":   "1130",
		"新富":   "1140",
		"北湖":   "1150",
		"湖口":   "1160",
		"新豐":   "1170",
		"竹北":   "1180",
		"北新竹":  "1190",
		"新竹":   "1210",
		"三姓橋":  "1220",
		"香山":   "1230",
		"崎頂":   "1240",
		"竹南":   "1250",
		"造橋":   "3140",
		"豐富":   "3150",
		"苗栗":   "3160",
		"南勢":   "3170",
		"銅鑼":   "3180",
		"三義":   "3190",
		"泰安":   "3210",
		"后里":   "3220",
		"豐原":   "3230",
		"栗林":   "3240",
		"潭子":   "3250",
		"頭家厝":  "3260",
		"松竹":   "3270",
		"太原":   "3280",
		"精武":   "3290",
		"臺中":   "3300",
		"五權":   "3310",
		"大慶":   "3320",
		"烏日":   "3330",
		"新烏日":  "3340",
		"成功":   "3350",
		"彰化":   "3360",
		"花壇":   "3370",
		"大村":   "3380",
		"員林":   "3390",
		"永靖":   "3400",
		"社頭":   "3410",
		"田中":   "3420",
		"二水":   "3430",
		"林內":   "3450",
		"石榴":   "3460",
		"斗六":   "3470",
		"斗南":   "3480",
		"石龜":   "3490",
		"大林":   "4050",
		"民雄":   "4060",
		"嘉北":   "4070",
		"嘉義":   "4080",
		"水上":   "4090",
		"南靖":   "4100",
		"後壁":   "4110",
		"新營":   "4120",
		"柳營":   "4130",
		"林鳳營":  "4140",
		"隆田":   "4150",
		"拔林":   "4160",
		"善化":   "4170",
		"南科":   "4180",
		"新市":   "4190",
		"永康":   "4200",
		"大橋":   "4210",
		"臺南":   "4220",
		"林森":   "4230",
		"南臺南":  "4240",
		"保安":   "4250",
		"仁德":   "4260",
		"中洲":   "4270",
		"大湖":   "4290",
		"路竹":   "4300",
		"岡山":   "4310",
		"橋頭":   "4320",
		"楠梓":   "4330",
		"新左營":  "4340",
		"左營":   "4350",
		"內惟":   "4360",
		"美術館":  "4370",
		"鼓山":   "4380",
		"三塊厝":  "4390",
		"高雄":   "4400",
		"民族":   "4410",
		"科工館":  "4420",
		"正義":   "4430",
		"鳳山":   "4440",
		"後庄":   "4450",
		"九曲堂":  "4460",
		"六塊厝":  "4470",
		"屏東":   "5000",
		"談文":   "2110",
		"大山":   "2120",
		"後龍":   "2130",
		"龍港":   "2140",
		"白沙屯":  "2150",
		"新埔":   "2160",
		"通霄":   "2170",
		"苑裡":   "2180",
		"日南":   "2190",
		"大甲":   "2200",
		"臺中港":  "2210",
		"清水":   "2220",
		"沙鹿":   "2230",
		"龍井":   "2240",
		"大肚":   "2250",
		"追分":   "2260",
		"暖暖":   "7390",
		"四腳亭":  "7380",
		"瑞芳":   "7360",
		"猴硐":   "7350",
		"三貂嶺":  "7330",
		"牡丹":   "7320",
		"雙溪":   "7310",
		"貢寮":   "7300",
		"福隆":   "7290",
		"石城":   "7280",
		"大里":   "7270",
		"大溪":   "7260",
		"龜山":   "7250",
		"外澳":   "7240",
		"頭城":   "7230",
		"頂埔":   "7220",
		"礁溪":   "7210",
		"四城":   "7200",
		"宜蘭":   "7190",
		"二結":   "7180",
		"中里":   "7170",
		"羅東":   "7160",
		"冬山":   "7150",
		"新馬":   "7140",
		"蘇澳新站": "7130",
		"蘇澳":   "7120",
		"永樂":   "7110",
		"東澳":   "7100",
		"南澳":   "7090",
		"武塔":   "7080",
		"漢本":   "7070",
		"和平":   "7060",
		"和仁":   "7050",
		"崇德":   "7040",
		"新城":   "7030",
		"景美":   "7020",
		"北埔":   "7010",
		"花蓮":   "7000",
		"吉安":   "6250",
		"志學":   "6240",
		"平和":   "6230",
		"壽豐":   "6220",
		"豐田":   "6210",
		"林榮新光": "6200",
		"南平":   "6190",
		"鳳林":   "6180",
		"萬榮":   "6170",
		"光復":   "6160",
		"大富":   "6150",
		"富源":   "6140",
		"瑞穗":   "6130",
		"三民":   "6120",
		"玉里":   "6110",
		"東里":   "6100",
		"東竹":   "6090",
		"富里":   "6080",
		"池上":   "6070",
		"海端":   "6060",
		"關山":   "6050",
		"瑞和":   "6040",
		"瑞源":   "6030",
		"鹿野":   "6020",
		"山里":   "6010",
		"臺東":   "6000",
		// 屏東-南迴線
		"歸來": "5010", "麟洛": "5020", "西勢": "5030",
		"竹田": "5040", "潮州": "5050", "崁頂": "5060", "南州": "5070",
		"鎮安": "5080", "林邊": "5090", "佳冬": "5100", "東海": "5110",
		"枋寮": "5120", "加祿": "5130", "內獅": "5140", "枋山": "5160",
		"大武": "5190", "瀧溪": "5200", "金崙": "5210", "太麻里": "5220",
		"知本": "5230", "康樂": "5240",

		// 平溪線
		"大華": "7331", "十分": "7332", "望古": "7333",
		"嶺腳": "7334", "平溪": "7335", "菁桐": "7336",

		// 內灣線
		"千甲": "1191", "新莊": "1192",
		"竹中": "1193", "上員": "1201", "榮華": "1202", "竹東": "1203",
		"橫山": "1204", "九讚頭": "1205", "合興": "1206", "富貴": "1207",
		"內灣": "1208",

		// 集集線
		"源泉": "3431", "濁水": "3432", "龍泉": "3433",
		"集集": "3434", "水里": "3435", "車埕": "3436",

		// 沙崙線
		"長榮大學": "4271", "沙崙": "4272",

		// 六家線
		"六家": "1194",

		// 深澳線
		"海科館": "7361", "八斗子": "7362",
	}
	return trainStations[station]
}