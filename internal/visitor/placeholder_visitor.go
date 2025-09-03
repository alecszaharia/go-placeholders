package visitor

import (
	parser "antrl/internal/grammar"
	"antrl/internal/model"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type PlaceholdersVisitor struct {
	*parser.BasePlaceholderParserVisitor
	Node *model.Node
}

func NewPlaceholdersVisitor() *PlaceholdersVisitor {
	return &PlaceholdersVisitor{
		BasePlaceholderParserVisitor: &parser.BasePlaceholderParserVisitor{},
	}
}

func (v *PlaceholdersVisitor) visitChildren(ctx antlr.RuleNode) {
	for _, c := range ctx.GetChildren() {
		if prc, ok := c.(antlr.ParserRuleContext); ok {
			prc.Accept(v)
		}
	}
}

func (v *PlaceholdersVisitor) VisitTemplate(ctx *parser.TemplateContext) interface{} {
	v.Node = model.NewRootNode(ctx.GetStart().GetTokenIndex(), ctx.GetStop().GetTokenIndex())
	v.visitChildren(ctx)
	return nil
}

func (v *PlaceholdersVisitor) VisitContent(ctx *parser.ContentContext) interface{} {
	v.visitChildren(ctx)
	return nil
}

func (v *PlaceholdersVisitor) VisitText(ctx *parser.TextContext) interface{} {
	node := model.NewTextNode(ctx.GetStart().GetTokenIndex(), ctx.GetStop().GetTokenIndex(), v.Node)
	node.Value = &model.Text{Text: ctx.GetText()}
	v.Node.AddChild(node)
	return nil
}

func (v *PlaceholdersVisitor) VisitAttribute(ctx *parser.AttributeContext) interface{} {
	placeholder, ok := v.Node.Value.(*model.ParsedPlaceholder)
	if !ok {
		// Handle the error - log, return error, or panic with context
		panic("expected *model.ParsedPlaceholder but got different type")
	}
	text := strings.Trim(ctx.GetValue().GetText(), "\"'")
	attr := &model.Attr{Name: ctx.GetName().GetText(), Value: text}
	placeholder.Attrs = append(placeholder.Attrs, attr)
	return nil
}

func (v *PlaceholdersVisitor) VisitPlaceholder(ctx *parser.PlaceholderContext) interface{} {
	v.visitChildren(ctx)
	return nil
}

func (v *PlaceholdersVisitor) VisitSimplePlaceholder(ctx *parser.SimplePlaceholderContext) interface{} {

	node := model.NewPlaceholderNode(ctx.GetStart().GetTokenIndex(), ctx.GetStop().GetTokenIndex())
	node.Value = &model.ParsedPlaceholder{
		Name:    ctx.GetPlaceholderName().GetText(),
		IsBlock: false,
	}

	v.Node.AddChild(node)

	// Move down to the new node

	if ctx.AllAttribute() != nil {
		v.Node = node
		attrCount := len(ctx.AllAttribute())
		node.Value.(*model.ParsedPlaceholder).Attrs = make([]*model.Attr, 0, attrCount)
		for _, attr := range ctx.AllAttribute() {
			attr.Accept(v)
		}
		// Pop back to the parent node
		v.Node = v.Node.Parent
	}

	return nil
}
func (v *PlaceholdersVisitor) VisitBlockPlaceholderStart(ctx *parser.BlockPlaceholderStartContext) interface{} {
	placeholder, ok := v.Node.Value.(*model.ParsedPlaceholder)
	if !ok {
		// Handle the error - log, return error, or panic with context
		panic("expected *model.ParsedPlaceholder but got different type")
	}
	placeholder.Name = ctx.GetPlaceholderName().GetText()

	if ctx.AllAttribute() != nil {
		attrCount := len(ctx.AllAttribute())
		v.Node.Value.(*model.ParsedPlaceholder).Attrs = make([]*model.Attr, 0, attrCount)
		for _, attr := range ctx.AllAttribute() {
			attr.Accept(v)
		}
	}

	return nil
}

func (v *PlaceholdersVisitor) VisitBlockPlaceholderContent(ctx *parser.BlockPlaceholderContentContext) interface{} {
	//fmt.Printf("VisitBlockPlaceholderContent: %s \n", ctx.GetText())
	placeholder, ok := v.Node.Value.(*model.ParsedPlaceholder)
	if !ok {
		// Handle the error - log, return error, or panic with context
		panic("expected *model.ParsedPlaceholder but got different type")
	}

	placeholder.ContentStart = ctx.GetStart().GetTokenIndex()
	placeholder.ContentEnd = ctx.GetStop().GetTokenIndex()

	v.visitChildren(ctx)

	return nil
}

func (v *PlaceholdersVisitor) VisitBlockPlaceholder(ctx *parser.BlockPlaceholderContext) interface{} {

	if ctx.BlockPlaceholderStart().GetPlaceholderName().GetText() != ctx.BlockPlaceholderEnd().GetPlaceholderName().GetText() {
		// Handle the error - log, return error, or panic with context
		panic("mismatched block placeholder names: |" + ctx.BlockPlaceholderStart().GetPlaceholderName().GetText() + "| != |" + ctx.BlockPlaceholderEnd().GetPlaceholderName().GetText() + "|")
	}

	node := model.NewBlockNode(ctx.GetStart().GetTokenIndex(), ctx.GetStop().GetTokenIndex())
	node.Value = &model.ParsedPlaceholder{
		IsBlock: true,
	}

	v.Node.AddChild(node)

	// Move down to the new node
	v.Node = node
	v.visitChildren(ctx)
	v.Node = v.Node.Parent

	return nil
}
