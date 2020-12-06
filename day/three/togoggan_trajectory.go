package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

const FILE_NAME = "./input.txt"

func parse_input_file(file_name string) *[][]bool {
  f, err := os.Open(file_name)

  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()

  scanner := bufio.NewScanner(f)

  forest := make([][]bool, 0)

  for scanner.Scan() {
    row_text := scanner.Text()

    row := make([]bool, len(row_text))

    for i, c := range row_text {
      if c == '.' {
        row[i] = false
      } else {
        row[i] = true
      }
    }

    forest = append(forest, row)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return &forest
}

func print_forest(forest *[][]bool) {
  for _, row := range(*forest) {
    for _, c := range(row) {
      if c {
        fmt.Print("#")
      } else {
        fmt.Print(".")
      }
    }
    fmt.Println()
  }
}

func sled(forest *[][]bool, right int, down int) int {
  column := 0

  trees_encountered := 0

  fmt.Println("Length of row:", len((*forest)[0]))

  for i := 0; i < len(*forest); i += down {
    row := (*forest)[i]

    if row[column] {
      trees_encountered++
    }

    column = (column + right) % len(row)
  }

  return trees_encountered
}

func main() {
  forest := parse_input_file(FILE_NAME)

  right_1_down_1 := sled(forest, 1, 1)
  right_3_down_1 := sled(forest, 3, 1)
  right_5_down_1 := sled(forest, 5, 1)
  right_7_down_1 := sled(forest, 7, 1)
  right_1_down_2 := sled(forest, 1, 2)

  fmt.Println("Right 1, Down 1:", right_1_down_1)
  fmt.Println("Right 3, Down 1:", right_3_down_1)
  fmt.Println("Right 5, Down 1:", right_5_down_1)
  fmt.Println("Right 7, Down 1:", right_7_down_1)
  fmt.Println("Right 1, Down 2:", right_1_down_2)

  fmt.Println(right_1_down_1 * right_3_down_1 * right_5_down_1 * right_7_down_1 * right_1_down_2)

}
