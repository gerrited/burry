# Burry URL Shortener

![Burry Logo](assets/logo.webp) 

**Burry** is a lightweight URL shortener written in Go. It allows you to shorten URLs, store them in a cache and use the short link to be redirected to the original URL. The service has been tested with Redis and Garnet as caching solutions.

## Features

- **Shorten URLs**: Create shorter versions of long URLs.
- **Cache Storage**: Store shortened URLs in a cache for quick retrieval.
- **Redirection**: Redirect to the original URL using the shortened link.
- **Deployable**: Can be run using Docker Compose or in a Kubernetes cluster.

## Endpoints

### 1. POST `/shorten`

Create a shortened URL.

**Request:**
```json
{
  "long_url": "https://example.com"
}
```

**Response:**
```json
{
  "short_url": "https://<your-domain>/<shortlink>"
}
```

### 2. GET `/<short_url>`

Redirect to the original URL.

## Getting Started

### Prerequisites

- Go 1.15+
- Docker
- Docker Compose
- Kubernetes (optional)

### Environment Variables

- `REDIS_ADDR`: The address of the cache (e.g. redis or garnet) server (e.g., `localhost:6379`).

### Running with Docker Compose

1. Clone the repository:
    ```sh
    git clone https://github.com/gerrited/burry.git
    cd burry
    ```

2. Start the services:
    ```sh
    docker-compose up --build
    ```

### Running in a Kubernetes Cluster

1. Clone the repository:
    ```sh
    git clone https://github.com/gerrited/burry.git
    cd burry
    ```

2. Set up the environment variable in your Kubernetes configuration:
    ```yaml
    env:
      - name: REDIS_ADDR
        value: "redis-service:6379"
    ```

3. Deploy the service:
    ```sh
    kubectl apply -f k8s/
    ```

## Usage

The app runs on port 8080.

### Shorten a URL

```sh
curl -X POST -H "Content-Type: application/json" -d '{"long_url": "https://example.com"}' https://<your-domain>/shorten
```

### Use a short link

Navigate to `https://<your-domain>/<shortlink>` in your browser.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
