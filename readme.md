# ShopSocial Monorepo

Welcome to **ShopSocial**, a monorepo containing:

1. **Backend (GoLang with Gin)**  
2. **Web Frontend (Next.js)**  
3. **Mobile App (React Native)**  

This `README.md` explains how to **configure** and **run** each part locally.

---

## Table of Contents
- [Project Structure](#project-structure)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Environment Variables](#environment-variables)
- [Local Development](#local-development)
    - [Backend (GoLang)](#backend-golang)
    - [Web (Next.js)](#web-nextjs)
    - [Mobile (React Native)](#mobile-react-native)
- [Common Commands](#common-commands)
- [Contributing](#contributing)
- [License](#license)

---

## Project Structure

```
shopsocial/
├── README.md
├── backend/          # Go (Gin) server
│   ├── cmd/          # Application entry point
│   ├── internal/     # Modules (Users, Products, etc.)
│   ├── pkg/          # Utilities (logging, errors, etc.)
│   ├── config/       # Environment and DB configuration
│   ├── go.mod
│   ├── go.sum
│   └── ...
├── web/              # Next.js web frontend
│   ├── package.json
│   ├── next.config.js
│   └── ...
├── mobile/           # React Native mobile app
│   ├── package.json
│   ├── App.js
│   └── ...
└── ...
```

1. **`backend/`**  
     - Contains **GoLang (Gin)** API code, **MongoDB** configuration, **JWT-based auth**, etc.

2. **`web/`**  
     - **Next.js** project for the **ShopSocial** web interface.

3. **`mobile/`**  
     - **React Native** project for the **ShopSocial** mobile application.

---

## Tech Stack

1. **Go (Gin Framework)**  
     - RESTful API  
     - JWT authentication  
     - MongoDB driver

2. **Next.js**  
     - React-based server-side rendering  
     - Styling (e.g., Tailwind CSS) optional

3. **React Native**  
     - Cross-platform mobile app  
     - Consumes the **ShopSocial** backend

4. **MongoDB**  
     - NoSQL database for storing data  
     - Running locally or via MongoDB Atlas

---

## Prerequisites

- **Go** (1.18+ recommended)
- **Node.js** (16+ recommended)
- **NPM** or **Yarn**  
- **MongoDB** (local or hosted)
- **Android/iOS** development environment (for React Native)

---

## Environment Variables

Create a `.env` file **in each** project folder (`backend/`, `web/`, `mobile/`) for environment-specific variables.

### Example: `backend/.env`

```bash
PORT=8080
MONGO_URI=mongodb://localhost:27017/shopsocial
JWT_SECRET=very-secret-key
ENV=development
LOG_LEVEL=info
```