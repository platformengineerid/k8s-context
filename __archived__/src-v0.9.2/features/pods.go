package features

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

//	func NewSortableTable(headers []string, rows [][]string, filter string) *tablewriter.Table {
//		table := tablewriter.NewWriter(os.Stdout)
//		table.SetHeader(headers)
//		if filter != "" {
//			filteredRows := [][]string{}
//			for _, row := range rows {
//				if strings.Contains(strings.ToLower(row[9]), strings.ToLower(filter)) {
//					filteredRows = append(filteredRows, row)
//				}
//			}
//			rows = filteredRows
//		}
//		for _, row := range rows {
//			table.Append(row)
//		}
//		table.SortBy([]tablewriter.SortBy{{Column: 4, Mode: tablewriter.DESC}})
//		return table
//	}

// func NewSortableTable(headers []string, rows [][]string, sortColumn int, sortMode int) *tablewriter.Table {
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetHeader(headers)
// 	for _, row := range rows {
// 		table.Append(row)
// 	}
// 	sortColumns := []tablewriter.SortableColumn{
// 		{
// 			Column:    sortColumn,
// 			Ascending: sortMode == tablewriter.Ascending,
// 		},
// 	}
// 	tablewriter.SortByColumns(sortColumns).Sort(table)
// 	table.Render()
// 	return table
// }

// func NewSortableTable(headers []string, rows [][]string, sortColumn int, sortMode int) *tablewriter.Table {
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetHeader(headers)
// 	for _, row := range rows {
// 		table.Append(row)
// 	}
// 	sortSettings := &tablewriter.SortBy{
// 		Columns: []tablewriter.SortColumn{
// 			{
// 				Column:    sortColumn,
// 				Ascending: sortMode == tablewriter.Ascending,
// 			},
// 		},
// 	}
// 	table.SortBy(sortSettings)
// 	table.Render()
// 	return table
// }

// func NewSortableTable(headers []string, rows [][]string, sortColumn int, sortMode int) *tablewriter.Table {
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetHeader(headers)
// 	for _, row := range rows {
// 		table.Append(row)
// 	}
// 	table.SetSortableBy([]tablewriter.SortableColumn{
// 		{
// 			ColumnIdx: sortColumn,
// 			Ascending: sortMode == tablewriter.Ascending,
// 			Comparator: func(a, b interface{}) int {
// 				strA := a.(string)
// 				strB := b.(string)
// 				return strings.Compare(strA, strB)
// 			},
// 		},
// 	})
// 	table.Render()
// 	return table
// }

// func NewSortableTable(headers []string, rows [][]string, sortColumn int, sortMode tablewriter.SortMode) *tablewriter.Table {
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetHeader(headers)
// 	for _, row := range rows {
// 		table.Append(row)
// 	}
// 	sortSettings := tablewriter.SortSettings{
// 		SortColumn: sortColumn,
// 		SortMode:   sortMode,
// 	}
// 	table.SetAutoMergeCellsByColumnIndex([]int{0})
// 	table.SetRowLine(true)
// 	table.SetCenterSeparator("|")
// 	table.SetColumnSeparator("|")
// 	table.SetBorder(false)
// 	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
// 	table.SetAlignment(tablewriter.ALIGN_LEFT)
// 	table.SetSortBy(sortSettings)
// 	table.Render()
// 	return table
// }

// func NewSortableTable(headers []string, rows [][]string) *tablewriter.SortableTable {
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetHeader(headers)
// 	for _, row := range rows {
// 		table.Append(row)
// 	}
// 	st := &tablewriter.SortableTable{
// 		Table: table,
// 		Columns: []*tablewriter.Column{
// 			{Title: headers[0]},
// 			{Title: headers[1], Numeric: true},
// 			{Title: headers[2]},
// 			{Title: headers[3], Numeric: true},
// 			{Title: headers[4], Numeric: true, DateFormat: "Jan 2 15:04:05 MST 2006"},
// 			{Title: headers[5]},
// 			{Title: headers[6]},
// 			{Title: headers[7]},
// 			{Title: headers[8]},
// 			{Title: headers[9]},
// 		},
// 	}
// 	st.SortBy([]tablewriter.SortBy{{Column: 4, Mode: tablewriter.DESC}})
// 	return st
// }

// func NewSortableTable(headers []string, rows [][]string) *tablewriter.Table {
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetHeader(headers)
// 	for _, row := range rows {
// 		table.Append(row)
// 	}
// 	table.Render()
// 	return table
// }

func NewSortableTable(headers []string, rows [][]string, filter int, keyword string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	for _, row := range rows {
		if strings.Contains(strings.ToLower(row[filter]), strings.ToLower(keyword)) {
			table.Append(row)
		}
	}
	table.Render()
	return table
}