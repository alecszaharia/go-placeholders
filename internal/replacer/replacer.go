package replacer

import (
	parser "antrl/internal/grammar"
	"antrl/internal/model"
	"antrl/internal/placeholder"
	"antrl/internal/visitor"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

type Replacer struct {
	group  *placeholder.Group       // group of placeholders
	tokens *antlr.CommonTokenStream // all tokens from the lexer
}

func New(group *placeholder.Group) *Replacer {
	return &Replacer{
		group: group,
	}
}

type bailErrorListener struct{ *antlr.DefaultErrorListener }

func (l bailErrorListener) SyntaxError(r antlr.Recognizer, s interface{}, line, col int, msg string, e antlr.RecognitionException) {
	panic(fmt.Errorf("syntax error at %d:%d: %s", line, col, msg))
}

func (r *Replacer) Replace(inputStr string) string {
	input := antlr.NewInputStream(inputStr)
	lexer := parser.NewPlaceholderLexer(input)
	r.tokens = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPlaceholderParser(r.tokens)
	v := visitor.NewPlaceholdersVisitor()
	p.RemoveErrorListeners() // remove default ConsoleErrorListener
	p.AddErrorListener(bailErrorListener{&antlr.DefaultErrorListener{}})

	// trigger parsing and visiting the parse tree
	// we start from the root rule which is 'template' in our grammar
	// this will populate the visitor's Node field with the parse tree
	template := p.Template()
	template.Accept(v)

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
				sc := &model.Context{CurrentNode: node, ParentContext: c}
				if content, err := r.replaceNodeWithPlaceholderValue(sc); err == nil {
					results[i] = content
				} else {
					results[i] = r.tokens.GetTextFromInterval(antlr.Interval{Start: node.Start, Stop: node.End})
				}
			}()

		case model.NodeBlock:
			wg.Add(1)
			go func() {
				defer wg.Done()
				sc := &model.Context{CurrentNode: node, ParentContext: c}
				if content, err := r.replaceBlockNodeWithPlaceholderValue(sc); err == nil {
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

func (r *Replacer) replaceNodeWithPlaceholderValue(c *model.Context) (string, error) {
	node := c.CurrentNode
	p := r.group.GetPlaceholderByName(node.GetPlaceholder().Name)

	if p == nil {
		return "", errors.New("placeholder not found") // or some error handling
	}

	content, err := p.Handler(c, nil)

	if err != nil && p.FallbackHandler != nil {
		content, err = p.FallbackHandler(c, nil)
	}
	return content, err
}

func (r *Replacer) replaceBlockNodeWithPlaceholderValue(c *model.Context) (string, error) {
	node := c.CurrentNode
	p := r.group.GetPlaceholderByName(node.GetPlaceholder().Name)

	if p == nil {
		return "", errors.New("placeholder not found") // or some error handling
	}

	strb := &strings.Builder{}

	_, err := p.Handler(c, func(c *model.Context) {
		r.replaceNodes(c, strb)
	})
	if err != nil {
		return "", err
	}

	return strb.String(), nil
}
