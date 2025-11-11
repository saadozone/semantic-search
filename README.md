# Semantic Search

A Go-based semantic search service built with the Gin web framework.

## Project Structure
 
```
semantic-search/
├── internal/
│   ├── api/                 # API handlers and routes
│   ├── index/               # Search indexing logic
│   ├── ingest/              # Data ingestion services
│   ├── embed/               # Embedding generation
│   ├── store/               # Data storage layer
│   ├── search/              # Search functionality
│   ├── security/            # Security and authentication
│   └── obs/                 # Observability and monitoring
├── configs/                 # Configuration files
├── scripts/                 # Utility scripts
├── data/                    # Data files and samples
├── main.go                  # Main application entry point
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
└── README.md               # This file
```

## Getting Started

### Prerequisites

- Go 1.19 or later
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/saad/semantic-search.git
cd semantic-search
```

2. Install dependencies:
```bash
go mod tidy
```

### Running the Server

Start the server:
```bash
go run main.go
```

The server will start on port 8080.

### Testing the Health Endpoint

Test the health check endpoint:

```bash
curl http://localhost:8080/v1/healthz
```

Expected response:
```json
{
  "status": "ok"
}
```

## API Endpoints

- `GET /v1/healthz` - Health check endpoint

## Development

To build the project:
```bash
go build -o semantic-search main.go
```

To run tests:
```bash
go test ./...
```
