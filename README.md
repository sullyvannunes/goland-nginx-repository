# Nginx Reverse Proxy with Go Backend

This repository is for **learning purposes only**. It demonstrates how to set up an Nginx reverse proxy that forwards requests to a simple Go backend.

## 🚀 Getting Started

### Prerequisites

- Docker
- Docker Compose

### Running the Application

1. Clone this repository:
   ```sh
   git clone <your-repository-url>
   cd <repository-folder>
   ```
2. Start the containers:
   ```sh
   docker-compose up --build
   ```

### Accessing the API

Once the containers are running, you can test the setup using **cURL**:

- **POST request to login endpoint**

  ```sh
  curl -X POST http://localhost/login
  ```

  **Expected Response:**

  ```
  Executado com sucesso
  ```

## 🛠 How It Works

- **Nginx** listens on port `80` and forwards requests to the Go backend running on port `8080`.
- Requests to `/login` are proxied to `http://go_app:8080/login`.
- The Go backend handles requests and returns appropriate responses.

## 📌 Notes

- This project is intended for **learning purposes only**.
- Ensure that no other services are using port `80` or `8080` on your machine.
- If needed, modify `docker-compose.yml` and `nginx.conf` to fit your setup.

## 🏗 Stopping the Application

To stop and remove the containers:

```sh
docker-compose down
```

Happy coding! 🚀
