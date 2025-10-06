// 代码生成时间: 2025-10-06 22:29:49
package main

import (
    "encoding/json"
    "github.com/revel/revel"
    "database/sql"
    _ "github.com/lib/pq" // Postgres driver
)

// Supplier represents a supplier in the system.
type Supplier struct {
    ID        int    "json:"id""
    Name      string "json:"name""
    Contact   string "json:"contact""
    CreatedAt string "json:"createdAt""
}

// Controller for handling supplier management operations.
type SupplierController struct {
    *revel.Controller
}

// IndexAction lists all suppliers in the system.
func (c SupplierController) IndexAction() revel.Result {
    suppliers := []Supplier{}
    err := c.Txn.WithTx(func(tx *sql.Tx) error {
        rows, err := tx.Query("SELECT id, name, contact, created_at FROM suppliers")
        if err != nil {
            return err
        }
        defer rows.Close()

        for rows.Next() {
            var s Supplier
            if err := rows.Scan(&s.ID, &s.Name, &s.Contact, &s.CreatedAt); err != nil {
                return err
            }
            suppliers = append(suppliers, s)
        }
        return rows.Err()
    })
    if err != nil {
        // Handle error here, possibly returning a 500 Internal Server Error response.
        return c.RenderError(err)
    }

    return c.Render(suppliers)
}

// AddSupplierAction adds a new supplier to the system.
func (c SupplierController) AddSupplierAction(name, contact string) revel.Result {
    result, err := c.Txn.NamedQuery("INSERT INTO suppliers (name, contact, created_at) VALUES (:name, :contact, CURRENT_TIMESTAMP) RETURNING id",
        map[string]interface{}{"name": name, "contact": contact})
    if err != nil {
        // Handle error here, possibly returning a 500 Internal Server Error response.
        return c.RenderError(err)
    }
    defer result.Close()

    var supplierID int
    if err := result.Scan(&supplierID); err != nil {
        return c.RenderError(err)
    }

    // Return the newly created supplier.
    return c.RenderJSON(Supplier{ID: supplierID, Name: name, Contact: contact, CreatedAt: revel.Now().Format("2006-01-02 15:04:05")})
}

// Error handling middleware to handle errors uniformly.
func RenderError(c *revel.Controller, err error) revel.Result {
    // Log the error
    revel.ERROR.Printf("%s", err)
    // Return a JSON response with the error message
    return c.RenderJSON(map[string]string{"error": err.Error()})
}
