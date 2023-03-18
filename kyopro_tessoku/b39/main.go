package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

var reader io.Reader
var writer io.Writer

const (
	MOD = 1000000007
	// ↓AtCoderがGo1.14.1を使用しているため、定義しておく
	intSize = 32 << (^uint(0) >> 63) // 32 or 64
	MaxInt  = 1<<(intSize-1) - 1
	MinInt  = -1 << (intSize - 1)
)

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

type job struct {
	// x日以降に有効になる
	x int
	// y=給料
	y int
	// enabled=一度仕事をしたらenabledがfalseになる
	enabled bool
}

func newJob(x, y int) *job {
	return &job{
		x:       x,
		y:       y,
		enabled: true,
	}
}

func (j *job) validDate(date int) bool {
	return date >= j.x && j.enabled
}

func (j *job) work() int {
	j.enabled = false
	return j.y
}

type jobs []*job

func (jobs jobs) String() string {
	ret := "\n"

	for _, job := range jobs {
		ret += fmt.Sprintf("job: x(%d) y(%d) enabled(%v)\n", job.x, job.y, job.enabled)
	}

	return ret
}

func (jobs *jobs) sort() {
	sort.Slice(*jobs, func(i, j int) bool {
		a := (*jobs)[i]
		b := (*jobs)[j]

		if a.y > b.y {
			if a.y == b.y {
				return a.x < b.x
			}

			return a.y > b.y
		}

		return false
	})
}

// n 日目
// return 給料
func (jobs *jobs) searchAndWork(n int) int {
	for _, j := range *jobs {
		if j.validDate(n) {
			return j.work()
		}
	}

	return 0
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	// N=仕事の数
	// D=日数
	N, D := ni2(sc)

	jobs := make(jobs, N)
	for i := 0; i < N; i++ {
		x, y := ni2(sc)
		jobs[i] = newJob(x, y)
	}

	jobs.sort()

	answer := 0
	// D日間仕事をするので、D日分ループする
	// i日目
	for i := 1; i < D+1; i++ {
		answer += jobs.searchAndWork(i)
	}

	fmt.Fprint(writer, answer)
}

// ==================================================
// io
// ==================================================

func ni(sc *bufio.Scanner) int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func ns(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func ni2(sc *bufio.Scanner) (int, int) {
	return ni(sc), ni(sc)
}

func ni3(sc *bufio.Scanner) (int, int, int) {
	return ni(sc), ni(sc), ni(sc)
}

func ni4(sc *bufio.Scanner) (int, int, int, int) {
	return ni(sc), ni(sc), ni(sc), ni(sc)
}

func nis(sc *bufio.Scanner, n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni(sc)
	}
	return a
}

// n: 行数
// m: 要素数
func nir(sc *bufio.Scanner, n int, m int) [][]int {
	a := make([][]int, n)
	for i := range a {
		a[i] = nis(sc, m)
	}
	return a
}

// ==================================================
// util
// ==================================================

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func sumInts(s []int) int {
	var ret int
	for _, v := range s {
		ret += v
	}
	return ret
}

func maxInt(s []int) int {
	if len(s) <= 0 {
		return MinInt
	}

	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}

	return max
}
