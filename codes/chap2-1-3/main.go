package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
type IO struct {
  scanner *bufio.Scanner
  writer  *bufio.Writer
}

func newIO() *IO {
  return &IO{
    scanner: bufio.NewScanner(os.Stdin),
    writer:  bufio.NewWriter(os.Stdout),
  }
}

func (io *IO) nextLine() string {
  io.scanner.Scan()
  return io.scanner.Text()
}

func (io *IO) nextInt() int {
  i, e := strconv.Atoi(io.nextLine())
  if e != nil {
    panic(e)
  }
  return i
}

func (io *IO) printLn(a ...interface{}) {
  fmt.Fprintln(io.writer, a...)
}

var io = newIO()

func main() {
  io.scanner.Split(bufio.ScanWords)      // switch to separating by space
  io.scanner.Buffer([]byte{}, 100000007) // switch to read large size input
  defer io.writer.Flush()

  N := io.nextInt()
  res := program(N)
  io.printLn(res)
}

func program(N int) string {
  var res string
  val := N

  for val >= 1 {
    if (val % 2 == 0) { res = "0" + res }
    if (val % 2 == 1) { res = "1" + res }
    val = val / 2
  }
  return res
}