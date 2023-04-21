package format

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var httpr = regexp.MustCompile("https?://")

// Appends zero to integers less than 10 for some data format
// Eg:AppendZero(9) returns 09
func AppendZero(v int) string {
	if v < 10 && v > -10 {
		return "0" + strconv.Itoa(v)
	} else {
		return strconv.Itoa(v)
	}
}

func ParseDate(date string) time.Time {
	layout := "02-01-2006"
	t, err := time.Parse(layout, date)

	if err != nil {
		if date != "" {
			log.Println("Unable to parse date")
			log.Println(err)
		}
	}
	return t
}

func ParseDateFormat(date string, format string) time.Time {
	t, err := time.Parse(format, date)

	if err != nil {
		if date != "" {
			log.Println("Unable to parse date")
			log.Println(err)
		}
	}
	return t
}

func ParseFloat(num string) float64 {
	f, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0.0
	}
	return f
}

func ParseInt(num string) int {
	i, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}
	return i
}

func ParseBoolean(boo string) bool {
	b, err := strconv.ParseBool(boo)
	if err != nil {
		return false
	}
	return b
}

func CleanWebsiteURL(v string) string {
	cv := strings.Trim(v, "! ")
	if len(cv) == 0 {
		return cv
	}
	if len(httpr.FindStringIndex(v)) == 0 {
		cv = "https://" + v
	}
	return cv
}

func Slugify(str string) string {
	// Convert to lowercase
	str = strings.ToLower(str)

	// Remove all non-word characters and replace with "-"
	reg, err := regexp.Compile(`[\W]+`)
	if err != nil {
		panic(err)
	}
	str = reg.ReplaceAllString(str, "-")

	// Remove any leading or trailing "-"
	str = strings.Trim(str, "-")

	return str
}
