package command

import (
	"easydock/generate"
	"errors"
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"github.com/DrSmithFr/go-console/pkg/style"
	"os"
)

type Reader interface {
	Parse()
}

func Parse()  {
	io := style.
		NewConsoleStyler().
		AddInputArgument(
			argument.
				New("action", argument.REQUIRED),
		).
		AddInputArgument(
			argument.
				New("language", argument.OPTIONAL),
		).
		AddInputOption(
			option.
				New("help", option.OPTIONAL).
				SetDefault("0"),
		).
		ParseInput().
		ValidateInput()

	action := io.GetInput().GetArgument("action")
	language := io.GetInput().GetArgument("language")
	help := io.GetInput().GetOption("help")

	if len(action) == 0 && len(language) == 0 && help != "0" {
		io.GetOutput().Writeln(helpOutput())
		os.Exit(0)
	}

	err := validateArg(action, language)
	if err != nil {
		io.GetOutput().Writeln(err.Error())
		os.Exit(0)
	}
	if action == "create" {
		generate.CreateProject(language)
	}
}

func helpOutput() (output string) {
	output = `action   => create project
language => language use for current project (only php)`
	return output
}

func validateArg(action string, language string) error {
	if action != "create" {
		return errors.New("Argument: " + action + " is not available")
	}

	if language != "php" {
		return errors.New("Language: " + language + " is not available")
	}

	return nil
}

func dispatch(action string, language string)  {
	if action == "create" {
		// call generate package
	}
}