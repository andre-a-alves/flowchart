package flowchart

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func (f DirectionEnum) toMermaid() string {
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

func (l Link) toMermaid() string {
	if l.Target == nil || l.Origin == nil {
		return ""
	}

	line := l.LineType.toMermaidBidirectional()

	if l.LineType == LineTypeNone {
		return fmt.Sprintf("%s %s %s", removeSpaces(l.Origin.nodeName()), line, removeSpaces(l.Target.nodeName()))
	}

	if (l.OriginArrow || l.TargetArrow) && (l.LineType == LineTypeSolid || l.LineType == LineTypeThick) {
		line = line[:2]
	}
	originArrow := ""
	targetArrow := ""
	if l.TargetArrow {
		targetArrow = l.ArrowType.toMermaidTarget()
		// arrows cannot be only from target to origin
		if l.OriginArrow {
			originArrow = l.ArrowType.toMermaidOrigin()
		}
	}

	if l.Label != nil && *l.Label != "" {
		line = fmt.Sprintf("%s \"%s\" %s", l.LineType.toMermaidOrigin(), *l.Label, l.LineType.toMermaidTarget())
	}

	return fmt.Sprintf("%s %s%s%s %s", removeSpaces(l.Origin.nodeName()), originArrow, line, targetArrow, removeSpaces(l.Target.nodeName()))
}

func (n *Node) toMermaid(indents int) string {
	indentSpaces := strings.Repeat(" ", 4*indents)

	if n.Label == nil || *n.Label == "" {
		return fmt.Sprintf("%s%s;\n", indentSpaces, removeSpaces(n.name))
	}
	switch n.Type {
	case NodeTypeTerminator:
		return fmt.Sprintf("%s%s(\"%s\");\n", indentSpaces, removeSpaces(n.name), *n.Label)
	case NodeTypeProcess:
		return fmt.Sprintf("%s%s[\"%s\"];\n", indentSpaces, removeSpaces(n.name), *n.Label)
	case NodeTypeAlternateProcess:
		return fmt.Sprintf("%s%s([\"%s\"]);\n", indentSpaces, removeSpaces(n.name), *n.Label)
	case NodeTypeSubprocess:
		return fmt.Sprintf("%s%s[[\"%s\"]];\n", indentSpaces, removeSpaces(n.name), *n.Label)
	case NodeTypeDecision:
		return fmt.Sprintf("%s%s{\"%s\"};\n", indentSpaces, removeSpaces(n.name), *n.Label)
	case NodeTypeInputOutput:
		return fmt.Sprintf("%s%s[/\"%s\"/];\n", indentSpaces, removeSpaces(n.name), *n.Label)
	case NodeTypeConnector:
		return fmt.Sprintf("%s%s((\"%s\"));\n", indentSpaces, removeSpaces(n.name), *n.Label)
	case NodeTypeDatabase:
		return fmt.Sprintf("%s%s[(\"%s\")];\n", indentSpaces, removeSpaces(n.name), *n.Label)
	default:
		return fmt.Sprintf("%s%s(\"%s\");\n", indentSpaces, removeSpaces(n.name), *n.Label)
	}
}

func (f *Flowchart) toMermaid(indents int, subgraph bool) string {
	indentSpaces := strings.Repeat(" ", 4*indents)
	var sb strings.Builder

	if subgraph {
		// start subgraph
		if f.Title == nil || *f.Title == "" {
			panic("subgraph with no title")
		}
		sb.WriteString(fmt.Sprintf("%ssubgraph %s [%s];\n",
			indentSpaces,
			removeSpaces(*f.Title),
			*f.Title,
		))
		// subgraph direction - indented
		sb.WriteString(fmt.Sprintf("%s%sdirection %s;\n",
			indentSpaces,
			"    ",
			f.Direction.toMermaid(),
		))
	} else {
		if f.Title != nil && *f.Title != "" {
			sb.WriteString(fmt.Sprintf("---\ntitle: %s\n---\n", *f.Title))
		}
		sb.WriteString(fmt.Sprintf("flowchart %s;\n", f.Direction.toMermaid()))
	}

	// nodes
	for _, node := range f.Nodes {
		sb.WriteString(fmt.Sprintf("%s", node.toMermaid(indents+1)))
	}

	// subgraphs
	for _, subgraph := range f.Subgraphs {
		sb.WriteString(fmt.Sprintf("%s", subgraph.toMermaid(indents+1, true)))
	}

	if subgraph {
		// end subgraph
		sb.WriteString(fmt.Sprintf("%send;\n", indentSpaces))
	}

	if !subgraph {
		allLinks := getAllLinks(f)
		for _, link := range allLinks {
			sb.WriteString(fmt.Sprintf("    %s;\n", link.toMermaid()))
		}
	}

	return sb.String()
}

func (f *Flowchart) ToMermaid() (string, error) {
	if !hasValidMermaidNames(f) {
		return "", fmt.Errorf("flowchart contains invalid mermaid names")
	}

	return f.toMermaid(0, false), nil
}

func getAllLinks(f *Flowchart) []Link {
	var allLinks []Link

	allLinks = append(allLinks, f.Links...)
	for _, subgraph := range f.Subgraphs {
		allLinks = append(allLinks, getAllLinks(subgraph)...)
	}

	sort.Slice(allLinks, func(i, j int) bool {
		if allLinks[i].Origin.nodeName() == allLinks[j].Origin.nodeName() {
			return allLinks[i].Target.nodeName() < allLinks[j].Target.nodeName()
		}
		return allLinks[i].Origin.nodeName() < allLinks[j].Origin.nodeName()
	})

	return allLinks
}

func hasValidMermaidNames(f *Flowchart) bool {
	for _, node := range f.Nodes {
		if !isValidMermaidNodeName(node.name) {
			return false
		}
	}
	for _, subgraph := range f.Subgraphs {
		if subgraph.Title == nil || *subgraph.Title == "" || !hasValidMermaidNames(subgraph) {
			return false
		}
	}
	return true
}

func isValidMermaidNodeName(s string) bool {
	// Define regex for a valid Mermaid.js node name
	// Allows letters, digits, underscores, and dashes only
	re := regexp.MustCompile(`^[a-zA-Z0-9_\- ]+$`)
	return re.MatchString(s)
}
