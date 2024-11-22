import gleam/io
import gleam/list

pub type Matrix(default) {
  Matrix(cells: List(List(default)), width: Int, height: Int)
}

pub fn new_matrix(w: Int, h: Int, default: value) -> Matrix(value) {
  Matrix(cells: list.repeat(list.repeat(default, w), h), width: w, height: h)
}

pub fn to_string(matrix: Matrix(value)) {
  matrix.cells
  |> list.map(fn(row) { io.debug(row) })
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
          let values =
            list.index_map(current_row, fn(value, col_index) {
              let source_pos = case col_index + shift % matrix.width {
                n if n >= matrix.width -> n - matrix.width
                n -> n
              }
              #(col_index, source_pos, value)
            })

          list.index_map(current_row, fn(value, col_index) {
            case
              list.find(values, fn(triple) {
                let #(_, source, _) = triple
                source == col_index
              })
            {
              Ok(#(_, _, v)) -> v
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

pub fn shift_column_down(
  matrix: Matrix(value),
  column: Int,
  shift: Int,
) -> Matrix(value) {
  Matrix(
    cells: list.index_map(matrix.cells, fn(current_row, row_index) {
      let values =
        list.index_map(matrix.cells, fn(row, i) {
          let source_pos = case i + shift % matrix.height {
            n if n >= matrix.height -> n - matrix.height
            n -> n
          }

          let assert Ok(col_value) =
            list.index_map(row, fn(v, j) {
              case j == column {
                True -> v
                False -> v
              }
            })
            |> list.first
          #(i, source_pos, col_value)
        })

      list.index_map(current_row, fn(val, col_index) {
        case col_index == column {
          True -> {
            case
              list.find(values, fn(triple) {
                let #(_, source, _) = triple
                source == row_index
              })
            {
              Ok(#(_, _, v)) -> v
              Error(_) -> val
            }
          }
          False -> val
        }
      })
    }),
    height: matrix.height,
    width: matrix.width,
  )
}
