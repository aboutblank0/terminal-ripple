
# Ripple Effect Terminal Application
![alacritty_7mHt3PauSY](https://github.com/user-attachments/assets/5bbe0f59-598b-4619-ba30-a33afdc0cd6d)

A full-screen terminal application that simulates ripple effects on the screen when the user clicks, acting like a mini graphics library.

This project was created for fun and learning, and does not use any external libraries, only Go standard libraries.

## Features
- Optimized, somewhat. Only draw "dirty" cells.
- Supports Key/Mouse input.
- Easily extensible
- Easy to add any kind of rendering you want, provided you know exactly which cells (x, y) to change

## Usage

1. Run the application:
    ```bash
    go run main.go

    ```

2. Click on the terminal window to create ripple effects.

