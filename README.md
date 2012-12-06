humanize-bytes -- Utilities to convert "MiB" etc to raw numbers
===============================================================

Install with

  go get github.com/tv42/humanize-bytes/cmd/bytes2human github.com/tv42/humanize-bytes/cmd/human2bytes

Use like this:

  $ human2bytes 42GiB
  45097156608

  $ human2bytes 42GB
  42000000000

For example, sort your `du`, but still get readable results:

  $ du * | sort -nr | head -5 | bytes2human
