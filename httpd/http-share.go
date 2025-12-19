package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
    "syscall"
)

func main() {
    ip := flag.String("i", "0.0.0.0", "IP address to bind to")
    port := flag.Int("p", 80, "Port to listen on")
    path := flag.String("d", "", "Directory to serve (default current directory)")
    daemon := flag.Bool("daemon", false, "Run in background")
    logFile := flag.String("log", "", "Path to log file (optional)")
    flag.Parse()

    dir := *path
    if dir == "" {
        var err error
        dir, err = os.Getwd()
        if err != nil {
            log.Fatal(err)
        }
    }

    if !*daemon {
        args := append(os.Args, "-daemon")
        cmd := exec.Command(args[0], args[1:]...)
        cmd.Stdout = nil
        cmd.Stderr = nil
        cmd.Stdin = nil
        cmd.SysProcAttr = &syscall.SysProcAttr{
            Setsid: true,
        }
        if err := cmd.Start(); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Server started in background with PID %d\n", cmd.Process.Pid)
        os.Exit(0)
    }

    if *logFile != "" {
        f, err := os.OpenFile(*logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
        if err != nil {
            log.Fatalf("failed to open log file: %v", err)
        }
        log.SetOutput(f)
    }

    fmt.Printf("Serving directory %s on %s:%d\n", dir, *ip, *port)
    fs := http.FileServer(http.Dir(dir))
    http.Handle("/", fs)

    bind := fmt.Sprintf("%s:%d", *ip, *port)
    log.Fatal(http.ListenAndServe(bind, nil))
}
