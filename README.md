# sshell
A straightforward SSH server that delivers a shell script to clients.

> **Warning:** This tool should be used within a controlled environment. Although it's designed to be secure, we do not recommend using it outside of a controlled setting.

## How it works
The server listens for incoming connections. To handle these connections, it creates a new pseudo-terminal (PTY), executes a shell script located on the server, and pipes its output back to the client. The client also sends input to the server, which is then piped into the PTY.

## Usage
Build the tool using go:
```bash
go build -o sshell main.go
```

Run the server:
```bash
./sshell -port 2222 -script /path/to/script.sh -key /path/to/key
```

## Options
- `-port`: The port to listen on. Default is 2222.
- `-script`: The path to the shell script to execute.
- `-key`: The path to the private key to use. By default a new key will be used every time the server is started.