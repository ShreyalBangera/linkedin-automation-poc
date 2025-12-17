# LinkedIn Automation Assignment (Proof of Concept)

## Overview

This project is a Go-based proof-of-concept built for the **LinkedIn Automation Assignment**.  
It demonstrates browser automation concepts using the **Rod** library, focusing on:

- Authentication handling
- Cookie-based session persistence
- Search and targeting
- Human-like behavior simulation
- Clean and modular Go architecture

⚠️ This project is strictly for **educational and technical evaluation purposes only**.

---
## 🎥 Demonstration Video

A walkthrough demonstration of the project setup, execution, and key features is available here:

🔗 <https://drive.google.com/file/d/1sDVglD94vv1FJx3lAi4RJLy-S1bFzMsA/view?usp=sharing>


## Features Implemented

### Authentication System
- Login using environment variables
- Detection of login success and failure
- Cookie-based session persistence
- Seamless session reuse without re-entering credentials
- Graceful handling of LinkedIn “Welcome Back” account chooser screen

### Search & Targeting
- Search LinkedIn users by keyword
- Extract LinkedIn profile URLs
- Duplicate profile detection
- Clean and reusable search module

### Stealth & Human-like Behavior
- Randomized delays between actions
- Human-like typing simulation
- Session reuse to reduce automation footprint
- Graceful UI state detection to avoid forced interactions

---

## Project Structure

linkedin-automation/
├── cmd/
│ └── main.go
├── search/
│ └── search.go
├── stealth/
│ ├── typing.go
│ ├── delays.go
│ └── scroll.go
├── storage/
│ └── cookies.go
├── .env.example
├── .gitignore
├── README.md


---

## Configuration

Create a `.env` file locally (not committed to GitHub):

LINKEDIN_EMAIL=your_email_here
LINKEDIN_PASSWORD=your_password_here


---

## How to Run

1. Install Go (1.20+ recommended)
2. Clone the repository
3. Create `.env` using `.env.example`
4. Run:

```bash
go run cmd/main.go

The browser will open and reuse the existing LinkedIn session if cookies are available.

Known Limitations

LinkedIn UI changes frequently and may affect selectors

Pagination across search results is not implemented

Connection requests and messaging are intentionally excluded

Some UI interactions are intentionally not forced to ensure stability

These trade-offs were made to prioritize reliability and clean architecture.

Ethical Disclaimer

This project is developed only for educational and technical evaluation purposes.

Automating LinkedIn violates LinkedIn’s Terms of Service.
This tool must never be used in production or on live accounts.

Author Notes

This project focuses on architecture, automation patterns, and stealth concepts,
not on building a production-ready automation tool.
