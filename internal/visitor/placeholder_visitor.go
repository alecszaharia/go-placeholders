package visitor

import (
	parser "antrl/internal/grammar"
	"antrl/internal/model"

	"github.com/antlr4-go/antlr/v4"
)

type PlaceholdersVisitor struct {
	*parser.BasePlaceholderParserVisitor
	Nodes       []*model.Node
	currentNode *model.Node
}

func NewPlaceholdersVisitor() *PlaceholdersVisitor {
	return &PlaceholdersVisitor{
		BasePlaceholderParserVisitor: &parser.BasePlaceholderParserVisitor{},
		Nodes:                        make([]*model.Node, 0, 160),
	}
}

func (v *PlaceholdersVisitor) pushRootIfTop(n *model.Node) {
	if n.Parent == nil {
		v.Nodes = append(v.Nodes, n)
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
	v.visitChildren(ctx)

	return nil
}

func (v *PlaceholdersVisitor) VisitContent(ctx *parser.ContentContext) interface{} {
	v.visitChildren(ctx)
	return nil
}

func (v *PlaceholdersVisitor) VisitText(ctx *parser.TextContext) interface{} {
	node := model.MakeTextNode(ctx.GetStart().GetTokenIndex(), ctx.GetStop().GetTokenIndex(), v.currentNode)

	// temporary
	//node.Value.(*model.Text).Text = ctx.GetText()

	v.pushRootIfTop(node)

	return nil
}

func (v *PlaceholdersVisitor) VisitAttribute(ctx *parser.AttributeContext) interface{} {
	placeholder, ok := v.currentNode.Value.(*model.Placeholder)
	if !ok {
		// Handle the error - log, return error, or panic with context
		panic("expected *model.Placeholder but got different type")
	}
	attr := &model.Attr{Name: ctx.GetName().GetText(), Value: ctx.GetValue().GetText()}
	placeholder.Attrs = append(placeholder.Attrs, attr)
	return nil
}

func (v *PlaceholdersVisitor) VisitPlaceholder(ctx *parser.PlaceholderContext) interface{} {
	v.visitChildren(ctx)
	return nil
}

func (v *PlaceholdersVisitor) VisitSimplePlaceholder(ctx *parser.SimplePlaceholderContext) interface{} {

	node := model.MakePlaceholderNode(ctx.GetStart().GetTokenIndex(), ctx.GetStop().GetTokenIndex(), v.currentNode)
	v.currentNode = node

	node.Value = &model.Placeholder{
		Name:    ctx.GetPlaceholderName().GetText(),
		IsBlock: false,
	}

	if ctx.AllAttribute() != nil {
		attrCount := len(ctx.AllAttribute())

		node.Value.(*model.Placeholder).Attrs = make([]*model.Attr, 0, attrCount)

		for _, attr := range ctx.AllAttribute() {
			attr.Accept(v)
		}
	}

	// Pop back to the parent node
	v.currentNode = v.currentNode.Parent

	v.pushRootIfTop(node)

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

func (v *PlaceholdersVisitor) VisitBlockPlaceholderContent(ctx *parser.BlockPlaceholderContentContext) interface{} {
	//fmt.Printf("VisitBlockPlaceholderContent: %s \n", ctx.GetText())
	placeholder, ok := v.currentNode.Value.(*model.Placeholder)
	if !ok {
		// Handle the error - log, return error, or panic with context
		panic("expected *model.Placeholder but got different type")
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

	n := model.MakeBlockNode(ctx.GetStart().GetTokenIndex(), ctx.GetStop().GetTokenIndex(), v.currentNode)
	n.Value = &model.Placeholder{
		IsBlock: true,
	}

	v.currentNode = n
	v.visitChildren(ctx)
	v.pushRootIfTop(n)

	// Pop back to the parent node
	v.currentNode = v.currentNode.Parent

	return nil
}
