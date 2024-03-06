import gleam/io
import gleam/string
import gleam/list
import gleam/int
import simplifile

pub fn main() {
  let assert Ok(input) = simplifile.read("src/y2022/d01/input.txt")

  io.debug(part_1(input))
  io.debug(part_2(input))
}

fn part_1(input) {
  let elf = totals(input)
  |> list.sort(by: int.compare)
  |> list.last
}

fn part_2(input) {
  let elf = totals(input)
  |> list.sort(by: int.compare)
  |> list.reverse
  |> list.take(3)
  |> list.fold(0, fn(acc, num) {
    acc + num
    })
}

fn totals(input) {
  input
  |> string.split(on: "\n\n")
  |> list.map(fn(elf) {
    elf
    |> string.split(on: "\n")
    |> list.fold(0, fn(acc, str) {
      let assert Ok(num) = int.parse(str)
      acc + num
      })
    })
}
