#!/bin/bash

# Load Testing Script for Loan API
# Usage: ./scripts/load_test.sh [concurrent_users] [duration]

CONCURRENT_USERS=${1:-100}
DURATION=${2:-30s}
URL="http://localhost:8080/api/v1/loans"

echo "🚀 Starting Load Test"
echo "📊 Concurrent Users: $CONCURRENT_USERS"
echo "⏱️  Duration: $DURATION"
echo "🎯 Target URL: $URL"
echo "================================"

# Install hey if not exists
if ! command -v hey &> /dev/null; then
    echo "Installing hey load testing tool..."
    go install github.com/rakyll/hey@latest
fi

# Create test payload
cat > /tmp/loan_payload.json << EOF
{
  "fullName": "Load Test User",
  "monthlyIncome": 5000,
  "loanAmount": 10000,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "loadtest@example.com"
}
EOF

# Run load test
echo "🔥 Running load test..."
hey -n 10000 -c $CONCURRENT_USERS -z $DURATION -m POST -H "Content-Type: application/json" -D /tmp/loan_payload.json $URL

# Cleanup
rm -f /tmp/loan_payload.json

echo "✅ Load test completed!"