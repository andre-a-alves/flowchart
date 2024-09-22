package flowchart

import (
	"fmt"
	"slices"
)

type (
	DirectionEnum int
	NodeTypeEnum  int
	LineTypeEnum  int
	ArrowTypeEnum int
)

const (
	DirectionHorizontalRight DirectionEnum = iota
	DirectionHorizontalLeft
	DirectionVertical
)

const (
	NodeTypeTerminator NodeTypeEnum = iota
	NodeTypeProcess
	NodeTypeAlternateProcess
	NodeTypeSubprocess
	NodeTypeDecision
	NodeTypeInputOutput
	NodeTypeConnector
	NodeTypeDatabase
)

const (
	LineTypeNone LineTypeEnum = iota
	LineTypeSolid
	LineTypeDotted
	LineTypeThick
)

const (
	ArrowTypeNone ArrowTypeEnum = iota
	ArrowTypeNormal
	ArrowTypeCircle
	ArrowTypeCross
)

type Link struct {
	Origin      Linkable
	Target      Linkable
	LineType    LineTypeEnum
	ArrowType   ArrowTypeEnum
	OriginArrow bool
	TargetArrow bool
	Label       *string
}

type Linkable interface {
	nodeName() string
}

func (n *Node) nodeName() string {
	return n.name
}

func (f *Flowchart) nodeName() string {
	if f.Title == nil {
		return ""
	}
	return *f.Title
}

type Node struct {
	name  string
	Type  NodeTypeEnum
	Label *string
	Links []Link
}

type Flowchart struct {
	Direction DirectionEnum
	Title     *string
	Nodes     []*Node
	Subgraphs []*Flowchart
	Links     []Link
}

func (f *Flowchart) AddLink(link Link) error {
	if link.Target == (Linkable)(nil) {
		return fmt.Errorf("cannot add link with no target node")
	}
	if link.Origin == (Linkable)(nil) {
		return fmt.Errorf("cannot add link with no origin node")
	}
	f.Links = append(f.Links, link)
	return nil
}

func (f *Flowchart) allNames() []string {
	var names []string
	if f.Title != nil {
		names = append(names, *f.Title)
	}
	for _, n := range f.Nodes {
		names = append(names, n.name)
	}
	for _, s := range f.Subgraphs {
		names = append(names, s.allNames()...)
	}
	return names
}

func (f *Flowchart) containsName(name string) bool {
	return slices.Contains(f.allNames(), name)
}

func (f *Flowchart) AddNode(node *Node) error {
	if f.containsName(node.name) {
		return fmt.Errorf("cannot add node with non-unique name")
	}
	f.Nodes = append(f.Nodes, node)
	return nil
}

func (f *Flowchart) AddSubgraph(subgraph *Flowchart) error {
	if subgraph.Title == nil {
		return fmt.Errorf("cannot add subgraph with no title")
	}
	if f.containsName(*subgraph.Title) {
		return fmt.Errorf("cannot add subgraph with already existing title")
	}
	f.Subgraphs = append(f.Subgraphs, subgraph)
	return nil
}

func BlankLink(target Linkable, label *string) Link {
	return basicLink(target, label, LineTypeNone)
}

func SolidLink(target Linkable, label *string) Link {
	return basicLink(target, label, LineTypeSolid)
}

func DottedLink(target Linkable, label *string) Link {
	return basicLink(target, label, LineTypeDotted)
}

func ThickLink(target Linkable, label *string) Link {
	return basicLink(target, label, LineTypeThick)
}

func basicLink(target Linkable, label *string, lineType LineTypeEnum) Link {
	return Link{
		Target:      target,
		LineType:    lineType,
		ArrowType:   ArrowTypeNormal,
		OriginArrow: false,
		TargetArrow: true,
		Label:       label,
	}
}

func TerminatorNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeTerminator)
}

func ProcessNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeProcess)
}

func AlternateProcessNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeAlternateProcess)
}

func SubprocessNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeSubprocess)
}

func DecisionNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeDecision)
}

func InputOutputNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeInputOutput)
}

func ConnectorNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeConnector)
}

func DatabaseNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeDatabase)
}

func basicNode(name string, label *string, typ NodeTypeEnum) *Node {
	return &Node{
		name:  name,
		Type:  typ,
		Label: label,
		Links: make([]Link, 0),
	}
}

func VerticalFlowchart(title *string) *Flowchart {
	return basicFlowchart(title, DirectionVertical)
}

func LrFlowchart(title *string) *Flowchart {
	return basicFlowchart(title, DirectionHorizontalRight)
}

func RlFlowchart(title *string) *Flowchart {
	return basicFlowchart(title, DirectionHorizontalLeft)
}

func basicFlowchart(title *string, direction DirectionEnum) *Flowchart {
	return &Flowchart{
		Direction: direction,
		Title:     title,
		Nodes:     make([]*Node, 0),
		Subgraphs: make([]*Flowchart, 0),
	}
}
