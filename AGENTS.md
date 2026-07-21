# AGENTS.md

## Overview

**hmm** ("Humans Messaging Machines") — a bring-your-own-key AI chat wrapper built with [Wails v3](https://v3.wails.io/) (Go backend + React/TypeScript frontend).

## Project Layout

- **All Go code lives in `app/`**, not the repo root. The Go module is `app`.
- `app/main.go` — entry point; wires up Wails app, window, and services
- `app/*_service.go` — Go services exposed to the frontend via Wails bindings
- `app/db/` — SQLite database layer (pure Go via `modernc.org/sqlite`, no CGO)
- `app/frontend/` — React 19 + TypeScript + Mantine 9 + Vite
- `app/frontend/bindings/` — **generated** Wails TypeScript bindings (do not edit)

## Commands

| Action | Command |
|---|---|
| Dev (hot reload) | `just dev` |
| Build | `just app build` |
| Frontend typecheck | `cd app/frontend && npx tsc --noEmit` |

`just dev` runs `sqlc generate` then `wails3 dev` automatically.

## Code Generation (do not hand-edit generated files)

- **sqlc**: `db/db.go`, `db/models.go`, `db/*.sql.go` are generated from `db/queries/*.sql` and `db/migrations/*.sql` via `sqlc generate`. Config: `app/sqlc.yaml`.
- **Wails bindings**: `frontend/bindings/` are generated automatically during `wails3 dev`. Go service methods become typed TypeScript APIs.

## Database

- SQLite stored at `$XDG_DATA_HOME/hmm/hmm.db` (auto-created on first run).
- Migrations in `app/db/migrations/` are embedded and auto-applied at startup in sorted filename order.
- Queries in `app/db/queries/` use sqlc annotation syntax (`-- name: QueryName :one/:many/:exec/:one`).
- After adding or changing migrations/queries, run `sqlc generate` (or just `just dev`/`just app build`).

## Toolchain

- **mise** manages Go 1.26.5 and Node 26.5.0 (see `mise.toml`).
- **just** is the task runner; root `justfile` delegates to `app/justfile`.
- Frontend dev server runs on port **9245** (`WAILS_VITE_PORT`).
