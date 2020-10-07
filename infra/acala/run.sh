#!/bin/bash
ADDRESS=$1

# Start
cd /app
# make run
# cargo run -- --dev --execution=native -lruntime=debug --ws-external --rpc-external
SKIP_WASM_BUILD= cargo run -- --dev --execution=native -lruntime=debug --ws-external --rpc-external
sleep 10

# Print setup
echo "ACALA_ADDRESS=$ADDRESS"
