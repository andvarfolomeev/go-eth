# go-eth

🚀 **go-eth** — high-performance Ethereum blockchain indexer in Go with batched RPC support.

---

## Overview

This project asynchronously fetches, deserializes, and saves Ethereum blockchain data. It uses worker pools and RPC batching to maximize throughput and efficiency.

---

## Features

- Batched Ethereum RPC calls
- Parallel block and transaction processing with worker pools
- Graceful shutdown with goroutine lifecycle management
- Clean architecture separating stages: fetching → deserialization → saving
- Safe concurrency with channels and sync.WaitGroup

---

## Quick Start

### Install

```bash
git clone git@github.com:andvarfolomeev/go-eth.git
cd go-eth
make db # launch db and migrations
make # build and run
