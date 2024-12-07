import gleam/dict
import gleam/list

pub type Point =
  #(Int, Int)

pub type Grid(default) =
  dict.Dict(Point, default)

pub type Direction {
  Right
  Down
  DownRight
  DownLeft
  Left
  Up
  UpRight
  UpLeft
}

const all_directions = [
  Right,
  Down,
  DownRight,
  DownLeft,
  Left,
  Up,
  UpRight,
  UpLeft,
]

fn dxdy(direction: Direction) {
  case direction {
    Right -> #(0, 1)
    Down -> #(1, 0)
    DownRight -> #(1, 1)
    DownLeft -> #(1, -1)
    Left -> #(0, -1)
    Up -> #(-1, 0)
    UpRight -> #(-1, 1)
    UpLeft -> #(-1, -1)
  }
}

fn points_in_direction(
  origin: Point,
  direction: Direction,
  length: Int,
) -> List(Point) {
  let #(row, col) = origin
  let #(dx, dy) = dxdy(direction)

  list.range(0, length - 1)
  |> list.map(fn(i) { #(row + dx * i, col + dy * i) })
}

fn relative_points_in_direction(
  origin: Point,
  direction: Direction,
  radius: Int,
) -> List(Point) {
  let #(row, col) = origin
  let #(dx, dy) = dxdy(direction)

  list.range(-radius, radius)
  |> list.map(fn(i) { #(row + dx * i, col + dy * i) })
}

pub fn get_x_pattern(
  origin origin: Point,
  radius radius: Int,
) -> #(List(Point), List(Point)) {
  let diagonal1 = relative_points_in_direction(origin, DownRight, radius)
  let diagonal2 = relative_points_in_direction(origin, DownLeft, radius)
  #(diagonal1, diagonal2)
}

pub fn get_adjacent(
  origin origin: Point,
  distance distance: Int,
) -> List(List(Point)) {
  all_directions
  |> list.map(fn(dir) { points_in_direction(origin, dir, distance) })
}
