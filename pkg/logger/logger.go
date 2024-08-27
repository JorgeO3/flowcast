// Package logger wraps the zerolog logger and provides a middleware for logging HTTP requests.
package logger

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

// Interface define los métodos que debe implementar un logger.
type Interface interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

// Logger es la implementación del logger usando zerolog.
type Logger struct {
	logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

// New crea una nueva instancia de Logger con el nivel de log especificado.
func New(level string) Interface {
	zerolog.SetGlobalLevel(parseLogLevel(level))

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &Logger{
		logger: &logger,
	}
}

// parseLogLevel convierte el nivel de log de string a zerolog.Level.
func parseLogLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case "error":
		return zerolog.ErrorLevel
	case "warn":
		return zerolog.WarnLevel
	case "info":
		return zerolog.InfoLevel
	case "debug":
		return zerolog.DebugLevel
	default:
		return zerolog.InfoLevel
	}
}

// shortenPath recorta el path hasta el directorio del proyecto.
func shortenPath(file string, line int) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return file + ":" + strconv.Itoa(line)
}

// Debug logs a debug message
func (l *Logger) Debug(message string, args ...interface{}) {
	l.log(zerolog.DebugLevel, message, args...)
}

// Info logs an info message
func (l *Logger) Info(message string, args ...interface{}) {
	l.log(zerolog.InfoLevel, message, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(message string, args ...interface{}) {
	l.log(zerolog.WarnLevel, message, args...)
}

// Error logs an error message
func (l *Logger) Error(message string, args ...interface{}) {
	l.log(zerolog.ErrorLevel, message, args...)
}

// Fatal logs a fatal message and exits the program
func (l *Logger) Fatal(message string, args ...interface{}) {
	l.log(zerolog.FatalLevel, message, args...)
	os.Exit(1)
}

func (l *Logger) log(level zerolog.Level, message string, args ...interface{}) {
	event := l.logger.WithLevel(level)
	if len(args) > 0 {
		event = event.Fields(map[string]interface{}{
			"details": fmt.Sprintf(message, args...),
		})
		message = "Log entry with details"
	}
	event.Msg(message)
}

// ZerologMiddleware for logging HTTP requests
func ZerologMiddleware(logg Interface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap the response writer
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			// Process the request
			next.ServeHTTP(ww, r)

			// Log the response
			logg.(*Logger).logger.Info().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Str("remote_addr", r.RemoteAddr).
				Str("user_agent", r.UserAgent()).
				Int("status", ww.Status()).
				Int("bytes", ww.BytesWritten()).
				Dur("duration", time.Since(start)).
				Msg("Handled request")
		})
	}
}

// ErrorHandlingMiddleware recovers from panics and sends an appropriate error response.
func ErrorHandlingMiddleware(logg Interface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					logg.Error(fmt.Sprintf("Recovered from panic: %v", rec))
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
