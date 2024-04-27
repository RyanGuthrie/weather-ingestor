package temperature

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"weather-ingestor/domain/uom"
	"weather-ingestor/domain/weather_time"
)

type Value float64

func (v Value) AsFahrenheit(origUOM uom.UOM) float64 {
	switch origUOM {
	case uom.DegC:
		return float64(v)*(9/5) + 32
	default:
		panic(fmt.Sprintf("unknown unit of measurement [%s], unable to convert to Fahrenheit", origUOM))
	}
}

type Temperature struct {
	ValidTime weather_time.WeatherTime `json:"validTime"`
	Value     Value                    `json:"value"`
}

type temperatures []Temperature

func (t temperatures) BucketByDay() []temperatures {
	partitionedTemps := map[string]temperatures{}

	for _, temp := range t {
		day := temp.ValidTime.Day()
		partitionedTemps[day] = append(partitionedTemps[day], temp)
	}

	sortedBucketTemperatures := make([]temperatures, 0, len(partitionedTemps))
	for _, temps := range partitionedTemps {
		sortedBucketTemperatures = append(sortedBucketTemperatures, temps)
	}

	sort.Slice(sortedBucketTemperatures, func(i, j int) bool {
		return sortedBucketTemperatures[i].Day() < sortedBucketTemperatures[j].Day()
	})

	return sortedBucketTemperatures
}

func (t temperatures) Day() string {
	if len(t) == 0 {
		panic("Unexpected empty temperatures slice")
	}

	return t[0].ValidTime.Day()
}

type Cell struct {
	values []string
}

func (c Cell) RowValue(row int, padding int) string {
	return c.values[row] + strings.Repeat(" ", padding-len(c.values[row]))
}

func (c Cell) NumRowsInCell() int {
	return len(c.values)
}

type Cells []Cell

func (t temperatures) AsCells(uom uom.UOM) Cells {
	tempByHour := make(Cells, 24) // one cell for each hour of the day

	for i, _ := range tempByHour {
		midnight := t[0].ValidTime.Midnight()
		tempByHour[i].values = []string{
			fmt.Sprintf("%s", midnight.Add(time.Hour*time.Duration(i)).Format(time.Kitchen)),
			"",
		}
	}

	for _, temp := range t {
		hour := temp.ValidTime.Hour()
		for offset := 0; offset < int(temp.ValidTime.Duration().Hours()); offset++ {
			tempByHour[hour+offset].values[1] = fmt.Sprintf("%.2f F", temp.Value.AsFahrenheit(uom))
		}
	}

	return tempByHour
}

type Temperatures struct {
	Uom    uom.UOM      `json:"uom"`
	Values temperatures `json:"values"`
}
