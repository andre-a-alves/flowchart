package flowchart

import (
	"fmt"
	"slices"
)

type (
	FlowchartDirectionEnum int
	NodeTypeEnum           int
	LineTypeEnum           int
	ArrowTypeEnum          int
)

const (
	DirectionHorizontalRight FlowchartDirectionEnum = iota
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
	TargetNode  *Node
	LineType    LineTypeEnum
	ArrowType   ArrowTypeEnum
	OriginArrow bool
	TargetArrow bool
	Label       *string
}

type Node struct {
	name  string
	Type  NodeTypeEnum
	Label *string
	Links []Link
}

type Flowchart struct {
	Direction FlowchartDirectionEnum
	Title     *string
	Nodes     []*Node
	Subgraphs []*Flowchart
}

func (n *Node) AddLink(link Link) error {
	if link.TargetNode == nil {
		return fmt.Errorf("cannot add link with no target node")
	}
	n.Links = append(n.Links, link)
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

func BlankLink(targetNode *Node, label *string) Link {
	return basicLink(targetNode, label, LineTypeNone)
}

func SolidLink(targetNode *Node, label *string) Link {
	return basicLink(targetNode, label, LineTypeSolid)
}

func DottedLink(targetNode *Node, label *string) Link {
	return basicLink(targetNode, label, LineTypeDotted)
}

func ThickLink(targetNode *Node, label *string) Link {
	return basicLink(targetNode, label, LineTypeThick)
}

func basicLink(targetNode *Node, label *string, lineType LineTypeEnum) Link {
	return Link{
		TargetNode:  targetNode,
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

func basicFlowchart(title *string, direction FlowchartDirectionEnum) *Flowchart {
	return &Flowchart{
		Direction: direction,
		Title:     title,
		Nodes:     make([]*Node, 0),
		Subgraphs: make([]*Flowchart, 0),
	}
}
