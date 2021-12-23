package ast

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
)

func isConnected(node *html.Node) bool {
	p := node.Parent
	for p != nil {
		if p.Type == html.DocumentNode {
			return true
		}
		p = p.Parent
	}
	return false
}

func ownerDocument(node *html.Node) dom.Document {
	p := node.Parent
	for p != nil {
		if p.Type == html.DocumentNode {
			return &Document{node: p}
		}
		p = p.Parent
	}
	return nil
}
func parentNode(node *html.Node) dom.Node           { return htmlNodeToDomNode(node.Parent) }
func parentElement(node *html.Node) dom.Element     { return htmlNodeToDomElement(node.Parent) }
func hasChildNodes(node *html.Node) bool            { return node.FirstChild != nil }
func childNodes(node *html.Node) dom.NodeList       { return (*SiblingNodeList)(node.FirstChild) }
func firstChild(node *html.Node) dom.ChildNode      { return htmlNodeToDomChildNode(node.FirstChild) }
func lastChild(node *html.Node) dom.ChildNode       { return htmlNodeToDomChildNode(node.LastChild) }
func previousSibling(node *html.Node) dom.ChildNode { return htmlNodeToDomChildNode(node.PrevSibling) }
func nextSibling(node *html.Node) dom.ChildNode     { return htmlNodeToDomChildNode(node.NextSibling) }

func textContent(node *html.Node) string {
	var buf bytes.Buffer
	recursiveTextContent(&buf, node)
	return buf.String()
}

func recursiveTextContent(sw io.StringWriter, n *html.Node) {
	if n.Type == html.TextNode {
		_, err := sw.WriteString(n.Data)
		if err != nil {
			panic(err)
		}
	}
	c := n.FirstChild
	for c != nil {
		recursiveTextContent(sw, c)
		c = c.NextSibling
	}
}

func cloneNode(node *html.Node, deep bool) dom.Node {
	if deep {
		var buf bytes.Buffer
		err := html.Render(&buf, node)
		if err != nil {
			panic(err)
		}
		cl, err := html.ParseFragment(&buf, node.Parent)
		if err != nil {
			panic(err)
		}
		if len(cl) == 0 {
			return nil
		}
		return htmlNodeToDomNode(cl[0])
	}
	c := html.Node{
		Type:      node.Type,
		Namespace: node.Namespace,
		Data:      node.Data,
		DataAtom:  node.DataAtom,
	}
	if node.Attr != nil {
		c.Attr = make([]html.Attribute, len(node.Attr))
		for i, at := range node.Attr {
			c.Attr[i].Key = at.Key
			c.Attr[i].Val = at.Val
			c.Attr[i].Namespace = at.Namespace
		}
	}
	return htmlNodeToDomNode(&c)
}

func isSameNode(node *html.Node, other dom.Node) bool {
	if node == nil || other == nil {
		return false
	}
	n := domNodeToHTMLNode(other)
	return n != nil && node == n
}

func contains(node *html.Node, other dom.Node) bool {
	o := domNodeToHTMLNode(other)
	if o == nil {
		return false
	}

	found := false
	walkNodes(node, func(n *html.Node) bool {
		found = n == o
		return found
	})

	return found
}

func insertBefore(parent *html.Node, node, child dom.ChildNode) dom.ChildNode {
	n := domNodeToHTMLNode(node)
	c := domNodeToHTMLNode(child)
	if n.Parent != nil {
		n.Parent.RemoveChild(n)
	}
	parent.InsertBefore(n, c)
	return htmlNodeToDomChildNode(n)
}

func appendChild(parent *html.Node, node dom.ChildNode) dom.ChildNode {
	n := domNodeToHTMLNode(node)
	if n.Parent != nil {
		n.Parent.RemoveChild(n)
	}
	parent.AppendChild(n)
	return htmlNodeToDomChildNode(n)
}

func replaceChild(parent *html.Node, node, child dom.ChildNode) dom.ChildNode {
	n := domNodeToHTMLNode(node)
	c := domNodeToHTMLNode(child)
	if c.Parent != parent {
		panic("browser: ReplaceChild called for an attached child node")
	}
	if c.PrevSibling != nil {
		c.PrevSibling.NextSibling = n
	}
	if c.NextSibling != nil {
		c.NextSibling.PrevSibling = n
	}
	if parent.FirstChild == c {
		parent.FirstChild = n
	}
	if parent.LastChild == c {
		parent.LastChild = n
	}
	n.PrevSibling = c.PrevSibling
	n.NextSibling = c.NextSibling
	n.Parent = c.Parent

	c.PrevSibling = nil
	c.NextSibling = nil
	c.Parent = nil

	return htmlNodeToDomChildNode(c)
}

