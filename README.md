# PipeAI 🛠️🤖

<div align="center">
  <img src="https://img.shields.io/badge/Go-1.23+-blue.svg" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  <img src="https://img.shields.io/badge/Status-Active-brightgreen.svg" alt="Status">
</div>

PipeAI is a powerful CLI tool that transforms natural language into executable shell commands using AI models. Whether you're a developer looking to automate tasks or a system administrator needing quick command generation, PipeAI simplifies the process by leveraging **Gemini**, **OpenAI**, or **Ollama** to understand your intent and generate the appropriate shell commands.

## Features

- **Natural Language Processing**: Convert plain English descriptions into executable shell commands
- **Multiple AI Providers**: Support for Gemini, OpenAI, and Ollama with easy switching
- **Interactive Setup**: Simple configuration wizard for first-time setup
- **Dry Run Mode**: Preview generated commands before execution
- **Flexible Configuration**: Override settings via command-line flags or config file
- **Streaming Execution**: Direct execution with proper stdin/stdout/stderr handling
- **Cross-Platform**: Works on Linux, macOS, and Windows (WSL)

## Installation

To install PipeAI, ensure you have Go 1.23+ installed and run the following command:

```sh
go install github.com/your-handle/pipeai@latest
```

Alternatively, clone the repository and build from source:

```sh
git clone https://github.com/your-handle/pipeai.git
cd pipeai
go mod tidy
go build -o pipeai .
```

## Quick Start

1. **First-time setup** - Configure your AI provider:
   ```sh
   pipeai setup
   ```

2. **Generate and execute commands**:
   ```sh
   # Basic usage
   pipeai --ask "find all files larger than 100MB in the current directory"
   
   # With pipe input
   cat data.json | $(pipeai --ask "jq to select .[].payload.id != null and starts with 'u_'")
   
   # Preview command without execution
   pipeai --ask "compress all log files" --evaluate
   ```

## Configuration

PipeAI stores its configuration in `~/.pipeai/config.yaml`. The setup command will guide you through the configuration process.

### Configuration File Format

```yaml
llm: gemini          # gemini | openai | ollama
model: gemini-pro    # Model name for the selected provider
key: your-api-key    # API key (not needed for ollama)
```

### Supported Providers

<details>
<summary><strong>Google Gemini</strong></summary>

- **Provider**: `gemini` or `googleai`
- **Models**: `gemini-pro`, `gemini-pro-vision`, etc.
- **API Key**: Required from [Google AI Studio](https://makersuite.google.com/app/apikey)

```yaml
llm: gemini
model: gemini-pro
key: AIzaSyD-************************
```

</details>

<details>
<summary><strong>OpenAI</strong></summary>

- **Provider**: `openai`
- **Models**: `gpt-4o`, `gpt-4o-mini`, `gpt-3.5-turbo`, etc.
- **API Key**: Required from [OpenAI Platform](https://platform.openai.com/api-keys)

```yaml
llm: openai
model: gpt-4o
key: sk-************************
```

</details>

<details>
<summary><strong>Ollama (Local)</strong></summary>

- **Provider**: `ollama`
- **Models**: Any model available in your local Ollama installation
- **API Key**: Not required (runs locally)

```yaml
llm: ollama
model: llama3
key: ""  # Not needed for local Ollama
```

**Prerequisites**: Install and run [Ollama](https://ollama.ai/) locally on port 11434.

</details>

## Usage

PipeAI provides a simple command-line interface with the following options:

### Command Structure

```sh
pipeai [global-flags] --ask "your natural language request"
```

### Global Flags

- `--ask`: Natural language request (required)
- `--evaluate`: Print command instead of executing (dry run mode)
- `--llm`: Override LLM provider for this command
- `--model`: Override model for this command
- `--key`: Override API key for this command

### Examples

<details>
<summary><strong>File Operations</strong></summary>

```sh
# Find and list files
pipeai --ask "find all PDF files in the documents folder"

# File compression
pipeai --ask "compress all log files in /var/log to gzip format"

# File cleanup
pipeai --ask "remove all temporary files older than 7 days"

# File search
pipeai --ask "search for files containing 'password' in the current directory"
```

</details>

<details>
<summary><strong>Data Processing</strong></summary>

```sh
# JSON processing with jq
cat data.json | $(pipeai --ask "jq to select .[].payload.id != null and starts with 'u_'")

# CSV processing
pipeai --ask "convert CSV file to JSON format"

# Text processing
pipeai --ask "count lines, words, and characters in all .txt files"
```

</details>

<details>
<summary><strong>System Administration</strong></summary>

```sh
# Process management
pipeai --ask "find all processes using more than 1GB of memory"

# Network operations
pipeai --ask "check which ports are open on localhost"

# System monitoring
pipeai --ask "show disk usage for all mounted filesystems"

# Package management
pipeai --ask "update all installed packages on Ubuntu"
```

</details>

<details>
<summary><strong>Git Operations</strong></summary>

```sh
# Repository management
pipeai --ask "clone a repository and switch to a new branch"

# Git history
pipeai --ask "show git log for the last 10 commits with file changes"

# Git cleanup
pipeai --ask "remove all local branches that have been merged to main"
```

</details>

### Advanced Usage

<details>
<summary><strong>Dry Run Mode</strong></summary>

Preview commands before execution to ensure they match your expectations:

```sh
pipeai --ask "delete all files in the temp directory" --evaluate
```

This will output the generated command without executing it, allowing you to review and modify if needed.

</details>

<details>
<summary><strong>Provider Override</strong></summary>

Temporarily use a different provider or model for a specific command:

```sh
# Use OpenAI for this specific command
pipeai --ask "analyze this code" --llm openai --model gpt-4o

# Use local Ollama for privacy-sensitive operations
pipeai --ask "process sensitive data" --llm ollama --model llama3
```

</details>

<details>
<summary><strong>Pipeline Integration</strong></summary>

PipeAI works seamlessly with Unix pipelines:

```sh
# Process output from another command
ls -la | $(pipeai --ask "filter to show only directories")

# Chain multiple operations
find . -name "*.log" | $(pipeai --ask "compress each file with gzip")
```

</details>

## Error Handling

PipeAI provides comprehensive error handling:

- **Configuration Errors**: Clear messages when API keys or models are invalid
- **Network Issues**: Graceful handling of API connectivity problems
- **Command Generation**: Validation of generated commands before execution
- **Execution Errors**: Proper error propagation from shell command execution

## Security Considerations

- **API Keys**: Stored securely in `~/.pipeai/config.yaml` with restricted permissions (600)
- **Command Validation**: Review generated commands before execution, especially for destructive operations
- **Local Execution**: Commands run in your local shell environment with your user permissions
- **No Command History**: Generated commands are not stored or logged by default

## Acknowledgments

- Built with [LangChainGo](https://github.com/tmc/langchaingo) for AI provider integration
- Uses [urfave/cli](https://github.com/urfave/cli) for command-line interface
- Powered by [zerolog](https://github.com/rs/zerolog) for structured logging

---

**Note**: Always review generated commands before execution, especially for operations that modify files or system settings. PipeAI is a tool to assist with command generation, but you remain responsible for the commands executed on your system.

```bash
# first-time
pipeai setup

# later
cat data.json | $(pipeai --ask "jq to select .[].payload.id != null and starts with 'u_'")
```

`~/.pipeai/config.yaml`:

```yaml
llm: gemini
model: gemini-pro
key: AIzaSyD-************************
```

