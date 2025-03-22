# Commita-AI

Commita-AI is a CLI tool designed to streamline your Git commit process by leveraging AI to summarize changes and generate meaningful commit messages. It supports conventional commit types and allows for custom messages or AI-generated summaries.

## Features
- Automatically summarize staged changes using AI.
- Generate commit messages following conventional commit standards.
- Configure your OpenAI API key for AI-powered features.

## Installation
Clone the repository and build the binary:
```bash
go install github.com/guths/commita-ai
```

## Configuration
Before using Commita-AI, you need to configure your OpenAI API key.

1. Run the configuration command:
   ```bash
   commita-ai config
   ```

2. This will create a configuration file at:
   ```
   ~/.config/commitaai/config.yaml
   ```

3. Open the file and replace `YOUR-API-KEY` with your actual OpenAI API key:
   ```yaml
   open_api_key: YOUR-API-KEY
   ```

## Usage
Commita-AI provides commands to help you commit changes efficiently.

### Commit Command
To commit changes, use the `c` command:
```bash
./commita-ai c -t <type> [-m <custom-message>]
```

#### Options:
- `-t, --type`: Specify the commit type (e.g., `feat`, `fix`, `chore`, `docs`, `test`).
- `-m, --message`: Provide a custom commit message (optional) (do not use ia).

#### Example:
1. Commit with AI-generated summary:
   ```bash
   ./commita-ai c -t feat
   ```

2. Commit with a custom message:
   ```bash
   ./commita-ai c -t fix -m "Fixed a critical bug in the payment module"
   ```

### Diagram of Workflow
Here’s a simple representation of how Commita-AI works:

```
| Staged Changes |
        ↓
|   Git Diff      |
        ↓
| AI Summarization |
        ↓
| Commit Message   |
        ↓
|    Git Commit    |
```

## Supported Commit Types
- `feat`: New features
- `fix`: Bug fixes
- `chore`: Maintenance tasks
- `docs`: Documentation updates
- `test`: Adding or updating tests

## Contributing
Feel free to open issues or submit pull requests to improve Commita-AI.

## License
This project is licensed under the MIT License.