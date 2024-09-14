package main

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestBasicLink(t *testing.T) {
	tests := []struct {
		name        string
		targetNode  *Node
		label       *string
		expected    Link
		expectedErr error
	}{
		{
			name:       "Valid target node with nil label",
			targetNode: &Node{Name: "Node1", Shape: ShapeRectangle},
			label:      nil,
			expected: Link{
				TargetNode:  &Node{Name: "Node1", Shape: ShapeRectangle},
				LineType:    LineTypeSolid,
				OriginArrow: ArrowTypeNone,
				TargetArrow: ArrowTypeNormal,
				Label:       nil,
			},
			expectedErr: nil,
		},
		{
			name:       "Valid target node with label",
			targetNode: &Node{Name: "Node2", Shape: ShapeCircle},
			label:      pointTo("Link Label"),
			expected: Link{
				TargetNode:  &Node{Name: "Node2", Shape: ShapeCircle},
				LineType:    LineTypeSolid,
				OriginArrow: ArrowTypeNone,
				TargetArrow: ArrowTypeNormal,
				Label:       pointTo("Link Label"),
			},
			expectedErr: nil,
		},
		{
			name:        "Nil target node",
			targetNode:  nil,
			label:       pointTo("Some Label"),
			expected:    Link{},
			expectedErr: errors.New("targetNode is nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := BasicLink(tt.targetNode, tt.label)

			// Compare the error using cmp.Diff
			if diff := cmp.Diff(tt.expectedErr, err, cmp.Comparer(compareErrors)); diff != "" {
				t.Errorf("BasicLink() error mismatch (-want +got):\n%s", diff)
			}

			// Compare the result using cmp.Diff
			if diff := cmp.Diff(tt.expected, result, cmp.AllowUnexported(Node{})); diff != "" {
				t.Errorf("BasicLink() result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestBasicNode(t *testing.T) {
	tests := []struct {
		name     string
		nodeName string
		label    *string
		expected *Node
	}{
		{
			name:     "Valid node with nil label",
			nodeName: "TestNode1",
			label:    nil,
			expected: &Node{
				Name:  "TestNode1",
				Shape: ShapeRectangle,
				Label: nil,
				Links: []Link{},
			},
		},
		{
			name:     "Valid node with label",
			nodeName: "TestNode2",
			label:    pointTo("Node Label"),
			expected: &Node{
				Name:  "TestNode2",
				Shape: ShapeRectangle,
				Label: pointTo("Node Label"),
				Links: []Link{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BasicNode(tt.nodeName, tt.label)

			// Compare result using cmp.Diff
			if diff := cmp.Diff(tt.expected, result, cmp.AllowUnexported(Node{})); diff != "" {
				t.Errorf("BasicNode() result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestBasicFlowchart(t *testing.T) {
	expected := &Flowchart{
		Direction: DirectionVertical,
		Title:     nil,
		Nodes:     []*Node{},
		Subgraphs: []*Flowchart{},
	}

	t.Run("Basic flowchart creation", func(t *testing.T) {
		result := BasicFlowchart()

		// Compare result using cmp.Diff
		if diff := cmp.Diff(expected, result, cmp.AllowUnexported(Flowchart{})); diff != "" {
			t.Errorf("BasicFlowchart() result mismatch (-want +got):\n%s", diff)
		}
	})
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
