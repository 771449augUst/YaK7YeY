// 代码生成时间: 2025-09-24 01:02:40
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "revel"
    "sort"
)

// SortingService 是一个控制器，用于处理排序相关的请求
type SortingService struct {
    revel.Controller
}

// SortJSON 是一个动作，它接收一个 JSON 数组，对其排序，并返回排序后的结果
func (c SortingService) SortJSON() revel.Result {
    var numbers []int
    var buffer bytes.Buffer
    if err := json.NewDecoder(c.Request.Request.Body).Decode(&numbers); err != nil {
        // 如果发生错误，返回错误信息和状态码 400
        return c.RenderJSON(ErrorResponse{Error: err.Error(), Code: 400})
    }
    // 对数字数组进行排序
    sort.Ints(numbers)
    // 将排序后的结果编码为 JSON 并返回
    buffer.Reset()
    if err := json.NewEncoder(&buffer).Encode(numbers); err != nil {
        return c.RenderJSON(ErrorResponse{Error: err.Error(), Code: 500})
    }
    return c.RenderJSON(buffer.String())
}

// ErrorResponse 是一个用于返回错误信息的结构体
type ErrorResponse struct {
    Error string `json:"error"`
    Code  int    `json:"code"`
}

func init() {
    // 注册路由，将 /sort/json 映射到 SortJSON 动作
    revel.Router(
        (*SortingService)(nil),
        "SortingService",
        [1]string{"SortJSON"},
        "SortJSON",
        [2]string{"*"},
    )
}
