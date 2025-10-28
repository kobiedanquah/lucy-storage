# Fullstack File Storage Platform TODO List

**Tech Stack:**

- **Backend:** Go (Gin) + PostgreSQL + Goose for migrations
- **Frontend:** Solid.js + TypeScript + TailwindCSS
- **Storage:** Local FS / Google Cloud Storage
- **Auth:** JWT / Session-based

---

## Project Setup & Foundation

- [x] Initialize Go module and folder structure
- [x] Setup HTTP server
- [x] Setup configuration management
- [x] Integrate PostgreSQL connection
- [x] Setup migration tool (Goose)
- [x] Setup logging and request middleware
- [x] Setup Docker and docker compose
- [x] Setup graceful shutdown
- [x] Initialize Solid.js + TypeScript project
- [x] Configure TailwindCSS & UI components
- [x] Setup routing (`@solidjs/router`)
- [x] Create API client (fetch wrapper)

---

## Authentication & User Management

- [ ] **Backend**
  - [x] User model & migrations
  - [x] Registration (email + password)
  - [x] Login (JWT-based)
  - [x] Password hashing (bcrypt)
  - [x] Email verification flow
  - [ ] Forgot/reset password endpoints
  - [ ] Middleware for JWT authentication

- [ ] **Frontend**
  - [x] Sign up page
  - [x] Login page
  - [x] Email verification page (PIN/Code input)
  - [ ] Forgot password page
  - [ ] Profile settings page
  - [x] Persistent auth (JWT localStorage)
  - [ ] Logout functionality

---

## File Management Core

- [ ] **Backend**
  - [ ] File metadata model (`files` table)
  - [ ] Folder model & hierarchy support
  - [ ] File upload API (multipart/form-data)
  - [ ] File download API (with access checks)
  - [ ] File delete endpoint
  - [ ] Rename/move file endpoint
  - [ ] Versioning support (optional)
  - [ ] Quota & storage limit per user

- [ ] **Frontend**
  - [ ] Dashboard / Home page (list files & folders)
  - [ ] File upload UI (drag & drop + progress bar)
  - [ ] File preview (images, PDFs, text files)
  - [ ] Download button
  - [ ] File actions (rename, delete, move)
  - [ ] Create folder modal
  - [ ] Folder navigation (breadcrumb UI)
  - [ ] Context menu (right-click actions)

---

## Sharing & Collaboration

- [ ] **Backend**
  - [ ] Shared links model (tokenized access)
  - [ ] Public file access route (`/share/:token`)
  - [ ] Access permissions (read/write)
  - [ ] User-to-user sharing (via email/username)
  - [ ] Collaboration endpoints (add/remove users)

- [ ] **Frontend**
  - [ ] Share modal (generate link, copy to clipboard)
  - [ ] Permission selection (view/edit)
  - [ ] Invite user by email
  - [ ] Manage shared users list
  - [ ] View shared-with-me page

---

## Search, Sorting & Filtering

- [ ] **Backend**
  - [ ] Search endpoint (by file name, type, owner)
  - [ ] Sorting and pagination support
  - [ ] Indexing for performance

- [ ] **Frontend**
  - [ ] Search bar with debounce
  - [ ] Sort dropdown (by name/date/size)
  - [ ] Filters (file type, shared, recent)
  - [ ] Display search results dynamically

---

## Advanced Features (Enhancements)

- [ ] File previews (video/audio support)
- [ ] File version history
- [ ] Trash / Recycle bin system
  - [ ] Soft delete files
  - [ ] Restore from trash
  - [ ] Auto-clean after X days

- [ ] Activity logs (uploads, deletions, shares)
- [ ] Notifications (email + in-app)
- [ ] Tagging / favorites system
- [ ] File comments / notes (collaboration)

---

## Infrastructure & DevOps

- [ ] Dockerize backend and frontend
- [ ] Add reverse proxy (NGINX / Caddy)
- [ ] Setup MinIO / AWS S3 for file storage
- [ ] Setup PostgreSQL in Docker
- [ ] Environment-specific config (dev/staging/prod)
- [ ] CI/CD (GitHub Actions)
- [ ] Logging & monitoring (Prometheus / Grafana)
- [ ] Error tracking (Sentry)
- [ ] API rate limiting & CORS

---

## Polishing & Production

- [ ] Responsive UI polish (mobile/tablet support)
- [ ] Lazy loading & pagination optimization
- [ ] Caching layer (Redis)
- [ ] SEO-friendly share pages
- [ ] Accessibility improvements
- [ ] Unit & integration tests (Go + Vitest)
- [ ] Security audits (JWT expiry, CORS, uploads)
- [ ] Documentation (README, API docs, architecture overview)

---

## Optional Add-ons

- [ ] Real-time collaboration (WebSocket/WebRTC)
- [ ] File diff viewer for text/code
- [ ] AI search / file classification
- [ ] End-to-end encryption
