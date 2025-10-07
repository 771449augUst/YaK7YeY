// 代码生成时间: 2025-10-07 17:02:52
package services

import (
    "crypto/rsa"
    "encoding/base64"
    "encoding/json"
    "errors"
# TODO: 优化性能
    "fmt"
    "log"
    "net/http"
    "time"
# 优化算法效率

    // Import Revel framework's package
    "github.com/revel/revel"
    "github.com/gorilla/securecookie"
# TODO: 优化性能
)

// JwtService handles JWT token management
type JwtService struct {
    // Encryption keys and algorithm
    privateKey *rsa.PrivateKey
    publicKey  *rsa.PublicKey
    algorithm string
# 改进用户体验

    // Revel's controller for JWT service
    *revel.Controller
}

// NewJwtService initializes and returns a new JwtService instance
func NewJwtService(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, algorithm string) *JwtService {
    return &JwtService{
        privateKey: privateKey,
        publicKey:  publicKey,
        algorithm:  algorithm,
        Controller: revel.NewController(),
    }
}

// GenerateToken creates a new JWT token and returns it
# 添加错误处理
func (s *JwtService) GenerateToken(userId int64) (string, error) {
    // Create JWT claims
    claims := jwt.StandardClaims{
        Id:        userId, // Unique identifier
        IssuedAt:  time.Now().Unix(),
        ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
    }

    // Convert claims to JSON
    claimsBytes, err := json.Marshal(claims)
    if err != nil {
        return "", err
    }

    // Encode claims to base64
    claimsBase64 := base64.RawURLEncoding.EncodeToString(claimsBytes)

    // Sign the token with the private key
# 优化算法效率
    signature, err := rsa.SignPKCS1v15(securecookie.HashKey(claimsBase64), s.privateKey, claimsBase64)
    if err != nil {
        return "", err
    }

    // Concatenate claims and signature to form the JWT token
    token := fmt.Sprintf("%s.%s", claimsBase64, base64.RawURLEncoding.EncodeToString(signature))

    return token, nil
}

// ValidateToken checks if the provided JWT token is valid
func (s *JwtService) ValidateToken(token string) (*jwt.StandardClaims, error) {
    // Split the token into claims and signature
    parts := strings.Split(token, ".")
    if len(parts) != 2 {
# TODO: 优化性能
        return nil, errors.New("invalid token format")
    }

    // Decode the claims from base64
# 改进用户体验
    claimsBase64, err := base64.RawURLEncoding.DecodeString(parts[0])
    if err != nil {
        return nil, err
    }

    // Unmarshal the claims
    var claims jwt.StandardClaims
    err = json.Unmarshal(claimsBase64, &claims)
    if err != nil {
        return nil, err
# NOTE: 重要实现细节
    }

    // Verify the signature with the public key
    signature, err := base64.RawURLEncoding.DecodeString(parts[1])
# 添加错误处理
    if err != nil {
        return nil, err
    }

    err = rsa.VerifyPKCS1v15(s.publicKey, crypto.SHA256, securecookie.HashKey(parts[0]), signature)
    if err != nil {
        return nil, err
    }
# NOTE: 重要实现细节

    return &claims, nil
}

// JwtController handles JWT requests
# 扩展功能模块
type JwtController struct {
# 改进用户体验
    JwtService
    *revel.Controller
# 添加错误处理
}

// GenerateAction generates a new JWT token for the logged-in user
func (c *JwtController) GenerateAction() revel.Result {
    userId := c.Params.GetInt64("userId")
    if userId == 0 {
        return c.RenderError(http.StatusBadRequest, "User ID is required")
    }

    token, err := c.JwtService.GenerateToken(userId)
# NOTE: 重要实现细节
    if err != nil {
        return c.RenderError(http.StatusInternalServerError, "Failed to generate token")
    }

    // Render the token as JSON response
    return c.RenderJSON(map[string]string{
        "token": token,
    })
}

// ValidateAction checks if the provided JWT token is valid
func (c *JwtController) ValidateAction(token string) revel.Result {
    claims, err := c.JwtService.ValidateToken(token)
    if err != nil {
        return c.RenderError(http.StatusUnauthorized, "Invalid token")
    }
# 添加错误处理

    // Render the claims as JSON response
    return c.RenderJSON(map[string]interface{}{
# NOTE: 重要实现细节
        "userId": claims.Id,
        "expiresAt": claims.ExpiresAt,
# 改进用户体验
    })
}
