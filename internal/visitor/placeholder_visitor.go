package visitor

import (
	parser "antrl/internal/grammar"
	"math/rand"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/google/uuid"
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

func (p *Placeholder) GetValue() string {
	rand.Seed(time.Now().UnixNano())
	//time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond) // ensure different seed for each call
	uid := uuid.New()
	return "PLACEHOLDER_" + p.Name + "_" + uid.String() + "_END"
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

func (v *PlaceholdersVisitor) VisitAttribute(ctx *parser.AttributeContext) interface{} {
	//fmt.Printf("VisitAttribute: %s\n", ctx.GetText())

	v.Placeholder.Attrs = append(v.Placeholder.Attrs, Attr{Name: ctx.GetName().GetText(), Value: ctx.GetValue().GetText()})

	return v.VisitChildren(ctx) // Visit the first placeholder
}

func (v *PlaceholdersVisitor) VisitTemplate(ctx *parser.TemplateContext) interface{} {
	//fmt.Printf("Visiting Template: %s\n", ctx.GetText())
	ctx.Content().Accept(v)
	return v.VisitChildren(ctx) // Visit the first placeholder
}

func (v *PlaceholdersVisitor) VisitContent(ctx *parser.ContentContext) interface{} {
	//fmt.Printf("Visiting Content: %s\n", ctx.GetText())
	for _, child := range ctx.AllText() {
		child.Accept(v)
	}
	for _, child := range ctx.AllPlaceholder() {
		child.Accept(v)
	}

	return v.VisitChildren(ctx) // Visit the first placeholder
}

func (v *PlaceholdersVisitor) VisitPlaceholder(ctx *parser.PlaceholderContext) interface{} {

	if ctx.SimplePlaceholder() != nil {
		ctx.SimplePlaceholder().Accept(v)
	}
	if ctx.BlockPlaceholder() != nil {
		ctx.BlockPlaceholder().Accept(v)
	}

	return nil
}

func (v *PlaceholdersVisitor) VisitSimplePlaceholder(ctx *parser.SimplePlaceholderContext) interface{} {
	//fmt.Printf("VisitSimplePlaceholder: %s \n", ctx.GetText())

	parentPlaceholder := v.Placeholder

	b := Placeholder{
		Name:   ctx.GetPlaceholderName().GetText(),
		HasEnd: false,
		Parent: parentPlaceholder,
	}

	if ctx.AllAttribute() != nil {
		// Temporarily set v.Placeholder for attribute parsing
		v.Placeholder = &b
		for _, attr := range ctx.AllAttribute() {
			attr.Accept(v)
		}
		v.Placeholder = parentPlaceholder
	}

	if parentPlaceholder != nil {
		parentPlaceholder.Children = append(parentPlaceholder.Children, &b)
	}

	v.Placeholders = append(v.Placeholders, b)

	if b.Parent == nil {
		v.PlaceholderTrees = append(v.PlaceholderTrees, &b)
	}

	// Recursively visit content (children blocks)
	v.Rewriter.ReplaceDefault(
		ctx.GetStart().GetTokenIndex(),
		ctx.GetStop().GetTokenIndex(),
		b.GetValue(),
	)

	return v.VisitChildren(ctx)
}

func (v *PlaceholdersVisitor) VisitBlockPlaceholderStart(ctx *parser.BlockPlaceholderStartContext) interface{} {
	//fmt.Printf("VisitBlockPlaceholderStart: %s \n", ctx.GetText())
	v.Placeholder.Name = ctx.GetPlaceholderName().GetText()
	v.Placeholder.HasEnd = true

	return nil
}
func (v *PlaceholdersVisitor) VisitBlockPlaceholderEnd(ctx *parser.BlockPlaceholderEndContext) interface{} {
	//fmt.Printf("VisitBlockPlaceholderEnd: %s \n", ctx.GetText())
	return nil
}
func (v *PlaceholdersVisitor) VisitBlockPlaceholderContent(ctx *parser.BlockPlaceholderContentContext) interface{} {
	//fmt.Printf("VisitBlockPlaceholderContent: %s \n", ctx.GetText())
	v.Placeholder.Content = ctx.GetText()
	ctx.Content().Accept(v)
	return nil
}

func (v *PlaceholdersVisitor) VisitBlockPlaceholder(ctx *parser.BlockPlaceholderContext) interface{} {

	prevBlock := v.Placeholder

	v.Placeholder = &Placeholder{
		HasEnd: false,
		Parent: prevBlock,
	}

	v.Placeholders = append(v.Placeholders, *v.Placeholder)

	ctx.BlockPlaceholderStart().Accept(v)
	ctx.BlockPlaceholderContent().Accept(v)
	ctx.BlockPlaceholderEnd().Accept(v)

	if prevBlock != nil {
		prevBlock.Children = append(prevBlock.Children, v.Placeholder)
	}

	if v.Placeholder.Parent == nil {
		v.PlaceholderTrees = append(v.PlaceholderTrees, v.Placeholder)
	}

	v.Rewriter.ReplaceDefault(
		ctx.GetStart().GetTokenIndex(),
		ctx.GetStop().GetTokenIndex(),
		v.Placeholder.GetValue(),
	)

	v.Placeholder = prevBlock

	return nil
	//
	//parentPlaceholder := v.Placeholder
	//b := Placeholder{
	//	Name:     ctx.GetPlaceholderName().GetText(),
	//	HasEnd:   true,
	//	Content:  ctx.Content().GetText(),
	//	Parent:   parentPlaceholder,
	//	Children: make([]*Placeholder, 0, 100), // Preallocate space for children
	//}
	//
	//fmt.Printf("VisitBlockPlaceholder: %s \n", ctx.GetText())
	//
	//ctx.Content().Accept(v)
	//
	////if ctx.AttributeList() != nil {
	////	// Temporarily set v.Placeholder for attribute parsing
	////	ctx.AttributeList().Accept(v)
	////}
	//
	//// Restore visitor's block pointer for recursion
	//
	//if parentPlaceholder != nil {
	//	parentPlaceholder.Children = append(parentPlaceholder.Children, &b)
	//}
	//
	//v.Placeholders = append(v.Placeholders, b)
	//v.Placeholder = &b
	//
	//// Recursively visit content (children blocks)
	////v.Rewriter.ReplaceDefault(
	////	ctx.GetStart().GetTokenIndex(),
	////	ctx.GetStop().GetTokenIndex(),
	////	"",
	////)
	//
	//return v.VisitChildren(ctx)
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
