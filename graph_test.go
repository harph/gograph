package gograph


import (
    "testing"
    "fmt"
)


type testValue interface {}


// AddNode test.
func TestAddNode(t *testing.T) {
    graph := NewGraph()
    dummyValue1 := new(testValue)
    type dummyStruct struct{
        X int
        Y int
    }
    dummyValue2 := &dummyStruct{10, 20}
    testCases := []struct{
        input testValue
        output bool
    }{
        {1, true}, // An int node
        {1, false}, // An existing node
        {3.14, true}, // A float node
        {"foo", true}, // A string node
        {"foo", false}, // A existing string node
        {nil, true}, // A nil node
        {nil, false}, // An existing nil node
        {[]int{1, 2, 3}, true}, // An array node
        {[]int{1, 2, 3}, false}, // An existing array node
        {dummyValue1, true}, // An interface instance
        {dummyValue1, false}, // An existing interface
        {dummyValue2, true}, // A struct
        {dummyValue2, false}, // An existing struct
    }
    for _, testCase := range testCases {
        added, node := graph.AddNode(testCase.input)
        if added != testCase.output {
            t.Errorf(
                "graph.AddNode(%#v) returned \"%t\" when \"%t\" was expected.\n",
                testCase.input, added, testCase.output,
            )
        }
        if fmt.Sprintf("%#v", node.Value) != fmt.Sprintf("%#v", testCase.input) {
            t.Errorf(
                "graph.AddNode(%#v) returned a node with value \"%#v\" that " +
                "doesn't match the input value \"%#v\"",
                testCase.input, node.Value, testCase.input,
            )
        }
    }
}

// DeleteNode test.
func TestDeleteNode(t *testing.T) {
    graph := NewGraph()
    dummyValue1 := new(testValue)
    type dummyStruct struct{
        X int
        Y int
    }
    dummyValue2 := &dummyStruct{10, 20}
    testCases := []struct{
        input testValue
        edges []testValue
        addNode bool
        output bool
    }{
        {"A", []testValue{"B", "C"}, true, true},
        {3.14, []testValue{}, false, false},
        {"C", []testValue{"foo", "bar"}, false, true},
        {nil, []testValue{"B", dummyValue1}, true, true},
        {[]int{1, 2 , 3}, []testValue{dummyValue1,}, true, true},
        {dummyValue1, []testValue{}, false, true},
        {dummyValue2, []testValue{}, true, true},
    }
    for _, testCase := range testCases {
        nodeValue := testCase.input
        if testCase.addNode {
            graph.AddNode(nodeValue)
        }
        for _, nodeTo := range testCase.edges {
            graph.AddEdge(nodeValue, nodeTo)
        }
        deleted := graph.DeleteNode(nodeValue)
        if deleted != testCase.output {
            t.Errorf(
                "graph.DeleteNode(%#v) returned \"%t\" when \"%t\" was " +
                "expected.",
                testCase.input, deleted, testCase.output,
            )
        }
        if graph.HasNode(testCase.input) {
            t.Errorf(
                "graph.HasNode(%#v) returned \"true\" when the node was " + 
                "previously removed",
                testCase.input,
            )
        }
        for _, nodeTo := range testCase.edges {
            if graph.HasEdge(nodeValue, nodeTo) ||
                    graph.HasArc(nodeValue, nodeTo ) ||
                    graph.HasArc(nodeTo, nodeValue) {
                t.Errorf(
                    "There is still an edge/arcs between %#v and %#v after " +
                    "execute graph.DeleteNode(%#v)",
                    nodeValue, nodeTo, nodeValue,
                )
            }
        }
    }
}

// HasNode test.
func TestHasNode(t *testing.T) {
    graph := NewGraph();
    dummyValue1 := new(testValue)
    type dummyStruct struct{
        X int
        Y int
    }
    dummyValue2 := &dummyStruct{10, 20}
    dummyValue3 := []int{1, 2, 3}
    graph.AddNode(1);
    graph.AddNode(3.14);
    graph.AddNode("foo");
    graph.AddNode(nil);
    graph.AddNode(dummyValue1);
    graph.AddNode(dummyValue3);
    testCases := []struct{
        input testValue
        output bool
    }{
        {1, true},
        {2, false},
        {3.14, true},
        {2.71, false},
        {"foo", true},
        {"bar", false},
        {nil, true},
        {dummyValue1, true},
        {dummyValue2, false},
        {dummyValue3, true},
    }
    for _, testCase := range testCases {
        hasIt := graph.HasNode(testCase.input)
        if hasIt != testCase.output {
            t.Errorf(
                "graph.HasNode(%#v) returned \"%t\" when \"%t\" was expected.\n",
                testCase.input, hasIt, testCase.output,
            )
        }
    }
}

