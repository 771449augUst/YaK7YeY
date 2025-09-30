// 代码生成时间: 2025-09-30 21:58:54
package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"
)

// MockDataGenerator 结构体用于生成模拟数据
type MockDataGenerator struct {
    // 这里可以添加更多属性来自定义模拟数据生成器的行为
}

// NewMockDataGenerator 创建一个新的 MockDataGenerator 实例
func NewMockDataGenerator() *MockDataGenerator {
    return &MockDataGenerator{}
}

// GenerateString 生成一个随机的字符串
func (g *MockDataGenerator) GenerateString(length int) string {
    // 定义可能的字符
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

    // 生成随机字符串
    var b strings.Builder
    for i := 0; i < length; i++ {
        b.WriteByte(charset[seededRand.Intn(len(charset))])
    }
    return b.String()
}

// GenerateInt 生成一个随机的整数
func (g *MockDataGenerator) GenerateInt(min, max int) int {
    if min > max {
        log.Fatalf("max must be greater than min")
    }
    return rand.Intn(max-min) + min
}

// GenerateData 生成模拟数据
func (g *MockDataGenerator) GenerateData() map[string]interface{} {
    var mockData = make(map[string]interface{})
    mockData["id"] = g.GenerateInt(1, 100)
    mockData["name"] = g.GenerateString(5)
    mockData["email"] = g.GenerateString(10) + "@example.com"
    mockData["created_at"] = time.Now().Format(time.RFC3339)
    return mockData
}

func main() {
    // 创建模拟数据生成器实例
    generator := NewMockDataGenerator()

    // 生成并打印模拟数据
    mockData := generator.GenerateData()
    fmt.Printf("Generated Mock Data: %+v
", mockData)
}
