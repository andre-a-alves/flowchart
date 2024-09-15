# Flowchart Generation in Mermaid.js

This package provides a simple yet powerful API for generating flowcharts in Mermaid.js format, an open-source diagramming and charting tool. The library allows you to create nodes, links, and subgraphs, convert them to Mermaid.js code, and handle Mermaid.js-compatible flowchart directionality.

## Features

- **Flowchart Construction**: Create flowcharts with nodes, links, and subgraphs.
- **Flowchart Directionality**: Support for horizontal and vertical flowchart layouts.
- **Mermaid.js Export**: Converts your flowchart structure into a Mermaid.js string that can be rendered directly in compatible tools.
- **Node and Link Types**: Supports a variety of node shapes and link styles.
- **Easy Linking**: Connect nodes with arrows, labels, and various line styles.
- **Subgraph Support**: Group nodes into subgraphs for better organization.

## Example Usage

```go
package main

import (
    "fmt"
    "flowchart"
)

func main() {
    // Create basic nodes
    start := flowchart.BasicNode("Start", flowchart.pointTo("Start Process"))
    process := flowchart.BasicNode("Process", flowchart.pointTo("Execute Task"))
    end := flowchart.BasicNode("End", flowchart.pointTo("End Process"))

    // Create links between nodes
    link1, _ := flowchart.BasicLink(process, flowchart.pointTo("Next"))
    link2, _ := flowchart.BasicLink(end, flowchart.pointTo("Complete"))

    // Add links to nodes
    start.AddLink(link1)
    process.AddLink(link2)

    // Create a flowchart and add nodes
    chart := flowchart.BasicFlowchart()
    chart.AddNode(start)
    chart.AddNode(process)
    chart.AddNode(end)

    // Generate the Mermaid.js output
    fmt.Println(chart.ToMermaid())
}
```
## Installation
```
go get github.com/yourusername/flowchart
```
## License
This project is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for more details.
