package flowchart

import (
	"github.com/google/go-cmp/cmp"
	"regexp"
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

func TestNodeTypeEnum_toMermaidLeft(t *testing.T) {
	tests := []struct {
		name     string
		nodeType NodeTypeEnum
		expected string
	}{
		{
			name:     "Terminator left",
			nodeType: NodeTypeTerminator,
			expected: "(",
		},
		{
			name:     "Process left",
			nodeType: NodeTypeProcess,
			expected: "[",
		},
		{
			name:     "AlternateProcess left",
			nodeType: NodeTypeAlternateProcess,
			expected: "([",
		},
		{
			name:     "Subprocess left",
			nodeType: NodeTypeSubprocess,
			expected: "[[",
		},
		{
			name:     "Decision left",
			nodeType: NodeTypeDecision,
			expected: "{",
		},
		{
			name:     "InputOutput left",
			nodeType: NodeTypeInputOutput,
			expected: "[/",
		},
		{
			name:     "Connector left",
			nodeType: NodeTypeConnector,
			expected: "((",
		},
		{
			name:     "Database left",
			nodeType: NodeTypeDatabase,
			expected: "[(",
		},
		{
			name:     "Default left (invalid node type)",
			nodeType: NodeTypeEnum(-1), // Unknown NodeTypeEnum
			expected: "(",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.nodeType.toMermaidLeft()

			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaidLeft() mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestNodeTypeEnum_toMermaidRight(t *testing.T) {
	tests := []struct {
		name     string
		nodeType NodeTypeEnum
		expected string
	}{
		{
			name:     "Terminator right",
			nodeType: NodeTypeTerminator,
			expected: ")",
		},
		{
			name:     "Process right",
			nodeType: NodeTypeProcess,
			expected: "]",
		},
		{
			name:     "AlternateProcess right",
			nodeType: NodeTypeAlternateProcess,
			expected: "])",
		},
		{
			name:     "Subprocess right",
			nodeType: NodeTypeSubprocess,
			expected: "]]",
		},
		{
			name:     "Decision right",
			nodeType: NodeTypeDecision,
			expected: "}",
		},
		{
			name:     "InputOutput right",
			nodeType: NodeTypeInputOutput,
			expected: "/]",
		},
		{
			name:     "Connector right",
			nodeType: NodeTypeConnector,
			expected: "))",
		},
		{
			name:     "Database right",
			nodeType: NodeTypeDatabase,
			expected: ")]",
		},
		{
			name:     "Default right (invalid node type)",
			nodeType: NodeTypeEnum(-1), // Unknown NodeTypeEnum
			expected: ")",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.nodeType.toMermaidRight()

			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaidRight() mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestLink_toMermaid(t *testing.T) {
	fixtureTargetNode := Node{name: "Target Node"}

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
				TargetNode:  nil,
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
				TargetNode:  &fixtureTargetNode,
			},
			expected: "~~~ TargetNode",
		},
		{
			name: "solid line - no label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeSolid,
				TargetArrow: false,
				Label:       nil,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "--- TargetNode",
		},
		{
			name: "solid line - with label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeSolid,
				TargetArrow: false,
				Label:       pointTo("some label"),
				TargetNode:  &fixtureTargetNode,
			},
			expected: "-- some label -- TargetNode",
		},
		{
			name: "dotted line - no label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeDotted,
				TargetArrow: false,
				Label:       nil,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "-.- TargetNode",
		},
		{
			name: "dotted line - with label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeDotted,
				TargetArrow: false,
				Label:       pointTo("some label"),
				TargetNode:  &fixtureTargetNode,
			},
			expected: "-. some label .- TargetNode",
		},
		{
			name: "thick line - no label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeThick,
				TargetArrow: false,
				Label:       nil,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "=== TargetNode",
		},
		{
			name: "thick line - with label",
			link: Link{
				ArrowType:   ArrowTypeNone,
				OriginArrow: false,
				LineType:    LineTypeThick,
				TargetArrow: false,
				Label:       pointTo("some label"),
				TargetNode:  &fixtureTargetNode,
			},
			expected: "== some label == TargetNode",
		},
		{
			name: "origin arrow yes",
			link: Link{
				ArrowType:   ArrowTypeNormal,
				OriginArrow: true,
				LineType:    LineTypeSolid,
				TargetArrow: false,
				Label:       nil,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "<-- TargetNode",
		},
		{
			name: "target arrow yes",
			link: Link{
				ArrowType:   ArrowTypeNormal,
				OriginArrow: false,
				LineType:    LineTypeSolid,
				TargetArrow: true,
				Label:       nil,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "--> TargetNode",
		},
		{
			name: "both arrows",
			link: Link{
				ArrowType:   ArrowTypeNormal,
				OriginArrow: true,
				LineType:    LineTypeSolid,
				TargetArrow: true,
				Label:       nil,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "<--> TargetNode",
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

func TestNode_toMermaidNode(t *testing.T) {
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
				Links: nil,
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
				Links: nil,
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
				Links: nil,
			},
			indents:  2,
			expected: "        FirstNode[Node Label];\n",
		},
		{
			name: "Node with links and no label",
			node: &Node{
				name:  "First Node",
				Type:  NodeTypeProcess,
				Label: nil,
				Links: []Link{
					{
						ArrowType:   ArrowTypeNormal,
						OriginArrow: true,
						LineType:    LineTypeSolid,
						TargetArrow: false,
						Label:       nil,
						TargetNode: &Node{
							name:  "Target Node",
							Type:  NodeTypeProcess,
							Label: pointTo("ignored node label"),
							Links: nil,
						},
					},
				},
			},
			indents:  0,
			expected: "FirstNode;\nFirstNode <-- TargetNode;\n",
		},
		{
			name: "Node with links and label",
			node: &Node{
				name:  "First Node",
				Type:  NodeTypeProcess,
				Label: pointTo("Label"),
				Links: []Link{
					{
						ArrowType:   ArrowTypeCross,
						OriginArrow: false,
						LineType:    LineTypeDotted,
						TargetArrow: true,
						Label:       pointTo("Link Label"),
						TargetNode: &Node{
							name:  "Target Node",
							Type:  NodeTypeProcess,
							Label: nil,
							Links: nil,
						},
					},
				}},
			indents:  1,
			expected: "    FirstNode[Label];\n    FirstNode -. Link Label .-x TargetNode;\n",
		},
		{
			name: "Node with multiple links and label",
			node: &Node{
				name:  "First Node",
				Type:  NodeTypeProcess,
				Label: pointTo("Node Label"),
				Links: []Link{
					{
						ArrowType:   ArrowTypeNormal,
						OriginArrow: false,
						LineType:    LineTypeSolid,
						TargetArrow: true,
						Label:       nil,
						TargetNode: &Node{
							name:  "Target Node One",
							Type:  NodeTypeProcess,
							Label: nil,
							Links: nil,
						},
					},
					{
						ArrowType:   ArrowTypeCircle,
						OriginArrow: true,
						LineType:    LineTypeDotted,
						TargetArrow: true,
						Label:       pointTo("Link Label"),
						TargetNode: &Node{
							name:  "Target Node Two",
							Type:  NodeTypeProcess,
							Label: nil,
							Links: nil,
						},
					},
				},
			},
			indents:  2,
			expected: "        FirstNode[Node Label];\n        FirstNode --> TargetNodeOne;\n        FirstNode o-. Link Label .-o TargetNodeTwo;\n",
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
			expected: "node(a label);\n",
		},
		{
			name:     "process",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeProcess,
			label:    fixtureLabel,
			expected: "node[a label];\n",
		},
		{
			name:     "alternate process",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeAlternateProcess,
			label:    fixtureLabel,
			expected: "node([a label]);\n",
		},
		{
			name:     "subprocess",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeSubprocess,
			label:    fixtureLabel,
			expected: "node[[a label]];\n",
		},
		{
			name:     "decision",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeDecision,
			label:    fixtureLabel,
			expected: "node{a label};\n",
		},
		{
			name:     "input/output",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeInputOutput,
			label:    fixtureLabel,
			expected: "node[/a label/];\n",
		},
		{
			name:     "connector",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeConnector,
			label:    fixtureLabel,
			expected: "node((a label));\n",
		},
		{
			name:     "database",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeDatabase,
			label:    fixtureLabel,
			expected: "node[(a label)];\n",
		},
		{
			name:     "invalid",
			nodeName: fixtureNodeName,
			nodeType: NodeTypeEnum(-1),
			label:    fixtureLabel,
			expected: "node(a label);\n",
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

func TestFlowchart_toMermaidSubgraph(t *testing.T) {
	tests := []struct {
		name      string
		flowchart *Flowchart
		indents   int
		expected  string
	}{
		{
			name: "Flowchart with no title",
			flowchart: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     nil,
				Nodes:     []*Node{{name: "Node One"}},
				Subgraphs: nil,
			},
			indents:  1,
			expected: "    subgraph 123456;\n        direction LR;\n\n        NodeOne;\n    end;\n",
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
			indents:  1,
			expected: "    subgraph MainTitle [Main Title];\n        direction TB;\n\n        FirstNode;\n\n        SecondNode;\n    end;\n",
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.flowchart.toMermaidSubgraph(tt.indents)
			if tt.flowchart.Title == nil {
				if diff := cmp.Diff(maskUUIDsInSubgraph(tt.expected), maskUUIDsInSubgraph(got)); diff != "" {
					t.Errorf("toMermaidSubgraph() mismatch (-expected +got):\n%s", diff)
				}
			} else {
				if diff := cmp.Diff(tt.expected, got); diff != "" {
					t.Errorf("toMermaidSubgraph() mismatch (-expected +got):\n%s", diff)
				}
			}
		})
	}
}

func TestFlowchart_ToMermaid(t *testing.T) {
	tests := []struct {
		name      string
		flowchart *Flowchart
		expected  string
	}{
		{
			name: "Flowchart with no title",
			flowchart: &Flowchart{
				Direction: DirectionHorizontalRight,
				Title:     nil,
				Nodes:     []*Node{{name: "Node One"}},
				Subgraphs: nil,
			},
			expected: "flowchart LR;\n\n    NodeOne;\n",
		},
		{
			name:      "Full Flowchart with title",
			flowchart: fixtureFlowchart(),
			expected: `---
title: Test Title
---
flowchart LR;

    NodeOne;
    NodeOne --> NodeTwo;

    NodeTwo;
    NodeTwo --> NodeThree;
    NodeTwo --> NodeFour;

    NodeThree;

    NodeFour;

    subgraph 123456;
        direction TB;

        NodeFive;
        NodeFive --> NodeSix;

        NodeSix;

        subgraph 123456;
            direction RL;

            NodeSeven;
            NodeSeven --> NodeEight;

            NodeEight;
        end;
    end;
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.flowchart.ToMermaid()
			if diff := cmp.Diff(maskUUIDsInSubgraph(tt.expected), maskUUIDsInSubgraph(got)); diff != "" {
				t.Errorf("toMermaidSubgraph() mismatch (-expected +got):\n%s", diff)
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

	nodeOneLink := SolidLink(nodeTwo, nil)
	nodeTwoLinkOne := SolidLink(nodeThree, nil)
	nodeTwoLinkTwo := SolidLink(nodeFour, nil)
	nodeFiveLink := SolidLink(nodeSix, nil)
	nodeSevenLink := SolidLink(nodeEight, nil)

	nodeOne.Links = []Link{nodeOneLink}
	nodeTwo.Links = []Link{nodeTwoLinkOne, nodeTwoLinkTwo}
	nodeFive.Links = []Link{nodeFiveLink}
	nodeSeven.Links = []Link{nodeSevenLink}

	return &Flowchart{
		Direction: DirectionHorizontalRight,
		Title:     pointTo("Test Title"),
		Nodes:     []*Node{nodeOne, nodeTwo, nodeThree, nodeFour},
		Subgraphs: []*Flowchart{
			{
				Direction: DirectionVertical,
				Nodes:     []*Node{nodeFive, nodeSix},
				Subgraphs: []*Flowchart{
					{
						Direction: DirectionHorizontalLeft,
						Nodes:     []*Node{nodeSeven, nodeEight},
					},
				},
			},
		},
	}
}

// Helper function to mask UUIDs in a string after each subgraph
func maskUUIDsInSubgraph(s string) string {
	// Regular expression to match the "subgraph " followed by 6 characters (UUID)
	re := regexp.MustCompile(`subgraph \w{6}`)

	// Replace each "subgraph " followed by the 6-character UUID with "subgraph UUID"
	return re.ReplaceAllString(s, "subgraph UUID")
}
