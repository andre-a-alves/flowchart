package flowchart

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func TestMermaidFlowchartDirection(t *testing.T) {
	tests := []struct {
		name      string
		direction DirectionEnum
		expected  string
	}{
		{
			name:      "Direction horizontal - right",
			direction: DirectionHorizontalRight,
			expected:  "LR",
		},
		{
			name:      "Direction horizontal - left",
			direction: DirectionHorizontalLeft,
			expected:  "RL",
		},
		{
			name:      "Direction vertical",
			direction: DirectionVertical,
			expected:  "TB",
		},
		{
			name:      "Unknown direction (fallback)",
			direction: DirectionEnum(999), // Unknown value to test fallback
			expected:  "TB",               // Default fallback
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mermaidFlowchartDirection(tt.direction)
			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestRenderArrows(t *testing.T) {
	tests := []struct {
		name                string
		link                Link
		expectedOriginArrow string
		expectedTargetArrow string
	}{
		{
			name: "arrow type none",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				TargetArrow: false,
			},
			expectedOriginArrow: "",
			expectedTargetArrow: "",
		},
		{
			name: "target arrow no",
			link: Link{
				ArrowType:   ArrowTypeNormal,
				OriginArrow: true,
				TargetArrow: false,
			},
			expectedOriginArrow: "",
			expectedTargetArrow: "",
		},
		{
			name: "normal arrow, target arrow yes, origin arrow yes",
			link: Link{
				ArrowType:   ArrowTypeNormal,
				OriginArrow: true,
				TargetArrow: true,
			},
			expectedOriginArrow: "<",
			expectedTargetArrow: ">",
		},
		{
			name: "normal arrow, target arrow yes, origin arrow no",
			link: Link{
				ArrowType:   ArrowTypeNormal,
				OriginArrow: false,
				TargetArrow: true,
			},
			expectedOriginArrow: "",
			expectedTargetArrow: ">",
		},
		{
			name: "circle arrow, target arrow yes, origin arrow yes",
			link: Link{
				ArrowType:   ArrowTypeCircle,
				OriginArrow: true,
				TargetArrow: true,
			},
			expectedOriginArrow: "o",
			expectedTargetArrow: "o",
		},
		{
			name: "circle arrow, target arrow yes, origin arrow no",
			link: Link{
				ArrowType:   ArrowTypeCircle,
				OriginArrow: false,
				TargetArrow: true,
			},
			expectedOriginArrow: "",
			expectedTargetArrow: "o",
		},
		{
			name: "cross arrow, target arrow yes, origin arrow yes",
			link: Link{
				ArrowType:   ArrowTypeCross,
				OriginArrow: true,
				TargetArrow: true,
			},
			expectedOriginArrow: "x",
			expectedTargetArrow: "x",
		},
		{
			name: "cross arrow, target arrow yes, origin arrow no",
			link: Link{
				ArrowType:   ArrowTypeCross,
				OriginArrow: false,
				TargetArrow: true,
			},
			expectedOriginArrow: "",
			expectedTargetArrow: "x",
		},
		{
			name: "cross arrow, target arrow yes, origin arrow no",
			link: Link{
				ArrowType:   ArrowTypeEnum(-1),
				OriginArrow: true,
				TargetArrow: true,
			},
			expectedOriginArrow: "",
			expectedTargetArrow: "",
		},
		{
			name: "cross arrow, target arrow yes, origin arrow no",
			link: Link{
				ArrowType:   ArrowTypeEnum(-1),
				OriginArrow: false,
				TargetArrow: true,
			},
			expectedOriginArrow: "",
			expectedTargetArrow: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOriginArrow, gotTargetArrow := renderArrows(tt.link)

			if diff := cmp.Diff(tt.expectedOriginArrow, gotOriginArrow); diff != "" {
				t.Errorf("renderArrows() origin mismatch (-expected +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.expectedTargetArrow, gotTargetArrow); diff != "" {
				t.Errorf("renderArrows() target mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestRenderMermaidLink(t *testing.T) {
	fixtureOriginSubgraph := Flowchart{Title: pointTo("Origin")}
	fixtureTargetNode := Node{name: "Target"}

	tests := []struct {
		name     string
		link     Link
		expected string
	}{
		{
			name: "no target",
			link: Link{
				ArrowType:   ArrowTypeNormal,
				OriginArrow: false,
				LineType:    LineTypeSolid,
				TargetArrow: true,
				Label:       pointTo("some label"),
				Origin:      &fixtureTargetNode,
				Target:      nil,
			},
			expected: "",
		},
		{
			name: "no origin",
			link: Link{
				ArrowType:   ArrowTypeNormal,
				OriginArrow: false,
				LineType:    LineTypeSolid,
				TargetArrow: true,
				Label:       pointTo("some label"),
				Origin:      nil,
				Target:      &fixtureTargetNode,
			},
			expected: "",
		},
		{
			name: "no line",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeNone,
				TargetArrow: false,
				Label:       pointTo("some label"),
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "Origin ~~~ Target",
		},
		{
			name: "no line - node to subgraph",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeNone,
				TargetArrow: false,
				Label:       pointTo("some label"),
				Origin:      &fixtureTargetNode,
				Target:      &fixtureOriginSubgraph,
			},
			expected: "Target ~~~ Origin",
		},
		{
			name: "solid line - no label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeSolid,
				TargetArrow: false,
				Label:       nil,
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "Origin --- Target",
		},
		{
			name: "solid line - with label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeSolid,
				TargetArrow: false,
				Label:       pointTo("some label"),
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "Origin -- \"some label\" -- Target",
		},
		{
			name: "dotted line - no label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeDotted,
				TargetArrow: false,
				Label:       nil,
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "Origin -.- Target",
		},
		{
			name: "dotted line - with label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeDotted,
				TargetArrow: false,
				Label:       pointTo("some label"),
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "Origin -. \"some label\" .- Target",
		},
		{
			name: "thick line - no label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeThick,
				TargetArrow: false,
				Label:       nil,
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "Origin === Target",
		},
		{
			name: "thick line - no label with arrows",
			link: Link{
				ArrowType:   ArrowTypeCircle,
				OriginArrow: true,
				LineType:    LineTypeThick,
				TargetArrow: true,
				Label:       nil,
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "Origin o==o Target",
		},
		{
			name: "thick line - with label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeThick,
				TargetArrow: false,
				Label:       pointTo("some label"),
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "Origin == \"some label\" == Target",
		},
		{
			name: "only origin arrow yes",
			link: Link{
				ArrowType:   ArrowTypeNormal,
				OriginArrow: true,
				LineType:    LineTypeSolid,
				TargetArrow: false,
				Label:       nil,
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "Origin --- Target",
		},
		{
			name: "target arrow yes",
			link: Link{
				ArrowType:   ArrowTypeNormal,
				OriginArrow: false,
				LineType:    LineTypeSolid,
				TargetArrow: true,
				Label:       nil,
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "Origin --> Target",
		},
		{
			name: "both arrows",
			link: Link{
				ArrowType:   ArrowTypeNormal,
				OriginArrow: true,
				LineType:    LineTypeSolid,
				TargetArrow: true,
				Label:       nil,
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "Origin <--> Target",
		},
		{
			name: "unhandled line type with no label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeEnum(-1),
				TargetArrow: false,
				Label:       nil,
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "",
		},
		{
			name: "unhandled line type with label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeEnum(-1),
				TargetArrow: false,
				Label:       pointTo("label"),
				Origin:      &fixtureOriginSubgraph,
				Target:      &fixtureTargetNode,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderMermaidLink(tt.link)

			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestRenderMermaidNode(t *testing.T) {
	tests := []struct {
		name     string
		node     *Node
		indents  int
		expected string
	}{
		{
			name: "Node with no label and no links",
			node: &Node{
				name:  "First Node",
				Type:  NodeTypeProcess,
				Label: nil,
			},
			indents:  0,
			expected: "FirstNode;\n",
		},
		{
			name: "Node with empty label and no links",
			node: &Node{
				name:  "First Node",
				Type:  NodeTypeProcess,
				Label: pointTo(""),
			},
			indents:  1,
			expected: "    FirstNode;\n",
		},
		{
			name: "Node with label and no links",
			node: &Node{
				name:  "First Node",
				Type:  NodeTypeProcess,
				Label: pointTo("Node Label"),
			},
			indents:  2,
			expected: "        FirstNode[\"Node Label\"];\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderMermaidNode(tt.node, tt.indents)
			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
			}
		})
	}

	fixtureNodeName := "node"
	fixtureLabel := pointTo("a label")

	testTypes := []struct {
		name     string
		nodeName string
		nodeType NodeTypeEnum
		label    *string
		expected string
	}{
		{
			name:     "terminator",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeTerminator,
			label:    fixtureLabel,
			expected: "node(\"a label\");\n",
		},
		{
			name:     "process",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeProcess,
			label:    fixtureLabel,
			expected: "node[\"a label\"];\n",
		},
		{
			name:     "subprocess",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeSubprocess,
			label:    fixtureLabel,
			expected: "node[[\"a label\"]];\n",
		},
		{
			name:     "decision",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeDecision,
			label:    fixtureLabel,
			expected: "node{\"a label\"};\n",
		},
		{
			name:     "input/output",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeInputOutput,
			label:    fixtureLabel,
			expected: "node[/\"a label\"/];\n",
		},
		{
			name:     "connector",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeConnector,
			label:    fixtureLabel,
			expected: "node((\"a label\"));\n",
		},
		{
			name:     "database",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeDatabase,
			label:    fixtureLabel,
			expected: "node[(\"a label\")];\n",
		},
		{
			name:     "invalid",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeEnum(-1),
			label:    fixtureLabel,
			expected: "node(\"a label\");\n",
		},
	}
	for _, tt := range testTypes {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				name:  tt.nodeName,
				Type:  tt.nodeType,
				Label: tt.label,
			}
			got := renderMermaidNode(node, 0)
			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestRenderMermaidFlowchart(t *testing.T) {
	tests := []struct {
		name        string
		flowchart   *Flowchart
		indents     int
		expected    string
		expectPanic bool
	}{
		{
			name: "Flowchart with no title",
			flowchart: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     nil,
				Nodes:     []*Node{{name: "Node One"}},
				Subgraphs: nil,
			},
			indents:     1,
			expected:    "",
			expectPanic: true,
		},
		{
			name: "Flowchart with no subgraphs",
			flowchart: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("Main Title"),
				Nodes: []*Node{
					{name: "First Node"},
					{name: "Second Node"},
				},
				Subgraphs: []*Flowchart{},
			},
			indents:     1,
			expected:    "    subgraph MainTitle [Main Title];\n        direction TB;\n        FirstNode;\n        SecondNode;\n    end;\n",
			expectPanic: false,
		},
		{
			name: "Flowchart with subgraphs",
			flowchart: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("Main Title"),
				Nodes:     []*Node{{name: "First Node"}},
				Subgraphs: []*Flowchart{
					{
						Direction: DirectionHorizontalLeft,
						Title:     pointTo("Subgraph One"),
						Nodes:     []*Node{{name: "Second Node"}},
					},
					{
						Direction: DirectionHorizontalRight,
						Title:     pointTo("Subgraph Two"),
						Nodes:     []*Node{{name: "Third Node"}},
					},
				},
			},
			indents: 0,
			expected: `subgraph MainTitle [Main Title];
    direction TB;
    FirstNode;
    subgraph SubgraphOne [Subgraph One];
        direction RL;
        SecondNode;
    end;
    subgraph SubgraphTwo [Subgraph Two];
        direction LR;
        ThirdNode;
    end;
end;
`,
			expectPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Catch panic if expected
			defer func() {
				if r := recover(); r != nil {
					if !tt.expectPanic {
						t.Errorf("unexpected panic: %v", r)
					}
				} else if tt.expectPanic {
					t.Error("expected panic but none occurred")
				}
			}()
			got := renderMermaidFlowchart(tt.flowchart, tt.indents, true)
			if tt.flowchart.Title == nil {
				if diff := cmp.Diff(tt.expected, got); diff != "" {
					t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
				}
			} else {
				if diff := cmp.Diff(tt.expected, got); diff != "" {
					t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
				}
			}
		})
	}
}

func TestRenderMermaid(t *testing.T) {
	tests := []struct {
		name        string
		flowchart   *Flowchart
		expected    string
		expectedErr bool
	}{
		{
			name: "Flowchart with no title",
			flowchart: &Flowchart{
				Direction: DirectionHorizontalRight,
				Nodes:     []*Node{{name: "Node One"}},
			},
			expected:    "flowchart LR;\n    NodeOne;\n",
			expectedErr: false,
		},
		{
			name:      "Full Flowchart with title",
			flowchart: fixtureFlowchart(),
			expected: `---
title: Test Title
---
flowchart LR;
    NodeOne;
    NodeTwo;
    NodeThree;
    NodeFour;
    subgraph SubgraphOne [Subgraph One];
        direction TB;
        NodeFive;
        NodeSix;
        NodeSeven;
        NodeEight;
    end;
    NodeFive --> NodeSix;
    NodeOne --> NodeTwo;
    NodeSeven --> NodeEight;
    NodeTwo --> NodeFour;
    NodeTwo --> NodeThree;
`,
			expectedErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RenderMermaid(GetMermaidFriendlyFlowchart(tt.flowchart))
			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
			}
			if (err == nil) == tt.expectedErr {
				t.Errorf("ToMermaid() error = %v, expected %v", err, tt.expectedErr)
			}
		})
	}

	invalidTests := []struct {
		name        string
		flowchart   *Flowchart
		expected    string
		expectedErr bool
	}{
		{
			name: "invalid mermaid name",
			flowchart: &Flowchart{
				Direction: DirectionHorizontalRight,
				Nodes:     []*Node{{name: "("}},
			},
			expected:    "",
			expectedErr: true,
		},
	}
	for _, tt := range invalidTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RenderMermaid(tt.flowchart)
			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
			}
			if (err == nil) == tt.expectedErr {
				t.Errorf("ToMermaid() error = %v, expected %v", err, tt.expectedErr)
			}
		})
	}
}

func TestValidateMermaid(t *testing.T) {
	// Helper nodes and subgraphs
	validNode := &Node{name: "ValidNode", Type: NodeTypeProcess, Label: pointTo("Valid Node")}
	invalidNode := &Node{name: "Invalid@Node", Type: NodeTypeProcess, Label: pointTo("Invalid Node")}
	duplicateNode := &Node{name: "DuplicateNode", Type: NodeTypeProcess, Label: pointTo("Duplicate Node")}
	subgraphWithNested := &Flowchart{
		Title:     pointTo("SubgraphWithNested"),
		Direction: DirectionVertical,
		Nodes:     []*Node{{name: "NestedNode"}},
		Subgraphs: []*Flowchart{
			{
				Title:     pointTo("NestedSubgraph"),
				Direction: DirectionHorizontalRight,
				Nodes:     []*Node{{name: "AnotherNestedNode"}},
			},
		},
	}
	subgraphWithoutNested := &Flowchart{
		Title:     pointTo("SubgraphWithoutNested"),
		Direction: DirectionHorizontalLeft,
		Nodes:     []*Node{{name: "SubgraphNode"}},
		Subgraphs: []*Flowchart{},
	}

	tests := []struct {
		name          string
		flowchart     *Flowchart
		expectedError string
	}{
		{
			name: "Valid flowchart with unique, valid names and no nested subgraphs",
			flowchart: &Flowchart{
				Title:     pointTo("MainFlowchart"),
				Direction: DirectionVertical,
				Nodes:     []*Node{validNode},
				Subgraphs: []*Flowchart{subgraphWithoutNested},
			},
			expectedError: "",
		},
		{
			name: "Invalid flowchart with invalid node names",
			flowchart: &Flowchart{
				Title:     pointTo("MainFlowchart"),
				Direction: DirectionVertical,
				Nodes:     []*Node{validNode, invalidNode},
				Subgraphs: []*Flowchart{subgraphWithoutNested},
			},
			expectedError: "flowchart contains violations: contains invalid mermaid names",
		},
		{
			name: "Invalid flowchart with nested subgraphs",
			flowchart: &Flowchart{
				Title:     pointTo("MainFlowchart"),
				Direction: DirectionHorizontalRight,
				Nodes:     []*Node{validNode},
				Subgraphs: []*Flowchart{subgraphWithNested},
			},
			expectedError: "flowchart contains violations: contains nested subgraphs",
		},
		{
			name: "Invalid flowchart with duplicate node names",
			flowchart: &Flowchart{
				Title:     pointTo("MainFlowchart"),
				Direction: DirectionHorizontalLeft,
				Nodes:     []*Node{duplicateNode, duplicateNode},
				Subgraphs: []*Flowchart{subgraphWithoutNested},
			},
			expectedError: "flowchart contains violations: contains repeated node and/or subgraph names",
		},
		{
			name: "Invalid flowchart with duplicate subgraph titles",
			flowchart: &Flowchart{
				Title:     pointTo("MainFlowchart"),
				Direction: DirectionVertical,
				Nodes:     []*Node{validNode},
				Subgraphs: []*Flowchart{
					subgraphWithoutNested,
					{
						Title:     pointTo("SubgraphWithoutNested"),
						Direction: DirectionVertical,
						Nodes:     []*Node{{name: "AnotherNode"}},
						Subgraphs: []*Flowchart{},
					},
				},
			},
			expectedError: "flowchart contains violations: contains repeated node and/or subgraph names",
		},
		{
			name: "Invalid flowchart with multiple violations: invalid names and nested subgraphs",
			flowchart: &Flowchart{
				Title:     pointTo("MainFlowchart"),
				Direction: DirectionVertical,
				Nodes:     []*Node{validNode, invalidNode},
				Subgraphs: []*Flowchart{subgraphWithNested},
			},
			expectedError: "flowchart contains violations: contains invalid mermaid names, contains nested subgraphs",
		},
		{
			name: "Invalid flowchart with all three violations: invalid names, nested subgraphs, and duplicate names",
			flowchart: &Flowchart{
				Title:     pointTo("MainFlowchart"),
				Direction: DirectionHorizontalRight,
				Nodes:     []*Node{validNode, invalidNode, duplicateNode, duplicateNode},
				Subgraphs: []*Flowchart{subgraphWithNested},
			},
			expectedError: "flowchart contains violations: contains invalid mermaid names, contains nested subgraphs, contains repeated node and/or subgraph names",
		},
		{
			name: "Valid flowchart with multiple unique subgraphs and nodes",
			flowchart: &Flowchart{
				Title:     pointTo("MainFlowchart"),
				Direction: DirectionVertical,
				Nodes:     []*Node{validNode, duplicateNode}, // Assuming "duplicateNode" name is unique here
				Subgraphs: []*Flowchart{
					{
						Title:     pointTo("Subgraph1"),
						Direction: DirectionHorizontalRight,
						Nodes:     []*Node{{name: "Subgraph1Node1"}},
					},
					{
						Title:     pointTo("Subgraph2"),
						Direction: DirectionHorizontalLeft,
						Nodes:     []*Node{{name: "Subgraph2Node1"}},
					},
				},
			},
			expectedError: "",
		},
		{
			name: "Empty flowchart (no nodes, no subgraphs)",
			flowchart: &Flowchart{
				Title:     pointTo("EmptyFlowchart"),
				Direction: DirectionVertical,
				Nodes:     []*Node{},
				Subgraphs: []*Flowchart{},
			},
			expectedError: "",
		},
		{
			name: "Flowchart with no title but valid nodes and no subgraphs",
			flowchart: &Flowchart{
				Title:     nil,
				Direction: DirectionHorizontalRight,
				Nodes:     []*Node{validNode},
				Subgraphs: []*Flowchart{},
			},
			expectedError: "",
		},
		{
			name: "Subgraph without title",
			flowchart: &Flowchart{
				Title:     pointTo("MainFlowchart"),
				Direction: DirectionVertical,
				Nodes:     []*Node{validNode},
				Subgraphs: []*Flowchart{
					{
						Title:     nil,
						Direction: DirectionHorizontalRight,
						Nodes:     []*Node{{name: "SubgraphNode"}},
					},
				},
			},
			expectedError: "flowchart contains violations: contains invalid mermaid names",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateMermaid(tt.flowchart)
			if tt.expectedError == "" && err != nil {
				t.Errorf("validateMermaid() unexpected error: %v", err)
			} else if tt.expectedError != "" {
				if err == nil {
					t.Errorf("validateMermaid() expected error '%s', but got nil", tt.expectedError)
				} else if err.Error() != tt.expectedError {
					t.Errorf("validateMermaid() expected error '%s', got '%s'", tt.expectedError, err.Error())
				}
			}
		})
	}
}

func fixtureFlowchart() *Flowchart {
	nodeOne := ProcessNode("Node One", nil)
	nodeTwo := ProcessNode("Node Two", nil)
	nodeThree := ProcessNode("Node Three", nil)
	nodeFour := ProcessNode("Node Four", nil)
	nodeFive := ProcessNode("Node Five", nil)
	nodeSix := ProcessNode("Node Six", nil)
	nodeSeven := ProcessNode("Node Seven", nil)
	nodeEight := ProcessNode("Node Eight", nil)

	nodeOneLink := SolidLink(nodeOne, nodeTwo, nil)
	nodeTwoLinkOne := SolidLink(nodeTwo, nodeThree, nil)
	nodeTwoLinkTwo := SolidLink(nodeTwo, nodeFour, nil)
	nodeFiveLink := SolidLink(nodeFive, nodeSix, nil)
	nodeSevenLink := SolidLink(nodeSeven, nodeEight, nil)

	return &Flowchart{
		Direction: DirectionHorizontalRight,
		Title:     pointTo("Test Title"),
		Nodes:     []*Node{nodeOne, nodeTwo, nodeThree, nodeFour},
		Links:     []Link{nodeOneLink, nodeTwoLinkOne, nodeTwoLinkTwo},
		Subgraphs: []*Flowchart{
			{
				Direction: DirectionVertical,
				Nodes:     []*Node{nodeFive, nodeSix},
				Title:     pointTo("Subgraph One"),
				Links:     []Link{nodeFiveLink},
				Subgraphs: []*Flowchart{
					{
						Direction: DirectionHorizontalLeft,
						Title:     pointTo("Subgraph Two"),
						Nodes:     []*Node{nodeSeven, nodeEight},
						Links:     []Link{nodeSevenLink},
					},
				},
			},
		},
	}
}

func TestHasValidMermaidNames(t *testing.T) {
	testCases := []struct {
		name      string
		flowchart Flowchart
		expected  bool
	}{
		{
			name: "happy path",
			flowchart: Flowchart{
				Nodes: []*Node{
					{name: "ValidNode1"},
					{name: "Valid_Node_2"},
					{name: "Valid-Node-3"},
					{name: "Valid Node 4"}, // With space
				},
				Subgraphs: []*Flowchart{
					{
						Title: pointTo("Subgraph1"),
						Nodes: []*Node{
							{name: "SubNode1"},
							{name: "SubNode2"},
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Invalid node name with special character",
			flowchart: Flowchart{
				Nodes: []*Node{
					{name: "ValidNode"},
					{name: "Invalid@Node"},
				},
			},
			expected: false,
		},
		{
			name: "Subgraph with invalid node name",
			flowchart: Flowchart{
				Nodes: []*Node{
					{name: "MainNode"},
				},
				Subgraphs: []*Flowchart{
					{
						Title: pointTo("Subgraph1"),
						Nodes: []*Node{
							{name: "ValidSubNode"},
							{name: "Invalid!SubNode"},
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "Subgraph with no title",
			flowchart: Flowchart{
				Nodes: []*Node{
					{name: "MainNode"},
				},
				Subgraphs: []*Flowchart{
					{
						Title: nil,
						Nodes: []*Node{
							{name: "ValidSubNode"},
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "Empty subgraph title",
			flowchart: Flowchart{
				Nodes: []*Node{
					{name: "MainNode"},
				},
				Subgraphs: []*Flowchart{
					{
						Title: pointTo(""),
						Nodes: []*Node{
							{name: "ValidSubNode"},
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "nested subgraph",
			flowchart: Flowchart{
				Nodes: []*Node{
					{name: "MainNode"},
				},
				Subgraphs: []*Flowchart{
					{
						Title: pointTo("Valid Subgraph"),
						Subgraphs: []*Flowchart{{
							Title: pointTo("Also Valid Subgraph"),
							Nodes: []*Node{
								{name: "Inv@alid"},
							},
						}},
					},
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := hasValidMermaidNames(&tc.flowchart)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestIsValidMermaidNodeName(t *testing.T) {
	testCases := []struct {
		name     string
		testCase string
		expected bool
	}{
		{
			"extreme valid case",
			"Valid test-case_1",
			true,
		},
		{
			"parenthesis",
			"(",
			false,
		},
		{
			"random special character",
			"@",
			false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidMermaidNodeName(tt.testCase)
			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("isValidMermaidNodeName() mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenFlowchart(t *testing.T) {
	// Define helper nodes
	node1 := &Node{name: "Node1", Type: NodeTypeProcess, Label: pointTo("Process 1")}
	node2 := &Node{name: "Node2", Type: NodeTypeDecision, Label: pointTo("Decision 2")}
	node3 := &Node{name: "Node3", Type: NodeTypeProcess, Label: pointTo("Process 3")}
	node4 := &Node{name: "Node4", Type: NodeTypeDatabase, Label: pointTo("Database 4")}
	node5 := &Node{name: "Node5", Type: NodeTypeSubprocess, Label: pointTo("Start 5")}

	// Define subgraphs
	emptySubgraph := &Flowchart{
		Title:     pointTo("EmptySubgraph"),
		Direction: DirectionVertical,
		Nodes:     []*Node{},
		Subgraphs: []*Flowchart{},
		Links:     []Link{},
	}

	subgraphWithNodes := &Flowchart{
		Title:     pointTo("SubgraphWithNodes"),
		Direction: DirectionHorizontalRight,
		Nodes:     []*Node{node2},
		Subgraphs: []*Flowchart{},
		Links: []Link{
			{
				Origin:      node1,
				Target:      node2,
				LineType:    LineTypeSolid,
				ArrowType:   ArrowTypeNormal,
				OriginArrow: false,
				TargetArrow: true,
				Label:       pointTo("Link1"),
			},
		},
	}

	nestedSubgraph := &Flowchart{
		Title:     pointTo("NestedSubgraph"),
		Direction: DirectionVertical,
		Nodes:     []*Node{node3},
		Subgraphs: []*Flowchart{
			{
				Title:     pointTo("DeepNestedSubgraph"),
				Direction: DirectionHorizontalRight,
				Nodes:     []*Node{node4},
				Subgraphs: []*Flowchart{},
				Links: []Link{
					{
						Origin:      node3,
						Target:      node4,
						LineType:    LineTypeDotted,
						ArrowType:   ArrowTypeCross,
						OriginArrow: true,
						TargetArrow: false,
						Label:       pointTo("Link2"),
					},
				},
			},
		},
		Links: []Link{},
	}

	subgraphWithEmptySubgraph := &Flowchart{
		Title:     pointTo("SubgraphWithEmptySubgraph"),
		Direction: DirectionVertical,
		Nodes:     []*Node{node5},
		Subgraphs: []*Flowchart{emptySubgraph},
		Links:     []Link{},
	}

	// Define test cases
	tests := []struct {
		name         string
		flowchart    *Flowchart
		expectedFlow *Flowchart
	}{
		{
			name: "Empty flowchart",
			flowchart: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("EmptyFlowchart"),
				Nodes:     []*Node{},
				Subgraphs: []*Flowchart{},
				Links:     []Link{},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("EmptyFlowchart"),
				Nodes:     []*Node{},
				Subgraphs: nil,
				Links:     []Link{},
			},
		},
		{
			name: "Flowchart with nodes and links, no subgraphs",
			flowchart: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     pointTo("SimpleFlowchart"),
				Nodes:     []*Node{node1, node2},
				Subgraphs: []*Flowchart{},
				Links: []Link{
					{
						Origin:      node1,
						Target:      node2,
						LineType:    LineTypeSolid,
						ArrowType:   ArrowTypeNormal,
						OriginArrow: false,
						TargetArrow: true,
						Label:       pointTo("Link1"),
					},
				},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     pointTo("SimpleFlowchart"),
				Nodes:     []*Node{node1, node2},
				Subgraphs: nil,
				Links: []Link{
					{
						Origin:      node1,
						Target:      node2,
						LineType:    LineTypeSolid,
						ArrowType:   ArrowTypeNormal,
						OriginArrow: false,
						TargetArrow: true,
						Label:       pointTo("Link1"),
					},
				},
			},
		},
		{
			name: "Flowchart with empty subgraphs",
			flowchart: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("FlowchartWithEmptySubgraphs"),
				Nodes:     []*Node{node1},
				Subgraphs: []*Flowchart{emptySubgraph},
				Links:     []Link{},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("FlowchartWithEmptySubgraphs"),
				Nodes:     []*Node{node1},
				Subgraphs: nil,
				Links:     []Link{},
			},
		},
		{
			name: "Flowchart with subgraphs that have nodes and links",
			flowchart: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("FlowchartWithSubgraphs"),
				Nodes:     []*Node{node1},
				Subgraphs: []*Flowchart{subgraphWithNodes},
				Links:     []Link{},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("FlowchartWithSubgraphs"),
				Nodes:     []*Node{node1, node2},
				Subgraphs: nil,
				Links: []Link{
					{
						Origin:      node1,
						Target:      node2,
						LineType:    LineTypeSolid,
						ArrowType:   ArrowTypeNormal,
						OriginArrow: false,
						TargetArrow: true,
						Label:       pointTo("Link1"),
					},
				},
			},
		},
		{
			name: "Flowchart with nested subgraphs",
			flowchart: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     pointTo("FlowchartWithNestedSubgraphs"),
				Nodes:     []*Node{node1},
				Subgraphs: []*Flowchart{nestedSubgraph},
				Links:     []Link{},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     pointTo("FlowchartWithNestedSubgraphs"),
				Nodes:     []*Node{node1, node3, node4},
				Subgraphs: nil,
				Links: []Link{
					{
						Origin:      node3,
						Target:      node4,
						LineType:    LineTypeDotted,
						ArrowType:   ArrowTypeCross,
						OriginArrow: true,
						TargetArrow: false,
						Label:       pointTo("Link2"),
					},
				},
			},
		},
		{
			name: "Flowchart with subgraphs containing empty subgraphs",
			flowchart: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("FlowchartWithSubgraphsAndEmptySubgraphs"),
				Nodes:     []*Node{node1},
				Subgraphs: []*Flowchart{subgraphWithEmptySubgraph},
				Links:     []Link{},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("FlowchartWithSubgraphsAndEmptySubgraphs"),
				Nodes:     []*Node{node1, node5},
				Subgraphs: nil,
				Links:     []Link{},
			},
		},
		{
			name: "Flowchart with multiple subgraphs and links",
			flowchart: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("ComplexFlowchart"),
				Nodes:     []*Node{node1},
				Subgraphs: []*Flowchart{subgraphWithNodes, nestedSubgraph},
				Links: []Link{
					{
						Origin:      node1,
						Target:      node3,
						LineType:    LineTypeThick,
						ArrowType:   ArrowTypeCircle,
						OriginArrow: true,
						TargetArrow: true,
						Label:       pointTo("Link3"),
					},
				},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("ComplexFlowchart"),
				Nodes:     []*Node{node1, node2, node3, node4},
				Subgraphs: nil,
				Links: []Link{
					{
						Origin:      node1,
						Target:      node3,
						LineType:    LineTypeThick,
						ArrowType:   ArrowTypeCircle,
						OriginArrow: true,
						TargetArrow: true,
						Label:       pointTo("Link3"),
					},
					{
						Origin:      node1,
						Target:      node2,
						LineType:    LineTypeSolid,
						ArrowType:   ArrowTypeNormal,
						OriginArrow: false,
						TargetArrow: true,
						Label:       pointTo("Link1"),
					},
					{
						Origin:      node3,
						Target:      node4,
						LineType:    LineTypeDotted,
						ArrowType:   ArrowTypeCross,
						OriginArrow: true,
						TargetArrow: false,
						Label:       pointTo("Link2"),
					},
				},
			},
		},
	}

	// Execute tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := flattenFlowchart(tt.flowchart)
			if diff := cmp.Diff(tt.expectedFlow, result, cmp.AllowUnexported(Node{}), cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("flattenFlowchart() mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestRemoveNonMermaidNames(t *testing.T) {
	// Define helper nodes
	validNode1 := &Node{name: "ValidNode1", Type: NodeTypeProcess, Label: pointTo("Valid Node 1")}
	validNode2 := &Node{name: "ValidNode2", Type: NodeTypeDecision, Label: pointTo("Valid Node 2")}
	invalidNode := &Node{name: "Invalid@Node", Type: NodeTypeProcess, Label: pointTo("Invalid Node")}

	// Define subgraph titles
	validTitle := "ValidSubgraph"
	invalidTitle := "Invalid@Subgraph"
	emptyTitle := ""
	nilTitle := (*string)(nil)

	// Subgraph with valid title
	validSubgraph := &Flowchart{
		Title:     &validTitle,
		Direction: DirectionVertical,
		Nodes:     []*Node{validNode1},
		Subgraphs: []*Flowchart{},
		Links:     []Link{},
	}

	// Subgraph with invalid title
	invalidSubgraph := &Flowchart{
		Title:     &invalidTitle,
		Direction: DirectionHorizontalRight,
		Nodes:     []*Node{validNode2},
		Subgraphs: []*Flowchart{},
		Links:     []Link{},
	}

	// Subgraph with empty title
	emptyTitleSubgraph := &Flowchart{
		Title:     &emptyTitle,
		Direction: DirectionHorizontalRight,
		Nodes:     []*Node{invalidNode},
		Subgraphs: []*Flowchart{},
		Links:     []Link{},
	}

	// Subgraph with nil title
	nilTitleSubgraph := &Flowchart{
		Title:     nilTitle,
		Direction: DirectionVertical,
		Nodes:     []*Node{invalidNode},
		Subgraphs: []*Flowchart{},
		Links:     []Link{},
	}

	tests := []struct {
		name         string
		flowchart    *Flowchart
		expectedFlow *Flowchart
	}{
		{
			name: "Flowchart with valid nodes and valid subgraphs",
			flowchart: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("MainFlowchart"),
				Nodes:     []*Node{validNode1},
				Subgraphs: []*Flowchart{validSubgraph},
				Links:     []Link{},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("MainFlowchart"),
				Nodes:     []*Node{validNode1},
				Subgraphs: []*Flowchart{
					{
						Direction: validSubgraph.Direction,
						Title:     validSubgraph.Title,
						Nodes:     []*Node{validNode1},
						Subgraphs: []*Flowchart{},
						Links:     []Link{},
					},
				},
				Links: []Link{},
			},
		},
		{
			name: "Flowchart with invalid nodes and invalid subgraphs",
			flowchart: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     pointTo("MainFlowchart"),
				Nodes:     []*Node{invalidNode},
				Subgraphs: []*Flowchart{invalidSubgraph},
				Links:     []Link{},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     pointTo("MainFlowchart"),
				Nodes:     []*Node{},      // Invalid nodes are excluded
				Subgraphs: []*Flowchart{}, // Invalid subgraphs are excluded
				Links:     []Link{},
			},
		},
		{
			name: "Flowchart with nodes and subgraphs with nil and empty titles",
			flowchart: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("MainFlowchart"),
				Nodes:     []*Node{validNode1, invalidNode},
				Subgraphs: []*Flowchart{nilTitleSubgraph, emptyTitleSubgraph},
				Links: []Link{
					{
						Origin:      validNode1,
						Target:      invalidNode,
						LineType:    LineTypeSolid,
						ArrowType:   ArrowTypeNormal,
						OriginArrow: false,
						TargetArrow: true,
						Label:       pointTo("Link1"),
					},
				},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("MainFlowchart"),
				Nodes:     []*Node{validNode1}, // invalidNode is excluded
				Subgraphs: []*Flowchart{},      // Subgraphs with nil or empty titles are excluded
				Links:     []Link{},            // Link involving invalid node is excluded
			},
		},
		{
			name: "Flowchart with mixed valid and invalid nodes and subgraphs",
			flowchart: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     pointTo("MainFlowchart"),
				Nodes:     []*Node{validNode1, invalidNode},
				Subgraphs: []*Flowchart{validSubgraph, invalidSubgraph},
				Links: []Link{
					{
						Origin:      validNode1,
						Target:      invalidNode,
						LineType:    LineTypeSolid,
						ArrowType:   ArrowTypeNormal,
						OriginArrow: true,
						TargetArrow: false,
						Label:       pointTo("Link1"),
					},
					{
						Origin:      invalidNode,
						Target:      validNode1,
						LineType:    LineTypeDotted,
						ArrowType:   ArrowTypeCross,
						OriginArrow: false,
						TargetArrow: true,
						Label:       pointTo("Link2"),
					},
				},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     pointTo("MainFlowchart"),
				Nodes:     []*Node{validNode1},
				Subgraphs: []*Flowchart{
					{
						Direction: validSubgraph.Direction,
						Title:     validSubgraph.Title,
						Nodes:     []*Node{validNode1},
						Subgraphs: []*Flowchart{},
						Links:     []Link{},
					},
				},
				Links: []Link{}, // Links involving invalid nodes are excluded
			},
		},
		{
			name: "Flowchart with nested invalid subgraphs",
			flowchart: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("MainFlowchart"),
				Nodes:     []*Node{},
				Subgraphs: []*Flowchart{
					{
						Title:     &invalidTitle, // Invalid title
						Direction: DirectionHorizontalRight,
						Nodes:     []*Node{validNode2},
						Subgraphs: []*Flowchart{
							validSubgraph,
						},
						Links: []Link{},
					},
				},
				Links: []Link{},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("MainFlowchart"),
				Nodes:     []*Node{},
				Subgraphs: []*Flowchart{}, // Invalid subgraphs are excluded
				Links:     []Link{},
			},
		},
		{
			name: "Empty flowchart",
			flowchart: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     pointTo("EmptyFlowchart"),
				Nodes:     []*Node{},
				Subgraphs: []*Flowchart{},
				Links:     []Link{},
			},
			expectedFlow: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     pointTo("EmptyFlowchart"),
				Nodes:     []*Node{},
				Subgraphs: []*Flowchart{},
				Links:     []Link{},
			},
		},
	}

	// Execute tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeNonMermaidNames(tt.flowchart)
			if diff := cmp.Diff(tt.expectedFlow, result, cmp.AllowUnexported(Node{}), cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("removeNonMermaidNames() mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestGetMermaidFriendlyFlowchart(t *testing.T) {
	// Define helper nodes
	validNode1 := &Node{name: "ValidNode1", Type: NodeTypeProcess, Label: pointTo("Valid Node 1")}
	validNode2 := &Node{name: "ValidNode2", Type: NodeTypeDecision, Label: pointTo("Valid Node 2")}
	invalidNode := &Node{name: "Invalid@Node", Type: NodeTypeProcess, Label: pointTo("Invalid Node")}

	firstLink := Link{
		Origin:      validNode1,
		Target:      validNode1,
		LineType:    LineTypeSolid,
		ArrowType:   ArrowTypeNormal,
		OriginArrow: false,
		TargetArrow: true,
		Label:       nil,
	}
	secondLink := Link{
		Origin:      validNode1,
		Target:      validNode2,
		LineType:    LineTypeDotted,
		ArrowType:   ArrowTypeCross,
		OriginArrow: false,
		TargetArrow: true,
		Label:       nil,
	}

	originalFlowchart := &Flowchart{
		Direction: DirectionHorizontalLeft,
		Title:     pointTo("graph title"),
		Nodes:     []*Node{validNode1},
		Subgraphs: []*Flowchart{
			{
				Direction: DirectionHorizontalRight,
				Title:     pointTo("valid subgraph"),
				Nodes:     []*Node{invalidNode},
				Subgraphs: []*Flowchart{{
					Direction: DirectionVertical,
					Title:     nil,
					Nodes:     []*Node{validNode2},
					Subgraphs: nil,
					Links:     []Link{secondLink},
				}},
				Links: nil,
			},
		},
		Links: []Link{firstLink},
	}

	// Expected flowchart after processing
	expectedFlowchart := &Flowchart{
		Direction: DirectionHorizontalLeft,
		Title:     pointTo("graph title"),
		Nodes:     []*Node{validNode1},
		Subgraphs: []*Flowchart{{
			Title: pointTo("valid subgraph"),
			Nodes: []*Node{validNode2},
			Links: []Link{secondLink},
		}},
		Links: []Link{firstLink},
	}

	// Run the function
	result := GetMermaidFriendlyFlowchart(originalFlowchart)

	// Compare the result with the expected flowchart
	if diff := cmp.Diff(expectedFlowchart, result, cmp.AllowUnexported(Node{}), cmpopts.EquateEmpty()); diff != "" {
		t.Errorf("GetMermaidFriendlyFlowchart() mismatch (-expected +got):\n%s", diff)
	}
}
