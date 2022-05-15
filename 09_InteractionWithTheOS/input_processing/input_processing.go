package input_processing

import (
	"log"
	"microshell/shell"
	"microshell/shell/commands"
	"microshell/shell/parse"
	"time"
)

func ReadLine() {
	//reader := bufio.NewReader(os.Stdin)
	shl := shell.NewShell()
	for {
		//line, _, err := reader.ReadLine()
		//if err != nil {
		//	return
		//}
		//fmt.Println(string(line))
		//cmds, err := parse.CreateCommands(string(line), shl.Paths())
		cmds, err := parse.CreateCommands("ls -la | echo", shl.Paths())
		if err != nil {
			log.Fatal(err)
		}
		commands.ExecuteAll(cmds)
		time.Sleep(time.Second * 2)
		return
	}
}
