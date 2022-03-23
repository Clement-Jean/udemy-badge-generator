package main

var valueBgColors []string = []string{
	"#e05d44", //<1
	"#fe7d37", //1
	"#dfb317", //2
	"#a4a61d", //3
	"#97ca00", //4
	"#4c1",    //5
}

func getValueBgColor(percent float64) string {
	defaultColor := valueBgColors[0]
	cutoffs := []float64{1, 2, 3, 4, 5}

	for i := 1; i < len(valueBgColors); i++ {
		if percent < cutoffs[i-1] {
			break
		}

		defaultColor = valueBgColors[i]
	}

	return defaultColor
}
