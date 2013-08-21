package gograph

import (
    "fmt"
)

// nodeValue represents a generic node value.
type nodeValue interface {}

// getNodeKey returns a unique key associated to the node value.
func getNodeKey(nv nodeValue) string {
    return fmt.Sprintf("%b", nv)
}

// node represents a graph node.
type node struct {
    key string // Mapping/identifier node value
    Value nodeValue // Actual value of the node
    OutgoingArcs map[string] node
    IncomingArcs map[string] node
}


// newNode creates, initializes and return a node instance.
func newNode(key string, nv nodeValue) *node {
    return &node{
        key: key,
        Value: nv,
        OutgoingArcs: make(map[string] node),
        IncomingArcs: make(map[string] node),
    }
}

/*
addArcTo add an directed arc to "nodeTo". Returns true if the arc is created,
otherwide it returns false because the arc, already exists. It also returns
false if the arc points to itself.
*/
func (n *node) addArcTo(nodeTo node) bool {
    nodeToKey := nodeTo.key
    nodeFromKey := n.key
    if nodeToKey == nodeFromKey {
        return false
    }
    _, ok := n.OutgoingArcs[nodeToKey]
    if !ok {
        n.OutgoingArcs[nodeToKey] = nodeTo
        nodeTo.IncomingArcs[nodeFromKey] = *n
    }
    return !ok
}


// graph represents a graph data structure.
type graph struct {
    nodeMap map[string] node // Keeps track of the nodes
}

// NewGraph creates, initializes and returns a graph structure.
func NewGraph() *graph {
    return &graph{
        nodeMap: make(map[string] node),
    }
}

/*
HasNode checks if the node "nv" exists in the graph. Returns true if so else
returns false.
*/
func (g *graph) HasNode(nv nodeValue) bool{
    key := getNodeKey(nv)
    _, ok := g.nodeMap[key]
    return ok
}

/*
AddNode adds a node to the graph if it doesn't exist. If the node is added it
returns true, otherwise it returns false indicating that the node has not been
added because it already existed.
*/
func (g *graph) AddNode(nv nodeValue) bool {
    ok := g.HasNode(nv)
    if !ok {
        key := getNodeKey(nv)
        g.nodeMap[key] = *newNode(key, nv)
    }
    return !ok
}

/*
getNode gets the node that match the "nodeValue". Returns the node if it
exists, otherwise it returns nil.
*/
func (g *graph) getNode(nv nodeValue) *node {
    nodeKey := getNodeKey(nv)
    n, ok := g.nodeMap[nodeKey]
    if ok {
        return &n
    }
    return nil
}

/*
AddArc creates an arc (unidirectional) from "nodeFromValue" to "nodeToValue".
It returns true if the arc has been created and false if the arc already
existed or if "nodeFromValue" is equals to "nodeToValue".
*/
func (g *graph) AddArc(nodeFromValue, nodeToValue nodeValue) bool {
    g.AddNode(nodeFromValue)
    g.AddNode(nodeToValue)
    nodeFrom := g.getNode(nodeFromValue)
    nodeTo := g.getNode(nodeToValue)
    return nodeFrom.addArcTo(*nodeTo)
}

/*
AddEdge creates an edge (bidrectional) between "node1Value" and "node2Value".
It returns a boolean array with the values true if the arc was created or
false if it already existed. i.e: {true, false} means that the arc from
"node1Value" to "node2Value" has been created and that the arc from
"node2Value" to "node1Value" has not been created because it already existed.
It also returns {false, false} if both node values are the same.
*/
func (g *graph) AddEdge(node1Value, node2Value nodeValue) [2]bool{
    g.AddNode(node1Value)
    g.AddNode(node2Value)
    node1 := g.getNode(node1Value)
    node2 := g.getNode(node2Value)
    return [2]bool{
        node1.addArcTo(*node2),
        node2.addArcTo(*node1),
    }
}


func (graph *graph) PrintGraph() {
    fmt.Printf("%v\n", graph.nodeMap)
}
