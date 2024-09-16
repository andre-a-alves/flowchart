# Flowchart Modeling in Go
![GitHub License](https://img.shields.io/github/license/andre-a-alves/flowchart)
![GitHub Tag](https://img.shields.io/github/v/tag/andre-a-alves/flowchart)

Flowchart is a Go package designed to model flowcharts with support for various node types, link styles, and subgraphs. Currently, the package can export flowcharts in [Mermaid](https://mermaid-js.github.io/mermaid/) syntax, with plans for additional export options in the future.

## Features

- **Node Types**: Various node types like process, decision, database, and more.
- **Link Styles**: Support for different line styles such as solid, dotted, thick, and no-line.
- **Arrow Types**: Add arrows to the origin, target, or both sides of a link.
- **Subgraphs**: Create subgraphs to organize your flowchart hierarchically.
- **Flowchart Directions**: Control the flow direction, including left-to-right, right-to-left, or top-to-bottom.
- **Mermaid Export**: Generate Mermaid syntax to easily visualize flowcharts.

## Installation

To install the package, run:

```bash
go get github.com/andre-a-alves/flowchart
```

## Example Usage

```go
package main

import (
	"fmt"
	
	"github.com/andre-a-alves/flowchart"
)

func main() {
	// Create nodes
	processNode := flowchart.ProcessNode("Process1", nil)
	decisionNode := flowchart.DecisionNode("Decision1", nil)

	// Create links
	link := flowchart.SolidLink(decisionNode, nil)

	// Add link to node
	processNode.AddLink(link)

	// Create flowchart
	chart := flowchart.VerticalFlowchart(nil)
	chart.AddNode(processNode)

	// Output Mermaid syntax
	fmt.Println(chart.ToMermaid())
}
```

## Future Plans
- Support for additional diagram formats (e.g., Graphviz, PlantUML)
- Improved node and link customization options
- More advanced flowchart layout controls

## License
This project is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for more details.
