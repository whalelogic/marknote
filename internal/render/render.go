package render

import (
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// MarkdownToHTML converts markdown content to HTML with Mermaid support
func MarkdownToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags:          htmlFlags,
		RenderNodeHook: mermaidRenderHook,
	}

	renderer := html.NewRenderer(opts)
	return markdown.Render(doc, renderer)
}

func mermaidRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if codeBlock, ok := node.(*ast.CodeBlock); ok {
		if string(codeBlock.Info) == "mermaid" {
			if entering {
				io.WriteString(w, "<pre class=\"mermaid\">")
				w.Write(codeBlock.Literal)
				io.WriteString(w, "</pre>")
			}
			return ast.SkipChildren, true
		}
	}
	return ast.GoToNext, false
}
