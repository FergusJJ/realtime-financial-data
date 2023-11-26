package server

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

func InterceptorLogger(l *slog.Logger) logging.Logger {
  return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
    logEntry := map[string]interface{}{
      "timestamp": time.Now().Format(time.RFC3339),
      "level": slog.Level(lvl),
      "message":msg,
    }
    jsonLog, err := json.Marshal(logEntry)
    if err != nil {
      l.ErrorContext(ctx, "Failed to marshal log entry", err)
      return
    }
    l.Log(ctx, slog.Level(lvl), string(jsonLog))
    })
}
