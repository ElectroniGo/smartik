#!/usr/bin/env bash

# This script initializes the .env file for all the services
# It should be run before starting the services in development

# Loops through all .env.example files in the specified directories
# and copies them to .env
find . -type f -name ".env.example" ! -path "*/node_modules/*" ! -path "*/.turbo/*" | while read -r file; do
    dir=$(dirname "$file")
    
    cp "$file" "$dir/.env"
    
    echo "Initialized .env in $dir"
done