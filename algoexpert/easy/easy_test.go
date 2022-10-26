package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlusOne(t *testing.T) {
	assert.Equal(t, []int{2}, plusOne([]int{1}))
	assert.Equal(t, []int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	assert.Equal(t, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
	assert.Equal(t, []int{1,0}, plusOne([]int{9}))
	assert.Equal(t, []int{1}, plusOne([]int{0}))
	assert.Equal(t, []int{9,0,0}, plusOne([]int{8,9,9}))
	assert.Equal(t, []int{1,0,0,0}, plusOne([]int{9,9,9}))
}

func TestBranchSums(t *testing.T) {
	tree := NewBinaryTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	expected := []int{15, 16, 18, 10, 11}
	output := BranchSums(tree)
	require.Equal(t, expected, output)
}

func TestIsValidSubsequence(t *testing.T) {
	assert.True(t, IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{22, 25, 6}))
}

func TestTwoNumberSum(t *testing.T) {
	sl := []int{-1, 11}
	assert.Equal(t, TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10), sl)
}

func TestNode_Depth(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 2}
	root.Left.Left = &BinaryTree{Value: 4}
	root.Left.Left.Left = &BinaryTree{Value: 8}
	root.Left.Left.Right = &BinaryTree{Value: 9}
	root.Left.Right = &BinaryTree{Value: 5}
	root.Right = &BinaryTree{Value: 3}
	root.Right.Left = &BinaryTree{Value: 6}
	root.Right.Right = &BinaryTree{Value: 7}
	actual := NodeDepths(root)
	require.Equal(t, 16, actual)
}

func TestGetNthFib(t *testing.T) {
	expected := 5
	output := GetNthFib(6)
	require.Equal(t, expected, output)
}

func TestProductSum(t *testing.T) {
	input := SpecialArray{
		5, 2,
		SpecialArray{7, -1},
		3,
		SpecialArray{
			6,
			SpecialArray{
				-13, 8,
			},
			4,
		},
	}
	output := ProductSum(input)
	expected := 12
	require.Equal(t, expected, output)
}

func TestBinarySearch(t *testing.T) {
	expected := 3
	output := BinarySearch([]int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, 33)
	require.Equal(t, expected, output)

	expected = -1
	output = BinarySearch([]int{1, 5, 23, 111}, 120)
	require.Equal(t, expected, output)
}

func TestFindThreeLargestNumbers(t *testing.T) {
	expected := []int{18, 141, 541}
	output := FindThreeLargestNumbers([]int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7})
	require.Equal(t, expected, output)
}

func TestBubbleSort(t *testing.T) {
	expected := []int{2, 3, 5, 5, 6, 8, 9}
	output := BubbleSort([]int{8, 5, 2, 9, 5, 6, 3})
	require.Equal(t, expected, output)
}

func TestInsertionSort(t *testing.T) {
	expected := []int{2, 3, 5, 5, 6, 8, 9}
	output := InsertionSort([]int{8, 5, 2, 9, 5, 6, 3})
	require.Equal(t, expected, output)
}

func TestSelectionSort(t *testing.T) {
	expected := []int{2, 3, 5, 5, 6, 8, 9}
	output := SelectionSort([]int{8, 5, 2, 9, 5, 6, 3})
	require.Equal(t, expected, output)
}

func TestIsPalindrome(t *testing.T) {
	output := IsPalindrome("abcdcba")
	require.Equal(t, true, output)
}

func TestCaesarCipherEncryptor(t *testing.T) {
	expected := "zab"
	output := CaesarCipherEncryptor("xyz", 2)
	require.Equal(t, expected, output)
}

func TestRunLengthEncoding(t *testing.T) {
	//input := "AAAAAAAAAAAAABBCCCCDD"
	//expected := "9A4A2B4C2D"
	//actual := RunLengthEncoding(input)
	//require.Equal(t, expected, actual)

	//input := "aA"
	//expected := "1a1A"
	//actual := RunLengthEncoding(input)
	//require.Equal(t, expected, actual)

	input := "........______=========AAAA   AAABBBB   BBB"
	expected := "8.6_9=4A3 3A4B3 3B"
	actual := RunLengthEncoding(input)
	require.Equal(t, expected, actual)
}

func TestClassPhotos(t *testing.T) {
	redShirtHeights := []int{5, 8, 1, 3, 4}
	blueShirtHeights := []int{6, 9, 2, 4, 5}
	actual := ClassPhotos(redShirtHeights, blueShirtHeights)
	require.Equal(t, true, actual)
}

func TestGenerateDocument(t *testing.T) {
	characters := "Bste!hetsi ogEAxpelrt x "
	document := "AlgoExpert is the Best!"
	expected := true
	actual := GenerateDocument(characters, document)
	require.Equal(t, expected, actual)
}

func TestFirstNonRepeatingCharacter(t *testing.T) {
	input := "abcdcaf"
	expected := 1
	actual := FirstNonRepeatingCharacter(input)
	require.Equal(t, expected, actual)
}

func TestMinimumWaitingTime(t *testing.T) {
	queries := []int{3, 2, 1, 2, 6}
	expected := 17
	actual := MinimumWaitingTime(queries)
	require.Equal(t, expected, actual)
}

func TestNonConstructableChange(t *testing.T) {
	input := []int{5, 6, 1, 1, 2, 3, 4, 9}
	expected := 32
	actual := NonConstructableChange(input)
	require.Equal(t, expected, actual)
}

func TestRemoveDuplicatesFromLinkedList(t *testing.T) {
	input := addMany(&LinkedList{Value: 1}, []int{1, 3, 4, 4, 4, 5, 6, 6})
	expected := addMany(&LinkedList{Value: 1}, []int{3, 4, 5, 6})
	actual := RemoveDuplicatesFromLinkedList(input)
	require.Equal(t, getValues(expected), getValues(actual))
}

func addMany(linkedList *LinkedList, values []int) *LinkedList {
	current := linkedList
	for current.Next != nil {
		current = current.Next
	}
	for _, value := range values {
		current.Next = &LinkedList{Value: value}
		current = current.Next
	}
	return linkedList
}

func getValues(linkedList *LinkedList) []int {
	var values []int
	current := linkedList
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

func TestSortedSquaredArray(t *testing.T) {
	input := []int{1, 2, 3, 5, 6, 8, 9}
	expected := []int{1, 4, 9, 25, 36, 64, 81}
	actual := SortedSquaredArray(input)
	require.Equal(t, expected, actual)
}

func TestTournamentWinner(t *testing.T) {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	expected := "Python"
	actual := TournamentWinner(competitions, results)
	require.Equal(t, expected, actual)
}
