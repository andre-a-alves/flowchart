package main

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
