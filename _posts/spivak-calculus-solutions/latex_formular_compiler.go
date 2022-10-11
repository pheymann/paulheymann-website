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
	flag.Parse()

	if sourceFilepath == nil || *sourceFilepath == "" {
		panic("Missing path to source")
	}

	source, err := os.ReadFile(*sourceFilepath)
	if err != nil {
		panic(err)
	}

	compile(string(source))
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

func compile(source string) {
	state := &parseState{}
	lines := strings.Split(source, "\n")

	fmt.Println("<p>")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		switch trimmedLine {
		case "":

		case paragraph:
			closeBlock(state)

			state.blockTpe = paragraph
			state.blockLine = 0

		case formular:
			closeBlock(state)

			state.blockTpe = formular
			state.blockLine = 0

		case problem:
			closeBlock(state)

			state.blockTpe = problem
			fmt.Println("</p>")
			fmt.Println("")
			fmt.Println("<p>")

		default:
			switch state.blockTpe {
			case paragraph:
				latexLine := paragraphCommandsToLatex(trimmedLine)
				fmt.Println("  " + latexLine)
				state.blockLine++

			case formular:
				if state.blockLine == 0 {
					fmt.Println("  \\[")
					fmt.Println("    \\begin{align*}")
				}

				latexLine := forumlarCommandsToLatex(trimmedLine)
				fmt.Println("      " + latexLine + " \\\\")
				state.blockLine++
			}
		}
	}

	closeBlock(state)
	fmt.Println("</p>")
}

func closeBlock(state *parseState) {
	if state.blockLine > 0 && state.blockTpe == formular {
		fmt.Println("    \\end{align*}")
		fmt.Println("  \\]")
	}

	fmt.Println("")
}

var (
	boldRegex = regexp.MustCompile("\\*\\*(.*)\\*\\*")
	textRegex = regexp.MustCompile("\"([\\w !?\\(\\)\\.]*)\"")
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
