// Copyright alphaair 2016
// 这是一些与数据相关的有趣问题

package amusing

// PrimeNumber 求指定数范围内所有素数
func PrimeNumber(max int) []int {
    
    nums := []int{2}
    
    for i := 3; i <= max; i++ {
        
        if IsPrimeNumber(i) {
            nums = append(nums,i)
        }
        
    }
    
    return nums
    
}

// IsPrimeNumber验证指定的数字是否为素数 
func IsPrimeNumber(num int) bool {
    if num <=1 {
        return false
    }
    
    for i := 2; i<num; i++ {
        if num % i == 0 {
            return false
        }
    }
    
    return true
}