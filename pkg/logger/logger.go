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
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

// Logger es la implementación del logger usando zerolog.
type Logger struct {
	logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

// New crea una nueva instancia de Logger con el nivel de log especificado.
func New(level string) Interface {
	zerolog.SetGlobalLevel(parseLogLevel(level))

	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	// zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
	// 	return shortenPath(file, line)
	// }
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

// Debug registra un mensaje de depuración.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.msg("debug", message, args...)
}

// Info registra un mensaje informativo.
func (l *Logger) Info(message string, args ...interface{}) {
	l.log(message, args...)
}

// Warn registra un mensaje de advertencia.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.log(message, args...)
}

// Error registra un mensaje de error.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, args...)
	}
	l.msg("error", message, args...)
}

// Fatal registra un mensaje de error fatal y termina el programa.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg("fatal", message, args...)
	os.Exit(1)
}

func (l *Logger) log(message string, args ...interface{}) {
	l.logger.Info().Msgf(message, args...)
}

func (l *Logger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg))
	}
}

// ZerologMiddleware improves the format of the log entries by adding more information about the request.
func ZerologMiddleware(logg Interface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Log the request
			logEntry := logg.(*Logger).logger.With().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Str("remote_addr", r.RemoteAddr).
				Str("user_agent", r.UserAgent()).
				Logger()

			// Wrap the response writer
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			// Process the request
			next.ServeHTTP(ww, r)

			// Log the response
			logEntry.Info().
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