// GetNode Test
func GetNodeTest(t *testing.T) {
    graph := NewGraph()
    dummyValue1 := new(testValue)
    type dummyStruct struct{
        X int
        Y int
    }
    dummyValue2 := &dummyStruct{10, 20}
    testCases := []testValue {
        1, 3.14, "foo", nil, []int{1, 2 , 3}, dummyValue1, dummyValue2,
    }
    for _, testCase := range testCases {
        _, node := graph.AddNode(testCase)
        if fmt.Sprintf("%#v", node.Value) != fmt.Sprintf("%#v", testCase) {
            t.Errorf(
                "graph.GetNode(%#v) returned a node with value \"%#v\" that " +
                "doesn't match the input value \"%#v\"",
                testCase, node.Value, testCase,
            )
        }
    }
}

// AddArc test.
func TestAddArc(t *testing.T) {
    graph := NewGraph()
    type dummyStruct struct{
        X int
        Y int
    }
    dummyValue := new(testValue)
    dummyArray := []int{1, 2, 3}
    testCases := []struct{
        inputNodeFrom testValue
        inputNodeTo testValue
        outputAddedArc bool
    }{
        {"A", "A", false},
        {"A", "B", true},
        {"A", dummyArray, true},
        {dummyValue, dummyValue, false},
        {dummyValue, "A", true},
        {dummyValue, nil, true},
        {dummyArray, dummyArray, false},
        {dummyArray, "B", true},
        {nil, nil, false},
        {nil, dummyValue, true},
        {nil, dummyArray, true},
        {"A", "B", false},
        {"A", dummyArray, false},
        {"A", nil, true},
    }
    for _, testCase := range testCases {
        nodeFrom := testCase.inputNodeFrom
        nodeTo := testCase.inputNodeTo
        added := graph.AddArc(nodeFrom, nodeTo)
        expectedOuput := testCase.outputAddedArc
        // First check if both nodes exists in the graph
        if !graph.HasNode(nodeFrom) || !graph.HasNode(nodeTo) {
            t.Errorf(
                "After executing graph.addArc(%#v, %#v) " + 
                "at least one of the nodes is not in the graph.",
                nodeFrom, nodeTo,
            )
        } else if (added != expectedOuput) {
            t.Errorf(
                "graph.addArc(%#v, %#v) returned " +
                "\"%t\" when \"%t\" was expected",
                nodeFrom, nodeTo, added, expectedOuput,
            )
        }
    }
}

// DeleteArc test.
func TestDeleteArc(t *testing.T) {
    graph := NewGraph()
    type dummyStruct struct{
        X int
        Y int
    }
    dummyValue := new(testValue)
    dummyArray := []int{1, 2, 3}
    testCases := []struct{
        inputNode1 testValue
        inputNode2 testValue
        addArc bool
        output bool
    }{
        {"A", "B", true, true},
        {"B", "A", false, false},
        {"A", dummyArray, false, false},
        {dummyValue, dummyArray, true, true},
        {nil, dummyArray, true, true},
    }
    for _, testCase := range testCases {
        node1 := testCase.inputNode1
        node2 := testCase.inputNode2
        if testCase.addArc {
            graph.AddArc(node1, node2)
        }
        deleted := graph.DeleteArc(node1, node2)
        if deleted != testCase.output {
            t.Errorf(
                "graph.DeleteArc(%#v, %#v) returned \"%t\" when \"%t\" was " +
                "expected",
                node1, node2, deleted, testCase.output,
            )
        }
    }
}

// HasArc test.
func TestHasArc(t *testing.T) {
    graph := NewGraph()
    dummyValue := new(testValue)
    dummyArray := []int{1, 2, 3}
    testCases := []struct{
        inputNodeFrom testValue
        inputNodeTo testValue
        addArc bool
        addEdge bool
        output bool
    }{
        {"A", "B", true, false, true}, // Arc from A to B
        {"A", "A", true, true, false},  // Arc and Edge from A to A
        {"B", "B", false, false, false}, // No Arc or Edge from B to B
        {"B", dummyValue, false, false, false}, // No Arc or Edge
        {dummyArray, dummyValue, false, true, true}, // Edge from array-value
        {nil, nil, false, false, false}, // No Arc/Edge and node doesn't exists
    }
    for _, testCase := range(testCases) {
        nodeFrom := testCase.inputNodeFrom
        nodeTo := testCase.inputNodeTo
        if testCase.addArc {
            graph.AddArc(nodeFrom, nodeTo)
        }
        if testCase.addEdge {
            graph.AddEdge(nodeFrom, nodeTo)
        }
        hasArc := graph.HasArc(nodeFrom, nodeTo)
        if hasArc != testCase.output {
            t.Errorf(
                "graph.HasArc(%#v, %#v) returned \"%t\" when \"%t\" " + 
                "was expected",
                nodeFrom, nodeTo, hasArc, testCase.output,
            )
        }
    }
}

