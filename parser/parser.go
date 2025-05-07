package parser

import (
	"strconv"
	"strings"
)

type CommandType int /// With only three variants this could be a uint8

const (
	ARIT_COMMAND CommandType = 0 /// You could use iota here
	PUSH_COMMAND CommandType = 1
	POP_COMMAND  CommandType = 2
)

type Command struct {
	CommandType CommandType
	// Should arguments be typed better, and/or depending on the command type?
	// Arg1 string can only be certains strings, Arg2 is only defined for push/pop
	Arg1 string
	Arg2 int
}

// Use named returns because it wasn't clear what the bool was for otherwise
func ParseLine(line string) (command Command, isValid bool) {
	// comments and empty lines are ignored
	if strings.HasPrefix(line, "//") || line == "" {
		return Command{}, false
	}

	commandType := getCommandType(line)

	/// Here I would see if you could parse all args, irrespective of what command you're dealing with.
	/// Then if you get back two args instead of one, and it's an arithmetic command, you can return an error.

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
	/// I think strings.Fields is what you want here (and below)
	substrings := strings.Split(line, " ")
	if len(substrings) > 1 {
		return substrings[1]
	}
	// arithmentic commands
	return line
}

func getArg2(line string) (int, error) {
	substrings := strings.Split(line, " ")
	arg2, err := strconv.Atoi(substrings[2])
	if err != nil {
		return 0, err
	}

	return arg2, nil
}
