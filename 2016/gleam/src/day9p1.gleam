import gleam/int
import gleam/io
import gleam/list
import gleam/option
import gleam/regexp
import gleam/result
import gleam/string
import input

pub fn main() {
  input.line_by_line("src/d9mini.txt")
  |> list.fold("", expand)
  |> io.debug
}

fn expand(data: String, expanded: String) -> String {
  let options = regexp.Options(case_insensitive: False, multi_line: True)
  let assert Ok(pattern) =
    regexp.compile("\\A(.*?)\\((\\d+)x(\\d+)\\)(.)(.*)\\z", options)

  io.debug(regexp.scan(pattern, data))
  case regexp.scan(pattern, data) {
    [match] -> {
      case match.submatches {
        [prefix, chars, repeat, suffix, ..] -> {
          let chars = chars |> maybe_to_int
          let repeat = repeat |> maybe_to_int

          let uncompressed =
            case suffix {
              option.Some(suffix_str) -> suffix_str
              option.None -> ""
            }
            |> string.slice(0, chars)
            |> string.repeat(repeat)

          expand(
            string.slice(uncompressed, chars, string.length(uncompressed)),
            expanded <> option.unwrap(prefix, "") <> uncompressed,
          )
        }
        _ -> expanded <> data
      }
    }
    [] -> expanded <> data
    _ -> expanded <> data
  }
}

fn maybe_to_int(maybe_string: option.Option(String)) -> Int {
  case maybe_string {
    option.Some(string) -> int.parse(string) |> result.unwrap(0)
    _ -> 0
  }
}
