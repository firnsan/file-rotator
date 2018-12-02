# About file-rotator
This is a golang library helps to write content into file and automatically rotate the file. 
It implements `io.Writer` and `io.Closer`, so you can combine it with golang  standard `log` library, or others log library such as `logrus`.

These are it's features:
* Rotate according to file size
* Rotate according to file lines
* Rotate daily
* Remove files that are n days ago 


# Usage

This is a example shows how to create a `file rotator` and combine with the `logrus` and the golang standard `log` library.
```
import (
    log "github.com/Sirupsen/logrus"
    rotator "github.com/firnsan/file-rotator"
    "io"
    stdlog "log"
    "os"
)

func InitLog() error {
    var err error
    level, err := log.ParseLevel(gApp.Cnf.LogLevel)
    if err != nil {
        log.Errorf("Parse log level failed: %s", err)
        return err 
    }   
    // Create a file rotator
    fw, err := rotator.NewFileRotator(gApp.Cnf.LogDir + "/app.log")
    if err != nil {
        log.Errorf("Set log failed: %s", err)
        return err 
    }   

    log.SetOutput(fw)
    log.SetLevel(level)

    // Also need to set the golang standard log's output to this writer                                                                                   
    w := log.StandardLogger().Writer()
    stdlog.SetOutput(w)

    return nil 
}

func UninitLog() {
    // Close Writer
    w := log.StandardLogger().Out
    log.SetOutput(os.Stderr)
    if wc, ok := w.(io.Closer); ok {
        wc.Close()
    }
    log.Printf("Uninit log success")
}


```
