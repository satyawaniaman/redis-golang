# Redis Clone in Go - Task Checklist

## Setup
- [ ] Initialize Go project (`go mod init`)  
- [ ] Create folder structure:  
  - `cmd/server/main.go`  
  - `internal/resp/`  
  - `internal/store/`  
  - `internal/handlers/`  

## Core Server
- [ ] Implement TCP server (`net.Listen`, `Accept`, `conn.Read/Write`)  
- [ ] Spawn goroutines for multiple concurrent clients  

## RESP Protocol
- [ ] Implement RESP parser for:
  - Simple Strings (`+`)  
  - Bulk Strings (`$`)  
  - Arrays (`*`)  
- [ ] Implement command dispatcher (`map[string]func(args []string) string`)  

## Core Commands
- [ ] `PING` → return `PONG`  
- [ ] `SET key value` → store value in memory  
- [ ] `GET key` → retrieve value from memory  
- [ ] `DEL key` → delete key from memory  

## TTL / Expiry
- [ ] `EXPIRE key seconds` → set TTL for a key  
- [ ] Background goroutine to auto-delete expired keys  
- [ ] `TTL key` → check remaining time  

## Optional Extended Commands
- [ ] `INCR key` → increment integer value  
- [ ] `DECR key` → decrement integer value  
- [ ] `MSET key1 val1 key2 val2 ...` → set multiple keys  
- [ ] `MGET key1 key2 ...` → get multiple keys  

## Optional Persistence
- [ ] Implement AOF-like logging (append write commands to file)  
- [ ] Replay log on startup  

## Polish & Testing
- [ ] Test all commands via `redis-cli` or `telnet`  
- [ ] Add logging for connections and commands  
- [ ] Clean, modular code (`resp/`, `store/`, `handlers/`)  

## README / Documentation
- [ ] Project description  
- [ ] Features / Supported commands  
- [ ] Example usage with `redis-cli`  
- [ ] How to run (`go run cmd/server/main.go`)  
- [ ] Architecture diagram (optional)  
- [ ] What was learned (goroutines, RESP, TCP, TTL)  

## Final Steps
- [ ] Final tests with multiple clients  
- [ ] Push to GitHub  
- [ ] Optional: record short demo / GIF for portfolio  


