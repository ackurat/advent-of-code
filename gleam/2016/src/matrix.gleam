import gleam/dict
import gleam/list
import gleam/string

pub type Matrix(default) {
  Matrix(cells: List(List(default)), width: Int, height: Int)
}

pub fn new_matrix(w: Int, h: Int, default: value) -> Matrix(value) {
  Matrix(cells: list.repeat(list.repeat(default, w), h), width: w, height: h)
}

pub fn to_string(
  matrix: Matrix(value),
  to_string_fn: fn(value) -> String,
) -> String {
  matrix.cells
  |> list.map(fn(row) {
    row
    |> list.map(to_string_fn)
    |> string.join("")
  })
  |> string.join("\n")
}

pub fn set(
  matrix: Matrix(value),
  x: Int,
  y: Int,
  new_value: value,
) -> Matrix(value) {
  Matrix(
    cells: list.index_map(matrix.cells, fn(row, row_index) {
      case row_index == y {
        True ->
          list.index_map(row, fn(value, col_index) {
            case col_index == x {
              True -> new_value
              False -> value
            }
          })
        False -> row
      }
    }),
    width: matrix.width,
    height: matrix.height,
  )
}

pub fn shift_row_right(
  matrix: Matrix(value),
  row: Int,
  shift: Int,
) -> Matrix(value) {
  Matrix(
    cells: list.index_map(matrix.cells, fn(current_row, row_index) {
      case row_index == row {
        True -> {
          let vals =
            list.index_fold(current_row, dict.new(), fn(acc, val, idx) {
              let source_pos = case idx + shift % matrix.width {
                n if n >= matrix.width -> n - matrix.width
                n -> n
              }
              dict.insert(acc, source_pos, #(idx, val))
            })

          list.index_map(current_row, fn(value, col_index) {
            case dict.get(vals, col_index) {
              Ok(#(_, v)) -> v
              Error(_) -> value
            }
          })
        }
        False -> current_row
      }
    }),
    width: matrix.width,
    height: matrix.height,
  )
}

pub fn transpose_matrix(matrix: Matrix(value)) -> Matrix(value) {
  Matrix(
    cells: list.transpose(matrix.cells),
    width: matrix.height,
    height: matrix.width,
  )
}

pub fn shift_column_down(
  matrix: Matrix(value),
  column: Int,
  shift: Int,
) -> Matrix(value) {
  transpose_matrix(matrix)
  |> shift_row_right(column, shift)
  |> transpose_matrix
}
