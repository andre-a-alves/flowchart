package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestArrowTypeEnum_ToMermaidOriginString(t *testing.T) {
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

func TestArrowTypeEnum_ToMermaidTargetString(t *testing.T) {
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

func TestLineTypeEnum_ToMermaidOrigin(t *testing.T) {
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

func TestLineTypeEnum_ToMermaidTarget(t *testing.T) {
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

func TestLineTypeEnum_ToMermaidBidirectional(t *testing.T) {
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
	fixtureTargetNode := Node{Name: "TargetNode"}

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
	// Helper function to create nodes and links for testing
	createNode := func(name string, label *string, links []Link) *Node {
		return &Node{
			Name:  name,
			Type:  ShapeProcess, // Default type for simplicity
			Label: label,
			Links: links,
		}
	}

	tests := []struct {
		name     string
		node     *Node
		indents  int
		expected string
	}{
		{
			name:     "Node with no label and no links",
			node:     createNode("Node1", nil, nil),
			indents:  0,
			expected: "Node1\n",
		},
		{
			name:     "Node with empty label and no links",
			node:     createNode("Node2", pointTo(""), nil),
			indents:  1,
			expected: "    Node2\n",
		},
		{
			name:     "Node with label and no links",
			node:     createNode("Node3", pointTo("Label"), nil),
			indents:  2,
			expected: "        Node3[Label]\n",
		},
		{
			name: "Node with links and no label",
			node: createNode("Node4", nil, []Link{
				{OriginArrow: ArrowTypeNormal, LineType: LineTypeSolid, TargetArrow: ArrowTypeNone, Label: nil, TargetNode: createNode("Node5", nil, nil)},
			}),
			indents:  0,
			expected: "Node4-- TargetNode\n",
		},
		{
			name: "Node with links and label",
			node: createNode("Node6", pointTo("Label"), []Link{
				{OriginArrow: ArrowTypeNormal, LineType: LineTypeDotted, TargetArrow: ArrowTypeCross, Label: pointTo("LinkLabel"), TargetNode: createNode("Node7", nil, nil)},
			}),
			indents:  1,
			expected: "    Node6[Label]\n    Node6-. LinkLabel .-x TargetNode\n",
		},
		{
			name: "Node with multiple links and label",
			node: createNode("Node8", pointTo("Label"), []Link{
				{OriginArrow: ArrowTypeNone, LineType: LineTypeSolid, TargetArrow: ArrowTypeNormal, Label: nil, TargetNode: createNode("Node9", nil, nil)},
				{OriginArrow: ArrowTypeCircle, LineType: LineTypeDotted, TargetArrow: ArrowTypeCross, Label: pointTo("LinkLabel2"), TargetNode: createNode("Node10", nil, nil)},
			}),
			indents:  2,
			expected: "        Node8[Label]\n        Node8-- TargetNode9\n        Node8o-. LinkLabel2 .-x TargetNode10\n",
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
}

//
//import (
//	"errors"
//	"testing"
//
//	"github.com/google/go-cmp/cmp"
//)
//
//func TestToDirection(t *testing.T) {
//	tests := []struct {
//		name        string
//		input       FlowchartDirectionEnumV1
//		expected    *FlowchartV1
//		expectedErr error
//	}{
//		{
//			name:  "Valid input - DirectionLeftRight",
//			input: DirectionLeftRight,
//			expected: &FlowchartV1{
//				Direction: DirectionLeftRight,
//				Nodes:     nil,
//			},
//			expectedErr: nil,
//		},
//		{
//			name:  "Valid input - DirectionTopDown",
//			input: DirectionTopDown,
//			expected: &FlowchartV1{
//				Direction: DirectionTopDown,
//				Nodes:     nil,
//			},
//			expectedErr: nil,
//		},
//		{
//			name:  "Valid input - with node",
//			input: DirectionTopDown,
//			expected: &FlowchartV1{
//				Direction: DirectionTopDown,
//				Nodes: []NodeV1{
//					{},
//				},
//			},
//			expectedErr: nil,
//		},
//		{
//			name:        "Invalid input - negative enum",
//			input:       FlowchartDirectionEnumV1(-1),
//			expected:    nil,
//			expectedErr: errors.New("invalid FlowchartDirectionEnumV1"),
//		},
//		{
//			name:        "Invalid input - out of range enum (positive)",
//			input:       FlowchartDirectionEnumV1(2), // Out of range, since only 0 and 1 are valid
//			expected:    nil,
//			expectedErr: errors.New("invalid FlowchartDirectionEnumV1"),
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			result, err := tt.input.toDirection()
//
//			// Check for error
//			if diff := cmp.Diff(tt.expectedErr, err, cmp.Comparer(errComparer)); diff != "" {
//				t.Errorf("unexpected error (-expected +got):\n%s", diff)
//			}
//
//			// Check for correct result
//			if diff := cmp.Diff(tt.expected, result); diff != "" {
//				t.Errorf("unexpected result (-expected +got):\n%s", diff)
//			}
//		})
//	}
//}
//
//func TestToNode(t *testing.T) {
//	tests := []struct {
//		name        string
//		shapeEnum   ShapeEnumV1
//		nodeName    string
//		nodeText    *string
//		expected    *NodeV1
//		expectedErr error
//	}{
//		{
//			name:      "Valid input - ShapeRectangle",
//			shapeEnum: ShapeRectangle,
//			nodeName:  "RectangleNode",
//			nodeText:  pointTo("SomeText"),
//			expected: &NodeV1{
//				Enum:              ShapeRectangle,
//				NodeName:          "RectangleNode",
//				NodeText:          pointTo("SomeText"),
//				ShapeWrapperStart: "[",
//				ShapeWrapperEnd:   "]",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeRoundEdge",
//			shapeEnum: ShapeRoundEdge,
//			nodeName:  "RoundEdgeNode",
//			nodeText:  pointTo("OtherText"),
//			expected: &NodeV1{
//				Enum:              ShapeRoundEdge,
//				NodeName:          "RoundEdgeNode",
//				NodeText:          pointTo("OtherText"),
//				ShapeWrapperStart: "(",
//				ShapeWrapperEnd:   ")",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeStadium",
//			shapeEnum: ShapeStadium,
//			nodeName:  "StadiumNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeStadium,
//				NodeName:          "StadiumNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "([",
//				ShapeWrapperEnd:   "])",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeSubroutine",
//			shapeEnum: ShapeSubroutine,
//			nodeName:  "SubroutineNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeSubroutine,
//				NodeName:          "SubroutineNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "[[",
//				ShapeWrapperEnd:   "]]",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeCylinder",
//			shapeEnum: ShapeCylinder,
//			nodeName:  "CylinderNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeCylinder,
//				NodeName:          "CylinderNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "[(",
//				ShapeWrapperEnd:   ")]",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeCircle",
//			shapeEnum: ShapeCircle,
//			nodeName:  "CircleNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeCircle,
//				NodeName:          "CircleNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "((",
//				ShapeWrapperEnd:   "))",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeAsymmetric",
//			shapeEnum: ShapeAsymmetric,
//			nodeName:  "AsymmetricNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeAsymmetric,
//				NodeName:          "AsymmetricNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: ">",
//				ShapeWrapperEnd:   "]",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeDiamond",
//			shapeEnum: ShapeDiamond,
//			nodeName:  "DiamondNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeRhombus,
//				NodeName:          "DiamondNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "{",
//				ShapeWrapperEnd:   "}",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeRhombus",
//			shapeEnum: ShapeRhombus,
//			nodeName:  "RhombusNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeRhombus,
//				NodeName:          "RhombusNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "{",
//				ShapeWrapperEnd:   "}",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeHexagon",
//			shapeEnum: ShapeHexagon,
//			nodeName:  "HexagonNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeHexagon,
//				NodeName:          "HexagonNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "{{",
//				ShapeWrapperEnd:   "}}",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeParallelogram",
//			shapeEnum: ShapeParallelogram,
//			nodeName:  "ParallelogramNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeParallelogram,
//				NodeName:          "ParallelogramNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "[/",
//				ShapeWrapperEnd:   "/]",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeAltParallelogram",
//			shapeEnum: ShapeAltParallelogram,
//			nodeName:  "AltParallelogramNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeAltParallelogram,
//				NodeName:          "AltParallelogramNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "[\\",
//				ShapeWrapperEnd:   "\\]",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeTrapezoid",
//			shapeEnum: ShapeTrapezoid,
//			nodeName:  "TrapezoidNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeTrapezoid,
//				NodeName:          "TrapezoidNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "[/",
//				ShapeWrapperEnd:   "\\]",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeAltTrapezoid",
//			shapeEnum: ShapeAltTrapezoid,
//			nodeName:  "AltTrapezoidNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeAltTrapezoid,
//				NodeName:          "AltTrapezoidNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "[\\",
//				ShapeWrapperEnd:   "/]",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:      "Valid input - ShapeDoubleCircle",
//			shapeEnum: ShapeDoubleCircle,
//			nodeName:  "DoubleCircleNode",
//			nodeText:  pointTo("Text"),
//			expected: &NodeV1{
//				Enum:              ShapeDoubleCircle,
//				NodeName:          "DoubleCircleNode",
//				NodeText:          pointTo("Text"),
//				ShapeWrapperStart: "(((",
//				ShapeWrapperEnd:   ")))",
//			},
//			expectedErr: nil,
//		},
//		{
//			name:        "Invalid input - unrecognized enum",
//			shapeEnum:   ShapeEnumV1(100), // Arbitrary invalid value
//			nodeName:    "InvalidNode",
//			nodeText:    pointTo("Text"),
//			expected:    nil,
//			expectedErr: errors.New("invalid ShapeEnumV1"),
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			result, err := tt.shapeEnum.toNode(tt.nodeName, tt.nodeText)
//
//			// Check for error
//			if diff := cmp.Diff(tt.expectedErr, err, cmp.Comparer(func(x, y error) bool {
//				return (x == nil && y == nil) || (x != nil && y != nil && x.Error() == y.Error())
//			})); diff != "" {
//				t.Errorf("unexpected error (-expected +got):\n%s", diff)
//			}
//
//			// Check for correct result
//			if diff := cmp.Diff(tt.expected, result); diff != "" {
//				t.Errorf("unexpected result (-expected +got):\n%s", diff)
//			}
//		})
//	}
//}
//
//func pointTo[T any](value T) *T {
//	return &value
//}
//
//func errComparer(x, y error) bool {
//	return (x == nil && y == nil) || (x != nil && y != nil && x.Error() == y.Error())
//}
