# ☕ grpc-mcp-otel-boilerplate (Java 21 + gRPC + OpenTelemetry)

[![Java](https://img.shields.io/badge/Java-21-ED8B00?logo=openjdk)](https://openjdk.org/)
[![gRPC](https://img.shields.io/badge/gRPC-Protobuf-teal)](https://grpc.io/)
[![OpenTelemetry](https://img.shields.io/badge/OpenTelemetry-Enabled-blueviolet)](https://opentelemetry.io/)
[![MCP](https://img.shields.io/badge/MCP-Agent%20Tools-orange)](https://modelcontextprotocol.io/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

A minimal, beginner-friendly **Java 21** boilerplate showing how to expose a basic **gRPC Arithmetic Calculator service** (`Add`, `Subtract`, `Multiply`, `Divide`) as **Model Context Protocol (MCP) tools** for AI Agents, with **OpenTelemetry (`OTLP`)** trace propagation to Jaeger or OpenSearch.

---

## 📂 Project Structure

- `proto/calculator.proto` — Standard Protocol Buffer definition for `CalculatorService`.
- `src/main/java/com/priti/calculator/CalculatorMcpServer.java` — Minimal Java 21 arithmetic tool runner with OpenTelemetry span logging.
- `pom.xml` — Standard Maven configuration.

---

## 🚀 Quickstart

```bash
mvn compile exec:java -Dexec.mainClass="com.priti.calculator.CalculatorMcpServer"
```

---

## 📄 License

MIT License — Created by **[Priti Jha](https://priti24.github.io/portfolio/)**.
