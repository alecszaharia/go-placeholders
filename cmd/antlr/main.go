package main

import (
	parser "antrl/gen"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type PlaceholdersListener struct {
	*parser.BasePlaceholderParserListener
}

func (l *PlaceholdersListener) ExitBlock_placeholder(ctx *parser.Block_placeholderContext) {
	fmt.Println("Parsed ExitBlock_placeholder:", ctx.GetText())
}

func (l *PlaceholdersListener) ExitPlaceholder(ctx *parser.PlaceholderContext) {
	fmt.Println("Parsed ExitPlaceholder:", ctx.GetText())
}

func (s *PlaceholdersListener) ExitAttribute(ctx *parser.AttributeContext) {
	fmt.Println("Parsed ExitAttribute:", ctx.GetText())
}

func main() {
	input := antlr.NewInputStream("aaa{{A  ty='t' }}weg{{B  ty='t' }}{{A  ty='x' }}tyj{{end_K}}weg{{end_B}}weg{{end_A}}")
	lexer := parser.NewPlaceholderLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPlaceholderParser(tokens)

	antlr.ParseTreeWalkerDefault.Walk(&PlaceholdersListener{}, p.Template())
}
