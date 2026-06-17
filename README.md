# Ascii Art Web

This is a simple web application that generate ASCII characters by selecting from the sytles and the input the user enters. The user will a text and select a style and the application will generate and display the ASCII character(s) depending on the style selected.

## Description 

Ascii-Art-Web make use of user-friendly interface for generating ASCII art from the input entered by the user. The user can select from a list of banner styles to customize the appearance of their ASCII art.

## Author

Terungwa Terkimbi

## How to run the ASCII-Art-Web application

1. Have Golang installed on you local machine
2. Clone this repository to your local machine
3. Navigate to the project directory in your terminal.
4. Use this command "go run ." to run the application.
5. Use this link `http://localhost:8080/` on your browser to access the application.

## Web Interface

1. Enter text in the enter text field.
2. Select banner style from the drop-down sector.
3. Click on the Go button to generate the ASCII art.

## Implementation
1. **Input phase:** The user enters an input and is read via tha HTML form.
2. **Select Banner:** The style of the banner is been determined depending on the user selection.
3. **Input conversion:** The entered  input is then converted to ascii characters depending on the selected banner.
4. **Result Rendering:** The ASCII character(s) are displayed on the webpage.

## Structure
1. **main.go:** This is the funcion that run the application.
2. **Templates:** This folder contains all the html files.
3. **Source:** This file houses the server source code and the ascii-art source code.
