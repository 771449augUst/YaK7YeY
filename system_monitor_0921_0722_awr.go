// 代码生成时间: 2025-09-21 07:22:36
package main

import (
    "fmt"
    "os/exec"
    "revel"
    "strings"
)

// SystemMonitor struct contains methods for system performance monitoring.
type SystemMonitor struct {
    // Empty struct as Revel doesn't require any fields for controller
}

// Controller 负责处理系统性能监控相关的请求
type Controller struct {
    *revel.Controller
    systemMonitor SystemMonitor
}

// GetSystemInfo 获取系统信息
func (c Controller) GetSystemInfo() revel.Result {
    // 获取CPU信息
    cpuInfo := c.getCPUInfo()
    // 获取内存信息
    memInfo := c.getMemInfo()
    // 获取磁盘信息
    diskInfo := c.getDiskInfo()

    // 返回系统信息
    return c.RenderJSON(map[string]interface{}{
        "cpuInfo": cpuInfo,
        "memInfo": memInfo,
        "diskInfo": diskInfo,
    })
}

// getCPUInfo 获取CPU使用率信息
func (c *Controller) getCPUInfo() string {
    cmd := exec.Command("top", "-b", "-n", "1")
    output, err := cmd.CombinedOutput()
    if err != nil {
        revel.ERROR.Printf("Failed to get CPU info: %v", err)
        return ""
    }
    return strings.TrimSpace(string(output))
}

// getMemInfo 获取内存使用率信息
func (c *Controller) getMemInfo() string {
    cmd := exec.Command("free", "-m")
    output, err := cmd.CombinedOutput()
    if err != nil {
        revel.ERROR.Printf("Failed to get memory info: %v", err)
        return ""
    }
    return strings.TrimSpace(string(output))
}

// getDiskInfo 获取磁盘使用率信息
func (c *Controller) getDiskInfo() string {
    cmd := exec.Command("df", "-h")
    output, err := cmd.CombinedOutput()
    if err != nil {
        revel.ERROR.Printf("Failed to get disk info: %v", err)
        return ""
    }
    return strings.TrimSpace(string(output))
}

func init() {
    // Add the routes
    revel.Router(
        (*Controller).GetSystemInfo,
        "GET /system/monitor",
    )
}

func main() {
    revel.Run()
}
