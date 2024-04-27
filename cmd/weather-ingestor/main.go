package main

import (
	"fmt"
	"strings"
	"weather-ingestor/domain/forecast_office"
	"weather-ingestor/domain/grid_point"
	"weather-ingestor/domain/points"
	"weather-ingestor/domain/temperature"
)

// And the forecast: https://api.weather.gov/gridpoints/BOU/64,73/forecast

// Stations are located via https://api.weather.gov/stations?state=CO

func main() {
	point, err := points.New(40.0150, -105.2705)
	if err != nil {
		panic(err)
	}

	office, err := forecast_office.Instances.InstanceFromString(point.Properties.GridId)
	if err != nil {
		panic(err)
	}
	fmt.Println(office.ForecastOfficeName)

	gridPoint, err := grid_point.New(point.Properties.GridId, point.Properties.GridX, point.Properties.GridY)
	if err != nil {
		panic(err)
	}

	uom := gridPoint.Properties.Temperature.Uom
	for _, temperatures := range gridPoint.Properties.Temperature.Values.BucketByDay() {
		printCells(temperatures.Day(), temperatures.AsCells(uom))
	}
}

// Max width
// Max cells
// Header

// /===========================================================\
// |                        2024-04-07                         |
// +===========+===========+===========+===========+===========+
// |  2:00AM   |  3:00AM   |  5:00AM   |  6:00AM   |  7:00AM   |
// |  38.67°F  |  39.22°F  |  37.56°F  |  36.44°F  |  35.33°F  |
// \===========+===========+===========+===========+===========/
func printCells(header string, cells temperature.Cells) {
	maxCellWidth := 11
	maxNumCells := 12

	overallWidth := (maxCellWidth * maxNumCells) + (maxNumCells) + 1

	var rows []string

	rows = append(rows, tableHeaderStart(overallWidth))
	rows = append(rows, tableHeaderValue(header, overallWidth))

	rows = append(rows, rowSeparator(maxCellWidth, maxNumCells))

	cellsInRow := partition(cells, maxNumCells)

	for _, cellRow := range cellsInRow {
		numRowsInCells := cellRow[0].NumRowsInCell()
		cellRows := make([][]string, numRowsInCells)
		for row := range numRowsInCells {
			numCols := len(cellRow)
			cellRows[row] = make([]string, numCols)
		}

		for i, cell := range cellRow {
			col := i % maxNumCells
			for row := range cell.NumRowsInCell() {
				cellRows[row][col] = cell.RowValue(row, maxCellWidth-1)
			}
		}
		for _, row := range cellRows {
			rows = append(rows, "| "+strings.Join(row, "| ")+"|")
			rows = append(rows, rowSeparator(maxCellWidth, maxNumCells))
		}

	}

	fmt.Printf(strings.Join(rows, "\n") + "\n\n")
}

func partition(cells temperature.Cells, maxCellsPerRow int) []temperature.Cells {
	groups := make([]temperature.Cells, len(cells)/maxCellsPerRow)

	for i, cell := range cells {
		groups[i/maxCellsPerRow] = append(groups[i/maxCellsPerRow], cell)
	}

	return groups
}

func rowSeparator(maxCellWidth int, maxNumCells int) string {
	return strings.Repeat("+"+strings.Repeat("-", maxCellWidth), maxNumCells) + "+"
}

func tableHeaderValue(header string, overallWidth int) string {
	leftPadding := (overallWidth - 1 - len(header)) / 2
	rightPadding := overallWidth - 1 - len(header) - leftPadding - 1
	headerValue := "|" + strings.Repeat(" ", leftPadding) + header + strings.Repeat(" ", rightPadding) + "|"
	return headerValue
}

func tableHeaderStart(overallWidth int) string {
	return fmt.Sprintf("+" + strings.Repeat("-", overallWidth-2) + `+`)
}

type RowType int

const (
	RowTypeTop = iota
	RowTypeMiddle
	RowTypeBottom
)

// +-----------+
// |  17:00AM  |
// |  35.33°F  |
// \-----------+
