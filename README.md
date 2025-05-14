# PS-Keyboard Admin CLI

This is an admin tool for managing PS Keyboard verification requests for my [PS Keyboard webpage](https://ps-keyboard.vercel.app/).

It's repo can be found [here](https://github.com/bartektricks/ps-keyboard)

## Usage

The CLI supports both interactive and non-interactive modes.

### Setup

1. Create a `.env` file based on the .env.example file:
   ```bash
   cp .env.example .env
   ```

2. Set the `DATABASE_URL` in your `.env` file.

3. Build and run the app:
   ```bash
   go build -o admin
   ./admin
   ```

### Commands

#### Interactive Mode

Running the binary without any flags launches interactive mode:
```bash
./admin
```

#### Non-Interactive Commands

- Print all verification requests:
  ```bash
  ./admin -print
  ```

- Accept a verification request:
  ```bash
  ./admin -accept <request-id>
  ```

- Reject a verification request:
  ```bash
  ./admin -reject <request-id>
  ```

## Development

### Project Structure

- `main.go`: Entry point that sets up the application
- `internal/cli`: Command-line interface handling
- `internal/config`: Configuration loading from environment
- `internal/db`: Database connection and repository
- `internal/model`: Data models
- `internal/service`: Business logic
- `internal/ui`: User interface components

### Getting Started with Development

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file with your database configuration

4. Run the app:
   ```bash
   go run main.go
   ```

### Making Changes

1. The main command processing happens in `cli.ExecuteCommand` in the internal/cli/cli.go file.
2. Business logic is implemented in `service.VerificationService` in the internal/service/verification.go file.

### Important Notes

⚠️ **Asynchronous Operations**: The application currently doesn't handle asynchronous operations. This functionality needs to be implemented if you want to process multiple requests concurrently or handle long-running operations without blocking the UI.

### Planned Improvements

- Implement asynchronous request handling for better performance
- Replace the current flag-based command parsing with a more robust solution like Cobra for better help documentation, nested commands, and improved user experience