// AddEdge test.
func TestAddEdge(t *testing.T) {
    graph := NewGraph()
    type dummyStruct struct{
        X int
        Y int
    }
    dummyValue := new(testValue)
    dummyArray := []int{1, 2, 3}
    testCases := []struct{
        inputNode1 testValue
        inputNode2 testValue
        outputAddedEdges [2]bool
    }{
        {"A", "B", [2]bool{true, true}}, // An unexisting edge
        {"A", "B", [2]bool{false, false}}, // An existing edge
        {"B", "A", [2]bool{false, false}}, // An existing edge
        {"A", dummyArray, [2]bool{true, true}}, // An unexisting edge
        {dummyValue, dummyArray, [2]bool{true, true}}, // An unexisting edge
        {dummyArray, dummyValue, [2]bool{false, false}}, // An existing edge
        {dummyArray, dummyArray, [2]bool{false, false}}, // Edge to itself
        {nil, dummyArray, [2]bool{true, true}}, // An unexisting edge
    }
    for _, testCase := range testCases {
        node1 := testCase.inputNode1
        node2 := testCase.inputNode2
        added := graph.AddEdge(node1, node2)
        expectedOuput := testCase.outputAddedEdges
        // First check if both nodes exists in the graph
        if !graph.HasNode(node1) || !graph.HasNode(node2) {
            t.Errorf(
                "After executing graph.addEdge(%#v, %#v) " + 
                "at least one of the nodes is not in the graph.",
                node1, node2,
            )
        } else if (added != expectedOuput) {
            t.Errorf(
                "graph.addEdge(%#v, %#v) returned " +
                "\"%v\" when \"%v\" was expected",
                node1, node2, added, expectedOuput,
            )
        }
    }
}

// DeleteEdge test.
func TestDeleteEdge(t *testing.T) {
    graph := NewGraph()
    type dummyStruct struct{
        X int
        Y int
    }
    dummyValue := new(testValue)
    dummyArray := []int{1, 2, 3}
    testCases := []struct{
        inputNode1 testValue
        inputNode2 testValue
        addEdge bool
        output bool
    }{
        {"A", "B", true, true},
        {"B", "A", false, false},
        {"A", dummyArray, false, false},
        {dummyValue, dummyArray, true, true},
        {nil, dummyArray, true, true},
    }
    for _, testCase := range testCases {
        node1 := testCase.inputNode1
        node2 := testCase.inputNode2
        if testCase.addEdge {
            graph.AddEdge(node1, node2)
        }
        deleted := graph.DeleteEdge(node1, node2)
        if deleted != testCase.output {
            t.Errorf(
                "graph.DeleteEdge(%#v, %#v) returned \"%t\" when \"%t\" was " +
                "expected",
                node1, node2, deleted, testCase.output,
            )
        }
    }
}

// HasEdge test.
func TestHasEdge(t *testing.T) {
    graph := NewGraph()
    dummyValue := new(testValue)
    dummyArray := []int{1, 2, 3}
    testCases := []struct{
        inputNode1 testValue
        inputNode2 testValue
        addArc bool
        addEdge bool
        output bool
    }{
        {"A", "B", false, true, true}, // Arc from A to B
        {"B", "A", false, false, true}, // Inverting the previous case
        {"A", "A", true, true, false},  // Arc and Edge from A to A
        {"B", "B", false, false, false}, // No Arc or Edge from B to B
        {"B", dummyValue, true, false, false}, // Arc but no Edge
        {dummyArray, dummyValue, false, true, true}, // Edge from array-value
        {nil, nil, false, false, false}, // No Arc/Edge and node doesn't exists
    }
    for _, testCase := range(testCases) {
        node1 := testCase.inputNode1
        node2 := testCase.inputNode2
        if testCase.addArc {
            graph.AddArc(node1, node2)
        }
        if testCase.addEdge {
            graph.AddEdge(node1, node2)
        }
        hasEdge := graph.HasEdge(node1, node2)
        if hasEdge != testCase.output {
            t.Errorf(
                "graph.HasEdge(%#v, %#v) returned \"%t\" when \"%t\" " + 
                "was expected",
                node1, node2, hasEdge, testCase.output,
            )
        }
    }
}
