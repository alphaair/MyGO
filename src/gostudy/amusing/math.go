// Copyright alphaair 2016
// 这是一些与数据相关的有趣问题

package amusing

// PrimeNumber 求指定数范围内所有素数
func PrimeNumber(max int) []int {
    
    nums := []int{2}
    
    for i := 3; i <= max; i++ {
        
        isn := true
        for j:=2; j<i; j++ {
            if i % j == 0 {
                isn = false
                break
            }
        }
        
        if isn {
            nums = append(nums,i)
        }
    }
    
    return nums
    
}
