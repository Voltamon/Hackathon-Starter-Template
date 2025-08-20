# Starter Fullstack Template

This repository contains a minimal full-stack starter app:

- Frontend: Vite + Preact + TypeScript, serves on port 3000
- Backend: Go + Gin, serves API on port 8080
- Pocketbase: runs on port 8090

Getting started (Windows PowerShell):

1. Build and start all services:

```
docker-compose up --build
```

2. Frontend will be available at http://localhost:3000
3. Backend API: http://localhost:8080/api/hello
4. Pocketbase admin: http://localhost:8090

Notes:
- The frontend Dockerfile runs Vite in dev mode. For production, build the frontend and serve static files from a simple web server.
- The backend uses CORS allowing all origins for development. Lock this down in production.

Development tips

- From PowerShell you can run the whole stack with:

```
docker-compose up --build
```

- To only run frontend locally (hot-reload) without Docker:

```
cd frontend; npm install; npm run dev
```

- To build frontend for production and serve statically (example):

```
cd frontend; npm install; npm run build
# then serve the 'dist' folder with any static server
```

Requirements coverage

- Frontend: Vite with Preact and TypeScript -> Done (`frontend/package.json`, `vite.config.ts`, `tsconfig.json`, `src`)
- Frontend index.html linking Pico CSS and app entry -> Done (`frontend/index.html`)
- Frontend App with button fetching backend -> Done (`frontend/src/App.tsx`, `src/api.ts`)
- Backend: Go Gin server with CORS and GET /api/hello -> Done (`backend/main.go`)
- Backend Dockerfile -> Done (`backend/Dockerfile`)
- Pocketbase service -> Done (`docker-compose.yml`)
- Frontend Dockerfile and compose wiring -> Done (`frontend/Dockerfile`, `docker-compose.yml`)

Notes / assumptions

- The frontend Dockerfile runs Vite in dev mode for simplicity. For production, replace with a multi-stage build that builds assets and serves them with a static server.
- The compose file maps ports 3000, 8080, 8090 on the host. If these are occupied, change them in `docker-compose.yml`.
- This is a minimal example intended for local development.
