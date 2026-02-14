# go-cryptotool

## Prerequisites

### 1. Install [go](https://go.dev)

### 2. Build the project

```bash
go build .
```

### 3. That's it

## Usage

### 1. Encryption

```powershell
.\go-cryptotool.exe railfence encrypt --message "Canabis" --key 3
```

```powershell
.\go-cryptotool.exe railfence encrypt --input ./examples/SunPoem.txt --output ./examples/SunPoem.enc --key 5
```

### 2. Decryption

```powershell
.\go-cryptotool.exe railfence decrypt --message "inCasba" --key 3
```

```powershell
.\go-cryptotool.exe railfence decrypt --input ./examples/SunPoem.enc --output ./examples/SunPoem.txt --key 5
```
