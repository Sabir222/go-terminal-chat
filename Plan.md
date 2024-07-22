# Project Structure for Terminal-based Chat Application

# Overview

This document explains the directory and file structure for a terminal-based chat application built using Go, Gorilla for the web framework, and Bubble Tea for the TUI (text user interface).

## Directory Structure

```
chat-app/
├── cmd/
│   └── chat/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── server/
│   │   ├── routes.go
│   │   ├── server.go
│   │   └── handlers.go
│   ├── tui/
│   │   ├── model.go
│   │   ├── view.go
│   │   ├── update.go
│   │   ├── bubble_tea.go
│   └── chat/
│       ├── message.go
│       ├── client.go
│       └── hub.go
├── pkg/
│   └── utils/
│       └── utils.go
├── go.mod
└── go.sum
```

### `cmd/`

- **Purpose:** Contains the entry point(s) of the application. It separates the main package code from the library code.
- **Files:**
  - `cmd/chat/main.go`: Entry point for the chat application. Initializes the application, loads configurations, sets up the server and TUI, and starts them.

### `internal/`

- **Purpose:** Holds internal application code that should not be imported by other projects, enforcing encapsulation.

#### `config/`

- **Files:**
  - `config.go`: Manages configuration loading and settings, including reading environment variables and configuration files.

#### `server/`

- **Files:**
  - `server.go`: Sets up the HTTP server using Gorilla, configures middlewares, and starts the server.
  - `routes.go`: Defines the HTTP routes and maps them to their respective handlers.
  - `handlers.go`: Contains HTTP handler functions that respond to client requests.

#### `tui/`

- **Files:**
  - `model.go`: Defines the Bubble Tea model, which holds the state of the TUI application.
  - `view.go`: Contains the view logic that renders the TUI based on the current state.
  - `update.go`: Contains the update logic that defines how the state should change based on different messages/events.
  - `bubble_tea.go`: Initializes and runs the Bubble Tea program, integrating the model, view, and update logic.

#### `chat/`

- **Files:**
  - `message.go`: Defines the data structures for chat messages.
  - `client.go`: Manages individual chat clients, their connections, and interactions.
  - `hub.go`: Manages the central hub for chat communication, handling message broadcasting, client connections, and message delivery.

### `pkg/`

- **Purpose:** Contains reusable packages that are not specific to this application but might be useful in other projects as well.
- **Files:**
  - `utils.go`: Contains utility functions and helpers, such as logging, error handling, and formatting functions.

### `go.mod` and `go.sum`

- **Purpose:** Manage the module dependencies.
  - `go.mod`: Defines the module path and its dependencies.
  - `go.sum`: Holds checksums for module versions to ensure integrity and reproducibility.

## Detailed Workflow.

1. **Initialization (`cmd/chat/main.go`):**

   - Load configuration from `internal/config/config.go`.
   - Initialize the chat hub from `internal/chat/hub.go`.
   - Set up the HTTP server from `internal/server/server.go` and define routes in `internal/server/routes.go`.
   - Start the Bubble Tea TUI from `internal/tui/bubble_tea.go`.

2. **Server Setup (`internal/server/`):**

   - `server.go`: Initializes the HTTP server with necessary middlewares (e.g., logging, CORS).
   - `routes.go`: Defines endpoints such as `/ws` for WebSocket connections.
   - `handlers.go`: Includes functions to handle WebSocket connections and HTTP requests.

3. **Chat Logic (`internal/chat/`):**

   - `hub.go`: Manages the chat hub which handles all active connections, broadcasts messages, and handles joining/leaving clients.
   - `client.go`: Manages individual WebSocket connections, handling sending and receiving messages.
   - `message.go`: Defines the structure of chat messages.

4. **Terminal UI (`internal/tui/`):**
   - `model.go`: Defines the state model for the TUI.
   - `view.go`: Handles the rendering logic for displaying chat messages and user input in the terminal.
   - `update.go`: Contains the state update logic, reacting to user input and server messages.
   - `bubble_tea.go`: Ties together the model, view, and update functions to run the Bubble Tea program.
