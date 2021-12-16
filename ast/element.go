package template

import (
	"strings"

	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
)

type Element struct {
	Node html.Node
}

func (e *Element) NodeType() dom.NodeType {
	return nodeType(e.Node.Type)
}

func (e *Element) IsConnected() bool {
	//TODO implement me
	panic("implement me")
}

func (e *Element) OwnerDocument() dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) GetRootNode(composed bool) dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) ParentNode() dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) ParentElement() dom.Element {
	//TODO implement me
	panic("implement me")
}

func (e *Element) HasChildNodes() bool {
	//TODO implement me
	panic("implement me")
}

func (e *Element) ChildNodes() dom.NodeList {
	//TODO implement me
	panic("implement me")
}

func (e *Element) FirstChild() dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) LastChild() dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) PreviousSibling() dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) NextSibling() dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) NodeValue() string {
	//TODO implement me
	panic("implement me")
}

func (e *Element) TextContent() string {
	//TODO implement me
	panic("implement me")
}

func (e *Element) Normalize() {
	//TODO implement me
	panic("implement me")
}

func (e *Element) CloneNode(deep bool) dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) IsEqualNode(other dom.Node) bool {
	//TODO implement me
	panic("implement me")
}

func (e *Element) IsSameNode(other dom.Node) bool {
	//TODO implement me
	panic("implement me")
}

func (e *Element) CompareDocumentPosition() dom.DocumentPosition {
	//TODO implement me
	panic("implement me")
}

func (e *Element) Contains(other dom.Node) bool {
	//TODO implement me
	panic("implement me")
}

func (e *Element) InsertBefore(node, child dom.Node) dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) AppendChild(node dom.Node) dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) ReplaceChild(node, child dom.Node) dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) RemoveChild(node dom.Node) dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) Children() dom.NodeList {
	//TODO implement me
	panic("implement me")
}

func (e *Element) FirstElementChild() dom.Element {
	//TODO implement me
	panic("implement me")
}

func (e *Element) LastElementChild() dom.Element {
	//TODO implement me
	panic("implement me")
}

func (e *Element) ChildElementCount() int {
	//TODO implement me
	panic("implement me")
}

func (e *Element) Prepend(nodes ...dom.Node) dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) Append(nodes ...dom.Node) dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) ReplaceChildren(nodes ...dom.Node) dom.Node {
	//TODO implement me
	panic("implement me")
}

func (e *Element) GetElementsByTagName(name string) dom.ElementCollection {
	//TODO implement me
	panic("implement me")
}

func (e *Element) GetElementsByTagNameNS(namespace, name string) dom.ElementCollection {
	//TODO implement me
	panic("implement me")
}

func (e *Element) GetElementsByClassName(name string) dom.ElementCollection {
	//TODO implement me
	panic("implement me")
}

func (e *Element) QuerySelector(query string) dom.Element {
	//TODO implement me
	panic("implement me")
}

func (e *Element) QuerySelectorAll(query string) dom.NodeList {
	//TODO implement me
	panic("implement me")
}

func (e *Element) TagName() string { return strings.ToUpper(e.Node.Data) }

func (e *Element) ID() string {
	//TODO implement me
	panic("implement me")
}

func (e *Element) ClassName() string {
	//TODO implement me
	panic("implement me")
}

func (e *Element) GetAttribute(name string) string {
	//TODO implement me
	panic("implement me")
}

func (e *Element) GetAttributeNS(namespace, name string) string {
	//TODO implement me
	panic("implement me")
}

func (e *Element) SetAttribute(name, value string) {
	//TODO implement me
	panic("implement me")
}

func (e *Element) SetAttributeNS(namespace, name, value string) {
	//TODO implement me
	panic("implement me")
}

func (e *Element) RemoveAttribute(name string) {
	//TODO implement me
	panic("implement me")
}

func (e *Element) RemoveAttributeNS(namespace, name string) {
	//TODO implement me
	panic("implement me")
}

func (e *Element) ToggleAttribute(name string) bool {
	//TODO implement me
	panic("implement me")
}

func (e *Element) HasAttribute(name string) bool {
	//TODO implement me
	panic("implement me")
}

func (e *Element) HasAttributeNS(namespace, name string) bool {
	//TODO implement me
	panic("implement me")
}

func (e *Element) Closest(selector string) dom.Element {
	//TODO implement me
	panic("implement me")
}

func (e *Element) Matches(selector string) bool {
	//TODO implement me
	panic("implement me")
}

func (e *Element) SetInnerHTML(s string) {
	//TODO implement me
	panic("implement me")
}

func (e *Element) InnerHTML() string {
	//TODO implement me
	panic("implement me")
}

func (e *Element) SetOuterHTML(s string) {
	//TODO implement me
	panic("implement me")
}

func (e *Element) OuterHTML() string {
	//TODO implement me
	panic("implement me")
}

func (e *Element) SetInnerText(s string) {
	//TODO implement me
	panic("implement me")
}

func (e *Element) InnerText() string {
	//TODO implement me
	panic("implement me")
}
