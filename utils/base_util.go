package utils

import (
	"crypto/sha256"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"io"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Rand(n int) int {
	return rand.Intn(n)
}

//
func IntSlice2Str(a []int) (b string) {
	for k, v := range a {
		if k+1 == len(a) {
			b += strconv.Itoa(v)
		} else {
			b += strconv.Itoa(v) + ","
		}
	}

	return
}

func Str2IntSlice(a string) []int {
	b := make([]int, 0)
	if a == "" {
		return b
	}

	c := strings.Split(a, ",")
	for _, v := range c {
		d, _ := strconv.Atoi(v)
		b = append(b, d)
	}
	return b
}

// 打乱[]int元素的顺序
func ShuffleIntSlice(a []int) []int {
	if len(a) < 2 {
		return a
	}

	b := make([]int, 0)
	for len(a) > 0 {
		b = append(b, a[Rand(len(a))])
		a = IntSliceDiff(a, b)
	}

	return b
}

// 交集
func IntSliceIntersection(a, b []int) []int {
	c := make([]int, 0)
	for _, i := range a {
		for _, j := range b {
			if j == i {
				c = append(c, j)
			}
		}
	}

	return c
}

// 差集
func IntSliceDiff(a, b []int) []int {
	c := make([]int, 0)
	inter := IntSliceIntersection(a, b)
	for _, i := range a {
		if len(IntSliceIntersection(inter, []int{i})) == 0 {
			c = append(c, i)
		}
	}

	return c
}

// 采用二分法对 string 进行去重、排序
func BinaryAppend(a []string, n int, key string) []string {
	if len(a) == 0 {
		return []string{key}
	}

	var low, high, mid int
	low = 0
	high = n

	for low <= high {
		mid = (low + high) / 2

		if key == a[mid] {
			return a
		} else if key < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	a2 := append([]string{}, a[low:]...)
	return append(append(a[0:low], key), a2...)
}

// 去掉字符串中的指定的字符
func RidSymbols(rawString string, symbols ...string) (parsedString string) {
	parsedString = rawString
	for _, symbol := range symbols {
		parsedString = strings.Replace(parsedString, symbol, "", -1)
	}

	return
}

// 解析Excel的时间：string -> time.Time
func ParseTime(rawTime string) (parsedTime time.Time) {
	//当Excel该列为日期、时间格式的数据时会解析到距离1899-12-30 00:00:00的天数
	dayNum, err := strconv.ParseFloat(rawTime, 64)
	if err != nil {
		if len(rawTime) > 3 {

			if len(rawTime) == 8 {
				parsedTime, err = time.Parse("2006-01-02", "20"+rawTime[6:8]+"-"+rawTime[0:5])
				if err == nil {
					return
				}
			}

			rawTime = strings.Replace(rawTime, "-", "/", -1)
			parsedTime, err = time.Parse("2006/1/2", rawTime)
			if err != nil {
				parsedTime, _ = time.Parse("2006", rawTime[0:4]) //仅出版年份
			}
		}
	} else {
		startTime, _ := time.Parse("2006-01-02 15:04:05", "1899-12-30 00:00:00")
		parsedTime = startTime.Add(time.Duration(int(dayNum*24*60*60)) * time.Second)
	}

	return
}

// Encrypt use sha256 to hash the password
func Encrypt(password string) string {
	h := sha256.New()
	_, _ = io.WriteString(h, "dB924YD"+password+"xuLwLdj")
	return fmt.Sprintf("%x", h.Sum(nil))

}

func GetRealIp(req *http.Request) (ip string) {

	if ip = req.Header.Get("X-Real-IP"); ip != "" {
		return
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		return
	}

	return
}

func GetGeoWithIp(ip string) map[string]string {
	if !MatchIp(ip) {
		return nil
	}

	db, err := geoip2.Open("static/GeoLite2-City.mmdb")
	defer func() {
		_ = db.Close()
	}()

	if err != nil {
		return nil
	}

	g, err := db.City(net.ParseIP(ip))
	if err != nil {
		return nil
	}

	geo := make(map[string]string)
	geo["city"] = g.City.Names["zh-CN"]
	if geo["city"] == "" {
		geo["city"] = g.City.Names["en"]
	}

	geo["province"] = ""
	if len(g.Subdivisions) > 0 {
		geo["province"] = g.Subdivisions[0].Names["zh-CN"]

		if geo["province"] == "" {
			geo["province"] = g.Subdivisions[0].Names["en"]
		}
	}

	geo["country"] = g.Country.Names["zh-CN"]
	if geo["country"] == "" {
		geo["country"] = g.Country.Names["en"]
	}

	geo["continent"] = g.Continent.Names["zh-CN"]
	if geo["continent"] == "" {
		geo["continent"] = g.Continent.Names["en"]
	}

	return geo
}

func MatchIp(ip string) bool {
	if m, _ := regexp.MatchString(IpPattern, ip); !m {
		return false
	}
	return true
}