func removeChild(parent *html.Node, node dom.ChildNode) dom.ChildNode {
	n := domNodeToHTMLNode(node)
	parent.RemoveChild(n)
	return htmlNodeToDomChildNode(n)
}

func children(parent *html.Node) dom.ElementCollection {
	return SiblingElements{firstChild: parent.FirstChild}
}

func firstElementChild(node *html.Node) dom.Element {
	child := node.FirstChild
	for child != nil {
		if child.Type == html.ElementNode {
			return &Element{node: child}
		}
		child = child.NextSibling
	}
	return nil
}

func lastElementChild(node *html.Node) dom.Element {
	child := node.LastChild
	for child != nil {
		if child.Type == html.ElementNode {
			return &Element{node: child}
		}
		child = child.PrevSibling
	}
	return nil
}

func childElementCount(node *html.Node) int {
	var (
		result = 0
		child  = node.FirstChild
	)
	for child != nil {
		if child.Type == html.ElementNode {
			result++
		}
		child = child.NextSibling
	}
	return result
}

func prependNodes(node *html.Node, nodes []dom.ChildNode) {
	for i := range nodes {
		dn := nodes[len(nodes)-1-i]
		n := domNodeToHTMLNode(dn)

		fc := node.FirstChild
		if fc != nil {
			fc.PrevSibling = n
			n.NextSibling = fc
		}
		n.Parent = node
		node.FirstChild = n

		if node.LastChild == nil {
			node.LastChild = n
		}
	}
}

func appendNodes(parent *html.Node, nodes []dom.ChildNode) {
	for _, node := range nodes {
		n := domNodeToHTMLNode(node)
		parent.AppendChild(n)
	}
}

func replaceChildren(parent *html.Node, nodes []dom.ChildNode) {
	clearChildren(parent)
	for _, node := range nodes {
		n := domNodeToHTMLNode(node)
		parent.AppendChild(n)
	}
}

func clearChildren(node *html.Node) {
	if fc := node.FirstChild; fc != nil {
		fc.Parent = nil
	}
	if lc := node.LastChild; lc != nil {
		lc.Parent = nil
	}
	node.FirstChild = nil
	node.LastChild = nil
}

func getElementsByTagName(node *html.Node, name string) dom.ElementCollection {
	name = strings.ToUpper(name)
	var list ElementList
	walkNodes(node, func(n *html.Node) bool {
		if n != node && strings.ToUpper(n.Data) == name {
			list = append(list, n)
		}
		return false
	})
	return list
}

func getElementsByClassName(node *html.Node, name string) dom.ElementCollection {
	var list ElementList
	walkNodes(node, func(n *html.Node) bool {
		if n != node && hasClasses(getAttribute(n, "class"), name) {
			list = append(list, n)
		}
		return false
	})
	return list
}

func hasClasses(elementClassesStr, classesStr string) bool {
	elementClasses := strings.Fields(elementClassesStr)
	classes := strings.Fields(classesStr)

	set := make(map[string]struct{}, len(classesStr))
	for _, c := range classes {
		set[c] = struct{}{}
	}

	for _, c := range elementClasses {
		delete(set, c)
	}

	return len(set) == 0
}

func firstElementWithTag(node *html.Node, tag string) *html.Node {
	var result *html.Node
	tag = strings.ToUpper(tag)
	walkNodes(node, func(node *html.Node) (done bool) {
		if strings.ToUpper(node.Data) == tag {
			result = node
			return true
		}
		return false
	})
	return result
}

//func querySelector(node *html.Node, query string) dom.Element {
//	panic("implement me")
//}
//
//func querySelectorAll(node *html.Node, query string) dom.NodeList {
//	panic("implement me")
//}
//
//func closest(node *html.Node, selector string) dom.Element {
//	panic("implement me")
//}
//
//func matches(node *html.Node, selector string) bool {
//	panic("implement me")
//}
