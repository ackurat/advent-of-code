import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/string
import input
import parsing

pub fn main() {
  let rules = input.line_by_line("src/d5rules_mini.txt") |> parse_rules
  let input =
    input.line_by_line("src/d5short.txt") |> parsing.parse_list_of_numbers(",")

  part1(input, rules) |> io.debug
}

fn parse_rules(rules: List(String)) -> dict.Dict(Int, Int) {
  use dict, rule <- list.fold(rules, dict.new())
  case string.split(rule, "|") {
    [a, b] -> {
      let assert Ok(ap) = int.parse(a)
      let assert Ok(bp) = int.parse(b)
      dict.insert(dict, ap, bp)
    }
    _ -> dict
  }
}

fn parse_input(input) {
  use digits <- list.map(input)
  use chars, digits, idx <- list.index_fold(digits, dict.new())
  dict.insert(chars, digits, idx)
}

fn part1(input, rules) {
  use chars <- list.map(parse_input(input))
  chars
  |> dict.map_values(fn(char, pos) {
    case dict.has_key(rules, char) {
      True -> {
        io.debug("rules has key")
        io.debug(char)
        io.debug(rules)
        let assert Ok(opposing) = dict.get(rules, char)
        case dict.has_key(chars, opposing) {
          True -> {
            io.debug("chars has key")
            io.debug(chars)
            io.debug(opposing)
            let assert Ok(opposing_pos) = dict.get(chars, opposing)
            io.debug(opposing_pos)
            io.debug(pos)
            case opposing_pos > pos {
              True -> True
              _ -> False
            }
          }
          _ -> False
        }
      }
      _ -> False
    }
  })
}
