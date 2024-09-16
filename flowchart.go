package flowchart

import (
	"fmt"
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
	Name  string
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

func (f *Flowchart) AddNode(node *Node) error {
	for _, n := range f.Nodes {
		if n.Name == node.Name {
			return fmt.Errorf("cannot add duplicate subgraph")
		}
	}
	f.Nodes = append(f.Nodes, node)
	return nil
}

func (f *Flowchart) AddSubgraph(subgraph *Flowchart) error {
	for _, s := range f.Subgraphs {
		if s.Title != nil && subgraph.Title != nil && *s.Title == *subgraph.Title {
			return fmt.Errorf("cannot add duplicate subgraph")
		}
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

func BasicNode(name string, label *string) *Node {
	return &Node{
		Name:  name,
		Type:  TypeProcess,
		Label: label,
		Links: make([]Link, 0),
	}
}

func BasicFlowchart() *Flowchart {
	return &Flowchart{
		Direction: DirectionVertical,
		Title:     nil,
		Nodes:     make([]*Node, 0),
		Subgraphs: make([]*Flowchart, 0),
	}
}
