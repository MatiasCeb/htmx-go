# HTMX-Go Contact Manager

A modern web application for managing contacts built with Go and HTMX, featuring a clean, minimalist design inspired by Apple's interface philosophy.

## Live Demo

The application is live and accessible at: [https://htmx-go.erp.homme.ar/](https://htmx-go.erp.homme.ar/)

## Features

- **Contact Management**: Add, view, and delete contacts with real-time updates
- **Dynamic UI**: Powered by HTMX for seamless, server-driven interactions without JavaScript
- **Clean Design**: Minimalist interface with smooth animations and responsive layout
- **Form Validation**: Client-side and server-side validation for contact data
- **Docker Support**: Containerized deployment with multi-stage Dockerfile
- **Static Assets**: Served CSS and images for enhanced user experience

## Technologies Used

- **Backend**: Go 1.22.3 with Echo framework
- **Frontend**: HTML5, CSS3, HTMX 2.0.0
- **Styling**: Custom CSS with Apple-inspired design
- **Containerization**: Docker with Alpine Linux

## Prerequisites

- Go 1.22.3 or later
- Docker (optional, for containerized deployment)

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd htmx-go
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

## Running the Application

### Local Development

1. Start the server:
   ```bash
   go run main.go
   ```

2. Open your browser and navigate to `http://localhost:42069`

### Using Docker

1. Build the Docker image:
   ```bash
   docker build -t htmx-go-contact-manager .
   ```

2. Run the container:
   ```bash
   docker run -p 42069:42069 htmx-go-contact-manager
   ```

3. Access the application at `http://localhost:42069`

## Project Structure

```
htmx-go/
├── main.go              # Main application file with Echo server setup
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
├── Dockerfile           # Docker configuration for containerization
├── views/               # HTML templates
│   ├── index.html       # Home page template
│   ├── contacts.html    # Contacts list template
│   ├── add_contact.html # Add contact form template
│   └── blocks.html      # Reusable template blocks
├── css/                 # Stylesheets
│   └── index.css        # Main stylesheet
└── images/              # Static images
    └── Spinner.svg      # Loading spinner
```

## API Endpoints

- `GET /` - Home page
- `GET /contacts` - View all contacts
- `GET /add-contact` - Add new contact form
- `POST /contacts` - Create new contact
- `DELETE /contacts/:id` - Delete contact by ID
- `GET /blocks` - Paginated content blocks
- `POST /count` - Increment counter (demo endpoint)

## Usage

1. Navigate to the home page to see the welcome screen
2. Click "Add Contact" to create a new contact
3. Fill in the name and email fields
4. Submit the form to add the contact
5. View all contacts on the "Contacts" page
6. Delete contacts using the delete button (with confirmation)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgments

- [Echo](https://echo.labstack.com/) - High performance, extensible web framework for Go
- [HTMX](https://htmx.org/) - Access to AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML
- Inspired by Apple's design principles for clean, intuitive user interfaces