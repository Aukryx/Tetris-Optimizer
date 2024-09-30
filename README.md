# Tetromino Solver

This Go program solves the problem of fitting a set of tetrominos into the smallest possible square grid. It uses a backtracking algorithm to find an optimal solution.

## How It Works

1. **Input**: The program takes a file path as a command-line argument. This file should contain tetromino shapes represented by '#' (filled) and '.' (empty) characters.

2. **Validation**: The input file is validated to ensure it contains valid tetromino representations.

3. **Preprocessing**: 
   - Tetrominos are extracted from the input file.
   - Each tetromino is validated and formatted.
   - Empty columns are removed from each tetromino.

4. **Solving**:
   - The program starts with a minimum square size based on the number of tetrominos.
   - It uses a backtracking algorithm to try placing tetrominos in the grid.
   - If a solution is not found, it increases the grid size and tries again.

5. **Output**: 
   - If a solution is found, it prints the solved grid to the console.
   - Each tetromino is represented by a unique emoji.
   - Empty spaces are represented by '.'.

6. **Performance**: The program measures and reports its execution time.

## How to Run

1. Ensure you have Go installed on your system.

2. Clone this repository and navigate to the project directory.

3. Run the program using the following command:

   ```
   go run . <path_to_tetromino_file>
   ```

   Replace `<path_to_tetromino_file>` with the path to your input file containing tetromino shapes.

4. The program will output the solution (if found) and the execution time.

## Example

```
$ go run . examples/goodexample00.txt
🤠🤠😱😱
🤠🤠😱🥵
🤠😱😱🥵
🥵🥵🥵..
Program took 0.00 seconds to finish
```

## Error Handling

The program will display error messages for various invalid inputs:

- If no file path is provided
- If the file format is invalid
- If any tetromino is invalid

## Testing

A bash script (`test.sh`) is provided to run the solver with various example inputs. To use it:

1. Make the script executable:
   ```
   chmod +x test.sh
   ```

2. Run the script:
   ```
   ./test.sh
   ```

This will test the solver with both valid and invalid inputs, demonstrating its behavior in different scenarios.

## Implementation Details

The solver is implemented in Go and consists of several packages:

- `main`: Contains the main program logic.
- `algo`: Implements the solving algorithm and timing functions.
- `format`: Handles file parsing, tetromino validation, and formatting.

The core solving algorithm uses backtracking with optimizations to efficiently find a solution.