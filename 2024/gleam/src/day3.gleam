import gleam/int
import gleam/string
import gleam/list
import gleam/regexp
import gleam/io
import input

pub fn main() {
  let input =
    input.line("src/d3short.txt")

  part1(input) |> io.debug
  part2(input) |> io.debug
}

fn part1 (input) {
  input
  |> extract_muls
}

fn extract_muls (input) {
  let opts = regexp.Options(True, False)
  let assert Ok(reg) = regexp.compile("mul\\([0-9]+,[0-9]+\\)", opts)

  regexp.scan(reg, input)
  |> list.map(fn(match) { match.content})
  |> parse_mul
}

fn parse_mul (mul) {
  let opts = regexp.Options(True, False)
  let assert Ok(reg) = regexp.compile("[0-9]+,[0-9]+", opts)

  mul
  |> list.fold(0, fn(acc, mu) {
    let p =
    regexp.scan(reg, mu)
    |> list.map(fn(match) { match.content})
    |> list.map(fn(digits) {
      string.split(digits, ",")
      |> list.map(fn(digit) {
        let assert Ok(parsed) = int.parse(digit)
        parsed
      })
    })
    |> list.flatten
    |> int.product


    acc + p
  })
}

fn part2 (input) {
  input
  |> extract_muls_pt2
}

fn extract_muls_pt2 (input) {
  let opts = regexp.Options(True, False)
  let assert Ok(reg) = regexp.compile("do\\(\\)(mul\\([0-9]+,[0-9]+\\))+", opts)

  regexp.scan(reg, input)
  |> list.map(fn(match) { match.content})
  |> parse_mul
} 
