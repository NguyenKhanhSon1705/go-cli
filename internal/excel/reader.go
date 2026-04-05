package excel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

// ReadExcelFile đọc file Excel và trả về WorkbookData
func ReadExcelFile(filePath string) (*WorkbookData, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("không thể mở file: %w", err)
	}
	defer f.Close()

	sheetList := f.GetSheetList()
	if len(sheetList) == 0 {
		return nil, fmt.Errorf("file Excel không có sheet nào")
	}

	workbook := &WorkbookData{
		FileName: filepath.Base(filePath),
		Sheets:   make([]SheetData, len(sheetList)),
	}

	for i, sheetName := range sheetList {
		rows, err := f.Rows(sheetName)
		if err != nil {
			continue
		}
		
		// Tìm dimension để pre-allocate
		dim, _ := f.GetSheetDimension(sheetName)
		totalRowsInSheet := 0
		if dim != "" {
			var row int
			for k := len(dim) - 1; k >= 0; k-- {
				if dim[k] == ':' {
					_, row, _ = excelize.SplitCellName(dim[k+1:])
					break
				}
			}
			totalRowsInSheet = row
		}

		sheetData := SheetData{
			SheetName: sheetName,
			Data:      make([][]string, 0, totalRowsInSheet),
		}

		maxCols := 0
		rowCount := 0

		for rows.Next() {
			rowCount++
			row, err := rows.Columns()
			if err != nil {
				break
			}

			if len(row) > maxCols {
				maxCols = len(row)
			}
			sheetData.Data = append(sheetData.Data, row)

			// Báo cáo tiến độ mỗi 50 dòng
			if rowCount%50 == 0 {
				baseP := float64(i) / float64(len(sheetList)) * 100
				sheetWeight := 100.0 / float64(len(sheetList))
				var rowP float64
				if totalRowsInSheet > 0 {
					rowP = float64(rowCount) / float64(totalRowsInSheet) * sheetWeight
				} else {
					rowP = (float64(rowCount) / 1000.0) * sheetWeight // Fallback
				}
				
				p := int(baseP + rowP)
				if p > 95 { p = 95 }
				fmt.Fprintf(os.Stderr, "PROGRESS:%d\n", p)
			}
		}
		
		// Chuẩn hóa số cột sau khi đọc xong sheet
		if maxCols > 0 {
			for j := range sheetData.Data {
				if len(sheetData.Data[j]) < maxCols {
					padding := make([]string, maxCols-len(sheetData.Data[j]))
					sheetData.Data[j] = append(sheetData.Data[j], padding...)
				}
			}
		}

		workbook.Sheets[i] = sheetData
		
		// Xong 1 sheet, báo cáo max 95%
		fmt.Fprintf(os.Stderr, "PROGRESS:%d\n", int(float64(i+1)/float64(len(sheetList))*95))
	}

	return workbook, nil
}

// ToJSON chuyển WorkbookData thành JSON string (dùng Encoder để tối ưu RAM)
func ToJSON(data *WorkbookData) (string, error) {
	result := ImportResult{
		Success: true,
		Data:    data,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(result); err != nil {
		return "", fmt.Errorf("lỗi encode JSON: %w", err)
	}

	return buf.String(), nil
}

// ErrorJSON trả về JSON lỗi
func ErrorJSON(errMsg string) string {
	jsonBytes, _ := json.Marshal(ImportResult{
		Success: false,
		Error:   errMsg,
	})
	return string(jsonBytes)
}
