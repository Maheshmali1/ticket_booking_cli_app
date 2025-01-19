# Ticket Booking CLI Application 🎫

A simple Command-Line Interface (CLI) application written in Go for managing ticket bookings. This project is designed to demonstrate the fundamentals of Go, including user input handling, validation, concurrency, and more.

## Features

- **User-Friendly Greetings**: Displays a welcome message to users.
- **Interactive Input**: Captures user details like:
    - Name
    - Email
    - Number of tickets to book
- **Input Validations**: Ensures the validity of user inputs (e.g., email format, ticket availability).
- **Retries for Booking**: Allows users to retry the booking process in case of invalid inputs.
- **Remaining Ticket Count**: Tracks and updates the number of tickets left.
- **Concurrency with Goroutines**: Sends email confirmations to users using Goroutines for improved performance.

## Technologies Used

- **Programming Language**: [Go (Golang)](https://golang.org/)
- **Concurrency**: Achieved using Goroutines for sending emails asynchronously.

## How to Run the Application

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Maheshmali1/ticket_booking_cli_app.git
   cd ticket_booking_cli_app
   ```

2. **Build the Application**:
   Ensure you have Go installed on your machine. Then run:
   ```bash
   go build
   ```

3. **Run the Application**:
   Execute the binary to start the CLI:
   ```bash
   ./ticket_booking_cli_app
   ```

4. **Follow the Prompts**:
    - Enter your details as prompted by the application.
    - Receive feedback on your booking status.

## Sample Workflow

1. The application greets the user.
2. Prompts for user details like name, email, and ticket count.
3. Validates user inputs and processes the booking.
4. Sends an email confirmation asynchronously while updating the ticket count.
5. Displays success or failure messages to the user.

## Key Concepts Explored

- **Basic Syntax**: Understanding Go's syntax and structure.
- **User Input Handling**: Using Go's standard library for input/output operations.
- **Error Handling**: Validating and handling erroneous inputs gracefully.
- **Concurrency**: Employing Goroutines to manage tasks concurrently.
- **Retries**: Allowing users to retry the process in case of failures.

## Learnings

This project serves as an excellent starting point for beginners exploring Go. It covers:
- The Go programming model.
- Handling concurrency using Goroutines.
- Input validation and user interaction in CLI-based applications.

## Future Enhancements

- Add a persistent data store for bookings.
- Implement user authentication.
- Enhance email confirmation by integrating external email services like SendGrid.

## Contribution

Contributions are welcome! Feel free to fork the repository, raise issues, or submit pull requests.

## Author
The Ticket Booking CLI Application is developed and maintained by [Mahesh Mali](https://github.com/Maheshmali1).