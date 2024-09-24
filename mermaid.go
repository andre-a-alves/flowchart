package flowchart

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func mermaidFlowchartDirection(f DirectionEnum) string {
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

func renderArrows(l Link) (string, string) {
	if !l.TargetArrow || l.ArrowType == ArrowTypeNone {
		return "", ""
	}
	if l.OriginArrow {
		switch l.ArrowType {
		case ArrowTypeNormal:
			return "<", ">"
		case ArrowTypeCircle:
			return "o", "o"
		case ArrowTypeCross:
			return "x", "x"
		default:
			return "", ""
		}
	}
	switch l.ArrowType {
	case ArrowTypeNormal:
		return "", ">"
	case ArrowTypeCircle:
		return "", "o"
	case ArrowTypeCross:
		return "", "x"
	default:
		return "", ""
	}
}

func renderMermaidLink(l Link) string {
	if l.Target == nil || l.Origin == nil {
		return ""
	}

	if l.LineType == LineTypeNone {
		return fmt.Sprintf("%s ~~~ %s", removeSpaces(l.Origin.nodeName()), removeSpaces(l.Target.nodeName()))
	}

	originArrow, targetArrow := renderArrows(l)

	if l.Label == nil || *l.Label == "" {
		switch l.LineType {
		case LineTypeSolid:
			if len(targetArrow) > 0 {
				return fmt.Sprintf("%s %s--%s %s", removeSpaces(l.Origin.nodeName()), originArrow, targetArrow, removeSpaces(l.Target.nodeName()))
			}
			return fmt.Sprintf("%s --- %s", removeSpaces(l.Origin.nodeName()), removeSpaces(l.Target.nodeName()))
		case LineTypeDotted:
			return fmt.Sprintf("%s %s-.-%s %s", removeSpaces(l.Origin.nodeName()), originArrow, targetArrow, removeSpaces(l.Target.nodeName()))
		case LineTypeThick:
			if len(targetArrow) > 0 {
				return fmt.Sprintf("%s %s==%s %s", removeSpaces(l.Origin.nodeName()), originArrow, targetArrow, removeSpaces(l.Target.nodeName()))
			}
			return fmt.Sprintf("%s === %s", removeSpaces(l.Origin.nodeName()), removeSpaces(l.Target.nodeName()))
		default:
			return ""
		}
	}

	switch l.LineType {
	case LineTypeSolid:
		return fmt.Sprintf("%s %s-- \"%s\" --%s %s", removeSpaces(l.Origin.nodeName()), originArrow, *l.Label, targetArrow, removeSpaces(l.Target.nodeName()))
	case LineTypeDotted:
		return fmt.Sprintf("%s %s-. \"%s\" .-%s %s", removeSpaces(l.Origin.nodeName()), originArrow, *l.Label, targetArrow, removeSpaces(l.Target.nodeName()))
	case LineTypeThick:
		return fmt.Sprintf("%s %s== \"%s\" ==%s %s", removeSpaces(l.Origin.nodeName()), originArrow, *l.Label, targetArrow, removeSpaces(l.Target.nodeName()))
	default:
		return ""
	}
}

func renderMermaidNode(n *Node, indents int) string {
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

func renderMermaidFlowchart(f *Flowchart, indents int, subgraph bool) string {
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
			mermaidFlowchartDirection(f.Direction),
		))
	} else {
		if f.Title != nil && *f.Title != "" {
			sb.WriteString(fmt.Sprintf("---\ntitle: %s\n---\n", *f.Title))
		}
		sb.WriteString(fmt.Sprintf("flowchart %s;\n", mermaidFlowchartDirection(f.Direction)))
	}

	// nodes
	for _, node := range f.Nodes {
		sb.WriteString(fmt.Sprintf("%s", renderMermaidNode(node, indents+1)))
	}

	// subgraphs
	for _, subgraph := range f.Subgraphs {
		sb.WriteString(fmt.Sprintf("%s", renderMermaidFlowchart(subgraph, indents+1, true)))
	}

	if subgraph {
		// end subgraph
		sb.WriteString(fmt.Sprintf("%send;\n", indentSpaces))
	}

	if !subgraph {
		allLinks := getAllLinks(f)
		for _, link := range allLinks {
			sb.WriteString(fmt.Sprintf("    %s;\n", renderMermaidLink(link)))
		}
	}

	return sb.String()
}

func RenderMermaid(f *Flowchart) (string, error) {
	if !hasValidMermaidNames(f) {
		return "", fmt.Errorf("flowchart contains invalid mermaid names")
	}

	return renderMermaidFlowchart(f, 0, false), nil
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
