package main

import (
    "context"
    "log"
    "time"
    pb "48HW/proto" 

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()

    conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
    if err != nil {
        log.Fatalf("Serverga ulanishda xato: %v", err)
    }
    defer conn.Close()

    c := pb.NewStringServiceClient(conn)

    response, err := c.ToUpperCase(ctx, &pb.StringRequest{Input: "salom"})
    if err != nil {
        log.Fatalf("Operatsiyani bajarishda xato: %v", err)
    }

    log.Printf("Katta harflar bilan: %s", response.Result)
}
