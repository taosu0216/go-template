package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// TokenCache 用于存储已使用的token
var (
	usedTokens     = make(map[string]bool)
	usedTokensLock sync.RWMutex
	// 用于限制IP访问频率
	ipLimiter = make(map[string]*IPLimit)
	ipLock    sync.RWMutex
)

type IPLimit struct {
	count    int
	lastTime time.Time
}

// cleanupUsedTokens 清理过期的token
func cleanupUsedTokens() {
	usedTokensLock.Lock()
	defer usedTokensLock.Unlock()
	usedTokens = make(map[string]bool)
}

// isTokenUsed 检查token是否已被使用
func isTokenUsed(token string) bool {
	usedTokensLock.RLock()
	defer usedTokensLock.RUnlock()
	return usedTokens[token]
}

// markTokenAsUsed 标记token为已使用
func markTokenAsUsed(token string) {
	usedTokensLock.Lock()
	defer usedTokensLock.Unlock()
	usedTokens[token] = true
}

// checkIPLimit 检查IP访问频率
func checkIPLimit(ip string) bool {
	ipLock.Lock()
	defer ipLock.Unlock()

	now := time.Now()
	limit, exists := ipLimiter[ip]

	if !exists {
		ipLimiter[ip] = &IPLimit{
			count:    1,
			lastTime: now,
		}
		return true
	}

	// 如果距离上次请求超过1分钟，重置计数
	if now.Sub(limit.lastTime) > time.Minute {
		limit.count = 1
		limit.lastTime = now
		return true
	}

	// 每分钟限制30次请求
	if limit.count >= 30 {
		return false
	}

	limit.count++
	return true
}

// VideoAuthMiddleware 视频防盗链中间件
func VideoAuthMiddleware() app.HandlerFunc {
	// 每小时清理一次已使用的token
	go func() {
		for {
			time.Sleep(time.Hour)
			cleanupUsedTokens()
		}
	}()

	return func(c context.Context, ctx *app.RequestContext) {
		// 0. 检查用户是否已登录
		userID := ctx.GetString("userID")
		if userID == "" {
			ctx.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"code": 401,
				"msg":  "Unauthorized",
			})
			ctx.Abort()
			return
		}

		// 0.1 检查IP访问频率
		clientIP := string(ctx.ClientIP())
		if !checkIPLimit(clientIP) {
			ctx.JSON(consts.StatusTooManyRequests, map[string]interface{}{
				"code": 429,
				"msg":  "Too many requests",
			})
			ctx.Abort()
			return
		}
		// 1. 检查Referer
		referer := ctx.Request.Header.Get("Referer")
		if referer == "" {
			ctx.JSON(consts.StatusForbidden, map[string]interface{}{
				"code": 403,
				"msg":  "Invalid request",
			})
			ctx.Abort()
			return
		}

		// 解析Referer
		refererURL, err := url.Parse(referer)
		if err != nil || !isAllowedDomain(refererURL.Host) {
			ctx.JSON(consts.StatusForbidden, map[string]interface{}{
				"code": 403,
				"msg":  "Invalid referer",
			})
			ctx.Abort()
			return
		}

		// 2. 验证时间戳和token
		timestamp := ctx.Query("t")
		token := ctx.Query("token")

		// 检查时间戳是否有效
		t, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			ctx.JSON(consts.StatusForbidden, map[string]interface{}{
				"code": 403,
				"msg":  "Invalid timestamp",
			})
			ctx.Abort()
			return
		}

		// 检查链接是否过期(默认2小时)
		if time.Now().Unix()-t > 7200 {
			ctx.JSON(consts.StatusForbidden, map[string]interface{}{
				"code": 403,
				"msg":  "Link expired",
			})
			ctx.Abort()
			return
		}

		// 验证token
		path := string(ctx.Path())

		// 检查token是否已被使用
		if isTokenUsed(token) {
			ctx.JSON(consts.StatusForbidden, map[string]interface{}{
				"code": 403,
				"msg":  "Token already used",
			})
			ctx.Abort()
			return
		}

		if !validateToken(path, timestamp, token) {
			ctx.JSON(consts.StatusForbidden, map[string]interface{}{
				"code": 403,
				"msg":  "Invalid token",
			})
			ctx.Abort()
			return
		}

		// 标记token为已使用
		markTokenAsUsed(token)

		ctx.Next(c)
	}
}

// isAllowedDomain 检查域名是否在允许列表中
func isAllowedDomain(domain string) bool {
	// 这里可以配置允许的域名列表
	allowedDomains := []string{
		"localhost",
		"127.0.0.1",
		// 添加其他允许的域名
	}

	for _, d := range allowedDomains {
		if strings.Contains(domain, d) {
			return true
		}
	}
	return false
}

// validateToken 验证token
func validateToken(path, timestamp, token string) bool {
	// 密钥,实际应用中应该从配置文件读取
	secretKey := "your-secret-key"

	// 构造签名字符串
	signStr := path + timestamp + secretKey

	// 计算MD5
	hash := md5.New()
	hash.Write([]byte(signStr))
	expectedToken := hex.EncodeToString(hash.Sum(nil))

	return token == expectedToken
}
