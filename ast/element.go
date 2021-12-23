package ast

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
)

type Element struct {
	node *html.Node
}

// Node

func (e *Element) NodeType() dom.NodeType         { return nodeType(e.node.Type) }
func (e *Element) IsConnected() bool              { return isConnected(e.node) }
func (e *Element) OwnerDocument() dom.Document    { return ownerDocument(e.node) }
func (e *Element) ParentNode() dom.Node           { return parentNode(e.node) }
func (e *Element) ParentElement() dom.Element     { return parentElement(e.node) }
func (e *Element) PreviousSibling() dom.ChildNode { return previousSibling(e.node) }
func (e *Element) NextSibling() dom.ChildNode     { return nextSibling(e.node) }
func (e *Element) TextContent() string            { return textContent(e.node) }
func (e *Element) CloneNode(deep bool) dom.Node   { return cloneNode(e.node, deep) }
func (e *Element) IsSameNode(other dom.Node) bool { return isSameNode(e.node, other) }
func (e *Element) Length() int                    { return e.ChildNodes().Length() }

// ParentNode

func (e *Element) Children() dom.ElementCollection        { return children(e.node) }
func (e *Element) FirstElementChild() dom.Element         { return firstElementChild(e.node) }
func (e *Element) LastElementChild() dom.Element          { return lastElementChild(e.node) }
func (e *Element) ChildElementCount() int                 { return childElementCount(e.node) }
func (e *Element) Prepend(nodes ...dom.ChildNode)         { prependNodes(e.node, nodes) }
func (e *Element) Append(nodes ...dom.ChildNode)          { appendNodes(e.node, nodes) }
func (e *Element) ReplaceChildren(nodes ...dom.ChildNode) { replaceChildren(e.node, nodes) }
func (e *Element) GetElementsByTagName(name string) dom.ElementCollection {
	return getElementsByTagName(e.node, name)
}
func (e *Element) GetElementsByClassName(name string) dom.ElementCollection {
	return getElementsByClassName(e.node, name)
}

//func (e *Element) QuerySelector(query string) dom.Element     { return querySelector(e.node, query) }
//func (e *Element) QuerySelectorAll(query string) dom.NodeList { return querySelectorAll(e.node, query) }
//func (e *Element) Closest(selector string) dom.Element        { return closest(e.node, selector) }
//func (e *Element) Matches(selector string) bool               { return matches(e.node, selector) }

func (e *Element) HasChildNodes() bool          { return hasChildNodes(e.node) }
func (e *Element) ChildNodes() dom.NodeList     { return childNodes(e.node) }
func (e *Element) FirstChild() dom.ChildNode    { return firstChild(e.node) }
func (e *Element) LastChild() dom.ChildNode     { return lastChild(e.node) }
func (e *Element) Contains(other dom.Node) bool { return contains(e.node, other) }
func (e *Element) InsertBefore(node, child dom.ChildNode) dom.ChildNode {
	return insertBefore(e.node, node, child)
}
func (e *Element) AppendChild(node dom.ChildNode) dom.ChildNode { return appendChild(e.node, node) }
func (e *Element) ReplaceChild(node, child dom.ChildNode) dom.ChildNode {
	return replaceChild(e.node, node, child)
}
func (e *Element) RemoveChild(node dom.ChildNode) dom.ChildNode { return removeChild(e.node, node) }

// Element

func (e *Element) TagName() string                 { return strings.ToUpper(e.node.Data) }
func (e *Element) ID() string                      { return getAttribute(e.node, "id") }
func (e *Element) ClassName() string               { return getAttribute(e.node, "class") }
func (e *Element) GetAttribute(name string) string { return getAttribute(e.node, name) }

func getAttribute(node *html.Node, name string) string {
	name = strings.ToLower(name)
	for _, att := range node.Attr {
		if att.Key == name {
			return att.Val
		}
	}
	return ""
}

func (e *Element) SetAttribute(name, value string) {
	name = strings.ToLower(name)
	for index, att := range e.node.Attr {
		if att.Key == name {
			e.node.Attr[index].Val = value
		}
	}
	e.node.Attr = append(e.node.Attr, html.Attribute{
		Key: name, Val: value,
	})
}

func (e *Element) RemoveAttribute(name string) {
	name = strings.ToLower(name)
	filtered := e.node.Attr[:0]
	for _, att := range e.node.Attr {
		if att.Key == name {
			continue
		}
		filtered = append(filtered, att)
	}
	e.node.Attr = filtered
}

func (e *Element) ToggleAttribute(name string) bool {
	name = strings.ToLower(name)
	if e.HasAttribute(name) {
		e.RemoveAttribute(name)
		return false
	}
	e.SetAttribute(name, "")
	return true
}

func (e *Element) HasAttribute(name string) bool {
	name = strings.ToLower(name)
	for _, att := range e.node.Attr {
		if att.Key == name {
			return true
		}
	}
	return false
}

func (e *Element) isNamed(name string) bool {
	return isNamed(e.node, name)
}

func isNamed(node *html.Node, name string) bool {
	if node == nil || node.Type != html.ElementNode {
		return false
	}
	id := getAttribute(node, "id")
	nm := getAttribute(node, "name")
	return (id != "" && id == name) || (nm != "" && nm == name)
}

func (e *Element) SetInnerHTML(s string) {
	nodes, err := html.ParseFragment(strings.NewReader(s), &html.Node{Type: html.ElementNode})
	if err != nil {
		panic(err)
	}
	clearChildren(e.node)
	for _, n := range nodes {
		e.node.AppendChild(n)
	}
}

func (e *Element) InnerHTML() string {
	var buf bytes.Buffer
	c := e.node.FirstChild
	for c != nil {
		err := html.Render(&buf, c)
		if err != nil {
			panic(err)
		}
		c = c.NextSibling
	}
	return buf.String()
}

func (e *Element) SetOuterHTML(s string) {
	nodes, err := html.ParseFragment(strings.NewReader(s), &html.Node{Type: html.ElementNode})
	if err != nil {
		panic(err)
	}
	if len(nodes) == 0 {
		return
	}
	if e.node.Parent == nil {
		panic("browser: SetOuterHTML called on an unattached node")
	}
	for _, node := range nodes {
		e.node.Parent.InsertBefore(node, e.node)
	}
	e.node.Parent.RemoveChild(e.node)
}

func (e *Element) OuterHTML() string {
	var buf bytes.Buffer
	err := html.Render(&buf, e.node)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
