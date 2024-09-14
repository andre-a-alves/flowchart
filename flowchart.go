package main

import "errors"

type (
	DirectionEnum int
	ShapeEnum     int
	LineTypeEnum  int
	ArrowTypeEnum int
)

const (
	DirectionHorizontalRight DirectionEnum = iota
	DirectionHorizontalLeft
	DirectionVertical
)

const (
	ShapeRectangle ShapeEnum = iota
	ShapeRoundEdge
	ShapeStadium
	ShapeSubroutine
	ShapeCylinder
	ShapeCircle
	ShapeAsymmetric
	ShapeDiamond
	ShapeRhombus
	ShapeHexagon
	ShapeParallelogram
	ShapeAltParallelogram
	ShapeTrapezoid
	ShapeAltTrapezoid
	ShapeDoubleCircle
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
	Shape ShapeEnum
	Label *string
	Links []Link
}

type Flowchart struct {
	Direction DirectionEnum
	Title     *string
	Nodes     []*Node
	Subgraphs []*Flowchart
}

func (n Node) addLink(link Link) {
	n.Links = append(n.Links, link)
}

func (f Flowchart) addNode(node *Node) {
	f.Nodes = append(f.Nodes, node)
}

func (f Flowchart) addSubgraph(subgraph *Flowchart) {
	f.Subgraphs = append(f.Subgraphs, subgraph)
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
		Shape: ShapeRectangle,
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
