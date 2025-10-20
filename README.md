[![Star this repo](https://img.shields.io/badge/⭐_Star-This_repo-lightgrey?style=flat)](https://github.com/satyawaniaman/redis-golang)
[![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)](https://golang.org/)
[![Version](https://img.shields.io/badge/Version-1.0.0-blue?style=flat)](https://github.com/satyawaniaman/redis-golang/releases)

# Redis like cache implementation in Go v1.0.0

A high-performance Redis-compatible server written in Go, featuring both **asynchronous** (epoll-based) and **synchronous** implementations. This project demonstrates advanced Go networking concepts including non-blocking I/O, event-driven architecture, and RESP protocol parsing.

## 🚀 Features

- **Dual Server Implementations**:
  - **Async Server**: Linux epoll-based event loop for maximum performance
  - **Sync Server**: Traditional goroutine-per-connection model for cross-platform compatibility
- **RESP Protocol**: Full Redis Serialization Protocol support
- **Core Redis Commands**: `PING`, `SET`, `GET`, `DEL`, `TTL`
- **Key Expiration**: Automatic cleanup of expired keys with background goroutines
- **Multi-key Operations**: `DEL` command supports deleting multiple keys at once
- **Docker Support**: Containerized deployment for Linux epoll functionality
- **Cross-Platform**: Sync server runs natively on macOS, Windows, and Linux

## 🏗️ Architecture

### Asynchronous Server (Linux/Docker)
- **Event-driven**: Uses Linux epoll for non-blocking I/O
- **Single-threaded**: Handles thousands of connections efficiently
- **High performance**: Similar to Redis, Nginx, and Node.js architecture
- **Low resource usage**: Minimal memory and CPU overhead

### Synchronous Server (Cross-platform)
- **Goroutine-based**: One goroutine per client connection
- **Cross-platform**: Works on macOS, Windows, and Linux
- **Simple architecture**: Easy to understand and debug
- **Good performance**: Suitable for moderate loads

## 🚀 Quick Start

### Option 1: Docker (Recommended for Async Server)

1. **Clone and run with Docker**:
   ```bash
   git clone https://github.com/satyawaniaman/redis-golang.git
   cd redis-golang
   docker-compose up --build
   ```

2. **Test the server**:
   ```bash
   redis-cli -h localhost -p 6379
   > PING
   PONG
   > SET mykey "Hello Redis!"
   OK
   > GET mykey
   "Hello Redis!"
   ```

### Option 2: Native (Sync Server on macOS/Windows)

1. **Clone the repository**:
   ```bash
   git clone https://github.com/satyawaniaman/redis-golang.git
   cd redis-golang
   ```

2. **Switch to sync server** (edit `main.go`):
   ```go
   // Change this line in main.go:
   server.RunSyncTCPServer()  // instead of RunAsyncTCPServer()
   ```

3. **Run the server**:
   ```bash
   go run main.go
   ```

4. **Test with redis-cli**:
   ```bash
   redis-cli -h localhost -p 6379
   ```

## 📋 Supported Commands (v1.0.0)

| Command | Description | Example |
|---------|-------------|---------|
| `PING` | Test server connectivity | `PING` → `PONG` |
| `SET key value [EX seconds]` | Store a key-value pair with optional expiration | `SET name "John" EX 60` → `OK` |
| `GET key` | Retrieve value by key | `GET name` → `"John"` |
| `DEL key [key ...]` | Delete one or more keys | `DEL name age` → `(integer) 2` |
| `TTL key` | Check remaining TTL | `TTL name` → `(integer) 45` |

### ✨ Key Features
- **Automatic Expiration**: Keys with TTL are automatically cleaned up by background goroutines
- **Multi-key Operations**: `DEL` command can delete multiple keys in a single operation
- **SET with Expiration**: Use `SET key value EX seconds` to set a key with immediate expiration

## 🧪 Testing

### Using redis-cli (Recommended)
```bash
# Connect to the server
redis-cli -h localhost -p 6379

# Test basic commands
> PING
PONG
> SET user:1 "Alice"
OK
> SET user:2 "Bob" EX 10
OK
> GET user:1
"Alice"
> GET user:2
"Bob"
> TTL user:1
(integer) -1
> TTL user:2
(integer) 7
> DEL user:1 user:2
(integer) 2
```

### Using telnet (Manual RESP)
```bash
telnet localhost 6379
*1\r\n$4\r\nPING\r\n
+PONG\r\n
```

## 🐳 Docker Details

The project includes Docker configuration for running the high-performance async server:

- **Dockerfile**: Multi-stage build with Go 1.21 Alpine
- **docker-compose.yml**: Service configuration with port mapping
- **Automatic builds**: No external dependencies required

```bash
# View logs
docker-compose logs -f

# Stop the server
docker-compose down

# Rebuild after code changes
docker-compose up --build
```

## 🗂️ Project Structure

```
redis_golang/
├── main.go              # Entry point (switches between async/sync)
├── server/
│   ├── async_tcp.go     # Linux epoll-based server
│   └── sync_tcp.go      # Cross-platform goroutine server
├── core/
│   ├── resp.go          # RESP protocol parser
│   ├── cmd.go           # Command structures
│   ├── eval.go          # Command evaluation logic
│   └── comm.go          # Communication wrapper
├── test_client/
│   └── main.go          # Go client for testing
├── Dockerfile           # Container configuration
├── docker-compose.yml   # Docker service setup
└── go.mod              # Go module definition
```

## 🔧 Troubleshooting

### macOS/Windows: "undefined: syscall.EpollEvent"
This is expected! The async server uses Linux-specific epoll. Solutions:
1. **Use Docker** (recommended): `docker-compose up --build`
2. **Switch to sync server**: Change `main.go` to call `server.RunSyncTCPServer()`

### Connection Issues
```bash
# Check if server is running
docker ps

# Check port accessibility
telnet localhost 6379

# View server logs
docker-compose logs
```

### Performance Comparison
- **Async Server**: Handles 10,000+ concurrent connections efficiently
- **Sync Server**: Good for moderate loads (hundreds of connections)

## 🗺️ Roadmap

### v1.1.0 - Testing & Quality (Next Release)
- [ ] **Comprehensive Test Suite**
  - Unit tests for all core functions
  - Integration tests for client-server communication
  - Benchmark tests vs real Redis
  - Load testing with concurrent clients
- [ ] **Adding Performance Benchmarks**
### v1.2.0 - Data Persistence
- [ ] **RDB Snapshots**: Save/load data to/from disk
- [ ] **AOF (Append-Only File)**: Log all write operations
- [ ] **Configurable persistence strategies**
- [ ] **Data recovery mechanisms**

### v1.3.0 - Extended Commands
- [ ] **Numeric Operations**: `INCR`, `DECR`, `INCRBY`, `DECRBY`
- [ ] **List Operations**: `LPUSH`, `RPUSH`, `LPOP`, `RPOP`, `LLEN`
- [ ] **Hash Operations**: `HSET`, `HGET`, `HDEL`, `HKEYS`, `HVALS`
- [ ] **Set Operations**: `SADD`, `SREM`, `SMEMBERS`, `SISMEMBER`

## 🎯 Learning Outcomes

This project demonstrates:
- **Network Programming**: TCP servers, non-blocking I/O
- **Go Concurrency**: Goroutines, channels, event loops
- **Protocol Implementation**: RESP parsing and generation
- **System Programming**: Linux epoll, file descriptors
- **Containerization**: Docker, multi-stage builds
- **Cross-platform Development**: Platform-specific code handling  


## License

MIT License - feel free to use for commercial projects!

## Star This Repo

If this helped you build something cool, give it a star!✨ 

---

**Built with ❤️ by [Aman Satyawani](https://x.com/satyawani_aman) to learn Redis like implementation in Go**


