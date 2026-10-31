# 🔭 grpc-mcp-otel-boilerplate

[![OpenTelemetry](https://img.shields.io/badge/OpenTelemetry-Enabled-blueviolet)](https://opentelemetry.io/)
[![gRPC](https://img.shields.io/badge/gRPC-Protobuf-teal)](https://grpc.io/)
[![MCP](https://img.shields.io/badge/MCP-Agent%20Tools-orange)](https://modelcontextprotocol.io/)
[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

A clean, production-inspired starter kit for exposing **typed gRPC / Protocol Buffer services** as **Model Context Protocol (MCP) tools** for AI Agents—instrumented end-to-end with **OpenTelemetry (`OTLP`) distributed traces**, **bounded batch concurrency**, and **fail-closed capacity guardrails**.

---

## 🎯 Why This Boilerplate Exists

As engineering teams connect autonomous AI Agents (via **MCP**) to existing **gRPC microservices**, three production challenges immediately appear:

1. **Zero Trace Visibility Across Agent Tool Calls:** Standard MCP JSON-RPC tool invocations often lose W3C `traceparent` context before reaching downstream gRPC services.
2. **Unbounded Agent Fan-Out (Database DDoS):** An autonomous agent iterating over 500 records can fire 500 parallel gRPC lookups unless bounded concurrency and fail-closed capacity checks are enforced at the bridge layer.
3. **Schema Drift:** Manually writing JSON schemas for MCP tools duplicates existing `.proto` contracts.

This boilerplate solves all three in **under 250 lines of Go**.

---

## ✨ Key Features

- **Automatic MCP-to-gRPC Tool Bridging:** Exposes `TicketService` RPCs (`GetTicket`, `CheckQueueCapacity`) as typed MCP tools.
- **End-to-End OpenTelemetry Tracing:** Propagates W3C Trace Context from the MCP tool invocation into downstream gRPC metadata so every agent action appears as a unified trace waterfall in **Jaeger** or **OpenSearch Dashboards**.
- **Bounded Concurrency & Fail-Closed Guardrails:** Prevents runaway AI agent loops from overloading backend storage using a bounded worker semaphore (`maxConcurrency`) and fail-closed error handling.

---

## 🚀 Quickstart

### 1. Start Jaeger (or OpenSearch Data Prepper) for OTLP Traces

```bash
docker run -d --name jaeger \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  jaegertracing/all-in-one:latest
```

### 2. Run the Instrumented Bridge & Demo Batch Run

```bash
go run main.go
```

---

## 🛡️ Fail-Closed vs. Per-Item Batch Fault Tolerance Pattern

```go
// Capacity Guard: Fail-Closed on downstream RPC/storage errors
func checkCapacityFailClosed(resp *TicketResponse) error {
    if resp.HasStorageError || resp.CommandCount == 0 {
        // Never treat an empty/errored RPC response as "0 active items"
        return fmt.Errorf("capacity check failed closed for ticket %s", resp.TicketID)
    }
    return nil
}
```

---

## 📄 License

MIT License — Created by **[Priti Jha](https://priti24.github.io/portfolio/)**.
