package parser

import (
	"strconv"
	"strings"
)

type CommandType int

const (
	ARIT_COMMAND CommandType = 0
	PUSH_COMMAND CommandType = 1
	POP_COMMAND  CommandType = 2
)

type Command struct {
	CommandType CommandType
	Arg1        string
	Arg2        int
}

func ParseLine(line string) (Command, bool) {
	// comments and empty lines are ignored
	if strings.HasPrefix(line, "//") || line == "" {
		return Command{}, false
	}

	commandType := getCommandType(line)
	finalLine := Command{
		CommandType: commandType,
		Arg1:        getArg1(line),
	}

	if commandType != ARIT_COMMAND {
		arg2, err := getArg2(line)
		if err != nil {
			return Command{}, false
		}
		finalLine.Arg2 = arg2
	}

	return finalLine, true
}

func getCommandType(line string) CommandType {
	if strings.HasPrefix(line, "pop") {
		return POP_COMMAND
	}

	if strings.HasPrefix(line, "push") {
		return PUSH_COMMAND
	}

	return ARIT_COMMAND
}

func getArg1(line string) string {
	substrings := strings.Split(line, " ")
	return substrings[1]
}

func getArg2(line string) (int, error) {
	substrings := strings.Split(line, " ")
	arg2, err := strconv.Atoi(substrings[2])
	if err != nil {
		return 0, err
	}

	return arg2, nil
}
