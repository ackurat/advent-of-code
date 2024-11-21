import gleam/int
import gleam/io
import gleam/list
import gleam/string
import input
import matrix

type Instruction {
  Rect(width: Int, height: Int)
  RotateColumn(x: Int, by: Int)
  RotateRow(y: Int, by: Int)
}

const height = 1

const width = 7

pub fn main() {
  let matrix = matrix.new_matrix(width, height, False)

  input.line_by_line("src/d8short.txt")
  |> list.filter_map(parse_line)
  |> list.fold(matrix, handle_instruction)
  |> io.debug
}

fn parse_line(line: String) -> Result(Instruction, String) {
  case string.split(line, " ") {
    ["rect", dimensions] -> {
      case string.split(dimensions, "x") {
        [width, height] -> {
          case int.parse(width), int.parse(height) {
            Ok(w), Ok(h) -> Ok(Rect(w, h))
            _, _ -> Error("Invalid rect dimensions")
          }
        }
        _ -> Error("Invalid rect format")
      }
    }
    ["rotate", "column", x_part, "by", amount] -> {
      case string.drop_left(x_part, 2), int.parse(amount) {
        x_str, Ok(by) -> {
          case int.parse(x_str) {
            Ok(x) -> Ok(RotateColumn(x, by))
            _ -> Error("Invalid column number")
          }
        }
        _, _ -> Error("Invalid rotate column format")
      }
    }
    ["rotate", "row", y_part, "by", amount] -> {
      case string.drop_left(y_part, 2), int.parse(amount) {
        y_str, Ok(by) -> {
          case int.parse(y_str) {
            Ok(y) -> Ok(RotateRow(y, by))
            _ -> Error("Invalid row number")
          }
        }
        _, _ -> Error("Invalid rotate row format")
      }
    }
    _ -> Error("Unknown instruction")
  }
}

fn handle_instruction(
  matrix: matrix.Matrix(Bool),
  instruction: Instruction,
) -> matrix.Matrix(Bool) {
  case instruction {
    Rect(w, h) -> {
      list.range(0, w - 1)
      |> list.fold(matrix, fn(matrix, x) {
        list.range(0, h - 1)
        |> list.fold(matrix, fn(matrix, y) { matrix.set(matrix, x, y, True) })
      })
    }

    RotateColumn(_x, _by) -> {
      matrix
    }

    RotateRow(y, by) -> {
      matrix.shift_row_right(matrix, y, by)
    }
  }
}
