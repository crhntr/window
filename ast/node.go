package template

import (
	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
)

func nodeType(nodeType html.NodeType) dom.NodeType {
	switch nodeType {
	case html.TextNode:
		return dom.NodeTypeText
	case html.DocumentNode:
		return dom.NodeTypeDocument
	case html.ElementNode:
		return dom.NodeTypeElement
	case html.CommentNode:
		return dom.NodeTypeComment
	case html.DoctypeNode:
		return dom.NodeTypeDocumentType
	default:
		fallthrough
	case html.ErrorNode, html.RawNode:
		return dom.NodeTypeUnknown
	}
}
