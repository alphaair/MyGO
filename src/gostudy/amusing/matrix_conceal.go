// Copyright alphaair 2016
// 这是一个数字矩阵藏宝游戏，来源来最强大脑节点，矩阵藏有五个质数，但是只有

package amusing

import "math/rand"

// MatrixConceal 矩阵藏地图
type MatrixConceal struct {
    Result [][]int
}

// newNum生成一个指定范围内的随机数
func (m *MatrixConceal) newNum(min int, max int) int {
    
    n := rand.Intn(max+1)
    
    for ; n < min; {
        n = rand.Intn(max+1)
    }
    
    return n
}

// Init 初始化矩阵藏宝
func (m *MatrixConceal) Init() {
    rs := m.newNum(7, 13)
    cs := m.newNum(13, 19)
    
    m.Result = make([][]int, rs, rs)
    
    for i, col := range m.Result {
        
        col = make([]int, cs, cs)
        for i := range col {
            col[i] = m.newNum(766, 2434)
        }
        
        m.Result[i] = col
    }
}