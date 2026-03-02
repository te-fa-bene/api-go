# Te Fa Bene

**_MAMMA MIA!!!_**

## Project overview

Te Fa Bene is a restaurant management API built in Go, designed as a portfolio-ready MVP with production-minded fundamentals. It provides a solid backend foundation for core restaurant operations such as staff access control and table workflow, and it’s structured to scale into ordering, kitchen, and cashier modules.

The project uses Gin for HTTP routing, GORM for persistence, Goose for database migrations, and PostgreSQL as the primary datastore. It also includes Swagger/OpenAPI documentation for easy endpoint exploration, JWT (HS256) authentication, and RBAC authorization to enforce role-based permissions (e.g., waiter, kitchen, cashier, manager), ensuring each user can only access what they’re allowed to.

This repository focuses on clean project structure, predictable local development (Dockerized Postgres + hot reload), and an API that is easy to run, understand, and extend.

## Requirements

### Developer

- Docker
- Goose
- Swaggo
- Air
