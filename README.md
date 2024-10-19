# Gopherit
## Why I Built This Tool

Handling individual code snippets in go is a pain.

1. **Organizational Overhead:** Creating a new folder for each code snippet was cumbersome. The only way to run my snippets was creating unrelated unit tests, which added further BS to the code.
2. **Limited Flexibility:** The inability to store multiple `main` functions made it difficult to quickly run different snippets without constantly modifying the same file.

Gopherit helps you streamline the process of storing, managing, and running Go code snippets.

## Usage

### Clone the rep
 ```bash
    git clone https://github.com/tren03/gopherit.git
    cd gopherit
```
### To create a snippet 
You can give the snippet name with the .go extension or not, it works both ways.         
I have set this up to open neovim when the file created.
 ```bash
   go run main.go --create <snippet_name>
```
### To open a snippet
I have set this up to open neovim on selection
 ```bash
   go run main.go --open
```
### To execute a snippet by searching for it via fzf
 ```bash
   go run main.go --run
```

### Executing a snippet
PLEASE DO NOT GIVE THE .go SUFFIX TO THE SNIPPET NAME WHILE CALLING IT
 ```bash
   go run main.go <snippet_name>
```

## Contribution
This project is currently in development. I plan to add more features in the future, so feel free to contribute! If you find any bugs or issues, please let me know by opening an issue in the repository. Your feedback is greatly appreciated!
