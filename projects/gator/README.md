# Gator - RSS Feed Aggregator CLI

A command-line RSS feed aggregator built with Go and PostgreSQL.

## Prerequisites

- Go 1.23+
- PostgreSQL

## Installation

```bash
go install github.com/fc1g/go-learning/tree/main/projects/gator/cmd/gator
```

## Setup

1. Create a PostgreSQL database:

```bash
createdb gator
```

2. Create config file at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

3. Run migrations:

```bash
goose -dir sql/schema postgres "your_db_url" up
```

## Usage

```bash
# User commands
gator register <username>
gator login <username>

# Feed management
gator addfeed <name> <url>
gator feeds
gator follow <url>
gator following

# Aggregation (runs continuously)
gator agg <duration>     # e.g., gator agg 1m

# Browse posts
gator browse [limit]     # e.g., gator browse 10
```

## Example

```bash
gator register alice
gator addfeed "Hacker News" https://news.ycombinator.com/rss
gator follow https://news.ycombinator.com/rss
gator agg 1m            # In separate terminal
gator browse 20
```
