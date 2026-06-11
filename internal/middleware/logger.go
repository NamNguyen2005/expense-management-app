package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(data []byte) (n int, err error) {
	w.body.Write(data)
	return  w.ResponseWriter.Write(data)
	
}

func LoggingMiddleware() gin.HandlerFunc {
	logPath := "internal/logs/https.log"
	
	// Tạo thư mục logs nếu chưa tồn tại
	logDir := filepath.Dir(logPath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}
	
	logger := zerolog.New(&lumberjack.Logger{
    Filename:   logPath,
    MaxSize:    1, // megabytes
    MaxBackups: 5,
    MaxAge:     5, //days
    Compress:   true,// disabled by default
	LocalTime: true,
}).With().Caller().Logger()

	return func(ctx *gin.Context) {
		contentType := ctx.GetHeader("Content-Type")
		requestBody := make(map[string]interface{})
		var formFiles []map[string]any
		if strings.HasPrefix(contentType , "multipart/form-data") {
			log.Printf("Skipping logging for multipart/form-data request")
			if err := ctx.Request.ParseMultipartForm(32 << 20); err != nil {
				for key, values := range ctx.Request.MultipartForm.Value {
					if len(values) == 1 {
						requestBody[key] = values[0]
					} else {
						requestBody[key] = values
					}
				}
				for filed , files := range ctx.Request.MultipartForm.File {
					for _ , file := range files {
						formFiles = append(formFiles , map[string]any {
							"filename" : file.Filename ,
							"size" : formatFileSize(file.Size) ,
							"header" : file.Header ,
							"field" : filed ,
							"content_type" : file.Header.Get("Content-Type") ,
						})
					}
				}
				if len(formFiles) > 0 {
					requestBody["files"] = formFiles
				}
			}
			
		}else {
			bodyBytes, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				logger.Error().Err(err).Msg("Error reading request body")
			}
			fmt.Printf("Request Body: %s\n", string(bodyBytes))
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			if strings.HasPrefix(contentType , "application/json"){
				json.Unmarshal(bodyBytes , &requestBody)
			}else {
				values , _ := url.ParseQuery(string(bodyBytes))
				for key , value := range values {
					if len(value) == 1 {
						requestBody[key] = value[0]
					}else{
						requestBody[key] = value
					}
				}
			}
		}
		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			logger.Error().Err(err).Msg("Error reading request body")
		}
		fmt.Printf("Request Body: %s\n", string(bodyBytes))
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		startTime := time.Now()

		customWriter := &CustomResponseWriter{ResponseWriter: ctx.Writer, body: bytes.NewBufferString("")}
		ctx.Writer = customWriter
		ctx.Next()

		responseBody := customWriter.body.String()
		responseContentType := ctx.Writer.Header().Get("Content-Type")
		var responseBodyParse interface{}
		if strings.HasPrefix(responseContentType , "image/"){
			responseBodyParse = "[binary data]" 
		}else if strings.HasPrefix(responseContentType , "application/json") &&
			(strings.HasPrefix(strings.TrimSpace(responseBody), "{") ||
			 strings.HasPrefix(strings.TrimSpace(responseBody), "[")) {
			if err := json.Unmarshal([]byte(responseBody), &responseBodyParse); err != nil {
				responseBodyParse = responseBody
			}
		} else {
			responseBodyParse = responseBody
		}
		log.Printf("Response Body: %s\n", responseBody)


		duration := time.Since(startTime)
		logEvent := logger.Info()
		statusCode := ctx.Writer.Status()
		if statusCode >= 500 {
			logEvent = logger.Error()
		} else if statusCode >= 400 {
			logEvent = logger.Warn()
		}
		logEvent.
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("client_ip", ctx.ClientIP()).
			Interface("header", ctx.Request.Header).
			Str("user_agent", ctx.Request.UserAgent()).
			Str("protocol", ctx.Request.Proto).
			Str("host", ctx.Request.Host).
			Str("referer", ctx.Request.Referer()).
			Str("query", ctx.Request.URL.RawQuery).
			Int("status", statusCode).
			Interface("request_body", requestBody).
			Int64("duration_ms", duration.Milliseconds()).
			Interface("response_body", responseBodyParse).
			Msg("HTTP request logged")
	}
}

func formatFileSize(size int64) string {
	switch {
	case size < 1024:
		return fmt.Sprintf("%d B", size)
	case size < 1024*1024:
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	case size < 1024*1024*1024:
		return fmt.Sprintf("%.2f MB", float64(size)/(1024*1024))
	default:
		return fmt.Sprintf("%.2f GB", float64(size)/(1024*1024*1024))
	}
}