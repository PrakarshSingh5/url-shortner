# URL Shortener

A clean and modern URL shortener with a Next.js frontend and a standalone Go backend backed by SQLite. The UI supports light/dark themes, rich animations, and a polished experience out of the box.

## Features

- 🔗 Shorten long URLs into memorable short links
- 🎨 Beautiful dark/light UI with Framer Motion animations
- 📋 Manage and copy your generated short links
- 🪄 Automatic slug generation with validation and normalization
- 🗄️ SQLite storage with zero external dependencies

## Architecture Overview

```
url-shortner/
├── app/                  # Next.js 14 App Router frontend
├── components/           # Reusable UI components
├── backend/              # Go HTTP API using Gin & SQLite
│   ├── internal/
│   │   ├── config/       # Environment configuration loader
│   │   ├── database/     # SQLite connection & migrations
│   │   ├── handlers/     # HTTP handlers (shorten, list, redirect)
│   │   ├── middleware/   # Logging & CORS
│   │   ├── models/       # Domain models
│   │   └── repository/   # Data persistence helpers
│   └── main.go           # Go server entry point
├── lib/                  # Legacy Node database utilities (unused after Go migration)
├── package.json
└── tsconfig.json
```

## Prerequisites

- **Node.js** v18.17 or newer (required by Next.js 14)
- **Go** v1.22 or newer
- **npm** (or an alternative package manager)

## Environment Configuration

| File             | Purpose                                   |
| ---------------- | ----------------------------------------- |
| `.env.local`     | Frontend configuration (public variables) |
| `.env` (backend) | Backend server configuration              |

Suggested defaults:

```
# .env.local
NEXT_PUBLIC_API_URL=http://localhost:8080

# backend/.env
PORT=8080
DB_PATH=./urls.db
BASE_URL=http://localhost:8080
CORS_ORIGINS=http://localhost:3000
```

If environment files are absent, sensible fallbacks are used: the frontend targets `http://localhost:8080`, and the backend listens on port `8080` with CORS open to `http://localhost:3000`.

## Installation & Development

1. **Install frontend dependencies**

   ```bash
   npm install
   ```

2. **Install backend dependencies** (requires Go)

   ```bash
   cd backend
   go mod tidy
   cd ..
   ```

3. **Run the Go API server**

   ```bash
   cd backend
   go run ./...
   ```

4. **Run the Next.js frontend** (in a separate terminal)

   ```bash
   npm run dev
   ```

5. Open [http://localhost:3000](http://localhost:3000) to access the UI. Short URLs resolve via the Go server (default base: `http://localhost:8080`).

## API (Go Backend)

All endpoints are served from the Go API base (`http://localhost:8080` by default):

| Method | Path           | Description                  |
| ------ | -------------- | ---------------------------- |
| POST   | `/api/shorten` | Create a new shortened URL   |
| GET    | `/api/urls`    | Retrieve all shortened URLs  |
| GET    | `/:slug`       | Redirect to the original URL |

Sample request:

```bash
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com"}'
```

## Frontend Development Notes

- The frontend reads `NEXT_PUBLIC_API_URL` to determine the API base. If unset, it defaults to `http://localhost:8080`.
- All API calls originate from the browser; ensure the Go server has CORS enabled for the frontend domain in production.
- The legacy Next.js API routes were removed during the migration; all requests now flow through the Go service.

## Deployment Considerations

- Deploy the Go API (e.g., Fly.io, Render, Docker) and expose it over HTTPS.
- Deploy the Next.js frontend (e.g., Vercel, Netlify) and configure `NEXT_PUBLIC_API_URL` to point to the public Go endpoint.
- Optionally serve the Next.js static build via the Go server or place both behind a shared reverse proxy (Nginx/Traefik) for a unified domain.

## Testing Checklist

- POST `/api/shorten` with valid/invalid URLs (expect validation errors for malformed input).
- GET `/api/urls` after creating entries (expect list to include new records).
- Access `http://localhost:8080/{slug}` to confirm redirects.
- Verify CORS headers if accessing the API from a different origin.

## Next Steps

- Add authentication and user-specific link management
- Introduce analytics (click counts, geo, device)
- Export/import functionality for bulk link management

Enjoy building with Go and Next.js! 🎉
