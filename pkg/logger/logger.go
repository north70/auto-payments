package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var defaultWriters []io.Writer

const defaultTimeFormat = time.RFC3339Nano

// Capacity of writers depends on writers slice, and once on stdout write (3+1=4)
const writersCapacity = 2

const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Cyan = "\033[36m"
const DarkGray = "\033[91m"
const colorReset = "\033[0m"

type FilteredWriter struct {
	writer zerolog.LevelWriter
	levels []zerolog.Level
}

func (w *FilteredWriter) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *FilteredWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	for _, filteredLevel := range w.levels {
		if level == filteredLevel {
			return w.writer.WriteLevel(level, p)
		}
	}
	return len(p), nil
}

func InitLogger() zerolog.Logger {
	defaultWriters = make([]io.Writer, 0, writersCapacity)

	outWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: defaultTimeFormat}
	outWriter = configureOutputMessage(outWriter)

	errWriter := zerolog.MultiLevelWriter(os.Stderr)
	filteredLevels := []zerolog.Level{zerolog.ErrorLevel, zerolog.PanicLevel, zerolog.FatalLevel}
	errWriter = &FilteredWriter{errWriter, filteredLevels}
	//
	defaultWriters = append(defaultWriters, errWriter)
	defaultWriters = append(defaultWriters, outWriter)

	writer := zerolog.MultiLevelWriter(defaultWriters...)
	return zerolog.New(writer).With().Str("application", "hr-bot").Caller().Timestamp().Logger()
}

func configureOutputMessage(writer zerolog.ConsoleWriter) zerolog.ConsoleWriter {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	writer.FormatLevel = configureFormatLevel
	writer.FormatMessage = configureFormatMessage
	writer.FormatErrFieldName = func(i interface{}) string {
		return fmt.Sprintf(Red+"%s=", i)
	}
	// formatting caller string
	zerolog.CallerMarshalFunc = configureCaller

	return writer
}

func configureCaller(pc uintptr, file string, line int) string {
	// lengthRequired means how many symbols can contain caller info
	const lengthRequired = 20
	simpleFileName := filepath.Base(file)
	simpleLinePointer := strconv.Itoa(line)

	callerLength := len(fmt.Sprintf("%s", simpleFileName+":"+simpleLinePointer))
	if callerLength <= lengthRequired {
		callerLength = lengthRequired + len(":"+simpleLinePointer)
		return fmt.Sprintf("%*s", callerLength, simpleFileName+":"+simpleLinePointer)
	} else {
		filePath := fmt.Sprintf("%s", simpleFileName)
		filePath = filePath[:lengthRequired-len(":"+simpleLinePointer)] + "..."
		return fmt.Sprintf(filePath) + ":" + simpleLinePointer
	}
}

func configureFormatLevel(i interface{}) string {
	inputMessage, ok := i.(string)
	if !ok {
		fmt.Printf("WARNING! Cant convert interface to string")
	}
	level, _ := zerolog.ParseLevel(inputMessage)
	switch level {
	case zerolog.DebugLevel:
		return fmt.Sprintf(Cyan+" %-5s ➙"+colorReset, i)
	case zerolog.InfoLevel:
		return fmt.Sprintf(Green+" %-5s ➙"+colorReset, i)
	case zerolog.WarnLevel:
		return fmt.Sprintf(Yellow+" %-5s ➙"+colorReset, i)
	case zerolog.ErrorLevel, zerolog.FatalLevel, zerolog.PanicLevel:
		return fmt.Sprintf(Red+" %-5s ➙"+colorReset, i)
	case zerolog.NoLevel:
		return fmt.Sprintf(DarkGray+" %-5s ➙"+colorReset, i)
	default:
		return fmt.Sprintf(" %-5s ➙", i)
	}
}

func configureFormatMessage(i interface{}) string {
	// lengthRequired means how many symbols would be in message between code pointer and params
	// this was done for greater readability
	lengthRequired := 80
	messageLength := len(fmt.Sprintf("%s", i))
	if messageLength <= lengthRequired {
		return fmt.Sprintf("%-80s", i)
	} else {
		message := fmt.Sprintf("%s", i)
		// lengthRequired-3 it is 80 symbols string with ellipsis (80-3=77  len"..."=3)
		message = message[:lengthRequired-3] + "..."
		return fmt.Sprintf(message)
	}
}
