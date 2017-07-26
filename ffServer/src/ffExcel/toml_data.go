package ffExcel

import "fmt"

// table
var fmtSheetList = `[[%v]]
`
var fmtSheetMapStart = `[%v]
`
var fmtSheetMapKey = `	[%v.%v]
`
var fmtSheetStruct = `[%v]
`

// field
var fmtFieldList = `	%v = %v
`
var fmtFieldMap = `	%v = %v
`
var fmtFieldStruct = `%v = %v
`

var fmtFieldSplitMap = "\n"
var fmtSheetSplit = "\n"

func genTomlData(excel *excel, exportConfig *ExportConfig) string {
	result := ""
	for _, sheet := range excel.sheets {
		if sheet.sheetType == sheetTypeMap {
			result += fmt.Sprintf(fmtSheetMapStart, sheet.name)
		} else if sheet.sheetType == sheetTypeStruct {
			result += fmt.Sprintf(fmtSheetStruct, sheet.name)
		}

		for i, row := range sheet.content.rows {
			if sheet.sheetType == sheetTypeMap {
				if i > 0 {
					result += fmtFieldSplitMap
				}

				result += fmt.Sprintf(fmtSheetMapKey, sheet.name, row.rowData[sheetTypeMapKeyName].Value())
			} else if sheet.sheetType == sheetTypeList {
				if i > 0 {
					result += fmtFieldSplitMap
				}

				result += fmt.Sprintf(fmtSheetList, sheet.name)
			}

			exportedLines := make(map[string]bool, len(sheet.header.lines))
			for _, line := range sheet.header.lines {
				if line.exportToServer() && !line.isMapKey() {
					if _, ok := exportedLines[line.lineName]; !ok {
						exportedLines[line.lineName] = true
						result += fmt.Sprintf(fmtFieldList, line.lineName, row.rowData[line.lineName].ValueToml())
					}
				}
			}
		}

		result += fmtSheetSplit
	}

	return result
}
