# pokedexcli

A command-line Pokedex application built in Go.

This project was created as part of my learning journey with [boot.dev](https://boot.dev). It provided valuable hands-on experience with several key Go concepts, including:

* **Concurrency:** Utilizing mutexes and channels to manage concurrent operations efficiently.
* **Client-Side API Interaction:** Fetching and processing data from a remote API.
* **Testing:** Implementing unit tests to ensure code reliability.
* **Project Organization:** Structuring a Go project with multiple packages and files.

## Features

* Explore a list of locations and find Pokemon.
* Get detailed information about a specific Pokemon.
* Catch 'em all.

## Getting Started

1.  **Clone the repository:**

    ```bash
    git clone <your_repository_url>
    cd pokedexcli
    ```

2.  **Build the application:**

    ```bash
    go build -o pokedexcli
    ```

3.  **Run the application:**

    ```bash
    ./pokedexcli
    ```

## Usage

```bash
./pokedexcli <command> [arguments]
```

## Available commands:

* **help**: Displays available commands and their usage.
* **exit**: Exits the Pokedex CLI.
* **map**: Displays the next set of locations.
* **mapb**: Displays the previous set of locations.
* **explore <location_area>**: Lists the Pokemon found in the specified location area.
* **pokemon <pokemon_name>**: Displays detailed information about the given Pokemon.
* **catch <pokemon_name>**: Attempts to catch the specified Pokemon.
* **inspect <pokemon_name>**: Displays information about a Pokemon in your Pokedex.
* **pokedex**: Lists the Pokemon you have caught.

## Project Structure
The project is organized into the following packages:

* **pokedexcli**: (Root) Contains the main application logic and CLI interface (main.go, `cli_cmds.go`).
* **cache**: Implements a simple caching mechanism (cache.go, cache_test.go).
* **pokeapi**: Handles communication with the PokeAPI (client.go).
* **cli**: Defines command line functionality.
* **config**: Manages application configuration (config.go).
Acknowledgements
This project utilizes the [PokeAPI](https://pokeapi.co/) for its data.

Thank you for checking out my pokedexcli project!
