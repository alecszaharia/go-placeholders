package visitor

import (
	parser "antrl/gen"
)

type Block struct {
	Name    string
	Content string
	HasEnd  bool
}

func NewPlaceholdersVisitor() *PlaceholdersVisitor {
	return &PlaceholdersVisitor{
		BasePlaceholderParserVisitor: &parser.BasePlaceholderParserVisitor{},
	}
}

type PlaceholdersVisitor struct {
	*parser.BasePlaceholderParserVisitor
	Blocks []Block
}

func (v *PlaceholdersVisitor) VisitTemplate(ctx *parser.TemplateContext) interface{} {
	//fmt.Printf("Visiting Template: %s\n", ctx.GetText())
	ctx.Content().Accept(v)
	return v.VisitChildren(ctx) // Visit the first placeholder
}

func (v *PlaceholdersVisitor) VisitContent(ctx *parser.ContentContext) interface{} {
	//fmt.Printf("Visiting Content: %s\n", ctx.GetText())
	for _, child := range ctx.AllBlock() {
		child.Accept(v)
	}
	for _, child := range ctx.AllText() {
		child.Accept(v)
	}
	return v.VisitChildren(ctx) // Visit the first placeholder
}

func (v *PlaceholdersVisitor) VisitText(ctx *parser.TextContext) interface{} {
	//fmt.Printf("Visiting Text: %s\n", ctx.GetText())
	return v.VisitChildren(ctx) // Visit the first placeholder
}

func (v *PlaceholdersVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {

	//fmt.Printf("VisitBlock: %s\n", ctx.GetText())

	b := Block{HasEnd: false}
	b.Name = ctx.GetBlockName().GetText()

	if ctx.Content() != nil {
		//b.Content = ctx.Content().GetText()
		ctx.Content().Accept(v)
	}

	if ctx.GetBlockEndName() != nil {
		b.HasEnd = true
	}

	v.Blocks = append(v.Blocks, b)
	return v.VisitChildren(ctx)
}
