# Go Concurrency Mini Projects

A collection of classic concurrency problems implemented in Go, created while learning concurrent programming patterns.

> **Course**: [Working with Concurrency in Go (Golang)](https://www.udemy.com/course/working-with-concurrency-in-go-golang/)

## Projects

### 🍝 Dining Philosophers Problem
The classic synchronization problem where five philosophers sit at a table and must share forks to eat. This implementation demonstrates deadlock prevention using resource ordering.

**Key concepts:**
- Mutex locks for shared resources (forks)
- Deadlock prevention by ordering resource acquisition
- WaitGroups for goroutine synchronization
- Buffered channels for communication

**How it works:** Each philosopher needs two forks to eat. To avoid deadlock, philosophers pick up the lower-numbered fork first, except one who reverses the order to break the circular wait.

```bash
cd "Dining Philosophers"
go run .
```

### 🍕 Producer & Consumer (Pizzeria)
A pizzeria simulation where pizzas are produced and consumed concurrently, demonstrating the producer-consumer pattern.

**Key concepts:**
- Producer-consumer pattern
- Channel-based communication
- Graceful shutdown with quit channels
- Random success/failure simulation

**How it works:** The pizzeria (producer) makes pizzas and sends them through a channel to customers (consumer). Some orders fail randomly due to various issues (ingredients, cook quitting, etc.).

```bash
cd "Producer&Consumer"
go run .
```

### 💈 Sleeping Barber Problem
A barbershop simulation with multiple barbers and clients arriving at random intervals. Demonstrates handling of limited resources and concurrent workers.

**Key concepts:**
- Multiple worker goroutines (barbers)
- Buffered channels as waiting rooms
- Handling capacity limits
- Timed operations and graceful shutdown

**How it works:** Barbers sleep when there are no clients. When clients arrive, they wake barbers or wait in the waiting room (if there's space). The shop closes after a set time.

```bash
cd "Sleeping barber"
go run .
```

### 🎯 Final Project - Subscription Service
A full-featured web application demonstrating real-world concurrent programming with a subscription-based service. This is the course's capstone project combining all learned concepts.

**Key concepts:**
- Concurrent background workers (mail service, error logging)
- Channel-based communication for background tasks
- Graceful shutdown with signal handling
- WaitGroups for goroutine coordination
- Redis-backed sessions
- PostgreSQL database with connection pooling
- Email queue with concurrent processing

**Stack:**
- Go web server with concurrent request handling
- PostgreSQL for data persistence
- Redis for session management
- MailHog for email testing
- Docker Compose for service orchestration

**How it works:** Users can register, log in, and subscribe to plans. The app uses background goroutines to handle email sending asynchronously, with proper error handling and graceful shutdown. Multiple services run concurrently (web server, mail worker, error logger) and communicate via channels.

```bash
cd "Final Project"
docker-compose up -d  # Start services
make run              # Build and run the app
```

## Requirements

- Go 1.16 or higher
- Dependencies are managed with Go modules

## Notes

These are learning projects, not production code. They're meant to demonstrate concurrency concepts in Go through classic computer science problems.
