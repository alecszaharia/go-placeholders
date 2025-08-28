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
	group    *placeholder.Group           // group of placeholders
	tokens   *antlr.CommonTokenStream     // all tokens from the lexer
	rewriter *antlr.TokenStreamRewriter   // token rewriter to perform replacements
	visitor  *visitor.PlaceholdersVisitor // token rewriter to perform replacements
	template parser.ITemplateContext
}

func New(group *placeholder.Group) *Replacer {
	return &Replacer{
		group: group,
	}
}

func (r *Replacer) Prepare(inputStr string) {
	input := antlr.NewInputStream(inputStr)
	lexer := parser.NewPlaceholderLexer(input)
	r.tokens = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	r.rewriter = antlr.NewTokenStreamRewriter(r.tokens)
	p := parser.NewPlaceholderParser(r.tokens)

	//p.RemoveErrorListeners()

	r.visitor = visitor.NewPlaceholdersVisitor()

	r.template = p.Template()
}

func (r *Replacer) Replace() string {

	r.template.Accept(r.visitor)

	v := r.visitor
	rewriter := r.rewriter
	tokens := r.tokens

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for _, node := range v.Nodes {
		switch node.Type {

		case model.NodePlaceholder:
			go func() {
				defer wg.Done()
				if content := r.replaceNodeWithPlaceholderValue(tokens, node); content != "" {
					mu.Lock()
					rewriter.ReplaceDefault(
						node.Start,
						node.End,
						content,
					)
					mu.Unlock()
				}
			}()
			wg.Add(1)

		case model.NodeBlock:
			go func() {
				defer wg.Done()
				if content := r.replaceBlockNodeWithPlaceholderValue(tokens, node); content != "" {
					mu.Lock()
					rewriter.ReplaceDefault(
						node.Start,
						node.End,
						content,
					)
					mu.Unlock()
				}
			}()
			wg.Add(1)
		}
	}

	wg.Wait()

	return rewriter.GetTextDefault()
}

func (r *Replacer) ReplaceSync() string {

	r.template.Accept(r.visitor)

	v := r.visitor
	rewriter := r.rewriter
	tokens := r.tokens

	mu := &sync.Mutex{}
	for _, node := range v.Nodes {
		switch node.Type {

		case model.NodePlaceholder:
			if content := r.replaceNodeWithPlaceholderValue(tokens, node); content != "" {
				mu.Lock()
				rewriter.ReplaceDefault(
					node.Start,
					node.End,
					content,
				)
				mu.Unlock()
			}

		case model.NodeBlock:
			if content := r.replaceBlockNodeWithPlaceholderValueSync(tokens, node); content != "" {
				mu.Lock()
				rewriter.ReplaceDefault(
					node.Start,
					node.End,
					content,
				)
				mu.Unlock()
			}
		}
	}

	return rewriter.GetTextDefault()
}

func (r *Replacer) replaceNodeWithPlaceholderValue(tokens *antlr.CommonTokenStream, node *model.Node) string {
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

func (r *Replacer) replaceBlockNodeWithPlaceholderValue(tokens *antlr.CommonTokenStream, node *model.Node) string {

	p := r.group.GetPlaceholderByName(node.Value.(*model.Placeholder).Name)

	if p == nil {
		return "" // or some error handling
	}

	interval := antlr.Interval{Start: node.Value.(*model.Placeholder).ContentStart, Stop: node.Value.(*model.Placeholder).ContentEnd}
	contentInput := tokens.GetTextFromInterval(interval)

	p.Handler(node)

	res := strings.Builder{}

	wg := &sync.WaitGroup{}

	repl := New(r.group)
	repl.Prepare(contentInput)
	mu := &sync.Mutex{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			res.WriteString(repl.Replace())
			mu.Unlock()
		}()
	}

	wg.Wait()

	return res.String()
}

func (r *Replacer) replaceBlockNodeWithPlaceholderValueSync(tokens *antlr.CommonTokenStream, node *model.Node) string {

	p := r.group.GetPlaceholderByName(node.Value.(*model.Placeholder).Name)

	if p == nil {
		return "" // or some error handling
	}

	interval := antlr.Interval{Start: node.Value.(*model.Placeholder).ContentStart, Stop: node.Value.(*model.Placeholder).ContentEnd}
	contentInput := tokens.GetTextFromInterval(interval)

	p.Handler(node)

	res := strings.Builder{}

	repl := New(r.group)
	repl.Prepare(contentInput)
	mu := &sync.Mutex{}
	for i := 0; i < 2; i++ {
		mu.Lock()
		res.WriteString(repl.ReplaceSync())
		mu.Unlock()
	}

	return res.String()

}
