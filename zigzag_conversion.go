package leetcode

func ZigZagConversion( s string, numRows int ) string {
	if numRows == 1 {
		return s
	}
	// define a matrix 
	matrix := make([][]byte, numRows)
	bottom, top := 0, numRows - 2
	for i :=0; i!=len(s);{
		if bottom != numRows {
			matrix[bottom] = append(matrix[bottom], byte(s[i]))
			i++
			bottom++
		} else if top > 0 {
			matrix[top] = append(matrix[top], byte(s[i]))
			top --
			i++
		} else {
			top = numRows - 2
			bottom = 0
		}
	}
	// convert the matrix to string
	var result string
	for _, v := range matrix {
		for _, c := range v {
			result += string(c)
		}
	}
	return result
}
