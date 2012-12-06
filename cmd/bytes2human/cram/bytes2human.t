  $ human2bytes 42MiB
  44040192

  $ human2bytes 42MB
  42000000

  $ human2bytes 5MiB 10kiB
  5242880
  10240

  $ human2bytes bork 42
  human2bytes: cannot convert line: strconv.ParseFloat: parsing "": invalid syntax
  [1]

  $ human2bytes -sloppy bork 42
  bork
  42

  $ printf '5MB foo\n2kiB bar\n34 partial' >input
  $ human2bytes <input
  5000000 foo
  2048 bar
  34 partial (no-eol)
