package flowchart

import (
	"errors"
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
	TypeTerminator NodeTypeEnum = iota
	TypeProcess
	TypeAlternateProcess
	TypeSubprocess
	TypeDecision
	TypeInputOutput
	TypeConnector
	TypeDatabase
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
	OriginArrow ArrowTypeEnum
	TargetArrow ArrowTypeEnum
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

func (n *Node) addLink(link Link) error {
	if link.TargetNode == nil {
		return fmt.Errorf("cannot add link with no target node")
	}
	n.Links = append(n.Links, link)
	return nil
}

func (f *Flowchart) addNode(node *Node) error {
	f.Nodes = append(f.Nodes, node)
	return nil
}

func (f *Flowchart) addSubgraph(subgraph *Flowchart) error {
	f.Subgraphs = append(f.Subgraphs, subgraph)
	return nil
}

func BasicLink(targetNode *Node, label *string) (Link, error) {
	if targetNode == nil {
		return Link{}, errors.New("targetNode is nil")
	}
	return Link{
		TargetNode:  targetNode,
		LineType:    LineTypeSolid,
		OriginArrow: ArrowTypeNone,
		TargetArrow: ArrowTypeNormal,
		Label:       label,
	}, nil
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
