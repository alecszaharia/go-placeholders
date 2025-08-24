package visitor

import (
	parser "antrl/internal/grammar"
	"antrl/internal/model"

	"github.com/antlr4-go/antlr/v4"
)

type PlaceholdersVisitor struct {
	*parser.BasePlaceholderParserVisitor
	Nodes              []model.Node
	currentNode        *model.Node
	currentPlaceholder *model.Placeholder
}

func NewPlaceholdersVisitor() *PlaceholdersVisitor {
	return &PlaceholdersVisitor{
		BasePlaceholderParserVisitor: &parser.BasePlaceholderParserVisitor{},
	}
}

func (v *PlaceholdersVisitor) VisitTemplate(ctx *parser.TemplateContext) interface{} {
	for _, actx := range ctx.GetChildren() {
		if child, ok := actx.(antlr.ParserRuleContext); ok {
			child.Accept(v)
		}
	}

	return nil
}

func (v *PlaceholdersVisitor) VisitContent(ctx *parser.ContentContext) interface{} {
	for _, actx := range ctx.GetChildren() {
		if child, ok := actx.(antlr.ParserRuleContext); ok {
			child.Accept(v)
		}
	}
	return nil
}

func (v *PlaceholdersVisitor) VisitText(ctx *parser.TextContext) interface{} {
	node := model.MakeTextNode(ctx.GetStart().GetTokenIndex(), ctx.GetStop().GetTokenIndex(), v.currentNode)

	// temporary
	node.Value.(*model.Text).Text = ctx.GetText()

	if v.currentNode == nil {
		v.Nodes = append(v.Nodes, *node)
	}

	return nil
}

func (v *PlaceholdersVisitor) VisitAttribute(ctx *parser.AttributeContext) interface{} {
	placeholder, ok := v.currentNode.Value.(*model.Placeholder)
	if !ok {
		// Handle the error - log, return error, or panic with context
		panic("expected *model.Placeholder but got different type")
	}
	attr := model.Attr{Name: ctx.GetName().GetText(), Value: ctx.GetValue().GetText()}
	placeholder.Attrs = append(placeholder.Attrs, attr)
	return nil
}

func (v *PlaceholdersVisitor) VisitPlaceholder(ctx *parser.PlaceholderContext) interface{} {

	for _, actx := range ctx.GetChildren() {
		if child, ok := actx.(antlr.ParserRuleContext); ok {
			child.Accept(v)
		}
	}

	return nil
}
func (v *PlaceholdersVisitor) VisitSimplePlaceholder(ctx *parser.SimplePlaceholderContext) interface{} {

	node := model.MakePlaceholderNode(ctx.GetStart().GetTokenIndex(), ctx.GetStop().GetTokenIndex(), v.currentNode)
	v.currentNode = node

	v.currentNode.Value = &model.Placeholder{
		Name:    ctx.GetPlaceholderName().GetText(),
		IsBlock: false,
	}

	if ctx.AllAttribute() != nil {
		// Temporarily set v.currentPlaceholder for attribute parsing
		for _, attr := range ctx.AllAttribute() {
			attr.Accept(v)
		}
	}

	// Pop back to the parent node
	v.currentNode = v.currentNode.Parent

	if v.currentNode == nil {
		v.Nodes = append(v.Nodes, *node)
	}

	return nil
}
func (v *PlaceholdersVisitor) VisitBlockPlaceholderStart(ctx *parser.BlockPlaceholderStartContext) interface{} {
	placeholder, ok := v.currentNode.Value.(*model.Placeholder)
	if !ok {
		// Handle the error - log, return error, or panic with context
		panic("expected *model.Placeholder but got different type")
	}
	placeholder.Name = ctx.GetPlaceholderName().GetText()

	return nil
}
func (v *PlaceholdersVisitor) VisitBlockPlaceholderEnd(ctx *parser.BlockPlaceholderEndContext) interface{} {
	//fmt.Printf("VisitBlockPlaceholderEnd: %s \n", ctx.GetText())
	return nil
}
func (v *PlaceholdersVisitor) VisitBlockPlaceholderContent(ctx *parser.BlockPlaceholderContentContext) interface{} {
	//fmt.Printf("VisitBlockPlaceholderContent: %s \n", ctx.GetText())
	placeholder, ok := v.currentNode.Value.(*model.Placeholder)
	if !ok {
		// Handle the error - log, return error, or panic with context
		panic("expected *model.Placeholder but got different type")
	}
	placeholder.ContentStart = ctx.GetStart().GetTokenIndex()
	placeholder.ContentEnd = ctx.GetStop().GetTokenIndex()

	for _, actx := range ctx.GetChildren() {
		if child, ok := actx.(antlr.ParserRuleContext); ok {
			child.Accept(v)
		}
	}
	return nil
}
func (v *PlaceholdersVisitor) VisitBlockPlaceholder(ctx *parser.BlockPlaceholderContext) interface{} {

	v.currentNode = model.MakeBlockNode(ctx.GetStart().GetTokenIndex(), ctx.GetStop().GetTokenIndex(), v.currentNode)

	v.currentNode.Value = &model.Placeholder{
		IsBlock: true,
	}

	for _, actx := range ctx.GetChildren() {
		if child, ok := actx.(antlr.ParserRuleContext); ok {
			child.Accept(v)
		}
	}

	if v.currentNode.Parent == nil {
		v.Nodes = append(v.Nodes, *v.currentNode)
	}

	// Pop back to the parent node
	v.currentNode = v.currentNode.Parent

	return nil
}
