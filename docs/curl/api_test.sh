#!/bin/bash

# Configuration
API_URL="https://books-api.ahmadmu.com/api/v1"
CONTENT_TYPE="Content-Type: application/json"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing Books API endpoints..."
echo "-----------------------------"

# 1. Create a book
echo -e "\n${GREEN}Creating a new book...${NC}"
CREATE_RESPONSE=$(curl -s -X POST \
  -H "${CONTENT_TYPE}" \
  -d '{
    "title": "The Go Programming Language",
    "author": "Alan A. A. Donovan",
    "year": 2015
  }' \
  "${API_URL}/books")
echo $CREATE_RESPONSE | jq '.'

# Extract book ID from response
BOOK_ID=$(echo $CREATE_RESPONSE | jq -r '.id')

# 2. Get the created book
echo -e "\n${GREEN}Getting book with ID ${BOOK_ID}...${NC}"
curl -s -X GET "${API_URL}/books/${BOOK_ID}" | jq '.'

# 3. List books
echo -e "\n${GREEN}Listing all books...${NC}"
curl -s -X GET "${API_URL}/books?page=1&size=10" | jq '.'

# 4. Update the book
echo -e "\n${GREEN}Updating book with ID ${BOOK_ID}...${NC}"
curl -s -X PUT \
  -H "${CONTENT_TYPE}" \
  -d '{
    "title": "The Go Programming Language",
    "author": "Alan A. A. Donovan & Brian W. Kernighan",
    "year": 2015
  }' \
  "${API_URL}/books/${BOOK_ID}" | jq '.'

# 5. Delete the book
echo -e "\n${GREEN}Deleting book with ID ${BOOK_ID}...${NC}"
curl -s -X DELETE "${API_URL}/books/${BOOK_ID}"

echo -e "\n${GREEN}API testing completed!${NC}"