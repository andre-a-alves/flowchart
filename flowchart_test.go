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
			node: Node{Name: "StartNode", Links: []Link{}},
			link: Link{
				TargetNode: &Node{Name: "TargetNode"},
				Label:      pointTo("LinkLabel"),
			},
			expectedErr: nil,
			expectedLinks: []Link{
				{
					TargetNode: &Node{Name: "TargetNode"},
					Label:      pointTo("LinkLabel"),
				},
			},
		},
		{
			name: "Add invalid link with nil TargetNode",
			node: Node{Name: "StartNode", Links: []Link{}},
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

			if diff := cmp.Diff(tt.expectedLinks, tt.node.Links); diff != "" {
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
			name:          "Add subgraph",
			flowchart:     &Flowchart{},
			node:          &Node{Name: "Singleton"},
			expectedErr:   nil,
			expectedNodes: []*Node{{Name: "Singleton"}},
		},
		{
			name:          "Add duplicate subgraph",
			flowchart:     &Flowchart{Nodes: []*Node{{Name: "Singleton"}}},
			node:          &Node{Name: "Singleton"},
			expectedErr:   fmt.Errorf("cannot add duplicate subgraph"),
			expectedNodes: []*Node{{Name: "Singleton"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.flowchart.AddNode(tt.node)

			if diff := cmp.Diff(tt.expectedErr, err, cmp.Comparer(compareErrors)); diff != "" {
				t.Errorf("AddLink() error mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.expectedNodes, tt.flowchart.Nodes); diff != "" {
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
			name:              "Add duplicate subgraph",
			flowchart:         &Flowchart{Subgraphs: []*Flowchart{{Title: pointTo("Singleton")}}},
			subgraph:          &Flowchart{Title: pointTo("Singleton")},
			expectedErr:       fmt.Errorf("cannot add duplicate subgraph"),
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

func TestBasicLink(t *testing.T) {
	tests := []struct {
		name       string
		targetNode *Node
		label      *string
		expected   Link
	}{
		{
			name:       "Nil target node",
			targetNode: nil,
			label:      pointTo("Link Label"),
			expected: Link{
				ArrowType:   ArrowTypeNormal,
				TargetNode:  nil,
				LineType:    LineTypeSolid,
				OriginArrow: false,
				TargetArrow: true,
				Label:       pointTo("Link Label"),
			},
		},
		{
			name:       "Nil label",
			targetNode: &Node{Name: "Node1"},
			label:      nil,
			expected: Link{
				ArrowType:   ArrowTypeNormal,
				TargetNode:  &Node{Name: "Node1"},
				LineType:    LineTypeSolid,
				OriginArrow: false,
				TargetArrow: true,
				Label:       nil,
			},
		},
		{
			name:       "Valid target node with label",
			targetNode: &Node{Name: "Node2"},
			label:      pointTo("Link Label"),
			expected: Link{
				ArrowType:   ArrowTypeNormal,
				TargetNode:  &Node{Name: "Node2"},
				LineType:    LineTypeSolid,
				OriginArrow: false,
				TargetArrow: true,
				Label:       pointTo("Link Label"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BasicLink(tt.targetNode, tt.label)

			if diff := cmp.Diff(tt.expected, got, cmp.AllowUnexported(Node{})); diff != "" {
				t.Errorf("BasicLink() got mismatch (-want +got):\n%s", diff)
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
				Type:  TypeProcess,
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
				Type:  TypeProcess,
				Label: pointTo("Node Label"),
				Links: []Link{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BasicNode(tt.nodeName, tt.label)

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
