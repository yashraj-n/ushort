# The UShort URL Shortener

A high-performance, modern URL shortener service built with Go, featuring rate limiting, Redis caching, and a clean API design.

![Go version](https://img.shields.io/badge/Go-1.23.4-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## ✨ Features

* 100% **free** and **open source**
* **High-performance** URL shortening service built with Go
* **Redis-backed** caching for optimal performance
* **Rate limiting** support to prevent abuse
* **Clean API design** for easy integration
* **Cross-platform** compatibility
* **Production-ready** with proper error handling
* **Easy to deploy** with minimal configuration

## ⚡️ Quick start

First, make sure you have **Go** installed. Version `1.23.4` (or higher) is required.

1. Clone the repository:
```bash
git clone https://github.com/yashraj-n/ushort.git
cd ushort
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Run the server:
```bash
go run cmd/main.go
```

That's it! 🔥 Your URL shortener service is now running.

## 🛠️ Technology Stack

* **Backend**: Go 1.23.4
* **Router**: Chi
* **Caching**: Redis
* **Rate Limiting**: httprate with Redis support
* **Environment**: godotenv for configuration

## 📦 Project Structure

```
ushort/
├── api/         # API handlers and routes
├── cmd/         # Main application entrypoints
├── repository/  # Data access layer
├── services/    # Business logic
└── tmp/         # Temporary files (Generated by Air)
```

## 🔧 Configuration

The service can be configured using environment variables:

* `PORT` - Server port (default: 8080)
* `REDIS_URL` - Redis connection URL

## 🚀 API Endpoints

* `GET /` - Serves the main web interface
* `GET /static/*` - Serves static files (CSS, JavaScript, images)
* `POST /submit` - Create a shortened URL
* `GET /{hash}` - Redirect to original URL

## 📖 API Documentation

### Shorten URL
```http
POST /submit
Content-Type: application/json

{
    "link": "https://example.com/very-long-url"
}
```

Response:
```json
"abc123"  // The hash of the shortened URL
```

To use the shortened URL, simply append the hash to the base URL:
```
http://localhost:8080/abc123
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
Made with ❤️ by [yashraj-n](https://github.com/yashraj-n) 