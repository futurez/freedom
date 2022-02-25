package flog

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
func StatusCodeColor(code int) string {
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

func MethodColor(method string) string {
	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		cost := time.Since(start)
		statusCode := c.Writer.Status()
		statusColor := StatusCodeColor(c.Writer.Status())
		methodColor := MethodColor(c.Request.Method)
		switch {
		case statusCode >= 400 && statusCode <= 499:
			{
				supar.Warnf("[GIN] |%s %3d %s| %10v | %15s |%s %-7s %s %s %s\n%s",
					statusColor, c.Writer.Status(), reset,
					cost,
					c.ClientIP(),
					methodColor, c.Request.Method, reset,
					path,
					query,
					c.Errors.ByType(gin.ErrorTypePrivate).String())
			}
		case statusCode >= 500:
			{
				supar.Errorf("[GIN] |%s %3d %s| %10v | %15s |%s %-7s %s %s %s\n%s",
					statusColor, c.Writer.Status(), reset,
					cost,
					c.ClientIP(),
					methodColor, c.Request.Method, reset,
					path,
					query,
					c.Errors.ByType(gin.ErrorTypePrivate).String())
			}
		default:
			supar.Infof("[GIN] |%s %3d %s| %10v | %15s |%s %-7s %s %s %s\n%s",
				statusColor, c.Writer.Status(), reset,
				cost,
				c.ClientIP(),
				methodColor, c.Request.Method, reset,
				path,
				query,
				c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					supar.Errorf("[GIN] | %s | %+v | %s", c.Request.URL.Path, err, string(httpRequest))
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					supar.Errorf("[GIN] | [Recovery from panic] | %+v | %s | %s",
						err, string(httpRequest), string(debug.Stack()))
				} else {
					supar.Errorf("[GIN] | [Recovery from panic] %+v | %s", err, string(httpRequest))
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
