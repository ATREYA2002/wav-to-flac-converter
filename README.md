1.To get started with your audio streaming service, you first need to download the GitHub repository I created.
2. Functionality Breakdown
main.go
•	Purpose: The main application entry point.
•	Functionality:
o	Initializes the web server and router (e.g., using the Gin framework).
o	Sets up the routes defined in routes/routes.go.
o	Handles the application’s startup logic and configuration.
go.mod
•	Purpose: The go.mod file defines the module and its dependencies.
•	Functionality:
o	Lists all required packages and their versions needed for the application to compile and run.
o	Facilitates dependency management in Go projects.
go.sum
•	Purpose: This file contains checksums for the dependencies listed in go.mod.
•	Functionality:
o	Ensures the integrity of the dependencies by verifying that they haven't been tampered with.
o	Automatically updated when dependencies are added or removed.
routes/
•	Directory Purpose: Contains route handlers that define the API endpoints for the application.
•	routes.go:
o	Functionality:
	Sets up the API routes for the application (e.g., /convert for file uploads).
	Defines handlers for incoming HTTP requests and routes them to the appropriate services.
	Provides basic responses for health checks and service status.
•	websocket.go:
o	Functionality:
	Manages WebSocket connections for real-time audio streaming (if implemented).
	Defines handlers to accept connections and manage data transmission over WebSocket.
services/
•	Directory Purpose: Contains the business logic for the application, specifically related to audio processing.
•	audio.go:
o	Functionality:
	Implements the core logic for converting WAV files to FLAC using a library (e.g., ffmpeg).
	Provides functions to handle file uploads, conversions, and saving the output.
	Includes error handling and validation of input files.
•	BAK.WAV : wav file which we tested locally on our system and convert it to FLAC 

Dockerfile
•	Purpose: The Dockerfile is used to build a Docker image for the application.
•	Functionality:
o	Defines the base image and environment for running the application (e.g., Golang).
o	Copies application source files into the container.
o	Installs necessary dependencies and builds the application.
o	Specifies the command to run the application when the container starts.
deployment.yaml
•	Purpose: This file contains the Kubernetes deployment configuration for the application. It defines how to deploy the application to a Kubernetes cluster.
•	Functionality:
o	Specifies the number of replicas to run.
o	Defines the container image to use, along with its ports and resource limits.
o	Configures the application to automatically restart if it fails.
service.yaml
•	Purpose: This file defines the Kubernetes service configuration for the application.
•	Functionality:
o	Specifies how to expose the application to the network (e.g., LoadBalancer, NodePort).
o	Maps the service to the deployment, allowing access to the application from external sources.
Tests
•	tests/audio_test.go: Contains unit tests for the audio conversion logic. It validates the functionality of the methods in services/audio.go to ensure they work as expected.
•	tests/BAK.wav: A WAV audio file used for testing purposes.
•	tests/output.flac: The expected output FLAC file generated from converting a specific WAV file, in this case, sample.wav.
•	tests/sample.wav: A sample WAV audio file used for testing the conversion functionality. It serves as an input for the conversion tests.
•	tests/websocket_test.go: Contains tests for the WebSocket functionality, ensuring that the WebSocket connections and data streaming work correctly.
3) Now that you have the WAV file BAK.wav, we can use it to test our Go application
4) We Make sure your websocket_test.go file sets up a WebSocket connection to the server correctly. This means ensuring the server is running and listening on the correct port (in this case, :8080).
5) Navigate to Your Project Directory
cd ~/wav-to-flac-converter
Run your Go application 
This command will start your web server, which typically runs on http://localhost:8080 (or another port if configured differently).
go run main.go
Run the Tests by opening a second terminal window and navigate to the project directory:
cd ~/wav-to-flac-converter
go test -v ./tests

7) Open your web browser and navigate to http://localhost:8080 to view your converter interface.

8) Test the Application
Once you have the page open in your browser, test your WAV to FLAC converter by selecting a WAV file and clicking the convert button. Monitor the status messages to ensure everything is working as intended.
To make it scalable 
1)First, create a Dockerfile to containerize your Go application. This enables easy deployment across different environments and facilitates scaling.
 
2)Login to docker in our local system and Build Docker image and push it to Docker Desktop 
 
Here docker username:pranidock

 
 
3)Set Up a Kubernetes Deployment
With the application containerized, create Kubernetes configuration files to deploy the service on a cluster. Kubernetes Deployment Configuration called deployment.yaml:
 

Set Up a Load Balancer Service
To expose the WebSocket service outside the Kubernetes cluster, create a Service resource.
Kubernetes Service Configuration called service.yaml
 


4) Deploy To Kubernetes
 


5)Setup Azure CLI and and create azure AKS with my student account
az aks create --resource-group PranidhanResource --name PraniAKSCluster --node-count 1 --enable-addons monitoring --generate-ssh-keys
Once the cluster is created, you need to configure your kubectl command-line tool to connect to your new AKS cluster:
az aks get-credentials --resource-group PranidhanResource --name PraniAKSCluster

6) Deploy Your Application
Now you can proceed with deploying your WAV to FLAC converter application using the deployment.yaml and service.yaml files as described earlier. Make sure to create these files in your working directory and then apply them using the kubectl command.
kubectl apply -f deployment.yaml // Apply Deployment:
kubectl apply -f service.yaml // Apply Service:



7)Access your Application
After deploying, run the following command to get the external IP of your service
kubectl get services 
You can access your application by following these methods:
A. Using a Web Browser
•	Open your web browser (Chrome, Firefox, etc.).
•	Enter the external IP address in the address bar:
Here : http://4.157.225.49


With this you can now acces the service from anywhere any system as its hosted on Kubernetes cluster we can do changes and make replicas to scale it too 











