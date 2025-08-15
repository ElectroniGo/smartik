#!/usr/bin/env bash

##################################################################
#
# This scripts sets up Ollama models (qwen2.5:0.5b) for development.
#
##################################################################

set -e

echo "Setting up Ollama models..."

# Wait for Ollama to be ready
echo "Waiting for Ollama server to start..."
sleep 10

# Check if Ollama is running
if ! docker exec ollama ollama list > /dev/null 2>&1; then
  echo "Ollama server is not ready. Check logs: \`docker compose logs ollama\`"
  exit 1
fi

# Pull the qwen2.5:0.5b model
echo "Downloading qwen2.5:0.5b model (this may take a while)..."
docker exec ollama ollama pull qwen2.5:0.5b

# Verify the model is available
echo "Verifying qwen2.5:0.5b installation..."
docker exec ollama ollama list

echo "Ollama models setup complete."
echo "The Python \`text-extraction\` service can now use the qwen2.5:0.5b model"