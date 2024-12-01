import gleam/list
import gleam/order
import gleam/string

pub type Matrix(default) {
  Matrix(cells: List(List(default)))
}

pub fn new_matrix(w: Int, h: Int, default: value) -> Matrix(value) {
  Matrix(cells: list.repeat(list.repeat(default, w), h))
}

pub fn empty_matrix(_default: value) -> Matrix(value) {
  Matrix(cells: [])
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

pub fn push(matrix: Matrix(value), new_row: List(value)) -> Matrix(value) {
  let new_cells = list.append(matrix.cells, [new_row])
  Matrix(cells: new_cells)
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
          current_row
          |> list.split(list.length(current_row) - shift)
          |> fn(splitted) { list.append(splitted.1, splitted.0) }
        }
        False -> current_row
      }
    }),
  )
}

pub fn transpose_matrix(matrix: Matrix(value)) -> Matrix(value) {
  Matrix(cells: list.transpose(matrix.cells))
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

pub fn sort(
  matrix: Matrix(value),
  compare_fn: fn(value, value) -> order.Order,
) -> Matrix(value) {
  let sorted_cells =
    list.map(matrix.cells, fn(row) { list.sort(row, compare_fn) })
  Matrix(cells: sorted_cells)
}

pub fn get_cells(matrix: Matrix(value)) -> List(List(value)) {
  matrix.cells
}
