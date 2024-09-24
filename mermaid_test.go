package flowchart

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestFlowchartDirectionEnum_toMermaid(t *testing.T) {
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
			got := tt.direction.toMermaid()
			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestArrowTypeEnum_toMermaidOriginString(t *testing.T) {
	tests := []struct {
		name     string
		arrow    ArrowTypeEnum
		expected string
	}{
		{
			name:     "Normal arrow",
			arrow:    ArrowTypeNormal,
			expected: "<",
		},
		{
			name:     "Cross arrow",
			arrow:    ArrowTypeCross,
			expected: "x",
		},
		{
			name:     "Circle arrow",
			arrow:    ArrowTypeCircle,
			expected: "o",
		},
		{
			name:     "No arrow",
			arrow:    ArrowTypeNone,
			expected: "",
		},
		{
			name:     "Invalid arrow",
			arrow:    ArrowTypeEnum(999), // Some invalid value
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.arrow.toMermaidOrigin()

			// Compare the result using cmp.Diff
			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf("toMermaidOrigin() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestArrowTypeEnum_toMermaidTargetString(t *testing.T) {
	tests := []struct {
		name     string
		arrow    ArrowTypeEnum
		expected string
	}{
		{
			name:     "Normal arrow",
			arrow:    ArrowTypeNormal,
			expected: ">",
		},
		{
			name:     "Cross arrow",
			arrow:    ArrowTypeCross,
			expected: "x",
		},
		{
			name:     "Circle arrow",
			arrow:    ArrowTypeCircle,
			expected: "o",
		},
		{
			name:     "No arrow",
			arrow:    ArrowTypeNone,
			expected: "",
		},
		{
			name:     "Invalid arrow",
			arrow:    ArrowTypeEnum(999), // Some invalid value
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.arrow.toMermaidTarget()

			// Compare the result using cmp.Diff
			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf("toMermaidTarget() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestLineTypeEnum_toMermaidOrigin(t *testing.T) {
	tests := []struct {
		name     string
		lineType LineTypeEnum
		expected string
	}{
		{
			name:     "Dotted line",
			lineType: LineTypeDotted,
			expected: "-.",
		},
		{
			name:     "Solid line",
			lineType: LineTypeSolid,
			expected: "--",
		},
		{
			name:     "Thick line",
			lineType: LineTypeThick,
			expected: "==",
		},
		{
			name:     "No line",
			lineType: LineTypeNone,
			expected: "",
		},
		{
			name:     "Invalid line type",
			lineType: LineTypeEnum(999), // Invalid value
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.lineType.toMermaidOrigin()

			// Compare result using cmp.Diff
			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf("toMermaidOrigin() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestLineTypeEnum_toMermaidTarget(t *testing.T) {
	tests := []struct {
		name     string
		lineType LineTypeEnum
		expected string
	}{
		{
			name:     "Dotted line",
			lineType: LineTypeDotted,
			expected: ".-",
		},
		{
			name:     "Solid line",
			lineType: LineTypeSolid,
			expected: "--",
		},
		{
			name:     "Thick line",
			lineType: LineTypeThick,
			expected: "==",
		},
		{
			name:     "No line",
			lineType: LineTypeNone,
			expected: "",
		},
		{
			name:     "Invalid line type",
			lineType: LineTypeEnum(999), // Invalid value
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.lineType.toMermaidTarget()

			// Compare result using cmp.Diff
			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf("toMermaidTarget() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestLineTypeEnum_toMermaidBidirectional(t *testing.T) {
	tests := []struct {
		name     string
		lineType LineTypeEnum
		expected string
	}{
		{
			name:     "Solid line",
			lineType: LineTypeSolid,
			expected: "---",
		},
		{
			name:     "Dotted line",
			lineType: LineTypeDotted,
			expected: "-.-",
		},
		{
			name:     "Thick line",
			lineType: LineTypeThick,
			expected: "===",
		},
		{
			name:     "No line",
			lineType: LineTypeNone,
			expected: "~~~",
		},
		{
			name:     "Invalid line type",
			lineType: LineTypeEnum(999), // Invalid value
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.lineType.toMermaidBidirectional()

			// Compare result using cmp.Diff
			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf("toMermaidBidirectional() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestLink_toMermaid(t *testing.T) {
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
			expected: "Origin -- Target",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.link.toMermaid()

			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestNode_toMermaid(t *testing.T) {
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
			got := tt.node.toMermaid(tt.indents)
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
			name:     "alternate process",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeAlternateProcess,
			label:    fixtureLabel,
			expected: "node([\"a label\"]);\n",
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
			got := node.toMermaid(0)
			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestFlowchart_toMermaid(t *testing.T) {
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
			got := tt.flowchart.toMermaid(tt.indents, true)
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

func TestFlowchart_ToMermaid(t *testing.T) {
	tests := []struct {
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
        subgraph SubgraphTwo [Subgraph Two];
            direction RL;
            NodeSeven;
            NodeEight;
        end;
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
			got, err := tt.flowchart.ToMermaid()
			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaid() mismatch (-expected +got):\n%s", diff)
			}
			if (err == nil) == tt.expectedErr {
				t.Errorf("ToMermaid() error = %v, expected %v", err, tt.expectedErr)
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
