# Ascii-Art-Web

## Description
Ascii-Art-Web is a web application that allows users to generate ASCII art using different banners. It provides a graphical user interface (GUI) where users can input text, select a banner style, and generate the corresponding ASCII art. The application is built using Go and follows good coding practices.

## Authors
- [Théo Ichane](https://github.com/TheoIchane)
- [Jaouhar Benromdhane](https://zone01normandie.org/git/jbenromd)
- [Alexandre Casimir](https://zone01normandie.org/git/acasimir)

## Usage: How to Run
1. **Clone the repository**:
    ```sh
    git clone https://github.com/jbenromd/ascii-art-web.git
    cd ascii-art-web
    ```

2. **Run the server**:
    ```sh
    go run main.go
    ```

3. **Open your browser** and navigate to `http://localhost:8080` to access the application.

## Implementation Details: Algorithm
The application uses predefined ASCII art banners (`shadow`, `standard`, `thinkertoy`) to transform input text into ASCII art. It provides a web interface for users to input their text and select a banner style. Upon submission, the server processes the request and generates the ASCII art, which is then displayed on the webpage.

### HTTP Endpoints

#### GET /
- **Description**: Sends the main HTML page.
- **Response**: 
    - HTML content displaying the main page with text input, banner selection, and submit button.
    - **Status Codes**: 
        - 200 OK if the page loads successfully.
        - 404 Not Found if the template is not found.

#### POST /ascii-art
- **Description**: Processes the text and banner selection, generates the ASCII art, and displays the result.
- **Request Body**: Form data containing:
    - `text`: The input text to be converted to ASCII art.
    - `banner`: The selected banner style.
- **Response**:
    - Displays the generated ASCII art on the webpage.
    - **Status Codes**:
        - 200 OK if the ASCII art is generated successfully.
        - 400 Bad Request if the input is invalid.
        - 500 Internal Server Error for unhandled errors.

## Instructions
1. **Navigate to the main page**:
    - Enter your text in the text input field.
    - Select a banner style using the provided options.
    - Click the submit button to generate the ASCII art.

2. **View the result**:
    - The generated ASCII art will be displayed on the same page or on a new page depending on the implementation.
      
3. **Export the result**:
    - Click the 'Export' button to generate and download a txt file with the output.
   
## Project Structure
- `main.go`: The main Go file containing the HTTP server implementation.
- `templates/`: Directory containing HTML templates for the web pages.

## Good Practices
- The code follows Go's standard coding practices.
- Only standard Go packages are used.

## Allowed Packages
- Only standard Go packages are allowed to be used in this project.


## Algorithm

To correctly serve our website we need to import some packages: 
- [net/http](https://pkg.go.dev/net/http)
- [log](https://pkg.go.dev/log)
- [html/template](https://pkg.go.dev/html/template)

In the first place we define a Handler Function to serve and parse HTML templates.
This function collect the http request sent by the user (text and banner) and use it to create the Ascii Art. 

This function handle also some errors like bad request and any others server errors.
After that it writes the parsed template with the data to the HTTP response.

In the second place we configure the HTTP server to handle static files.

And finally it start the configured server on a specified port (:8080 in our case).
