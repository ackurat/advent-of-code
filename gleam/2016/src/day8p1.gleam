import gleam/int
import gleam/io
import gleam/list
import gleam/string
import input

type Instruction {
  Rect(width: Int, height: Int)
  RotateColumn(x: Int, by: Int)
  RotateRow(y: Int, by: Int)
}

type Matrix {
  Matrix(cells: List(List(String)))
}

const height = 6

const width = 50

fn new_matrix(w, h) -> Matrix {
  Matrix(cells: list.repeat(list.repeat(".", w), h))
}

fn to_string(matrix: Matrix) -> String {
  matrix.cells
  |> list.map(fn(row) { string.join(row, "") })
  |> string.join("\n")
}

pub fn main() {
  let matrix = new_matrix(width, height)

  to_string(matrix)
  |> io.println

  input.line_by_line("src/d8short.txt")
  |> list.filter_map(parse_line)
  |> list.map(fn(instruction) {
    case instruction {
      Rect(w, h) -> {
        io.println(
          "Create rectangle " <> int.to_string(w) <> "x" <> int.to_string(h),
        )
      }

      RotateColumn(x, by) ->
        io.println(
          "Rotate column " <> int.to_string(x) <> " by " <> int.to_string(by),
        )
      RotateRow(y, by) ->
        io.println(
          "Rotate row " <> int.to_string(y) <> " by " <> int.to_string(by),
        )
    }
  })
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
// fn handle_instruction(instruction: Instruction) {
//   todo
// }
