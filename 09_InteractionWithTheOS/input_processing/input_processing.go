package input_processing

import (
	"log"
	"microshell/shell"
	"microshell/shell/commands"
	"microshell/shell/parse"
	"time"
)

func ReadLine() {
	var i int
	//reader := bufio.NewReader(os.Stdin)
	shl := shell.NewShell()
	for i < 1 {
		//line, _, err := reader.ReadLine()
		//if err != nil {
		//	return
		//}
		//cmds, err := parse.CreateCommands(string(line), shl.Paths())
		cmds, err := parse.CreateCommands("ls; ls -la | cat -e; cat -e | cat -e | cat -e", shl.Paths())
		if err != nil {
			log.Fatal(err)
		}
		commands.ExecuteAll(cmds)
		time.Sleep(time.Second * 2)
		i++
	}
}
