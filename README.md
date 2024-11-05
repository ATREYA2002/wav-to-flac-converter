
# WAV to FLAC Converter Service

## Overview

This service allows users to upload WAV files, converts them to FLAC format, and streams the output back over WebSocket. The project leverages Go and the Gin framework for routing and WebSocket handling. A Docker container and Kubernetes configurations are included to simplify deployment and scalability.

## Features

- Real-time WAV to FLAC conversion.
- WebSocket support for continuous streaming.
- Containerized setup for deployment in any environment.
- Kubernetes support for horizontal scalability.

## Getting Started

### Prerequisites
- Go v1.23.2 or later
- FFmpeg (for audio processing)
- Docker (for containerization)
- Kubernetes & Azure CLI (for deployment on AKS)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/wav-to-flac-converter.git
   cd wav-to-flac-converter
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Project Structure

```plaintext
wav-to-flac-converter/
├── main.go            # Entry point of the application
├── go.mod             # Go module dependencies
├── go.sum             # Go dependencies checksums
├── routes/            # API and WebSocket routes
│   ├── routes.go      # HTTP routes for the application
│   └── websocket.go   # WebSocket handlers
├── services/          # Business logic for audio processing
│   └── audio.go       # WAV to FLAC conversion functions
├── Dockerfile         # Dockerfile for containerizing the app
├── deployment.yaml    # Kubernetes deployment configuration
├── service.yaml       # Kubernetes service configuration
└── tests/             # Unit and integration tests
    ├── audio_test.go  # Tests for audio conversion logic
    ├── websocket_test.go # Tests for WebSocket functionality
    ├── BAK.wav        # Sample WAV file for testing
    └── sample.wav     # Additional sample WAV for testing
```

## Dependencies

| Dependency             | Description                                                  | Installation                                    |
|------------------------|--------------------------------------------------------------|-------------------------------------------------|
| **Gin**                | HTTP framework for Go routing and request handling.          | `go get -u github.com/gin-gonic/gin`            |
| **WebSocket**          | Library for WebSocket connections for real-time streaming.   | `go get -u github.com/gorilla/websocket`        |
| **ffmpeg-go**          | Go bindings for FFmpeg, enabling audio file conversion.      | `go get -u github.com/u2takey/ffmpeg-go`        |
| **Testify**            | Assertion library for testing.                               | `go get -u github.com/stretchr/testify`         |

### Other Requirements
Ensure the following software is installed on your system:
- **Go** v1.23.2 (or later)
- **FFmpeg** (install using `sudo apt-get install ffmpeg` on Linux or `brew install ffmpeg` on macOS)

## Running the Application

1. **Run the Server**  
   Start the application locally on port 8080:
   ```bash
   go run main.go
   ```

2. **Access the Converter Interface**  
   Open your browser and go to `http://localhost:8080` to view the interface.

3. **Testing with Sample WAV File**  
   Upload a sample WAV file (e.g., `BAK.wav`) and click the Convert button to test.

## Getting Started

### Step 1: Download the Repository
Clone this GitHub repository to your local machine:

```bash
git clone https://github.com/your-username/wav-to-flac-converter.git
cd wav-to-flac-converter
```

---

