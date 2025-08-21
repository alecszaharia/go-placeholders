package visitor

import (
	parser "antrl/gen"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type ReplaceContext struct {
	ParentContext *ReplaceContext
	Placeholder   *Placeholder
}

type Placeholder struct {
	Name     string
	Content  string
	HasEnd   bool
	Attrs    []Attr
	Parent   *Placeholder
	Children []*Placeholder
}

type Attr struct {
	Name  string
	Value string
}

func NewPlaceholdersVisitor(tokens *antlr.CommonTokenStream) *PlaceholdersVisitor {
	return &PlaceholdersVisitor{
		BasePlaceholderParserVisitor: &parser.BasePlaceholderParserVisitor{},
		Rewriter:                     antlr.NewTokenStreamRewriter(tokens),
	}
}

type PlaceholdersVisitor struct {
	*parser.BasePlaceholderParserVisitor
	Rewriter         *antlr.TokenStreamRewriter
	Placeholders     []Placeholder
	PlaceholderTrees []*Placeholder
	Placeholder      *Placeholder
}

func (v *PlaceholdersVisitor) VisitTemplate(ctx *parser.TemplateContext) interface{} {
	fmt.Printf("Visiting Template: %s\n", ctx.GetText())
	ctx.Content().Accept(v)
	return v.VisitChildren(ctx) // Visit the first placeholder
}

func (v *PlaceholdersVisitor) VisitContent(ctx *parser.ContentContext) interface{} {
	fmt.Printf("Visiting Content: %s\n", ctx.GetText())
	for _, child := range ctx.AllPlaceholder() {
		child.Accept(v)
	}

	return v.VisitChildren(ctx) // Visit the first placeholder
}

//func (v *PlaceholdersVisitor) VisitAttributeList(ctx *parser.AttributeListContext) interface{} {
//	fmt.Printf("Visiting VisitAttributeList: %s\n", ctx.GetText())
//	attrs := make([]Attr, 0, 20) // let's assume we have at most 20 attributes and preallocate the slice
//	for _, child := range ctx.AllAttribute() {
//		//fmt.Printf("Attribute: %s = %s\n", child.ID().GetText(), child.AttrValue().GetText())
//		attrs = append(attrs, Attr{Name: child.ID().GetText(), Value: child.AttrValue().GetText()})
//	}
//	v.Placeholder.Attrs = attrs
//	return v.VisitChildren(ctx) // Visit the first placeholder
//}

//func (v *PlaceholdersVisitor) VisitAttribute(ctx *parser.AttributeContext) interface{} {
//	fmt.Printf("Visiting VisitAttribute: %s = %s\n", ctx.ID().GetText(), ctx.AttrValue().GetText())
//	return v.VisitChildren(ctx) // Visit the first placeholder
//}

//func (v *PlaceholdersVisitor) VisitAttrValue(ctx *parser.AttrValueContext) interface{} {
//	fmt.Printf("Visiting VisitAttrValue: %s\n", ctx.GetText())
//	return v.VisitChildren(ctx) // Visit the first placeholder
//}

//func (v *PlaceholdersVisitor) VisitText(ctx *parser.TextContext) interface{} {
//	//fmt.Printf("Visiting Text: %s\n", ctx.GetText())
//	return v.VisitChildren(ctx) // Visit the first placeholder
//}

func (v *PlaceholdersVisitor) VisitPlaceholder(ctx *parser.PlaceholderContext) interface{} {

	parentPlaceholder := v.Placeholder
	b := Placeholder{
		Name:     ctx.GetPlaceholderName().GetText(),
		HasEnd:   ctx.GetPlaceholderEndName() != nil,
		Parent:   parentPlaceholder,
		Children: make([]*Placeholder, 0, 30), // Preallocate for up to 30 children
	}

	if ctx.Content() != nil {
		b.Content = ctx.Content().GetText()
		ctx.Content().Accept(v)
	}

	fmt.Printf("VisitPlaceholder: %s and content: %s \n", ctx.GetText(), b.Content)
	v.Placeholder = &b

	//

	//if ctx.AttributeList() != nil {
	//	// Temporarily set v.Placeholder for attribute parsing
	//	ctx.AttributeList().Accept(v)
	//}

	// Restore visitor's block pointer for recursion
	if parentPlaceholder != nil {
		parentPlaceholder.Children = append(parentPlaceholder.Children, &b)
	}

	v.Placeholders = append(v.Placeholders, b)
	if b.Parent == nil {
		v.PlaceholderTrees = append(v.PlaceholderTrees, &b)
	}

	// Recursively visit content (children blocks)

	//v.Rewriter.ReplaceDefault(
	//	ctx.GetStart().GetTokenIndex(),
	//	ctx.GetStop().GetTokenIndex(),
	//	"",
	//)

	return v.VisitChildren(ctx)
}
