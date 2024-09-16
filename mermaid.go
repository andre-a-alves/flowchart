package flowchart

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func (f FlowchartDirectionEnum) toMermaid() string {
	switch f {
	case DirectionHorizontalRight:
		return "LR"
	case DirectionHorizontalLeft:
		return "RL"
	case DirectionVertical:
		return "TB"
	}
	return "TB"
}

func (a ArrowTypeEnum) toMermaidOrigin() string {
	if a == ArrowTypeNormal {
		return "<"
	}
	return a.toMermaidBidirectional()
}

func (a ArrowTypeEnum) toMermaidTarget() string {
	if a == ArrowTypeNormal {
		return ">"
	}
	return a.toMermaidBidirectional()
}

func (a ArrowTypeEnum) toMermaidBidirectional() string {
	switch a {
	case ArrowTypeNone:
		return ""
	case ArrowTypeCross:
		return "x"
	case ArrowTypeCircle:
		return "o"
	default:
		return ""
	}
}

func (l LineTypeEnum) toMermaidOrigin() string {
	if l == LineTypeNone {
		return ""
	}
	if l == LineTypeDotted || l == LineTypeSolid || l == LineTypeThick {
		return l.toMermaidBidirectional()[:2]
	}
	return l.toMermaidBidirectional()
}

func (l LineTypeEnum) toMermaidTarget() string {
	if l == LineTypeNone {
		return ""
	}
	if l == LineTypeDotted || l == LineTypeSolid || l == LineTypeThick {
		return l.toMermaidBidirectional()[1:]
	}
	return l.toMermaidBidirectional()
}

func (l LineTypeEnum) toMermaidBidirectional() string {
	switch l {
	case LineTypeNone:
		return "~~~"
	case LineTypeSolid:
		return "---"
	case LineTypeDotted:
		return "-.-"
	case LineTypeThick:
		return "==="
	default:
		return ""
	}
}

func (n NodeTypeEnum) toMermaidLeft() string {
	switch n {
	case TypeTerminator:
		return "("
	case TypeProcess:
		return "["
	case TypeAlternateProcess:
		return "(["
	case TypeSubprocess:
		return "[["
	case TypeDecision:
		return "{"
	case TypeInputOutput:
		return "[/"
	case TypeConnector:
		return "(("
	case TypeDatabase:
		return "[("
	default:
		return "("
	}
}

func (n NodeTypeEnum) toMermaidRight() string {
	switch n {
	case TypeTerminator:
		return ")"
	case TypeProcess:
		return "]"
	case TypeAlternateProcess:
		return "])"
	case TypeSubprocess:
		return "]]"
	case TypeDecision:
		return "}"
	case TypeInputOutput:
		return "/["
	case TypeConnector:
		return "))"
	case TypeDatabase:
		return ")]"
	default:
		return "}"
	}
}

func (l Link) toMermaid() string {
	if l.TargetNode == nil {
		return ""
	}

	line := l.LineType.toMermaidBidirectional()

	if l.LineType == LineTypeNone {
		return fmt.Sprintf("%s %s", line, removeSpaces(l.TargetNode.Name))
	}

	if (l.OriginArrow || l.TargetArrow) && (l.LineType == LineTypeSolid || l.LineType == LineTypeThick) {
		line = line[:2]
	}
	originArrow := ""
	if l.OriginArrow {
		originArrow = l.ArrowType.toMermaidOrigin()
	}
	targetArrow := ""
	if l.TargetArrow {
		targetArrow = l.ArrowType.toMermaidTarget()
	}

	if l.Label != nil && *l.Label != "" {
		line = fmt.Sprintf("%s %s %s", l.LineType.toMermaidOrigin(), *l.Label, l.LineType.toMermaidTarget())
	}

	return fmt.Sprintf("%s%s%s %s", originArrow, line, targetArrow, removeSpaces(l.TargetNode.Name))
}

func (n *Node) toMermaidNode(indents int) string {
	indentSpaces := strings.Repeat(" ", 4*indents)
	var sb strings.Builder

	if n.Label == nil || *n.Label == "" {
		sb.WriteString(fmt.Sprintf("%s%s;\n", indentSpaces, removeSpaces(n.Name)))
	} else {
		sb.WriteString(fmt.Sprintf("%s%s%s%s%s;\n",
			indentSpaces,
			removeSpaces(n.Name),
			n.Type.toMermaidLeft(),
			*n.Label,
			n.Type.toMermaidRight(),
		))
	}
	for _, link := range n.Links {
		sb.WriteString(fmt.Sprintf("%s%s %s;\n",
			indentSpaces,
			removeSpaces(n.Name),
			link.toMermaid(),
		))
	}

	return sb.String()
}

func (f *Flowchart) toMermaid(indents int) string {
	var sb strings.Builder

	// nodes
	for _, node := range f.Nodes {
		sb.WriteString(fmt.Sprintf("\n%s", node.toMermaidNode(indents)))
	}

	// subgraphs
	for _, subgraph := range f.Subgraphs {
		sb.WriteString(fmt.Sprintf("\n%s", subgraph.toMermaidSubgraph(indents)))
	}

	return sb.String()
}

func (f *Flowchart) toMermaidSubgraph(indents int) string {
	indentSpaces := strings.Repeat(" ", 4*indents)
	var sb strings.Builder

	// start subgraph
	if f.Title == nil {
		sb.WriteString(fmt.Sprintf("%ssubgraph %s;\n",
			indentSpaces,
			uuid.New().String()[0:6],
		))
	} else {
		sb.WriteString(fmt.Sprintf("%ssubgraph %s [%s];\n",
			indentSpaces,
			removeSpaces(*f.Title),
			*f.Title,
		))
	}

	// subgraph direction - indented
	sb.WriteString(fmt.Sprintf("%s%sdirection %s;\n",
		indentSpaces,
		"    ",
		f.Direction.toMermaid(),
	))

	sb.WriteString(f.toMermaid(indents + 1))

	// end subgraph
	sb.WriteString(fmt.Sprintf("%send;\n", indentSpaces))

	return sb.String()
}

func (f *Flowchart) ToMermaid() string {
	var sb strings.Builder

	if f.Title != nil && *f.Title != "" {
		sb.WriteString(fmt.Sprintf("---\ntitle: %s\n---\n", *f.Title))
	}

	sb.WriteString(fmt.Sprintf("flowchart %s;\n", f.Direction.toMermaid()))
	sb.WriteString(f.toMermaid(1))

	return sb.String()
}
