package main

//
//import "errors"
//
//type (
//	FlowchartDirectionEnumV1 int
//	ShapeEnumV1              int
//)
//
//const (
//	DirectionLeftRight FlowchartDirectionEnumV1 = iota
//	DirectionTopDown
//)
//
//// TODO: Implement
//type FlowchartV1 struct {
//	Direction FlowchartDirectionEnumV1
//	Nodes     []NodeV1
//}
//
//// TODO: Remove struct
//// TODO: Add FlowChartDirectionEnum.toString
//type FlowchartDirection struct {
//	Enum FlowchartDirectionEnumV1
//	Text string
//}
//
//type NodeV1 struct {
//	Enum     ShapeEnumV1
//	NodeName string
//	NodeText *string
//	//TODO: Remove these two
//	//TODO: Add NodeV1.toMermaid
//	ShapeWrapperStart string
//	ShapeWrapperEnd   string
//	//TODO: Add unit tests and adder function(maybe)
//	Links []*LinkTo
//}
//
//type LinkTo struct {
//	LineTypeEnum int
//	Text         *string
//	Target       NodeV1
//	//TODO: Add LinkToEnum.toLinkTo
//	//TODO: Add Unit Test
//	//TODO: Add LinkTo.toMermaid
//}
//
//func (d FlowchartDirectionEnumV1) toFlowchart(nodes []NodeV1) (*FlowchartV1, error) {
//	switch d {
//	case DirectionLeftRight:
//		return &FlowchartV1{
//			Direction: DirectionLeftRight,
//			Nodes:     nodes,
//		}, nil
//	case DirectionTopDown:
//		return &FlowchartV1{
//			Direction: DirectionTopDown,
//			Nodes:     nodes,
//		}, nil
//	default:
//		return nil, errors.New("invalid FlowchartDirectionEnumV1")
//	}
//}
//
//func (d FlowchartDirectionEnumV1) toDirection() (*FlowchartDirection, error) {
//	switch d {
//	case DirectionLeftRight:
//		return &FlowchartDirection{
//			Enum: DirectionLeftRight,
//			Text: "LR",
//		}, nil
//	case DirectionTopDown:
//		return &FlowchartDirection{
//			Enum: DirectionTopDown,
//			Text: "TD",
//		}, nil
//	default:
//		return nil, errors.New("invalid FlowchartDirectionEnumV1")
//	}
//}
//
//func (s ShapeEnumV1) toNode(name string, text *string) (*NodeV1, error) {
//	switch s {
//	case ShapeRectangle:
//		return &NodeV1{
//			Enum:              ShapeRectangle,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "[",
//			ShapeWrapperEnd:   "]",
//		}, nil
//	case ShapeRoundEdge:
//		return &NodeV1{
//			Enum:              ShapeRoundEdge,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "(",
//			ShapeWrapperEnd:   ")",
//		}, nil
//	case ShapeStadium:
//		return &NodeV1{
//			Enum:              ShapeStadium,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "([",
//			ShapeWrapperEnd:   "])",
//		}, nil
//	case ShapeSubroutine:
//		return &NodeV1{
//			Enum:              ShapeSubroutine,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "[[",
//			ShapeWrapperEnd:   "]]",
//		}, nil
//	case ShapeCylinder:
//		return &NodeV1{
//			Enum:              ShapeCylinder,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "[(",
//			ShapeWrapperEnd:   ")]",
//		}, nil
//	case ShapeCircle:
//		return &NodeV1{
//			Enum:              ShapeCircle,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "((",
//			ShapeWrapperEnd:   "))",
//		}, nil
//	case ShapeAsymmetric:
//		return &NodeV1{
//			Enum:              ShapeAsymmetric,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: ">",
//			ShapeWrapperEnd:   "]",
//		}, nil
//	case ShapeDiamond:
//		return &NodeV1{
//			Enum:              ShapeRhombus,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "{",
//			ShapeWrapperEnd:   "}",
//		}, nil
//	case ShapeRhombus:
//		return &NodeV1{
//			Enum:              ShapeRhombus,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "{",
//			ShapeWrapperEnd:   "}",
//		}, nil
//	case ShapeHexagon:
//		return &NodeV1{
//			Enum:              ShapeHexagon,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "{{",
//			ShapeWrapperEnd:   "}}",
//		}, nil
//	case ShapeParallelogram:
//		return &NodeV1{
//			Enum:              ShapeParallelogram,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "[/",
//			ShapeWrapperEnd:   "/]",
//		}, nil
//	case ShapeAltParallelogram:
//		return &NodeV1{
//			Enum:              ShapeAltParallelogram,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "[\\",
//			ShapeWrapperEnd:   "\\]",
//		}, nil
//	case ShapeTrapezoid:
//		return &NodeV1{
//			Enum:              ShapeTrapezoid,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "[/",
//			ShapeWrapperEnd:   "\\]",
//		}, nil
//	case ShapeAltTrapezoid:
//		return &NodeV1{
//			Enum:              ShapeAltTrapezoid,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "[\\",
//			ShapeWrapperEnd:   "/]",
//		}, nil
//	case ShapeDoubleCircle:
//		return &NodeV1{
//			Enum:              ShapeDoubleCircle,
//			NodeName:          name,
//			NodeText:          text,
//			ShapeWrapperStart: "(((",
//			ShapeWrapperEnd:   ")))",
//		}, nil
//	default:
//		return nil, errors.New("invalid ShapeEnumV1")
//	}
//}
