package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// TicketStatus represents a simplified gRPC response payload.
type TicketStatus struct {
	TicketID    string
	Status      string
	ProductArea string
}

// BatchResult tracks resolved tickets separately from transient RPC lookup errors
// so batch processors never miscount transient failures as closed tickets.
type BatchResult struct {
	ResolvedCount    int
	OpenCount        int
	LookupErrorCount int
}

// ExecuteBoundedMCPBatch demonstrates bounded async concurrency + OpenTelemetry-style
// span lifecycle for an AI Agent MCP tool invoking downstream gRPC lookups.
func ExecuteBoundedMCPBatch(ctx context.Context, ticketIDs []string, maxConcurrency int) BatchResult {
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var result BatchResult

	for _, id := range ticketIDs {
		wg.Add(1)
		sem <- struct{}{} // Acquire bounded concurrency slot

		go func(ticketID string) {
			defer wg.Done()
			defer func() { <-sem }()

			status, err := fetchTicketFailClosed(ctx, ticketID)
			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				// Per-item fault tolerance: log and increment LookupErrorCount
				// without aborting the entire batch or miscounting as RESOLVED.
				log.Printf("[WARN] [otel.span=grpc.GetTicket] lookup failed for %s: %v", ticketID, err)
				result.LookupErrorCount++
				return
			}

			if status.Status == "RESOLVED" {
				result.ResolvedCount++
			} else {
				result.OpenCount++
			}
		}(id)
	}

	wg.Wait()
	return result
}

// fetchTicketFailClosed simulates a fail-closed gRPC backend lookup.
func fetchTicketFailClosed(ctx context.Context, ticketID string) (*TicketStatus, error) {
	time.Sleep(15 * time.Millisecond)
	if ticketID == "TICKET-ERR-503" {
		return nil, errors.New("downstream gRPC status UNAVAILABLE: failing closed")
	}
	if ticketID == "TICKET-101" || ticketID == "TICKET-102" {
		return &TicketStatus{TicketID: ticketID, Status: "RESOLVED", ProductArea: "cloud-infra"}, nil
	}
	return &TicketStatus{TicketID: ticketID, Status: "OPEN", ProductArea: "browser-runtime"}, nil
}

func main() {
	ctx := context.Background()
	tickets := []string{"TICKET-101", "TICKET-102", "TICKET-103", "TICKET-ERR-503", "TICKET-104"}

	fmt.Println("🔭 Starting gRPC-MCP-OTel Bridge Demo (Bounded Concurrency = 2)...")
	res := ExecuteBoundedMCPBatch(ctx, tickets, 2)

	fmt.Printf("✅ Batch Complete | Resolved: %d | Open: %d | Lookup Errors (Isolated): %d\n",
		res.ResolvedCount, res.OpenCount, res.LookupErrorCount)
}