## Functionality Breakdown
![image](https://github.com/user-attachments/assets/6b24cec1-4afe-45c9-8284-6966237e9a55)


### `main.go`
- **Purpose:** Entry point for the application.
- **Functionality:**
  - Initializes the web server and sets up routes using a framework like Gin.
  - Defines routes in `routes/routes.go`.
  - Manages startup logic and configuration.

### `go.mod`
- **Purpose:** Manages dependencies for the application.
- **Functionality:**
  - Lists required packages and their versions.
  - Enables dependency management in Go.

### `go.sum`
- **Purpose:** Contains checksums for dependencies in `go.mod`.
- **Functionality:**
  - Ensures dependency integrity by verifying them against checksums.
  - Updated automatically when dependencies change.

---

### `routes/`
- **Directory Purpose:** Manages API endpoints for the application.
- **Files:**
  - `routes.go`: Defines API routes (e.g., `/convert` for file uploads), handles HTTP requests, and provides health checks.
  - `websocket.go`: Manages WebSocket connections, including data transmission for real-time audio streaming.

### `services/`
- **Directory Purpose:** Contains business logic for audio processing.
- **Files:**
  - `audio.go`: Implements WAV to FLAC conversion using libraries like `ffmpeg`, handles file uploads, conversions, and error handling.

### `BAK.wav`
- **Purpose:** Sample WAV file for local testing.

---

### `Dockerfile`
- **Purpose:** Defines steps to build a Docker image.
- **Functionality:**
  - Specifies the base image, copies source files, installs dependencies, builds the application, and sets the entry command.

### `deployment.yaml`
- **Purpose:** Kubernetes deployment configuration.
- **Functionality:**
  - Defines replicas, container image, ports, and restart policies.

### `service.yaml`
- **Purpose:** Configures Kubernetes Service.
- **Functionality:**
  - Defines network exposure (e.g., LoadBalancer, NodePort) and links the service to the deployment.

---

### `tests/`
- `tests/audio_test.go`: Unit tests for audio conversion.
- `tests/BAK.wav`: WAV file for testing.
- `tests/output.flac`: Expected output FLAC file generated from `sample.wav`.
- `tests/sample.wav`: Input WAV file for testing.
- `tests/websocket_test.go`: Tests WebSocket functionality.

---

## Running the Application Locally

### Step 1: Test the Application with `BAK.wav`
Place `BAK.wav` in your project directory to test conversion.

### Step 2: Start the Server
Navigate to your project directory and start the application:

```bash
cd ~/wav-to-flac-converter
go run main.go
```
![image](https://github.com/user-attachments/assets/513b9f8c-5683-4b37-b91e-006d48d5ca59)


The server will run on `http://localhost:8080`.

### Step 3: Run Tests
In a second terminal, navigate to the project directory and run tests:

```bash
cd ~/wav-to-flac-converter
go test -v ./tests
```
![image](https://github.com/user-attachments/assets/6c86c575-795d-430b-879e-5d9f4970d43a)


### Step 4: Access the Converter Interface
Open a browser and navigate to `http://localhost:8080`. Upload a WAV file and test the conversion.

---

## Deployment and Scalability with Kubernetes and Azure AKS

### Step 1: Containerize with Docker
Use Docker to containerize the application. Ensure you have a Docker Hub account (e.g., `docker username: pranidock`).

Build and push the Docker image:

```bash
docker login
docker build -t pranidock/wav-to-flac-converter .
docker push pranidock/wav-to-flac-converter
```
![image](https://github.com/user-attachments/assets/c9a1bb1d-9f55-4095-8e83-596b1d271f0e) 
  build docker image
![image](https://github.com/user-attachments/assets/5b247050-93c9-412a-b93a-09e62d03ffe6)
  push docker image to docker desktop

### Step 2: Kubernetes Deployment
Create Kubernetes configuration files:

1. **Deployment:** `deployment.yaml` for setting up application replicas.
   ![image](https://github.com/user-attachments/assets/9e3d70cb-8638-4071-9aff-6e3fe56be7f1)

2. **Service:** `service.yaml` to expose the WebSocket service.
   ![image](https://github.com/user-attachments/assets/b49f9b12-29d7-448e-a69b-50a4e486b8b0)


### Step 3: Deploy to Azure AKS
Using your university's free cloud credits, set up an AKS cluster.

1. **Create the AKS cluster:**
   ```bash
   az aks create --resource-group PranidhanResource --name PraniAKSCluster --node-count 1 --enable-addons monitoring --generate-ssh-keys
   ```

2. **Connect `kubectl` to AKS:**
   ```bash
   az aks get-credentials --resource-group PranidhanResource --name PraniAKSCluster
   ```

3. **Deploy the application:**
   ```bash
   kubectl apply -f deployment.yaml
   kubectl apply -f service.yaml
   ```

### Step 4: Access the Application
Get the external IP:

```bash
kubectl get services
```
![image](https://github.com/user-attachments/assets/6db294b9-5e25-4e0b-a6bc-412431463c99)


Open your browser and enter the IP address (e.g., `http://4.157.225.49`) to access the service.

---

## Scalable Access and Future Modifications

With the application hosted on a Kubernetes cluster, you can scale it by adjusting the replica count in `deployment.yaml`.



Here’s how you can document your API endpoints and the WebSocket functionality for your WAV-to-FLAC converter service in the README file. This will provide a comprehensive guide for users and developers to understand how to interact with your application.

## API Documentation

### API Endpoints

#### 1. Health Check Endpoint
- **Method:** `GET`
- **Endpoint:** `/`
- **Description:** Returns a welcome message to confirm the service is running.
- **Usage:**
  ```bash
  curl http://<External-IP>/
  ```
- **Response:**
  ```json
  {
    "message": "Welcome to the WAV to FLAC conversion service!"
  }
  ```

#### 2. File Conversion Endpoint
- **Method:** `POST`
- **Endpoint:** `/convert`
- **Description:** Accepts a WAV file upload, converts it to FLAC, and returns a downloadable link for the converted file.
- **Request Body:**
  - Form-data containing the WAV file.
- **Usage:**
  ```bash
  curl -X POST http://<External-IP>/convert -F 'file=@/path/to/your/file.wav'
  ```
- **Response:**
  Upon successful conversion, the response will include a link to download the converted FLAC file.
  ```json
  {
    "message": "File converted successfully!",
    "download_link": "http://<External-IP>/path/to/converted/file.flac"
  }
  ```

### WebSocket Endpoint

#### Test the WebSocket Endpoint
- **Method:** `GET`
- **Endpoint:** `/ws`
- **Description:** Establishes a WebSocket connection for streaming audio.
- **Usage:**
  You can test the WebSocket connection using a tool like `websocat` or any WebSocket client software. Here’s an example using `websocat`:
  ```bash
  websocat ws://<External-IP>/ws
  ```



