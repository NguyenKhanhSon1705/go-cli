package root

import (
	"fmt"
	help "go-cli/cmd/helps"
	version "go-cli/cmd/versions"
	cmdConst "go-cli/const"
	"os"
)

func Execute() {

	if len(os.Args) < 2 {
		help.PrintHelp()
		os.Exit(0)
	}
	cmd := os.Args[1]
	fmt.Println("COMMAND:", cmd)
	switch cmd {
	// hướng dẫn sử dụng
	case cmdConst.CmdHelp:
		help.PrintHelp()
	// kiểm tra version
	case cmdConst.CmdVersion,
		cmdConst.CmdShortVersion:
		version.PrintVersion()

	case cmdConst.CmdImport:
		// xử lý import
		fmt.Println("Importing data...")
	case cmdConst.CmdExport:
		fmt.Println("Exporting data...")
	// hướng dẫn sử dụng
	default:
		help.PrintHelp()
	}
}
