package flowchart

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNode_AddLink(t *testing.T) {
	tests := []struct {
		name          string
		node          Node
		link          Link
		expectedErr   error
		expectedLinks []Link
	}{
		{
			name: "Add valid link",
			node: Node{name: "StartNode", Links: []Link{}},
			link: Link{
				TargetNode: &Node{name: "TargetNode"},
				Label:      pointTo("LinkLabel"),
			},
			expectedErr: nil,
			expectedLinks: []Link{
				{
					TargetNode: &Node{name: "TargetNode"},
					Label:      pointTo("LinkLabel"),
				},
			},
		},
		{
			name: "Add invalid link with nil TargetNode",
			node: Node{name: "StartNode", Links: []Link{}},
			link: Link{
				TargetNode: nil,
				Label:      nil,
			},
			expectedErr:   fmt.Errorf("cannot add link with no target node"),
			expectedLinks: []Link{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.node.AddLink(tt.link)

			if diff := cmp.Diff(tt.expectedErr, err, cmp.Comparer(compareErrors)); diff != "" {
				t.Errorf("AddLink() error mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.expectedLinks, tt.node.Links, cmp.AllowUnexported(Node{})); diff != "" {
				t.Errorf("AddLink() Links mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlowchart_AddNode(t *testing.T) {
	tests := []struct {
		name          string
		flowchart     *Flowchart
		node          *Node
		expectedErr   error
		expectedNodes []*Node
	}{
		{
			name:          "Add node",
			flowchart:     &Flowchart{},
			node:          &Node{name: "Singleton"},
			expectedErr:   nil,
			expectedNodes: []*Node{{name: "Singleton"}},
		},
		{
			name:          "Add duplicate node",
			flowchart:     &Flowchart{Nodes: []*Node{{name: "Singleton"}}},
			node:          &Node{name: "Singleton"},
			expectedErr:   fmt.Errorf("cannot add node with non-unique name"),
			expectedNodes: []*Node{{name: "Singleton"}},
		},
		{
			name:          "Add node with subgraph names",
			flowchart:     &Flowchart{Subgraphs: []*Flowchart{LrFlowchart(pointTo("Singleton"))}},
			node:          &Node{name: "Singleton"},
			expectedErr:   fmt.Errorf("cannot add node with non-unique name"),
			expectedNodes: nil,
		},
		{
			name: "Add node with duplicate in subgraph",
			flowchart: &Flowchart{
				Subgraphs: []*Flowchart{{
					Direction: DirectionVertical,
					Nodes:     []*Node{{name: "Singleton"}},
				}},
			},
			node:          &Node{name: "Singleton"},
			expectedErr:   fmt.Errorf("cannot add node with non-unique name"),
			expectedNodes: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.flowchart.AddNode(tt.node)

			if diff := cmp.Diff(tt.expectedErr, err, cmp.Comparer(compareErrors)); diff != "" {
				t.Errorf("AddLink() error mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.expectedNodes, tt.flowchart.Nodes, cmp.AllowUnexported(Node{})); diff != "" {
				t.Errorf("AddNode() Nodes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlowchart_AddSubgraph(t *testing.T) {
	tests := []struct {
		name              string
		flowchart         *Flowchart
		subgraph          *Flowchart
		expectedErr       error
		expectedSubgraphs []*Flowchart
	}{
		{
			name:              "Add subgraph",
			flowchart:         &Flowchart{},
			subgraph:          &Flowchart{Title: pointTo("Singleton")},
			expectedErr:       nil,
			expectedSubgraphs: []*Flowchart{{Title: pointTo("Singleton")}},
		},
		{
			name:              "Add subgraph with no title",
			flowchart:         &Flowchart{},
			subgraph:          &Flowchart{},
			expectedErr:       fmt.Errorf("cannot add subgraph with no title"),
			expectedSubgraphs: nil,
		},
		{
			name:              "Add duplicate subgraph",
			flowchart:         &Flowchart{Subgraphs: []*Flowchart{{Title: pointTo("Singleton")}}},
			subgraph:          &Flowchart{Title: pointTo("Singleton")},
			expectedErr:       fmt.Errorf("cannot add subgraph with already existing title"),
			expectedSubgraphs: []*Flowchart{{Title: pointTo("Singleton")}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.flowchart.AddSubgraph(tt.subgraph)

			if diff := cmp.Diff(tt.expectedErr, err, cmp.Comparer(compareErrors)); diff != "" {
				t.Errorf("AddLink() error mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.expectedSubgraphs, tt.flowchart.Subgraphs); diff != "" {
				t.Errorf("AddNode() Nodes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlowchart_allNames(t *testing.T) {
	tests := []struct {
		name      string
		flowchart Flowchart
		expected  []string
	}{
		{
			name: "Flowchart with title, nodes, and subgraphs",
			flowchart: Flowchart{
				Title: stringPtr("Main Flowchart"),
				Nodes: []*Node{
					{name: "Node1"},
					{name: "Node2"},
				},
				Subgraphs: []*Flowchart{
					{
						Title: stringPtr("Subgraph1"),
						Nodes: []*Node{
							{name: "SubNode1"},
							{name: "SubNode2"},
						},
					},
				},
			},
			expected: []string{"Main Flowchart", "Node1", "Node2", "Subgraph1", "SubNode1", "SubNode2"},
		},
		{
			name: "Flowchart without title",
			flowchart: Flowchart{
				Nodes: []*Node{
					{name: "Node1"},
				},
			},
			expected: []string{"Node1"},
		},
		{
			name: "Flowchart with nested subgraphs",
			flowchart: Flowchart{
				Title: stringPtr("Main Flowchart"),
				Subgraphs: []*Flowchart{
					{
						Title: stringPtr("Subgraph1"),
						Subgraphs: []*Flowchart{
							{
								Title: stringPtr("NestedSubgraph"),
								Nodes: []*Node{
									{name: "NestedNode1"},
								},
							},
						},
					},
				},
			},
			expected: []string{"Main Flowchart", "Subgraph1", "NestedSubgraph", "NestedNode1"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.flowchart.allNames()
			if !cmp.Equal(result, test.expected) {
				t.Errorf("unexpected result. got: %v, expected: %v", result, test.expected)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}

func TestLinks(t *testing.T) {
	fixtureLink := func(mods ...func(l *Link)) Link {
		link := &Link{
			TargetNode:  nil,
			LineType:    LineTypeSolid,
			ArrowType:   ArrowTypeNormal,
			OriginArrow: false,
			TargetArrow: true,
			Label:       nil,
		}
		for _, mod := range mods {
			mod(link)
		}
		return *link
	}
	fixtureLabel := pointTo("Link Label")
	fixtureTarget := &Node{name: "Node1"}

	testBasicLink := []struct {
		name       string
		targetNode *Node
		label      *string
		lineType   LineTypeEnum
		expected   Link
	}{
		{
			name:       "Nil target node",
			targetNode: nil,
			label:      fixtureLabel,
			lineType:   LineTypeSolid,
			expected: fixtureLink(func(l *Link) {
				l.Label = fixtureLabel
			}),
		},
		{
			name:       "Nil label",
			targetNode: fixtureTarget,
			label:      nil,
			lineType:   LineTypeSolid,
			expected: fixtureLink(func(l *Link) {
				l.TargetNode = fixtureTarget
			}),
		},
		{
			name:       "Valid target node with label - dotted",
			targetNode: fixtureTarget,
			label:      fixtureLabel,
			lineType:   LineTypeDotted,
			expected: fixtureLink(func(l *Link) {
				l.TargetNode = fixtureTarget
				l.Label = fixtureLabel
				l.LineType = LineTypeDotted
			}),
		},
	}
	for _, tt := range testBasicLink {
		t.Run(tt.name, func(t *testing.T) {
			got := basicLink(tt.targetNode, tt.label, tt.lineType)

			if diff := cmp.Diff(tt.expected, got, cmp.AllowUnexported(Node{})); diff != "" {
				t.Errorf("basicLink() got mismatch (-want +got):\n%s", diff)
			}
		})
	}

	testLinkTypes := []struct {
		name     string
		function func(*Node, *string) Link
		expected Link
	}{
		{
			name:     "blank link",
			function: BlankLink,
			expected: fixtureLink(func(l *Link) {
				l.LineType = LineTypeNone
			}),
		},
		{
			name:     "solid link",
			function: SolidLink,
			expected: fixtureLink(func(l *Link) {
				l.LineType = LineTypeSolid
			}),
		},
		{
			name:     "dotted link",
			function: DottedLink,
			expected: fixtureLink(func(l *Link) {
				l.LineType = LineTypeDotted
			}),
		},
		{
			name:     "thick link",
			function: ThickLink,
			expected: fixtureLink(func(l *Link) {
				l.LineType = LineTypeThick
			}),
		},
	}
	for _, tt := range testLinkTypes {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.function(nil, nil)

			if diff := cmp.Diff(tt.expected, got, cmp.AllowUnexported(Node{})); diff != "" {
				t.Errorf("got mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestNodes(t *testing.T) {
	fixtureNodeName := "node"
	fixtureNode := func(mods ...func(l *Node)) *Node {
		node := &Node{
			name:  fixtureNodeName,
			Type:  NodeTypeProcess,
			Label: nil,
			Links: []Link{},
		}
		for _, mod := range mods {
			mod(node)
		}
		return node
	}

	testBasicNode := []struct {
		name     string
		nodeName string
		label    *string
		typ      NodeTypeEnum
		expected *Node
	}{
		{
			name:     "nil label",
			nodeName: fixtureNodeName,
			label:    nil,
			typ:      NodeTypeProcess,
			expected: fixtureNode(),
		},
		{
			name:     "label",
			nodeName: fixtureNodeName,
			label:    pointTo("Node Label"),
			typ:      NodeTypeDatabase,
			expected: fixtureNode(func(l *Node) {
				l.Label = pointTo("Node Label")
				l.Type = NodeTypeDatabase
			}),
		},
	}
	for _, tt := range testBasicNode {
		t.Run(tt.name, func(t *testing.T) {
			got := basicNode(tt.nodeName, tt.label, tt.typ)

			if diff := cmp.Diff(tt.expected, got, cmp.AllowUnexported(Node{})); diff != "" {
				t.Errorf("basicNode() got mismatch (-want +got):\n%s", diff)
			}
		})
	}

	testNodeTypes := []struct {
		name     string
		function func(string, *string) *Node
		expected *Node
	}{
		{
			name:     "terminator",
			function: TerminatorNode,
			expected: fixtureNode(func(l *Node) {
				l.Type = NodeTypeTerminator
			}),
		},
		{
			name:     "process",
			function: ProcessNode,
			expected: fixtureNode(func(l *Node) {
				l.Type = NodeTypeProcess
			}),
		},
		{
			name:     "alternate process",
			function: AlternateProcessNode,
			expected: fixtureNode(func(l *Node) {
				l.Type = NodeTypeAlternateProcess
			}),
		},
		{
			name:     "subprocess",
			function: SubprocessNode,
			expected: fixtureNode(func(l *Node) {
				l.Type = NodeTypeSubprocess
			}),
		},
		{
			name:     "decision",
			function: DecisionNode,
			expected: fixtureNode(func(l *Node) {
				l.Type = NodeTypeDecision
			}),
		},
		{
			name:     "input/output",
			function: InputOutputNode,
			expected: fixtureNode(func(l *Node) {
				l.Type = NodeTypeInputOutput
			}),
		},
		{
			name:     "connector",
			function: ConnectorNode,
			expected: fixtureNode(func(l *Node) {
				l.Type = NodeTypeConnector
			}),
		},
		{
			name:     "database",
			function: DatabaseNode,
			expected: fixtureNode(func(l *Node) {
				l.Type = NodeTypeDatabase
			}),
		},
	}
	for _, tt := range testNodeTypes {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.function(fixtureNodeName, nil)

			if diff := cmp.Diff(tt.expected, got, cmp.AllowUnexported(Node{})); diff != "" {
				t.Errorf("got mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestBasicFlowchart(t *testing.T) {
	fixtureFlowchart := func(mods ...func(f *Flowchart)) *Flowchart {
		flowchart := &Flowchart{
			Direction: DirectionVertical,
			Title:     nil,
			Nodes:     []*Node{},
			Subgraphs: []*Flowchart{},
		}
		for _, mod := range mods {
			mod(flowchart)
		}
		return flowchart
	}
	fixtureTitle := pointTo("Flowchart Title")

	testBasicFlowchart := []struct {
		name      string
		direction FlowchartDirectionEnum
		title     *string
		expected  *Flowchart
	}{
		{
			name:      "nil title",
			direction: DirectionVertical,
			title:     nil,
			expected:  fixtureFlowchart(),
		},
		{
			name:      "title",
			direction: DirectionVertical,
			title:     fixtureTitle,
			expected: fixtureFlowchart(func(f *Flowchart) {
				f.Title = fixtureTitle
			}),
		},
	}
	for _, tt := range testBasicFlowchart {
		t.Run(tt.name, func(t *testing.T) {
			got := basicFlowchart(tt.title, tt.direction)

			if diff := cmp.Diff(tt.expected, got, cmp.AllowUnexported(Node{})); diff != "" {
				t.Errorf("basicLink() got mismatch (-want +got):\n%s", diff)
			}
		})
	}

	testFlowchartTypes := []struct {
		name     string
		function func(*string) *Flowchart
		expected *Flowchart
	}{
		{
			name:     "vertical",
			function: VerticalFlowchart,
			expected: fixtureFlowchart(),
		},
		{
			name:     "left-right",
			function: LrFlowchart,
			expected: fixtureFlowchart(func(f *Flowchart) {
				f.Direction = DirectionHorizontalRight
			}),
		},
		{
			name:     "right-left",
			function: RlFlowchart,
			expected: fixtureFlowchart(func(f *Flowchart) {
				f.Direction = DirectionHorizontalLeft
			}),
		},
	}
	for _, tt := range testFlowchartTypes {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.function(nil)

			if diff := cmp.Diff(tt.expected, got, cmp.AllowUnexported(Flowchart{})); diff != "" {
				t.Errorf("got mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// Helper function to compare errors
func compareErrors(err1, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	}
	if err1 != nil && err2 != nil {
		return err1.Error() == err2.Error()
	}
	return false
}
