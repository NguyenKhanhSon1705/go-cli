package helps

import "fmt"

func PrintHelp() {
	fmt.Println(`Usage:
  myapp <command> [flags]

Commands:
  fill        Điền dữ liệu Excel vào Word template
  version     Hiển thị phiên bản

Flags:
  fill:
    --excel      File Excel (bắt buộc)
    --template   Word template (bắt buộc)
    --out        File output (bắt buộc)

Examples:
  myapp fill --excel=data.xlsx --template=mau.docx --out=ketqua.docx
  myapp version`)
}
