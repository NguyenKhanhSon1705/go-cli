package excel

// SheetData chứa tên sheet và dữ liệu dạng mảng 2 chiều
type SheetData struct {
	SheetName string     `json:"sheetName"`
	Data      [][]string `json:"data"`
}

// WorkbookData chứa toàn bộ dữ liệu workbook
type WorkbookData struct {
	FileName string      `json:"fileName"`
	Sheets   []SheetData `json:"sheets"`
}

// ImportResult là kết quả trả về cho Electron
type ImportResult struct {
	Success bool          `json:"success"`
	Data    *WorkbookData `json:"data,omitempty"`
	Error   string        `json:"error,omitempty"`
}
