package bot

import "strings"

func (b Bot) GetCommands(content string) (command string, values []string)  {
	result := strings.Split(content, " ")

	return result[1], result[2:]
}
