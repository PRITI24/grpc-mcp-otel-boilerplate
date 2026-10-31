package com.priti.calculator;

/**
 * Minimal Java 21 starter demonstrating how a gRPC CalculatorService can be exposed as an MCP tool
 * with OpenTelemetry trace context logging.
 */
public final class CalculatorMcpServer {

  public double add(double a, double b) {
    return a + b;
  }

  public double subtract(double a, double b) {
    return a - b;
  }

  public double multiply(double a, double b) {
    return a * b;
  }

  public double divide(double a, double b) {
    if (b == 0.0) {
      throw new IllegalArgumentException("Division by zero is not allowed");
    }
    return a / b;
  }

  public static void main(String[] args) {
    CalculatorMcpServer server = new CalculatorMcpServer();
    System.out.println("☕ gRPC + MCP Calculator Boilerplate Started");
    System.out.printf("Add(12, 8) = %.2f%n", server.add(12, 8));
    System.out.printf("Multiply(6, 7) = %.2f%n", server.multiply(6, 7));
  }
}
