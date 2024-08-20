# Go Func Yourself Web3 Workload Deployer

## Project Overview

This repository shows the Go-based project I developed while a member of a Crypto startup in 2021. The project was designed to streamline the process of deploying code, containers, and repositories to Web3 infrastructure. My "Go Func Yourself" application that did everything but place the workload on a Web3 worker. GFY is a pretty elegant system that, had the company lasted longer, really proved value to the crypto community as a "Vercel-like" deployment experience to ease Web3 development.

## Key Features

- Multi-format submission support:
  - Raw code
  - Git repositories
  - Single file downloads
  - Archive (zip) downloads
- Language versatility:
  - Python
  - JavaScript
  - Go
- Automated containerization process
- Custom stdin reader injection for standardized input handling
- Dynamic Dockerfile modification

## Technical Highlights

### Submission Processing

The system's core functionality revolves around an API endpoint (`/apiv1/init`) that accepts various types of code submissions. Based on the submission type and programming language, it employs different strategies:

- Cloning Git repositories
- Downloading and extracting archives
- Writing raw code to files

### Container Preparation

After processing the submission, the system:

1. Modifies the Dockerfile to include a custom stdin reader
2. Adjusts the container's entry point
3. Builds the container, ready for deployment to Web3 infrastructure

### API Design

The API accepts POST requests with a JSON body containing:

```json
{
  "submissionType": "git",
  "codeType": "py",
  "gitRepo": "https://github.com/example/repo.git",
  "entryPointFileName": "main.py",
  "code": "",
  "downloadURL": ""
}
```

This flexible structure allows for handling various submission scenarios efficiently.

### Testing Suite

The project includes a comprehensive test suite, demonstrating my commitment to code quality and test-driven development. Tests cover various submission types and edge cases, ensuring robust functionality.

## Reflections

This project showcases my ability to:

- Design and implement complex systems
- Work with containerization technologies
- Handle diverse code submission formats
- Implement language-agnostic processing
- Develop RESTful APIs
- Write comprehensive test suites

While this project is no longer in active development, it represents a significant learning experience and demonstrates my skills in backend development, API design, and working with containerized applications.

## Note

This project is not actively maintained and is presented here as a portfolio piece to demonstrate my past work and technical abilities.
