package cmd

import (
	"fmt"
	help "go-cli/cmd/helps"
	version "go-cli/cmd/versions"
	cmdConst "go-cli/const"
	"go-cli/internal/excel"
	"os"
)

func Execute() {

	if len(os.Args) < 2 {
		help.PrintHelp()
		os.Exit(0)
	}
	cmd := os.Args[1]

	switch cmd {
	// hướng dẫn sử dụng
	case cmdConst.CmdHelp:
		help.PrintHelp()

	// kiểm tra version
	case cmdConst.CmdVersion,
		cmdConst.CmdShortVersion:
		version.PrintVersion()

	// xử lý import Excel
	case cmdConst.CmdImport:
		handleImport()

	case cmdConst.CmdExport:
		fmt.Println("Exporting data...")

	// hướng dẫn sử dụng
	default:
		help.PrintHelp()
	}
}

func handleImport() {
	// Parse arguments: import --file <path>
	filePath := ""
	for i := 2; i < len(os.Args); i++ {
		if os.Args[i] == "--file" && i+1 < len(os.Args) {
			filePath = os.Args[i+1]
			break
		}
	}

	if filePath == "" {
		fmt.Println(excel.ErrorJSON("thiếu tham số --file <đường dẫn file>"))
		os.Exit(1)
	}

	// Đọc file Excel
	workbook, err := excel.ReadExcelFile(filePath)
	if err != nil {
		fmt.Println(excel.ErrorJSON(err.Error()))
		os.Exit(1)
	}

	// Chuyển sang JSON và in ra stdout
	jsonStr, err := excel.ToJSON(workbook)
	if err != nil {
		fmt.Println(excel.ErrorJSON(err.Error()))
		os.Exit(1)
	}

	fmt.Println(jsonStr)
}
