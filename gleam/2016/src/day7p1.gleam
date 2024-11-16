import gleam/io
import gleam/list
import gleam/regex
import gleam/string
import input

pub fn main() {
  input.line_by_line("src/d7.txt")
  |> list.filter(supports_tls)
  |> list.length
  |> io.debug
}

fn supports_tls(input) {
  let opts = regex.Options(True, False)
  let assert Ok(reg) = regex.compile("\\[(.*?)\\]", opts)

  let bracketed =
    regex.scan(reg, input)
    |> list.map(fn(match) { match.content })
    |> list.map(string.to_graphemes)
    |> list.map(find_palindromes)
    |> list.any(fn(l) { !list.is_empty(l) })

  let outside_brackets =
    regex.split(reg, input)
    |> list.map(string.to_graphemes)
    |> list.map(find_palindromes)
    |> list.any(fn(l) { !list.is_empty(l) })

  outside_brackets && !bracketed
}

fn find_palindromes(input: List(String)) -> List(String) {
  input
  |> list.window(4)
  |> list.filter(fn(quad) {
    case quad {
      [a, b, c, d] -> a == d && b == c && a != b
      _ -> False
    }
  })
  |> list.map(fn(quad) {
    case quad {
      [a, b, c, d] -> a <> b <> c <> d
      _ -> ""
    }
  })
}
