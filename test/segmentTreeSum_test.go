package test

import (
	"testing"

	"github.com/bejohi/SegmentTree_Go/segTree"
)

func getMockTree() *[]float64 {
	tree := []float64{1, 3, 5, 7, 9, 11}
	return &tree
}

func TestConstructTree_InputArrayIsNull_NoPanicAndErrorReturn(t *testing.T) {
	// Arrange
	tree := segTree.SegmentTreeSum{}

	// Act
	err := tree.Construct(nil)
	innerTree, innerTreeError := tree.GetTree()

	// Assert
	if err == nil || innerTreeError == nil {
		t.Error("No error was returned")
	}

	if innerTree != nil {
		t.Error("A tree was returned")
	}
}

func TestGetTree_TreeNotConstructed_ErrorReturned(t *testing.T) {
	// Arrange
	tree := segTree.SegmentTreeSum{}

	// Act
	result, err := tree.GetTree()

	// Assert
	if result != nil || err == nil {
		t.Error("Wrong result for unconstructed tree.")
	}
}

func TestGetTree_TreeConstructed_CorrectTreeReturned(t *testing.T) {
	// Arrange
	tree := segTree.SegmentTreeSum{}
	treeArr := getMockTree()
	tree.Construct(treeArr)

	// Act
	result, err := tree.GetTree()

	// Assert
	if err != nil {
		t.Error("Error returned.")
	}

	if (*result)[0] != 36 || (*result)[1] != 9 ||
		(*result)[2] != 27 || (*result)[3] != 4 ||
		(*result)[4] != 5 || (*result)[5] != 16 ||
		(*result)[6] != 11 || (*result)[7] != 1 ||
		(*result)[8] != 3 {
		t.Error("The tree was not constructed correct.")

	}
}
