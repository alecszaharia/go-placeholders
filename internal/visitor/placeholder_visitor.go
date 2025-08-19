package visitor

import (
	parser "antrl/gen"
)

type Block struct {
	Name    string
	Content string
	HasEnd  bool
	Attrs   []Attr
}

type Attr struct {
	Name  string
	Value string
}

func NewPlaceholdersVisitor() *PlaceholdersVisitor {
	return &PlaceholdersVisitor{
		BasePlaceholderParserVisitor: &parser.BasePlaceholderParserVisitor{},
	}
}

type PlaceholdersVisitor struct {
	*parser.BasePlaceholderParserVisitor
	Blocks []Block
	Block  *Block
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
	//for _, child := range ctx.AllText() {
	//	child.Accept(v)
	//}
	return v.VisitChildren(ctx) // Visit the first placeholder
}

func (v *PlaceholdersVisitor) VisitAttributeList(ctx *parser.AttributeListContext) interface{} {
	//fmt.Printf("Visiting Text: %s\n", ctx.GetText())
	attrs := make([]Attr, 0, 20) // let's assume we have at most 20 attributes and preallocate the slice
	for _, child := range ctx.AllAttribute() {
		//fmt.Printf("Attribute: %s = %s\n", child.ID().GetText(), child.AttrValue().GetText())
		attrs = append(attrs, Attr{Name: child.ID().GetText(), Value: child.AttrValue().GetText()})
	}
	v.Block.Attrs = attrs
	return v.VisitChildren(ctx) // Visit the first placeholder
}

//func (v *PlaceholdersVisitor) VisitAttribute(ctx *parser.AttributeContext) interface{} {
//	fmt.Printf("Visiting VisitAttribute: %s = %s\n", ctx.ID().GetText(), ctx.AttrValue().GetText())
//	return v.VisitChildren(ctx) // Visit the first placeholder
//}

//func (v *PlaceholdersVisitor) VisitAttrValue(ctx *parser.AttrValueContext) interface{} {
//	fmt.Printf("Visiting VisitAttrValue: %s\n", ctx.GetText())
//	return v.VisitChildren(ctx) // Visit the first placeholder
//}

func (v *PlaceholdersVisitor) VisitText(ctx *parser.TextContext) interface{} {
	//fmt.Printf("Visiting Text: %s\n", ctx.GetText())
	return v.VisitChildren(ctx) // Visit the first placeholder
}

func (v *PlaceholdersVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {

	//fmt.Printf("VisitBlock: %s\n", ctx.GetText())

	b := Block{HasEnd: false}
	v.Block = &b
	b.Name = ctx.GetBlockName().GetText()

	if ctx.Content() != nil {
		b.Content = ctx.Content().GetText()
		ctx.Content().Accept(v)
	}

	if ctx.GetBlockEndName() != nil {
		b.HasEnd = true
	}

	if ctx.AttributeList() != nil {
		ctx.AttributeList().Accept(v)
	}

	v.Blocks = append(v.Blocks, b)
	return v.VisitChildren(ctx)
}
