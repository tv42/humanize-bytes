  $ bytes2human 42
  42B

  $ bytes2human 2134124234 8866
  2.0GiB
  8.7KiB

  $ bytes2human -si 2134124234 8866
  2.1GB
  8.9KB

  $ bytes2human bork 42
  bytes2human: cannot convert line: strconv.ParseUint: parsing "bork": invalid syntax
  [1]

  $ bytes2human -sloppy bork 42
  bork
  42B

  $ printf '5000 foo\n99999999 bar\n34 partial' >input
  $ bytes2human <input
  4.9KiB foo
  95MiB bar
  34B partial (no-eol)
