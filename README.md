<div align="center">
  <img src="doc/logo.png" alt="PipeAI Logo" width="200">
  <h1>PipeAI</h1>
  <p>
    <img src="https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat-square&logo=go" alt="Go Version">
    <img src="https://img.shields.io/badge/OS-Linux%20%7C%20macOS%20%7C%20Windows-darkblue?style=flat-square&logo=windows" alt="OS Support">
    <img src="https://img.shields.io/badge/AI-Gemini%20%7C%20OpenAI%20%7C%20Ollama-purple?style=flat-square&logo=openai" alt="AI Providers">
    <img src="https://img.shields.io/badge/License-MIT-green?style=flat-square" alt="License">
  </p>
</div>

Natural language → shell commands. Cloud or local AI.

## Install

### From Release
```sh
# Linux/macOS (AMD64)
curl -L https://github.com/brasa-ai/pipeai/releases/download/stable/pipeai-linux-amd64.tar.gz -o pipeai-linux-amd64.tar.gz
tar -xvf pipeai-linux-amd64.tar.gz pipeai-linux-amd64
chmod +x pipeai-linux-amd64
sudo mv pipeai-linux-amd64 /usr/local/bin/pipeai

# Linux/macOS (ARM64)
curl -L https://github.com/brasa-ai/pipeai/releases/download/stable/pipeai-linux-arm64.tar.gz -o pipeai-linux-arm64.tar.gz
tar -xvf pipeai-linux-arm64.tar.gz pipeai-linux-arm64
chmod +x pipeai-linux-arm64
sudo mv pipeai-linux-arm64 /usr/local/bin/pipeai

# Windows (PowerShell)
Invoke-WebRequest -Uri https://github.com/brasa-ai/pipeai/releases/download/stable/pipeai-windows-amd64.zip -OutFile pipeai-windows-amd64.zip
Expand-Archive -Path pipeai-windows-amd64.zip -DestinationPath .
Move-Item -Path pipeai-windows-amd64/pipeai-windows-amd64.exe -Destination pipeai.exe
```

### From Source
```sh
# Clone repository
git clone https://github.com/brasa-ai/pipeai.git
cd pipeai

# Build binary
go build -o pipeai .

# Install globally (Linux/macOS)
sudo mv pipeai /usr/local/bin/

# Or use go install
go install github.com/brasa-ai/pipeai@latest
```

## Usage
```sh
# Generate only
pipeai --ask "list PDF files"

# Generate + execute
pipeai --act "list PDF files"

# Configure
pipeai setup
```

## Configuration

PipeAI supports multiple AI providers through an interactive setup process:

```sh
# Run interactive setup
pipeai setup

# Example session:
LLM provider (gemini/openai/ollama) []: gemini
API key (skip for ollama) []: YOUR_GEMINI_KEY
Model []: gemini-pro
```

The configuration is stored in `~/.pipeai/config.yaml`:

```yaml
# Example config for different providers:

# Gemini
llm: gemini
model: gemini-pro
key: YOUR_GEMINI_KEY

# OpenAI
llm: openai
model: gpt-4o
key: YOUR_OPENAI_KEY

# Ollama (local)
llm: ollama
model: llama2
key: ""  # not needed
```

You can also override config values via flags:
```sh
pipeai --llm ollama --model codellama --ask "find large files"
```

### Cross-Platform Commands
```sh
# Same command, different OS:
pipeai --ask "show memory usage"
# → Windows: wmic OS get FreePhysicalMemory,TotalVisibleMemorySize /Value
# → Linux: free -h
# → macOS: vm_stat && system_profiler SPHardwareDataType | grep Memory
```

## Examples

### Kubernetes
```sh
# Pod management
pipeai --act "get all pods not running in namespace prod"
# → kubectl get pods -n prod --field-selector status.phase!=Running
```

### Docker & Containers
```sh
# Cleanup
pipeai --act "remove all stopped containers and unused images"
# → docker rm $(docker ps -aq) && docker image prune -af
```

### System Operations
```sh
# File operations
pipeai --act "find large log files modified in last 24h"
# → find /var/log -type f -mtime -1 -size +100M -exec ls -lh {} \;
```

### Pipeline Processing
```sh
# File processing
echo "*.pdf" | pipeai --act "find all matching files recursively"
# → find . -type f -name "*.pdf"

# Process filtering
ps aux | pipeai --act "show processes using more than 1GB RAM"
# → awk 'NR>1 && $6>1000000 {print}'

# JSON/YAML handling
kubectl get pods -o json | pipeai --act "extract container images"
# → jq -r '.items[].spec.containers[].image' | sort | uniq
```

### Git Operations
```sh
# History analysis
pipeai --act "show commits touching kubernetes files in last week"
# → git log --since="1 week ago" --pretty=format:"%h %s" -- "**/k8s/*.yaml"

pipeai --act "find who changed this line in git history"
# → git log -L '/pattern/',+1:file.txt

# Branch management
pipeai --act "list branches not merged to main"
# → git branch --no-merged main | grep -v '^*'
```

## Providers
- **Gemini**: `pipeai --llm gemini --model gemini-pro --key YOUR_KEY`
- **OpenAI**: `pipeai --llm openai --model gpt-4o --key YOUR_KEY`
- **Ollama**: `pipeai --llm ollama --model llama2` (local, [install](https://ollama.ai))

## Debug
```sh
pipeai --act "list files" --debug
