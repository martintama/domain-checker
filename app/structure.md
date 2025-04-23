app/
├── cmd/
│   ├── root.go          # Contains the root command definition
│   ├── serve.go         # Command to serve the application
│   └── version.go       # Command to display version information
├── internal/            # Private application code
│   ├── config/          # Configuration management
│   ├── handlers/        # Request handlers
│   └── models/          # Data models
├── pkg/                 # Public packages that can be imported by other applications
│   └── utils/           # Utility functions
├── main.go              # Application entry point
├── go.mod               # Go modules file
└── go.sum               # Go modules checksum file