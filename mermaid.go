package flowchart

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// mermaidFlowchartDirection converts a DirectionEnum to a Mermaid.js flowchart direction string.
// Valid values are:
// - "LR" for left-to-right horizontal direction.
// - "RL" for right-to-left horizontal direction.
// - "TB" for top-to-bottom vertical direction (default if direction is not recognized).
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

// renderArrows generates the arrow characters for a Mermaid.js link based on the Link's ArrowType and
// whether the OriginArrow and TargetArrow flags are set.
// It returns two strings representing the arrows for the origin and target, respectively.
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

// renderMermaidLink generates a Mermaid.js representation of a Link between two Nodes.
// It returns a string that defines the link, including the line type, any arrows, and optional labels.
// Returns an empty string if either the origin or target nodes are nil.
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

// renderMermaidNode generates a Mermaid.js representation of a Node based on its type and label.
// It returns a string with the proper indentation for the node's position in the flowchart.
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

// renderMermaidFlowchart generates a Mermaid.js representation of a Flowchart.
// It recursively renders subgraphs and their nodes, as well as links between nodes.
// If the flowchart is a subgraph, it starts and ends the subgraph block.
// It returns a string that defines the entire flowchart in Mermaid.js syntax.
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

// RenderMermaid generates a Mermaid.js flowchart string for the given Flowchart object.
// It validates that all nodes in the flowchart have valid names, returning an error if any names are invalid.
// It returns the Mermaid.js representation of the flowchart or an error if validation fails.
func RenderMermaid(f *Flowchart) (string, error) {
	if !hasValidMermaidNames(f) {
		return "", fmt.Errorf("flowchart contains invalid mermaid names")
	}

	return renderMermaidFlowchart(f, 0, false), nil
}

// getAllLinks collects all Links from the flowchart, including links from subgraphs.
// It returns a sorted slice of all links in the flowchart, ordered first by origin and then by target.
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

// hasValidMermaidNames checks if all nodes and subgraphs in the flowchart have valid Mermaid.js names.
// It returns true if all names are valid; otherwise, it returns false.
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

// isValidMermaidNodeName checks if a string is a valid Mermaid.js node name.
// A valid node name contains only letters, digits, underscores, dashes, and spaces.
// It returns true if the name is valid; otherwise, it returns false.
func isValidMermaidNodeName(s string) bool {
	// Define regex for a valid Mermaid.js node name
	// Allows letters, digits, underscores, and dashes only
	re := regexp.MustCompile(`^[a-zA-Z0-9_\- ]+$`)
	return re.MatchString(s)
}
