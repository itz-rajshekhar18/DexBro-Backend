@echo off
echo ================================
echo DexBro Workshop Backend Setup
echo ================================
echo.

REM Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Go is not installed. Please install Go 1.21 or higher.
    exit /b 1
)

echo [OK] Go is installed
go version
echo.

REM Check if .env exists
if not exist .env (
    echo [INFO] Creating .env file from .env.example...
    copy .env.example .env
    echo [OK] .env file created. Please update it with your MongoDB credentials.
) else (
    echo [OK] .env file already exists
)

REM Install dependencies
echo [INFO] Installing Go dependencies...
go mod download
go mod tidy

echo.
echo ================================
echo Setup complete!
echo ================================
echo.
echo Next steps:
echo 1. Update .env file with your MongoDB connection details
echo 2. Make sure MongoDB is running
echo 3. Run the server with: go run main.go
echo.
echo Available commands:
echo   go run main.go    - Run the server
echo   go build          - Build the binary
echo   go test ./...     - Run tests
echo.

pause
