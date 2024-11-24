# SQL Injection Demonstration

This project demonstrates the concept of SQL injection, its risks, and how to prevent it. It includes a basic login system, and a products query page to showcase both a vulnerable implementation and a secure implementation.


## What is SQL Injection?

SQL Injection is a security vulnerability that allows attackers to manipulate SQL queries by injecting malicious input. This can lead to unauthorized data access, data corruption, or even full system compromise.


## Key Points

- **Vulnerable Implementation**: Shows how improper handling of user input can allow an attacker to subvert application logic, and reveal database contents.
- **Secure Implementation**: Demonstrates how to protect against SQL injection by using prepared statements.

## How to Run

1. Rename .env.default to .env
2. Populate .env with database details, and JWT Password
3. Compile with "go build -o sqli ./cmd"
4. Run with "./sqli"