#!/bin/bash

# Load the API key from config/config-local.yaml
API_KEY=$(grep 'ApiKey' ./config/config-local.yaml | awk '{print $2}' | tr -d '"')

# Configuration
API_URL="http://localhost:8088"
NUM_STREAMS=4         # Maximum number of streams
NUM_MESSAGES=1000     # Total number of messages to send
MESSAGE="Round Robin Test Message"
HEADER="X-API-Key: $API_KEY"
STREAM_IDS=()
SUCCESS_COUNT=0
FAILURE_COUNT=0
START_TIME=$(date +%s)
CONSUMER_PIDS=()

# Function to create a new stream and store the stream_id
create_stream() {
    response=$(curl -s -X POST "$API_URL/stream/start" -H "$HEADER" -H "Content-Type: application/json")
    stream_id=$(echo "$response" | jq -r '.data.stream_id')
    if [[ -n "$stream_id" && "$stream_id" != "null" ]]; then
        STREAM_IDS+=("$stream_id")
        echo "Created stream with stream_id: $stream_id"
    else
        echo "Failed to create stream: $response"
    fi
}

# Function to produce messages to a specified stream
produce_messages() {
    local stream_id=$1
    for ((i = 0; i < NUM_MESSAGES; i++)); do
        response=$(curl -s -X POST "$API_URL/stream/$stream_id/send" \
            -H "$HEADER" \
            -H "Content-Type: application/json" \
            -d "{\"message\": \"$MESSAGE\"}")

        status=$(echo "$response" | jq -r '.status')

        if [[ "$status" == "success" ]]; then
            ((SUCCESS_COUNT++))
            echo "Message sent successfully to stream_id: $stream_id"
        else
            ((FAILURE_COUNT++))
            echo "Failed to send message to stream_id: $stream_id, response: $response"
        fi
    done
    echo "Finished producing messages to stream_id: $stream_id"
}

# Function to consume messages from a specified stream using WebSocket
consume_messages() {
    local stream_id=$1
    wscat -c "$API_URL/stream/$stream_id/results" -H "$HEADER" | while read -r message; do
        echo "Received message from stream_id $stream_id: $message"
    done &
    CONSUMER_PIDS+=($!)
}

# Function to monitor resource utilization using `docker stats`
monitor_resources() {
    echo "Starting resource monitoring..."
    docker stats --no-stream --format "table {{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.NetIO}}" > docker_stats.txt &
    MONITOR_PID=$!
}

# Cleanup function to terminate all background processes
cleanup() {
    echo "Cleaning up..."
    for pid in "${CONSUMER_PIDS[@]}"; do
        kill "$pid" 2>/dev/null
    done
    kill "$MONITOR_PID" 2>/dev/null
    wait
    echo "Cleanup completed."
}

# Trap SIGINT (Ctrl+C) and SIGTERM to run the cleanup function
trap cleanup SIGINT SIGTERM

# Main script
echo "Starting setup for round-robin message production and consumption with up to $NUM_STREAMS streams"
echo "API Key: $API_KEY"

# Create up to 4 streams
for ((i = 0; i < NUM_STREAMS; i++)); do
    create_stream
done

# Check if we have exactly 4 streams
if [[ ${#STREAM_IDS[@]} -ne 4 ]]; then
    echo "Failed to create 4 streams. Exiting."
    exit 1
fi

# Start resource monitoring
monitor_resources

# Start producing and consuming messages concurrently for each stream
for stream_id in "${STREAM_IDS[@]}"; do
    produce_messages "$stream_id" &
    consume_messages "$stream_id" &
done

# Wait for all background processes to complete
wait

# Calculate performance metrics
END_TIME=$(date +%s)
TOTAL_TIME=$((END_TIME - START_TIME))
AVERAGE_TIME=$(echo "scale=2; $TOTAL_TIME / $NUM_MESSAGES" | bc)

# Output performance results
echo "Round-robin message production and consumption completed for $NUM_STREAMS streams"
echo "Total time taken: $TOTAL_TIME seconds"
echo "Average time per request: $AVERAGE_TIME seconds"

# Display resource utilization summary
echo "Resource utilization during the test:"
cat docker_stats.txt

# Run cleanup function to ensure all connections are closed
cleanup
