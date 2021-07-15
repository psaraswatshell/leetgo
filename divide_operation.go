package leetcode

func Divide(dividend int, divisor int) int {
    
    sign:= -1
    if dividend < 0 && divisor < 0 {
        sign = 1
    }
    if dividend > 0 && divisor > 0 {
        sign = 1
    }
    dvd := abs(dividend)
    dvs := abs(divisor)
    quotient :=0
    temp :=0
    for i := 31; i > -1 ; i-- {
        if (temp+ (dvs << i )) <= dvd {
            temp = temp + dvs << i
            quotient = quotient | 1 << i
        }
    }
    return sign*quotient
    
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
