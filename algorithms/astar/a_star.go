package astar

import (
	"math"

	"github.com/shomali11/go-interview/datastructures/linkedlists/singlylinkedlists"
	"github.com/shomali11/go-interview/datastructures/priorityqueues"
	"github.com/shomali11/go-interview/datastructures/sets/hashsets"
)

var (
	positiveInfinity = math.Inf(1)
)

// Node a node interface
type Node interface {
	EstimateDistanceToGoal(goal Node) float64
	ActualDistanceToNeighbor(neighbor Node) float64
	GetNeighbors() []Node
}

type wrapper struct {
	h      float64
	g      float64
	node   Node
	parent *wrapper
}

// New new factory
func New() *AStar {
	compare := func(one, two interface{}) bool {
		f1 := one.(*wrapper).g + one.(*wrapper).h
		f2 := two.(*wrapper).g + two.(*wrapper).h

		// Tie Breaker ... Maximum G Value
		if f1 == f2 {
			g1 := one.(*wrapper).g
			g2 := two.(*wrapper).g

			// Maximum G Value
			return g1 > g2
		}

		// Minimum F Value
		return f1 < f2
	}
	return &AStar{openPQ: priorityqueues.New(compare), closedSet: hashsets.New(), nodeWrapperMap: make(map[Node]*wrapper)}
}

// AStar a star
type AStar struct {
	openPQ         *priorityqueues.PriorityQueue
	closedSet      *hashsets.HashSet
	nodeWrapperMap map[Node]*wrapper
}

// Search search
func (s *AStar) Search(start Node, goal Node) ([]Node, float64) {
	defer s.clear()

	startWrapper := s.getWrapper(start)
	startWrapper.h = start.EstimateDistanceToGoal(goal)

	goalWrapper := s.getWrapper(goal)

	s.openPQ.Push(startWrapper)

	// While not Empty ...
	for !s.openPQ.IsEmpty() {
		// Extract Top of the Heap ...
		currentElement, _ := s.openPQ.Pop()
		currentWrapper := currentElement.(*wrapper)

		// Did we reach Goal ?
		if currentWrapper == goalWrapper {
			// Shortest Path ...
			return getPath(goalWrapper), goalWrapper.g
		}

		// Set Node to Visited ...
		s.closedSet.Add(currentWrapper)

		// Get List of neighbors ..
		neighbors := currentWrapper.node.GetNeighbors()

		// Traverse the neighbors ...
		for _, neighborNode := range neighbors {
			neighborWrapper := s.getWrapper(neighborNode)

			// We only care about the ones we did not visit ...
			if !s.closedSet.Contains(neighborWrapper) {
				// Is not in the Heap ..
				if !s.openPQ.Contains(neighborWrapper) {
					neighborWrapper.g = positiveInfinity
					neighborWrapper.parent = nil
				}

				s.updateVertex(currentWrapper, neighborWrapper, goalWrapper)
			}
		}
	}
	return make([]Node, 0), 0
}

func (s *AStar) updateVertex(currentWrapper *wrapper, neighborWrapper *wrapper, goalWrapper *wrapper) {
	// Check for Consistency ..
	totalCostToNeighbor := currentWrapper.g + currentWrapper.node.ActualDistanceToNeighbor(neighborWrapper.node)
	if totalCostToNeighbor < neighborWrapper.g {
		// Update the G Function
		neighborWrapper.g = totalCostToNeighbor
		neighborWrapper.parent = currentWrapper

		// If Min Heap Has the Node ..
		if s.openPQ.Contains(neighborWrapper) {
			// Remove it ..
			s.openPQ.Remove(neighborWrapper)
		}

		// Add the state with the new function value ..
		neighborWrapper.h = neighborWrapper.node.EstimateDistanceToGoal(goalWrapper.node)
		s.openPQ.Push(neighborWrapper)
	}
}

func (s *AStar) clear() {
	s.nodeWrapperMap = make(map[Node]*wrapper)
	s.closedSet.Clear()
	s.openPQ.Clear()
}

func (s *AStar) getWrapper(node Node) *wrapper {
	w, ok := s.nodeWrapperMap[node]
	if ok {
		return w
	}

	w = &wrapper{node: node}
	s.nodeWrapperMap[node] = w
	return w
}

func getPath(goalWrapper *wrapper) []Node {
	path := singlylinkedlists.New()
	currentWrapper := goalWrapper
	for currentWrapper != nil {
		path.InsertAt(0, currentWrapper.node)
		currentWrapper = currentWrapper.parent
	}

	nodeList := make([]Node, 0)
	for _, value := range path.GetValues() {
		nodeList = append(nodeList, value.(Node))
	}
	return nodeList
}
