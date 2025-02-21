# sshell
A straightforward SSH server that pipes the io to a [bubbletea](https://github.com/charmbracelet/bubbletea) program per ssh session.

> [!WARNING]
> This tool should be used within a controlled environment. Although it's designed to be secure, we do not recommend using it outside of a controlled setting.

## How it works
The server listens for incoming connections. To handle these connections, it creates a new pseudo-terminal (PTY), connected to a bubbletea instance and piped back to the client. 

## Usage
Build the tool using go:
```bash
go build -o sshell .
```

Run the server:
```bash
./sshell -port 2222 -key /path/to/key
```

Connect to the server:
```bash
ssh -p 2222 user@localhost
```

## Options
| Option   | Description                              | Default | Required |
| -------- | ---------------------------------------- | ------- | -------- |
| `port`   | The port to listen on.                   | 2222    | No       |
| `key`    | The path to the private key to use.      | None    | No       |

