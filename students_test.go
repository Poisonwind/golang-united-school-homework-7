package coverage

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLen(t *testing.T) {

	testCases := [3]People{
		make(People, 0),
		make(People, 100),
	}

	for _, val := range(testCases) {
		assert.Equal(t, len(val), val.Len())		
	}

}

func TestLess(t *testing.T) {

	testPeople := People{
		{"John", "Doe", time.Date(1980, 1, 1, 0, 0, 0, 0, time.Local)},
		{"John", "Brown", time.Date(1980, 1, 1, 0, 0, 0, 0, time.Local)},
		{"Joe", "Doe", time.Date(1980, 1, 1, 0, 0, 0, 0, time.Local)},
		{"John", "Doe", time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)},
	}

	testCases := []struct {
		first int
		second int
		result bool
	}{
		{0, 1, false},
		{0, 2, false},
		{0, 3, false},
		{1, 0, true},
		{2, 0, true},
		{3, 0, true},
	}

	for _, val := range(testCases) {
		assert.Equal(t, val.result, testPeople.Less(val.first, val.second))
	}
}

func TestSwap(t *testing.T) {

	peopleBefore := People{
		{"John", "Doe", time.Date(1980, 1, 1, 0, 0, 0, 0, time.Local)},
		{"John", "Brown", time.Date(1980, 1, 1, 0, 0, 0, 0, time.Local)},
	}

	peopleAfter := People{
		{"John", "Brown", time.Date(1980, 1, 1, 0, 0, 0, 0, time.Local)},
		{"John", "Doe", time.Date(1980, 1, 1, 0, 0, 0, 0, time.Local)},
	}

	peopleBefore.Swap(0, 1)

	assert.Equal(t, peopleBefore[0], peopleAfter[0])
}

func TestNewMatrix(t *testing.T) {

	TestCases := []struct {
		input string
		res *Matrix
	}{
		{"1 2 3\n4 5 6", &Matrix{2, 3, []int{1, 2, 3, 4, 5, 6}}},
		{"1 2 3\n4 5 6\n7 8 9", &Matrix{3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}},
	}

	for _, val := range(TestCases) {

		newMatrix, _ := New(val.input)
		assert.Equal(t, val.res, newMatrix)
	}

	NilCases := []string{
		"1 2 3\n4 5",
		"1 b c\n2 3 q",
	}

	for _, val := range(NilCases) {
		newMatrix, _ := New(val)
		assert.Nil(t, newMatrix)
	}
}

func TestRow(t *testing.T){

	testCases := []struct{
		input Matrix
		res [][]int
	}{
		{Matrix{3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		{Matrix{2, 3, []int{1, 2, 3, 4, 5, 6}}, [][]int{{1, 2, 3}, {4, 5, 6}}},
	}

	for _, val := range(testCases){
		rows := val.input.Rows()
		assert.Equal(t, val.res, rows) 
	}

}

func TestCols(t *testing.T){

	testCases := []struct{
		input Matrix
		res [][]int
	}{
		{Matrix{3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{Matrix{2, 3, []int{1, 2, 3, 4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}},
	}

	for _, val := range(testCases){
		cols := val.input.Cols()
		assert.Equal(t, val.res, cols) 
	}



}

func TestSet(t *testing.T){

	testCases := []struct{
		input *Matrix
		row, col, value int
	}{
		{&Matrix{3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 0, 0, 11},
		{&Matrix{2, 3, []int{1, 2, 3, 4, 5, 6}}, 1, 1, 22},
	}

	for _, val := range(testCases){
		val.input.Set(val.row, val.col, val.value)
		actual, _ := GetVal(val.input, val.row, val.col)
		assert.Equal(t, val.value, actual)
	}
}

func GetVal(m *Matrix, row, col int) (int, error) {

	if row < 0 || m.rows <= row || col < 0 || m.cols <= col {
		return 0, fmt.Errorf("index out of bound")
	}

	index := row*m.cols+col
	return m.data[index], nil
}