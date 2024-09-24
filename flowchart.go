// Package flowchart provides a set of types and methods to represent and manipulate flowcharts
// with various node types, links, and subgraphs.
package flowchart

import (
	"fmt"
	"slices"
)

type (

	// DirectionEnum represents the flowchart direction (e.g., horizontal or vertical).
	DirectionEnum int

	// NodeTypeEnum represents the type of node in a flowchart (e.g., process, decision).
	NodeTypeEnum int

	// LineTypeEnum represents the type of line connecting nodes in a flowchart (e.g., solid, dotted).
	LineTypeEnum int

	// ArrowTypeEnum represents the type of arrow used in flowchart links (e.g., normal, circle).
	ArrowTypeEnum int
)

// Constants for flowchart directions.
const (
	DirectionHorizontalRight DirectionEnum = iota // Left to right
	DirectionHorizontalLeft                       // Right to left
	DirectionVertical                             // Top to bottom
)

// Constants for various node types in a flowchart.
const (
	NodeTypeTerminator       NodeTypeEnum = iota // Start/End node
	NodeTypeProcess                              // Standard process node
	NodeTypeAlternateProcess                     // Alternate process node
	NodeTypeSubprocess                           // Subprocess node
	NodeTypeDecision                             // Decision node
	NodeTypeInputOutput                          // Input/Output node
	NodeTypeConnector                            // Connector node
	NodeTypeDatabase                             // Database node
)

// Constants for line types in flowchart links.
const (
	LineTypeNone   LineTypeEnum = iota // No line
	LineTypeSolid                      // Solid line
	LineTypeDotted                     // Dotted line
	LineTypeThick                      // Thick line
)

// Constants for arrow types in flowchart links.
const (
	ArrowTypeNone   ArrowTypeEnum = iota // No arrow
	ArrowTypeNormal                      // Normal arrow
	ArrowTypeCircle                      // Circle arrow
	ArrowTypeCross                       // Cross arrow
)

// Link represents a connection between two nodes in a flowchart.
type Link struct {
	Origin      Linkable      // The origin node of the link
	Target      Linkable      // The target node of the link
	LineType    LineTypeEnum  // Type of line connecting the nodes
	ArrowType   ArrowTypeEnum // Type of arrow used in the link
	OriginArrow bool          // Whether the link has an arrow at the origin
	TargetArrow bool          // Whether the link has an arrow at the target
	Label       *string       // Optional label for the link
}

// Linkable represents an object that can be linked in a flowchart.
type Linkable interface {
	nodeName() string
}

// nodeName returns the name of the node as a string.
// It is used to uniquely identify a Node within the flowchart.
func (n *Node) nodeName() string {
	return n.name
}

// nodeName returns the title of the flowchart if available, or an empty string if no title is set.
// It is used to uniquely identify a Flowchart when linking or creating subgraphs.
func (f *Flowchart) nodeName() string {
	if f.Title == nil {
		return ""
	}
	return *f.Title
}

// Node represents a node in the flowchart.
type Node struct {
	name  string       // Internal name of the node
	Type  NodeTypeEnum // Type of the node
	Label *string      // Optional label for the node
}

// Flowchart represents a flowchart with nodes, subgraphs, and links.
type Flowchart struct {
	Direction DirectionEnum // Flow direction (LR, RL, TB)
	Title     *string       // Title of the flowchart
	Nodes     []*Node       // List of nodes in the flowchart
	Subgraphs []*Flowchart  // List of subgraphs
	Links     []Link        // List of links between nodes
}

// AddLink adds a link to the flowchart.
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

// allNames returns all the names (nodes and subgraphs) within the flowchart.
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

// containsName checks if a given name is present in the flowchart (either in nodes or subgraphs).
func (f *Flowchart) containsName(name string) bool {
	return slices.Contains(f.allNames(), name)
}

// AddNode adds a node to the flowchart, ensuring it has a unique name.
func (f *Flowchart) AddNode(node *Node) error {
	if f.containsName(node.name) {
		return fmt.Errorf("cannot add node with non-unique name")
	}
	f.Nodes = append(f.Nodes, node)
	return nil
}

// AddSubgraph adds a subgraph to the flowchart, ensuring it has a unique title.
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

// BlankLink creates a link with no line between two nodes.
func BlankLink(origin, target Linkable, label *string) Link {
	return basicLink(origin, target, label, LineTypeNone)
}

// SolidLink creates a solid link between two nodes.
func SolidLink(origin, target Linkable, label *string) Link {
	return basicLink(origin, target, label, LineTypeSolid)
}

// DottedLink creates a dotted link between two nodes.
func DottedLink(origin, target Linkable, label *string) Link {
	return basicLink(origin, target, label, LineTypeDotted)
}

// ThickLink creates a thick link between two nodes.
func ThickLink(origin, target Linkable, label *string) Link {
	return basicLink(origin, target, label, LineTypeThick)
}

// basicLink is a helper function to create a link with the specified origin, target, label, and line-type.
func basicLink(origin, target Linkable, label *string, lineType LineTypeEnum) Link {
	return Link{
		Origin:      origin,
		Target:      target,
		LineType:    lineType,
		ArrowType:   ArrowTypeNormal,
		OriginArrow: false,
		TargetArrow: true,
		Label:       label,
	}
}

// TerminatorNode creates a terminator node (start/end) with the specified name and label.
func TerminatorNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeTerminator)
}

// ProcessNode creates a process node with the specified name and label.
func ProcessNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeProcess)
}

// AlternateProcessNode creates an alternate process node with the specified name and label.
func AlternateProcessNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeAlternateProcess)
}

// SubprocessNode creates a subprocess node with the specified name and label.
func SubprocessNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeSubprocess)
}

// DecisionNode creates a decision node with the specified name and label.
func DecisionNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeDecision)
}

// InputOutputNode creates an input/output node with the specified name and label.
func InputOutputNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeInputOutput)
}

// ConnectorNode creates a connector node with the specified name and label.
func ConnectorNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeConnector)
}

// DatabaseNode creates a database node with the specified name and label.
func DatabaseNode(name string, label *string) *Node {
	return basicNode(name, label, NodeTypeDatabase)
}

// basicNode is a helper function to create a node with the specified name, label, and type.
func basicNode(name string, label *string, typ NodeTypeEnum) *Node {
	return &Node{
		name:  name,
		Type:  typ,
		Label: label,
	}
}

// VerticalFlowchart creates a flowchart with vertical direction.
func VerticalFlowchart(title *string) *Flowchart {
	return basicFlowchart(title, DirectionVertical)
}

// LrFlowchart creates a flowchart with left-to-right direction.
func LrFlowchart(title *string) *Flowchart {
	return basicFlowchart(title, DirectionHorizontalRight)
}

// RlFlowchart creates a flowchart with right-to-left direction.
func RlFlowchart(title *string) *Flowchart {
	return basicFlowchart(title, DirectionHorizontalLeft)
}

// basicFlowchart is a helper function to create a flowchart with the specified title and direction.
func basicFlowchart(title *string, direction DirectionEnum) *Flowchart {
	return &Flowchart{
		Direction: direction,
		Title:     title,
		Nodes:     make([]*Node, 0),
		Subgraphs: make([]*Flowchart, 0),
	}
}
