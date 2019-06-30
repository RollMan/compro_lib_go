package compro_lib

import (
  "bufio"
  "strconv"
  "os"
)

var sc = bufio.NewScanner(os.Stdin)

func Init() {
  sc.Split(bufio.ScanWords)
}

func NextLine() (string, bool) {
  nonstop := sc.Scan()
  return sc.Text(), nonstop
}

func NextWord() (string, bool) {
  nonstop := sc.Scan()
  return sc.Text(), nonstop
}

func NextInt() (int, bool) {
  nonstop := sc.Scan()
  res, _ := strconv.ParseInt(sc.Text(), 10, 32)
  return int(res), nonstop
}

func NextLong() (int64, bool) {
  nonstop := sc.Scan()
  res, _ := strconv.ParseInt(sc.Text(), 10, 64)
  return int64(res), nonstop
}

func NextFloat() (float32, bool) {
  nonstop := sc.Scan()
  res, _ := strconv.ParseFloat(sc.Text(), 32)
  return float32(res), nonstop
}

func NextDouble() (float64, bool) {
  nonstop := sc.Scan()
  res, _ := strconv.ParseFloat(sc.Text(), 64)
  return float64(res), nonstop
}
