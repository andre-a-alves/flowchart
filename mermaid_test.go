package main

import (
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestFlowchartDirectionEnum_toMermaid(t *testing.T) {
	tests := []struct {
		name      string
		direction FlowchartDirectionEnum
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
			expected:  "TD",
		},
		{
			name:      "Unknown direction (fallback)",
			direction: FlowchartDirectionEnum(999), // Unknown value to test fallback
			expected:  "TD",                        // Default fallback
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
			expected: "--",
		},
		{
			name:     "Dotted line",
			lineType: LineTypeDotted,
			expected: "-.-",
		},
		{
			name:     "Thick line",
			lineType: LineTypeThick,
			expected: "==",
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
			nodeType: ShapeTerminator,
			expected: "(",
		},
		{
			name:     "Process left",
			nodeType: ShapeProcess,
			expected: "[",
		},
		{
			name:     "AlternateProcess left",
			nodeType: ShapeAlternateProcess,
			expected: "([",
		},
		{
			name:     "Subprocess left",
			nodeType: ShapeSubprocess,
			expected: "[[",
		},
		{
			name:     "Decision left",
			nodeType: ShapeDecision,
			expected: "{",
		},
		{
			name:     "InputOutput left",
			nodeType: ShapeInputOutput,
			expected: "[/",
		},
		{
			name:     "Connector left",
			nodeType: ShapeConnector,
			expected: "((",
		},
		{
			name:     "Database left",
			nodeType: ShapeDatabase,
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
			nodeType: ShapeTerminator,
			expected: ")",
		},
		{
			name:     "Process right",
			nodeType: ShapeProcess,
			expected: "]",
		},
		{
			name:     "AlternateProcess right",
			nodeType: ShapeAlternateProcess,
			expected: "])",
		},
		{
			name:     "Subprocess right",
			nodeType: ShapeSubprocess,
			expected: "]]",
		},
		{
			name:     "Decision right",
			nodeType: ShapeDecision,
			expected: "}",
		},
		{
			name:     "InputOutput right",
			nodeType: ShapeInputOutput,
			expected: "/[",
		},
		{
			name:     "Connector right",
			nodeType: ShapeConnector,
			expected: "))",
		},
		{
			name:     "Database right",
			nodeType: ShapeDatabase,
			expected: ")]",
		},
		{
			name:     "Default right (invalid node type)",
			nodeType: NodeTypeEnum(-1), // Unknown NodeTypeEnum
			expected: "}",
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
	fixtureLabelText := "label"
	fixtureEmptyLabel := ""
	fixtureTargetNode := Node{Name: "Target Node"}

	tests := []struct {
		name     string
		link     Link
		expected string
	}{
		{
			name: "Basic link with no label",
			link: Link{
				OriginArrow: ArrowTypeNone,
				LineType:    LineTypeSolid,
				TargetArrow: ArrowTypeNormal,
				Label:       nil,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "--> TargetNode",
		},
		{
			name: "Link with empty label",
			link: Link{
				OriginArrow: ArrowTypeNone,
				LineType:    LineTypeSolid,
				TargetArrow: ArrowTypeNormal,
				Label:       &fixtureEmptyLabel,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "--> TargetNode",
		},
		{
			name: "Link with label",
			link: Link{
				OriginArrow: ArrowTypeNone,
				LineType:    LineTypeSolid,
				TargetArrow: ArrowTypeNormal,
				Label:       &fixtureLabelText,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "-- label --> TargetNode",
		},
		{
			name: "Link with different arrows and no label",
			link: Link{
				OriginArrow: ArrowTypeCircle,
				LineType:    LineTypeDotted,
				TargetArrow: ArrowTypeCross,
				Label:       nil,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "o-.-x TargetNode",
		},
		{
			name: "Link with different arrows and a label",
			link: Link{
				OriginArrow: ArrowTypeCircle,
				LineType:    LineTypeDotted,
				TargetArrow: ArrowTypeCross,
				Label:       &fixtureLabelText,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "o-. label .-x TargetNode",
		},
		{
			name: "Link with no line",
			link: Link{
				OriginArrow: ArrowTypeNormal,
				LineType:    LineTypeNone,
				TargetArrow: ArrowTypeNormal,
				Label:       &fixtureLabelText,
				TargetNode:  &fixtureTargetNode,
			},
			expected: "~~~ TargetNode",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.link.toMermaid()

			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaidNode() mismatch (-expected +got):\n%s", diff)
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
				Name:  "First Node",
				Type:  ShapeProcess,
				Label: nil,
				Links: nil,
			},
			indents:  0,
			expected: "FirstNode;\n",
		},
		{
			name: "Node with empty label and no links",
			node: &Node{
				Name:  "First Node",
				Type:  ShapeProcess,
				Label: pointTo(""),
				Links: nil,
			},
			indents:  1,
			expected: "    FirstNode;\n",
		},
		{
			name: "Node with label and no links",
			node: &Node{
				Name:  "First Node",
				Type:  ShapeProcess,
				Label: pointTo("Node Label"),
				Links: nil,
			},
			indents:  2,
			expected: "        FirstNode[Node Label];\n",
		},
		{
			name: "Node with links and no label",
			node: &Node{
				Name:  "First Node",
				Type:  ShapeProcess,
				Label: nil,
				Links: []Link{
					{
						OriginArrow: ArrowTypeNormal,
						LineType:    LineTypeSolid,
						TargetArrow: ArrowTypeNone,
						Label:       nil,
						TargetNode: &Node{
							Name:  "Target Node",
							Type:  ShapeProcess,
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
				Name:  "First Node",
				Type:  ShapeProcess,
				Label: pointTo("Label"),
				Links: []Link{
					{
						OriginArrow: ArrowTypeNormal,
						LineType:    LineTypeDotted,
						TargetArrow: ArrowTypeCross,
						Label:       pointTo("Link Label"),
						TargetNode: &Node{
							Name:  "Target Node",
							Type:  ShapeProcess,
							Label: nil,
							Links: nil,
						},
					},
				}},
			indents:  1,
			expected: "    FirstNode[Label];\n    FirstNode <-. Link Label .-x TargetNode;\n",
		},
		{
			name: "Node with multiple links and label",
			node: &Node{
				Name:  "First Node",
				Type:  ShapeProcess,
				Label: pointTo("Node Label"),
				Links: []Link{
					{
						OriginArrow: ArrowTypeNone,
						LineType:    LineTypeSolid,
						TargetArrow: ArrowTypeNormal,
						Label:       nil,
						TargetNode: &Node{
							Name:  "Target Node One",
							Type:  ShapeProcess,
							Label: nil,
							Links: nil,
						},
					},
					{
						OriginArrow: ArrowTypeCircle,
						LineType:    LineTypeDotted,
						TargetArrow: ArrowTypeCross,
						Label:       pointTo("Link Label"),
						TargetNode: &Node{
							Name:  "Target Node Two",
							Type:  ShapeProcess,
							Label: nil,
							Links: nil,
						},
					},
				},
			},
			indents:  2,
			expected: "        FirstNode[Node Label];\n        FirstNode --> TargetNodeOne;\n        FirstNode o-. Link Label .-x TargetNodeTwo;\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.node.toMermaidNode(tt.indents)
			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("toMermaidNode() mismatch (-expected +got):\n%s", diff)
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
				Nodes:     []*Node{{Name: "Node One"}},
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
					{Name: "First Node"},
					{Name: "Second Node"},
				},
				Subgraphs: []*Flowchart{},
			},
			indents:  1,
			expected: "    subgraph MainTitle [Main Title];\n        direction TD;\n\n        FirstNode;\n\n        SecondNode;\n    end;\n",
		},
		{
			name: "Flowchart with subgraphs",
			flowchart: &Flowchart{
				Direction: DirectionVertical,
				Title:     pointTo("Main Title"),
				Nodes:     []*Node{{Name: "First Node"}},
				Subgraphs: []*Flowchart{
					{
						Direction: DirectionHorizontalLeft,
						Title:     pointTo("Subgraph One"),
						Nodes:     []*Node{{Name: "Second Node"}},
					},
					{
						Direction: DirectionHorizontalRight,
						Title:     pointTo("Subgraph Two"),
						Nodes:     []*Node{{Name: "Third Node"}},
					},
				},
			},
			indents: 0,
			expected: `subgraph MainTitle [Main Title];
    direction TD;

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
				if diff := cmp.Diff(tt.expected, got, cmp.Comparer(ignoreUUIDInSubgraph)); diff != "" {
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

func ignoreUUIDInSubgraph(x, y string) bool {
	xIndex := strings.Index(x, "subgraph ")
	yIndex := strings.Index(y, "subgraph ")

	return xIndex == yIndex && x[0:xIndex+1] == y[0:yIndex+1] && x[xIndex+16:] == y[yIndex+16:]
}
