# Custom Error Library

`custom_error` is a Go package designed to simplify the creation and management of custom RESTful error responses. It provides a standardized way to handle errors in web applications by including structured error information such as a message, timestamp, HTTP status code, and potential causes. This library is especially useful for APIs, enabling detailed error information to be returned in a consistent format.

## Features

- **Consistent error structure**: Return detailed error responses with a message, timestamp, HTTP status code, and optional causes.
- **Predefined error constructors**: Quickly create common REST errors, such as `BadRequest`, `Unauthorized`, `NotFound`, and `InternalServerError`.
- **Customizable causes**: Specify error causes with field-level details for more precise debugging.

## Installation

To install this package, use the following command:

```bash
go get -u github.com/yourusername/custom_error
```

## Usage
### Creating a Custom Error

Each error created by the custom_error package returns a structured JSON error response. Here’s a quick example:

```go
err := custom_error.NewRestBadRequestError(
  "Invalid input data",
  []custom_error.Cause{
    {
      Field: "username",
      Message: "Username is required"
    },
  }
)
```

### Available Error Types

The following functions provide ready-to-use error constructors for common HTTP errors:

- `NewRestErr(message string, code int, causes []Cause)`: General-purpose error constructor.
- `NewRestBadRequestError(message string, causes []Cause)`: Creates a 400 Bad Request error.
- `NewRestUnauthorizedError(message string, causes []Cause)`: Creates a 401 Unauthorized error.
- `NewRestNotFoundError(message string, causes []Cause)`: Creates a 404 Not Found error.
- `NewRestInternalServerError(message string, causes []Cause)`: Creates a 500 Internal Server Error.

## Error Structure

Each RestErr object returned has the following structure:

```js
{
  "message": "Error message",
  "timestamp": "2024-10-26T15:30:00Z",
  "code": 400,
  "causes": [
    {
      "field": "field_name",
      "message": "Detailed error message"
    }
  ]
}
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing
Feel free to submit issues or create pull requests. We welcome any improvements and suggestions to make custom_error better!
