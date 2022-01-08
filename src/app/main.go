package main

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
    "io/ioutil"
	"local.data/data"
	"strings"
	"strconv"
    "log"
)

func form (room string, f string) bool {
	if f == "online" && room == "ｵﾝﾗｲﾝ授業" || f == "real" && room != "ｵﾝﾗｲﾝ授業" || len(f) == 0 {
		return true
	} else {
		return false 
	}
}

func season (gakki string, s string) bool {
	if gakki == s || len(s) == 0 {
		return true
	} else {
		return false 
	}
}

func days (niti string, d string) bool {
	if strings.Contains(d, niti) || len(d) == 0 {
		return true
	} else {
		return false 
	}
}

func times (gen int, t string) bool {
	var s string
	s = strconv.Itoa(gen)

	if strings.Contains(t, s) || len(t) == 0 {
		return true
	} else {
		return false 
	}
}


func main() {
	// JSONファイル読み込み
	jsonSyllabus, err := ioutil.ReadFile("syllabus.json")
	jsonTimetable, err := ioutil.ReadFile("timetable.json")
	if err != nil {
        log.Fatal(err)
    }

	// JSONデコード
    var syllabusdata []data.Syllabus
	var timetabledata []data.Timetable

	if err := json.Unmarshal(jsonSyllabus, &syllabusdata); err != nil {
        log.Fatal(err)
    }

	if err := json.Unmarshal(jsonTimetable, &timetabledata); err != nil {
        log.Fatal(err)
    }
	

	// fmt.Println(reflect.TypeOf(syllabusdata))   

	// gin GET
	r := gin.Default()


	// 時間割検索     /search?value=小渡　悟&form=ｵﾝﾗｲﾝ授業&season=後期&days=火
	r.GET("/search", func(c *gin.Context) {

		var query =  make(map[string]string, 5)
		var q[5] string = [5]string {"value", "form", "season", "days", "times"}
		var value []data.Timetable

		for _, q := range q {
			query[q] = c.Query(q)
		}


		for _, v := range timetabledata {
			
			if form(v.Kyoushitsu, query["form"]) && season(v.Gakki, query["season"]) && days(v.Niti, query["days"]) && times(v.Gen, query["times"]) &&
			(strings.Contains(v.Gakubu, query["value"]) || strings.Contains(v.Gakka, query["value"]) ||
			strings.Contains(v.Kamoku, query["value"]) || strings.Contains(v.Tantou, query["value"]) ) {
				
				value = append(value, v)
			} else {
				continue
			}
			
		}

		// フィルターの結果を出力
		c.JSON(200, value)
	})

	// シラバス呼び出し /syllabus?day=金&time=１&season=後期&courseTitle=専門演習基礎&teacher=小渡　悟
	r.GET("/syllabus", func(d *gin.Context) {

		var sQuery =  make(map[string]string, 5)
		var sq[5] string = [5]string {"day", "time", "season", "courseTitle", "teacher"}
		// var upperTime[7] string = [7]string {"１", "２", "３", "４", "５", "６", "７"}

		for _, sq := range sq {
			sQuery[sq] = d.Query(sq)
		}

		//upperTime := width.Narrow.String(sQuery["time"])

		for _, s := range syllabusdata {
			if strings.Contains(s.WeekdayPeriod, sQuery["day"]) && strings.Contains(s.WeekdayPeriod, sQuery["time"]) &&
				s.Semester == sQuery["season"] && s.CourseTitle == sQuery["courseTitle"] && strings.Contains(s.Teacher, sQuery["teacher"]) {

				// シラバスを出力
				d.JSON(200, s)
				break
			}
		}

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

