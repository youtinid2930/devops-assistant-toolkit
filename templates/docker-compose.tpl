version: "3.9"

services:
  {{PROJECT_NAME}}:
    build: .
    ports:
      - "8000:8000"
    volumes:
      - .:/app
      