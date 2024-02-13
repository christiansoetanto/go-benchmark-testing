package main

import (
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/api/doFibo", func(c *gin.Context) {
		memo, _ := strconv.ParseBool(c.Query("memo"))
		n, _ := strconv.ParseInt(c.Query("n"), 10, 64)
		c.JSON(http.StatusOK, gin.H{"result": doFibo(int(n), memo)})
	})
	r.GET("/api/doloop", func(c *gin.Context) {
		n, _ := strconv.ParseInt(c.Query("n"), 10, 64)
		_ = doloop(int(n))
		c.JSON(http.StatusOK, gin.H{"result": "done"})
	})

	r.GET("/api/doSort", func(c *gin.Context) {
		n, _ := strconv.ParseInt(c.Query("n"), 10, 64)
		c.JSON(http.StatusOK, gin.H{"result": doSort(int(n))})
	})
	r.GET("/api/doBinarySearch", func(c *gin.Context) {
		n, _ := strconv.ParseInt(c.Query("n"), 10, 64)
		c.JSON(http.StatusOK, gin.H{"result": doBinarySearch(int(n))})
	})

	return r
}
func doFibo(n int, memo bool) int {
	if memo {
		memo := make(map[int]int)
		return doFiboMemo(n, memo)
	}
	return doFiboNormal(n)
}
func doFiboMemo(n int, memo map[int]int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = doFiboMemo(n-1, memo) + doFiboMemo(n-2, memo)
	return memo[n]
}
func doFiboNormal(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return doFiboNormal(n-1) + doFiboNormal(n-2)
}
func doloop(n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomNumber() int {
	return rng.Intn(1000) + 1
}

func generateRandomNumbers(n int) []int {
	var nums []int
	for i := 0; i < n; i++ {
		nums = append(nums, generateRandomNumber())
	}
	return nums
}
func doSort(n int) []int {
	nums := generateRandomNumbers(n)
	sort.Ints(nums)
	return nums
}
func doBinarySearch(n int) int {
	var nums []int
	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}
	return binarySearch(nums, 0, len(nums)-1, 0)
}
func binarySearch(arr []int, l int, r int, x int) int {
	if r >= l {
		mid := l + (r-l)/2
		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		}
		return binarySearch(arr, mid+1, r, x)
	}
	return -1
}
func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
