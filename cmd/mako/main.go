package mako

import (
	"flag"
	"fmt"
	"os"
	//"github.com/dylandibeneditto/mako/internal/parser"
	//"github.com/dylandibeneditto/mako/internal/evaluator"
)

func main() {
	force := flag.Bool("force", false, "allow overwriting the input file")
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Usage:\n    mako <input file> <output file> [-force]\n    mako run <header file> <target file> [output file])")
		os.Exit(1)
	}

	if args[0] == "run" {
		if len(args) < 3 || len(args) > 4 {
			fmt.Println("Usage: mako run <header file> <target file> [output file]")
			os.Exit(1)
		}
		headerPath := args[1]
		targetPath := args[2]
		if len(args) == 4 {
			runWithHeader(headerPath, targetPath, args[3])
		} else {
			runWithHeader(headerPath, targetPath, targetPath)
		}
		return
	}

	inputPath := args[0]
	outputPath := args[1]

	if inputPath == outputPath && !*force {
		fmt.Println("Running this command would destroy the mako source. If this is desired please add the '-force' flag.")
		os.Exit(1)
	}

	run(inputPath, outputPath)
}

func run(inputPath, outputPath string) {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("Error while reading input:", err)
		os.Exit(1)
	}

	prog, err := parser.Parse(string(content))
	if err != nil {
		fmt.Println("Error while parsing file:", err)
		os.Exit(1)
	}

	output, err := evaluator.Execute(prog)
	if err != nil {
		fmt.Println("Error executing mako:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(outputPath, []byte(output), 0644); err != nil {
		fmt.Println("Error writing output:", err)
		os.Exit(1)
	}

}

func runWithHeader(headerPath, targetPath, outputPath string) {
	headerContent, err := os.ReadFile(headerPath)
	if err != nil {
		fmt.Println("Error reading header:", err)
		os.Exit(1)
	}

	targetContent, err := os.ReadFile(targetPath)
	if err != nil {
		fmt.Println("Error reading target:", err)
		os.Exit(1)
	}

	prog, err := parser.Parse(string(headerContent))
	if err != nil {
		fmt.Println("Error parsing header:", err)
		os.Exit(1)
	}

	output, err := evaluator.ExecuteOnContent(prog, string(targetContent))
	if err != nil {
		fmt.Println("Error executing mako:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(outputPath, []byte(output), 0644); err != nil {
		fmt.Println("Error writing output:", err)
		os.Exit(1)
	}
}
