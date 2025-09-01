package replacer

import (
	parser "antrl/internal/grammar"
	"antrl/internal/model"
	"antrl/internal/placeholder"
	"antrl/internal/visitor"
	"strings"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

type Replacer struct {
	group    *placeholder.Group       // group of placeholders
	tokens   *antlr.CommonTokenStream // all tokens from the lexer
	template parser.ITemplateContext
}

func New(group *placeholder.Group) *Replacer {
	return &Replacer{
		group: group,
	}
}

func (r *Replacer) Replace(inputStr string) string {
	input := antlr.NewInputStream(inputStr)
	lexer := parser.NewPlaceholderLexer(input)
	r.tokens = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPlaceholderParser(r.tokens)
	v := visitor.NewPlaceholdersVisitor()

	// trigger parsing and visiting the parse tree
	// we start from the root rule which is 'template' in our grammar
	// this will populate the visitor's Node field with the parse tree
	r.template = p.Template()
	r.template.Accept(v)

	ctx := &model.Context{CurrentNode: v.Node}
	b := &strings.Builder{}
	r.replaceNodes(ctx, b)

	return b.String()
}

func (r *Replacer) replaceNodes(c *model.Context, b *strings.Builder) {
	wg := &sync.WaitGroup{}
	results := make([]string, len(c.CurrentNode.Children))
	for i, node := range c.CurrentNode.Children {
		switch node.Type {

		case model.NodeText:
			results[i] = r.tokens.GetTextFromInterval(antlr.Interval{Start: node.Start, Stop: node.End})
		case model.NodePlaceholder:
			wg.Add(1)
			go func() {
				defer wg.Done()
				if content := r.replaceNodeWithPlaceholderValue(node); content != "" {
					results[i] = content
				} else {
					results[i] = r.tokens.GetTextFromInterval(antlr.Interval{Start: node.Start, Stop: node.End})
				}
			}()

		case model.NodeBlock:
			wg.Add(1)
			go func() {
				defer wg.Done()
				if content := r.replaceBlockNodeWithPlaceholderValue(c, node); content != "" {
					results[i] = content
				} else {
					results[i] = r.tokens.GetTextFromInterval(antlr.Interval{Start: node.Start, Stop: node.End})
				}
			}()

		}
	}
	wg.Wait()

	for _, r := range results {
		b.WriteString(r)
	}
}

func (r *Replacer) replaceNodeWithPlaceholderValue(node *model.Node) string {
	p := r.group.GetPlaceholderByName(node.Value.(*model.Placeholder).Name)

	if p == nil {
		return "" // or some error handling
	}

	content, err := p.Handler(node)

	if err != nil && p.FallbackHandler != nil {
		content, err = p.FallbackHandler(node)
	}

	return content
}

func (r *Replacer) replaceBlockNodeWithPlaceholderValue(c *model.Context, node *model.Node) string {

	sc := &model.Context{CurrentNode: node, ParentContext: c}

	p := r.group.GetPlaceholderByName(node.Value.(*model.Placeholder).Name)

	if p == nil {
		return "" // or some error handling
	}

	strb := &strings.Builder{}
	for i := 0; i < 5; i++ {
		r.replaceNodes(sc, strb)
	}

	return strb.String()
}
