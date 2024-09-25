package flowchart

import (
	"fmt"
	"regexp"
	"slices"
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
	err := validateMermaid(f)
	if err != nil {
		return "", err
	}

	return renderMermaidFlowchart(f, 0, false), nil
}

// validateMermaid validates the Flowchart structure to ensure it adheres to Mermaid.js requirements.
// It checks for the following violations:
// 1. All node and subgraph names must be valid according to Mermaid.js naming conventions.
// 2. The flowchart must not contain nested subgraphs.
// 3. All node and subgraph names must be unique within the flowchart.
//
// If any of these conditions are not met, it aggregates the corresponding violation messages
// and returns a single error detailing all violations. If no violations are found, it returns nil.
//
// Example:
//
//	err := validateMermaid(flowchart)
//	if err != nil {
//	    // Handle validation errors
//	}
func validateMermaid(f *Flowchart) error {
	violations := make([]string, 0, 3)

	if !hasValidMermaidNames(f) {
		violations = append(violations, "contains invalid mermaid names")
	}
	if haveSubgraphs(f.Subgraphs) {
		violations = append(violations, "contains nested subgraphs")
	}
	if !hasUniqueNodeAndSubgraphNames(f) {
		violations = append(violations, "contains repeated node and/or subgraph names")
	}

	if len(violations) > 0 {
		return fmt.Errorf("flowchart contains violations: %s", strings.Join(violations, ", "))
	}
	return nil
}

// haveSubgraphs checks whether any of the provided Flowcharts contain nested subgraphs.
// It iterates through each Flowchart in the slice and returns true if any Flowchart has one or more subgraphs.
//
// Parameters:
//   - charts: A slice of Flowchart pointers to be checked for nested subgraphs.
//
// Returns:
//   - bool: true if at least one Flowchart contains nested subgraphs; false otherwise.
//
// Example:
//
//	hasNested := haveSubgraphs(flowchart.Subgraphs)
//	if hasNested {
//	    // Handle nested subgraphs
//	}
func haveSubgraphs(charts []*Flowchart) bool {
	for _, chart := range charts {
		if chart.Subgraphs != nil && len(chart.Subgraphs) > 0 {
			return true
		}
	}
	return false
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

// hasUniqueNodeAndSubgraphNames checks whether all node names and subgraph titles within the Flowchart are unique.
// It ensures that there are no duplicate names among nodes and no duplicate titles among subgraphs.
// This function iterates through all nodes and subgraphs, collecting their names and titles and verifying their uniqueness.
//
// Parameters:
//   - f: A pointer to the Flowchart to be validated.
//
// Returns:
//   - bool: Returns true if all node names and subgraph titles are unique within the Flowchart.
//     Returns false if there are any duplicate names or titles.
//
// Example:
//
//	unique := hasUniqueNodeAndSubgraphNames(flowchart)
//	if !unique {
//	    // Handle duplicate names or titles
//	}
func hasUniqueNodeAndSubgraphNames(f *Flowchart) bool {
	var names []string

	for _, node := range f.Nodes {
		if slices.Contains(names, node.name) {
			return false
		}
		names = append(names, node.name)
	}
	for _, subgraph := range f.Subgraphs {
		if subgraph.Title != nil {
			if slices.Contains(names, *subgraph.Title) {
				return false
			}
			names = append(names, *subgraph.Title)
		}

	}
	return true
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

// isValidMermaidNodeName checks if a string is a valid Mermaid.js node name.
// A valid node name contains only letters, digits, underscores, dashes, and spaces.
// It returns true if the name is valid; otherwise, it returns false.
func isValidMermaidNodeName(s string) bool {
	// Define regex for a valid Mermaid.js node name
	// Allows letters, digits, underscores, and dashes only
	re := regexp.MustCompile(`^[a-zA-Z0-9_\- ]+$`)
	return re.MatchString(s)
}

// GetMermaidFriendlyFlowchart transforms a Flowchart into a Mermaid-friendly version.
// It performs the following steps:
// 1. Flattens all nested subgraphs into a single-level structure.
// 2. Removes any nodes, subgraphs, and links that do not conform to Mermaid.js naming conventions.
//
// Parameters:
// - f: A pointer to the original Flowchart to be transformed.
//
// Returns:
// - *Flowchart: A new Flowchart instance that is compatible with Mermaid.js rendering.
func GetMermaidFriendlyFlowchart(f *Flowchart) *Flowchart {
	var subgraphs []*Flowchart
	for _, subgraph := range f.Subgraphs {
		subgraphs = append(subgraphs, flattenFlowchart(subgraph))
	}

	return removeNonMermaidNames(&Flowchart{
		Direction: f.Direction,
		Title:     f.Title,
		Nodes:     f.Nodes,
		Subgraphs: subgraphs,
		Links:     f.Links,
	})
}

// removeNonMermaidNames filters out any nodes, subgraphs, and links that have names
// not compliant with Mermaid.js naming conventions. It ensures that only valid
// elements are retained in the Flowchart.
//
// Parameters:
// - f: A pointer to the Flowchart to be filtered.
//
// Returns:
// - *Flowchart: A new Flowchart instance with only Mermaid-compliant nodes, subgraphs, and links.
func removeNonMermaidNames(f *Flowchart) *Flowchart {
	var nodes []*Node
	for _, node := range f.Nodes {
		if isValidMermaidNodeName(node.name) {
			nodes = append(nodes, node)
		}
	}
	var subgraphs []*Flowchart
	for _, subgraph := range f.Subgraphs {
		if subgraph.Title != nil && *subgraph.Title != "" && isValidMermaidNodeName(*subgraph.Title) {
			subgraphs = append(subgraphs, removeNonMermaidNames(subgraph))
		}
	}
	var links []Link
	for _, link := range f.Links {
		if isValidMermaidNodeName(link.Origin.nodeName()) && isValidMermaidNodeName(link.Target.nodeName()) {
			links = append(links, link)
		}
	}

	return &Flowchart{
		Direction: f.Direction,
		Title:     f.Title,
		Nodes:     nodes,
		Subgraphs: subgraphs,
		Links:     links,
	}
}

// flattenFlowchart recursively flattens a Flowchart by aggregating all nodes and links
// from its subgraphs into a single-level structure.
func flattenFlowchart(f *Flowchart) *Flowchart {
	var nodes []*Node
	var links []Link

	nodes = append(nodes, f.Nodes...)
	links = append(links, f.Links...)

	for _, subgraph := range f.Subgraphs {
		flattenedSubgraph := flattenFlowchart(subgraph)
		nodes = append(nodes, flattenedSubgraph.Nodes...)
		links = append(links, flattenedSubgraph.Links...)
	}

	return &Flowchart{
		Direction: f.Direction,
		Title:     f.Title,
		Nodes:     nodes,
		Subgraphs: nil, // Subgraphs are flattened
		Links:     links,
	}
}
