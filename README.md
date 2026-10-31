# ☕ grpc-mcp-otel-boilerplate (Java 21 + Protobuf + OpenTelemetry)

[![Java](https://img.shields.io/badge/Java-21-ED8B00?logo=openjdk)](https://openjdk.org/)
[![gRPC](https://img.shields.io/badge/gRPC-Protobuf-teal)](https://grpc.io/)
[![OpenTelemetry](https://img.shields.io/badge/OpenTelemetry-Enabled-blueviolet)](https://opentelemetry.io/)
[![MCP](https://img.shields.io/badge/MCP-Agent%20Tools-orange)](https://modelcontextprotocol.io/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

A clean, production-inspired **Java 21** starter kit for exposing **typed gRPC / Protocol Buffer services** as **Model Context Protocol (MCP) tools** for AI Agents—instrumented with **OpenTelemetry (`OTLP`) distributed traces**, **partitioned bulk I/O batching (batches of 50)**, and **fail-closed capacity guardrails**.

---

## 🎯 Why This Architecture Exists

When enterprise Java backends expose **gRPC** microservices to autonomous **AI Agents (via MCP)**, three production reliability bottlenecks arise:

1. **The N+1 RPC Fan-Out Problem:** An AI agent enriching a queue of 200 items often triggers 200 sequential RPCs. Partitioning lookups into parallel batches of 50 and synthesizing states in a single in-memory pass keeps P99 latency overhead under **50ms**.
2. **Symmetric Fail-Closed vs. Per-Item Batch Fault Tolerance:**
   - **Capacity / Duplicate-Creation Guards must Fail Closed:** If a downstream query returns an error or empty command response, the guard throws `IllegalStateException` rather than mistaking an RPC failure for `0` open items.
   - **Batch Resolution Trackers must Isolate Per-Item Failures:** When scanning hundreds of items in a background job, per-item exceptions are caught, logged with context, and counted in `lookupErrorCount` separately from `resolvedCount` so one transient timeout never aborts a 24h batch run.
3. **End-to-End Trace Propagation:** Preserves W3C `traceparent` context from the MCP JSON-RPC tool call into downstream gRPC metadata.

---

## 📂 Project Structure

- `proto/ticket.proto` — Typed Protocol Buffer contract (`TicketService`, `BatchEnrichTickets`, `CheckQueueCapacity`).
- `src/main/java/com/priti/mcp/CaseEnricherBridge.java` — Java 21 implementation of partitioned bulk I/O enrichment, bounded async batch execution, and fail-closed capacity guards.
- `pom.xml` — Maven build configuration with gRPC, Protobuf, Guava, and OpenTelemetry SDK dependencies.

---

## 🚀 Quickstart

```bash
mvn compile exec:java -Dexec.mainClass="com.priti.mcp.CaseEnricherBridge"
```

---

## 📄 License

MIT License — Created by **[Priti Jha](https://priti24.github.io/portfolio/)**.
