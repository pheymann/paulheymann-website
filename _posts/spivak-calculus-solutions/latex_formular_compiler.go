package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	sourceFilepath := flag.String("source", "", "path to your source file")
	outputFile := flag.String("out", "", "path to the output file")
	flag.Parse()

	if sourceFilepath == nil || *sourceFilepath == "" {
		panic("Missing path to source")
	}

	source, err := os.ReadFile(*sourceFilepath)
	if err != nil {
		panic(err)
	}

	if outputFile != nil && *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			panic("couldn't open or create output file. Reason: " + err.Error())
		}

		compile(string(source), FilePrinter{file})
	} else {
		compile(string(source), CLIPrinter{})
	}
}

type blockType = string

const (
	paragraph blockType = "#p"
	formular  blockType = "#f"
	problem   blockType = "#--"
)

type parseState struct {
	blockTpe  blockType
	blockLine int8
}

func compile(source string, output Output) {
	state := &parseState{}
	lines := strings.Split(source, "\n")

	output.println("<p>")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		switch trimmedLine {
		case "":

		case paragraph:
			closeBlock(state, output)

			state.blockTpe = paragraph
			state.blockLine = 0

		case formular:
			closeBlock(state, output)

			state.blockTpe = formular
			state.blockLine = 0

		case problem:
			closeBlock(state, output)

			state.blockTpe = problem
			output.println("</p>")
			output.println("")
			output.println("<p>")

		default:
			switch state.blockTpe {
			case paragraph:
				latexLine := paragraphCommandsToLatex(trimmedLine)
				output.println("  " + latexLine)
				state.blockLine++

			case formular:
				if state.blockLine == 0 {
					output.println("  \\[")
					output.println("    \\begin{align*}")
				}

				latexLine := forumlarCommandsToLatex(trimmedLine)
				output.println("      " + latexLine + " \\\\")
				state.blockLine++
			}
		}
	}

	closeBlock(state, output)
	output.println("</p>")
}

func closeBlock(state *parseState, output Output) {
	if state.blockLine > 0 && state.blockTpe == formular {
		output.println("    \\end{align*}")
		output.println("  \\]")
	}

	output.println("")
}

var (
	boldRegex = regexp.MustCompile("\\*\\*(.*)\\*\\*")
	textRegex = regexp.MustCompile("\"([\\w !?\\(\\)\\.']*)\"")
)

func paragraphCommandsToLatex(line string) string {
	var latexLine = line

	latexLine = strings.ReplaceAll(latexLine, "#->", "\\rightarrow")
	latexLine = strings.ReplaceAll(latexLine, "#<-", "\\leftarrow")
	latexLine = boldRegex.ReplaceAllString(latexLine, "<strong>$1</strong>")

	return latexLine
}

func forumlarCommandsToLatex(line string) string {
	var latexLine = line

	latexLine = strings.ReplaceAll(latexLine, "#->", "\\rightarrow")
	latexLine = strings.ReplaceAll(latexLine, "#<-", "\\leftarrow")
	latexLine = strings.ReplaceAll(latexLine, "<", "&<")
	latexLine = strings.ReplaceAll(latexLine, ">", "&>")
	latexLine = strings.ReplaceAll(latexLine, "=", "&=")
	latexLine = textRegex.ReplaceAllString(latexLine, " && \\text{| $1}")

	return latexLine
}

type Output interface {
	println(string)
}

type CLIPrinter struct{}

func (printer CLIPrinter) println(str string) {
	fmt.Println(str)
}

type FilePrinter struct {
	file *os.File
}

func (printer FilePrinter) println(str string) {
	printer.file.WriteString(str + "\n")
}
