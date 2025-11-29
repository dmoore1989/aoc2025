# Advent of Code Go Template

A streamlined Go template for participating in [Advent of Code](https://adventofcode.com/). This template automatically downloads puzzle inputs, generates solution scaffolding, and provides convenient commands for running your solutions against sample and real inputs.

## Features

- Automatic puzzle input downloading from Advent of Code
- Auto-generated solution scaffolding for each day
- Makefile commands for quick testing and running
- Support for sample and real puzzle inputs
- Automatic browser opening to the puzzle page

## Prerequisites

- **Go** 1.24.5 or later ([Download Go](https://go.dev/dl/))
- **Make** (usually pre-installed on macOS/Linux)
- An **Advent of Code account** and session cookie

## Initial Setup

### 1. Get Your Advent of Code Session Cookie

To download puzzle inputs, you need your session cookie from the Advent of Code website:

1. Log in to [Advent of Code](https://adventofcode.com/)
2. Open your browser's Developer Tools (F12 or Right Click → Inspect)
3. Go to the **Application** or **Storage** tab
4. Navigate to **Cookies** → `https://adventofcode.com`
5. Find the cookie named `session` and copy its **Value**

### 2. Set Environment Variables

Set the following environment variables (add to your `~/.zshrc`, `~/.bashrc`, or `~/.bash_profile`):

```bash
export AOC_COOKIE="your_session_cookie_here"
export EMAIL="your.email@example.com"  # Used for User-Agent header
```

After adding these, reload your shell:
```bash
source ~/.zshrc  # or ~/.bashrc
```

### 3. Update the Year

Open `main.go` and update the `YEAR` constant to the current Advent of Code year:

```go
// UPDATE THE YEAR HERE TO PULL THE CORRECT DATA
const YEAR string = "2025"  // Change to your target year
```

### 4. Verify Go Installation

```bash
go version
```

## Usage

### Downloading a Puzzle & Setting Up a Day

To fetch puzzle input and create solution scaffolding for a specific day:

```bash
make get_day <day_number>
```

**Example:**
```bash
make get_day 1
```

This command will:
1. Download the puzzle input to `lib/day1.txt`
2. Create `cmd/day1/main.go` with a solution template
3. Open the puzzle page in your browser

**Note:** If no day number is provided, it defaults to the current day.

### Running Your Solution

#### Against Sample Input

Test your solution with the example input from the puzzle description. First, create a `sample.txt` file in the project root with the sample data.

```bash
# Part 1
make run_pt1_sample <day_number>

# Part 2
make run_pt2_sample <day_number>
```

**Example:**
```bash
make run_pt1_sample 1
make run_pt2_sample 1
```

#### Against Real Input

Run your solution with the actual puzzle input:

```bash
# Part 1
make run_pt1_real <day_number>

# Part 2
make run_pt2_real <day_number>
```

**Example:**
```bash
make run_pt1_real 1
make run_pt2_real 1
```

## Project Structure

```
aoc_template_go/
├── main.go              # Setup script for downloading puzzles
├── Makefile             # Convenient commands for running solutions
├── go.mod               # Go module definition
├── sample.txt           # Sample input for testing (create this yourself)
├── lib/
│   ├── aoc_template.txt # Template for generating day solutions
│   └── day*.txt         # Downloaded puzzle inputs (gitignored)
└── cmd/
    └── day*/            # Generated solution directories
        └── main.go      # Your solution code
```

## Solution Template

Each generated solution (`cmd/dayN/main.go`) includes:

```go
package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    args := os.Args[1:]

    filePath := "sample.txt"
    if args[0] == "real" {
        filePath = "lib/dayN.txt"
    }

    fileArr, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }

    fileStr := strings.TrimSuffix(string(fileArr), "\n")

    var answer string
    if args[1] == "1" {
        answer = part1(fileStr)
    } else {
        answer = part2(fileStr)
    }
    fmt.Println(answer)
}

func part1(fileTxt string) string {
    // Implement Part 1 solution here
    return "Part 1: " + fileTxt
}

func part2(fileTxt string) string {
    // Implement Part 2 solution here
    return "Part 2: " + fileTxt
}
```

## Typical Workflow

Here's a typical workflow for solving a day's puzzle:

1. **Fetch the puzzle** (at midnight EST when it unlocks):
   ```bash
   make get_day 1
   ```

2. **Read the puzzle** in your browser (auto-opened)

3. **Copy the sample input** from the puzzle description to `sample.txt`

4. **Implement Part 1** in `cmd/day1/main.go` → `part1()` function

5. **Test against sample**:
   ```bash
   make run_pt1_sample 1
   ```

6. **Run against real input**:
   ```bash
   make run_pt1_real 1
   ```

7. **Submit your answer** on the Advent of Code website

8. **Repeat for Part 2** (steps 4-7 using `part2()` function and `run_pt2_*` commands)

## Troubleshooting

### "Invalid cookie" error

Your session cookie has expired. Log in to Advent of Code again and update the `AOC_COOKIE` environment variable with a fresh cookie value.

### "Not time for new puzzle" error

The puzzle hasn't been released yet. Puzzles unlock at midnight EST (UTC-5).

### "Day already exists" error

You've already downloaded this day. Check `cmd/dayN/` for your existing solution.

### Permission issues

If you encounter permission issues, ensure the files have the correct permissions:
```bash
chmod +x main.go
```

## Tips

- **Reuse code**: Create a `utils/` package for common functions (parsing, grid traversal, etc.)
- **Test thoroughly**: Always validate against sample inputs before running on real inputs
- **Version control**: Commit your solutions to track your progress
- **Read carefully**: Many bugs come from misreading puzzle requirements

## License

This template is free to use for participating in Advent of Code. Per the TOS of Advent of Code, do not commit file inputs (the inputs are restricted by this template's git ignore).

## Acknowledgments

- [Advent of Code](https://adventofcode.com/) created by [Eric Wastl](http://was.tl/)
