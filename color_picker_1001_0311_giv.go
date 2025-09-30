// 代码生成时间: 2025-10-01 03:11:18
package controllers

import (
    "encoding/json"
    "github.com/revel/revel"
)

// ColorPicker is a controller for handling color selection.
type ColorPicker struct {
    revel.Controller
}

// Index action returns a JSON response with available colors.
func (c ColorPicker) Index() revel.Result {
    colors := map[string]string{
        "Red": "#FF0000",
        "Green": "#00FF00",
        "Blue": "#0000FF",
        "Yellow": "#FFFF00",
        "Black": "#000000",
    }

    // Transform colors map to JSON
    colorJSON, err := json.Marshal(colors)
    if err != nil {
        // Handle encoding error
        return c.RenderError(err)
    }

    // Return JSON response
    return c.RenderJSON(colorJSON)
}
