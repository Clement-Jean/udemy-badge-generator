package main

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var url string = "https://www.udemy.com/instructor-api/v1/taught-courses/courses?fields[course]=id,rating"

func getRating(ids []string, udemyToken string) float64 {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)

	check(err)

	req.Header.Set("authorization", "bearer "+udemyToken)
	response, err := client.Do(req)

	check(err)
	defer response.Body.Close()

	courses := &CourseResult{}
	json.NewDecoder(response.Body).Decode(&courses)

	total := 0.0
	found := 0.0

	for _, id := range ids {
		for _, course := range courses.Results {
			if id == course.Id {
				found++
				total += course.Rating
			}
		}
	}

	if found == 0 {
		return -1.0
	}

	return math.Ceil(total/found*100) / 100
}

var characterLengths map[string]interface{}
var kerningPairs map[string]interface{}

//go:embed assets/template.svg
var badgeTemplate string

//go:embed assets/default-widths.json
var defaultWidthsBytes []byte

func generate(b Badge, path string, filename string) {
	valueLength := calculateTextLength110(b.Value)
	labelLength := calculateTextLength110(b.Label)
	rightWidth := math.Ceil(valueLength/10) + 10
	leftWidth := math.Ceil(labelLength/10) + 10
	badgeWidth := leftWidth + rightWidth
	// The -10 below is for an exta buffer on right end of badge
	rightCenter := 10*leftWidth + rightWidth*5 - 10
	// The +10 below is for an extra buffer on left end of badge
	leftCenter := 10 + leftWidth*5
	dim := Dimensions{
		LabelLength: labelLength,
		ValueLength: valueLength,
		RightWidth:  rightWidth,
		LeftWidth:   leftWidth,
		BadgeWidth:  badgeWidth,
		LeftCenter:  leftCenter,
		RightCenter: rightCenter,
	}

	err := os.MkdirAll(path, os.ModePerm)
	check(err)

	fpath := filepath.Join(path, filename+".svg")
	f, err := os.Create(fpath)
	check(err)

	defer f.Close()
	w := bufio.NewWriter(f)
	t, err := template.New("badge").Parse(string(badgeTemplate))
	check(err)
	err = t.Execute(w, struct {
		Badge
		Dimensions
	}{b, dim})
	check(err)

	w.Flush()
}

func init() {
	var defaultWidths map[string]interface{}
	json.Unmarshal(defaultWidthsBytes, &defaultWidths)

	characterLengths = defaultWidths["character-lengths"].(map[string]interface{})
	kerningPairs = defaultWidths["kerning-pairs"].(map[string]interface{})
}

func trimSuffix(s string, substr string) (t string) {
	for {
		t = strings.TrimSuffix(s, substr)
		if t == s {
			break
		}
		s = t
	}
	return t
}

func formatRating(r float64) string {
	rating := fmt.Sprintf("%f", r)
	rating = trimSuffix(rating, "0")
	rating = trimSuffix(rating, ".")
	return rating
}

func main() {
	label := os.Args[1]
	labelBgColor := os.Args[2]
	labelFgColor := os.Args[3]
	valueFgColor := os.Args[4]
	badgeRadius, err := strconv.ParseFloat(os.Args[5], 64)
	check(err)
	generatePath := os.Args[6]
	filename := os.Args[7]

	udemyToken := os.Getenv("UDEMY_TOKEN")
	if udemyToken == "" {
		panic("Udemy token not set")
	}

	udemyCoursesId := os.Getenv("UDEMY_COURSE")
	if udemyCoursesId == "" {
		panic("Udemy course id should at least contain one element")
	}

	ids := strings.Split(udemyCoursesId, ",")

	value := getRating(ids, udemyToken)
	if value == -1.0 {
		panic("Couldn't find the course")
	}

	valueBgColor := getValueBgColor(value)

	badge := Badge{
		Label:        label,
		Value:        formatRating(value),
		BgColorLabel: labelBgColor,
		FgColorLabel: labelFgColor,
		FgColorValue: valueFgColor,
		BgColorValue: valueBgColor,
		Radius:       badgeRadius,
	}

	generate(badge, generatePath, filename)
}
