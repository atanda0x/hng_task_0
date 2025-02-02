# Public API for Basic Information

This API returns basic information in JSON format, including:
- Registered email address
- Current datetime in ISO 8601 format (UTC)
- GitHub URL of the project's codebase

## API Endpoint

**GET /**

**Response Format (200 OK):**
```json
{
    "email":"atanda0x@gmail.com",
    "current_datetime":"2025-02-02T00:08:31Z",
    "github_url":"https/github.com/atanda0x/hng_task_0"
}
