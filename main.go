// package main2

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// 	"sync"
// )

// var client = &http.Client{}

// func UnmarshalFootball(data []byte) (Football, error) {
// 	var r Football
// 	err := json.Unmarshal(data, &r)
// 	return r, err
// }

// func (r *Football) Marshal() ([]byte, error) {
// 	return json.Marshal(r)
// }

// type Football struct {
// 	Page       string  `json:"page"`
// 	PerPage    int64   `json:"per_page"`
// 	Total      int64   `json:"total"`
// 	TotalPages int64   `json:"total_pages"`
// 	Data       []Datum `json:"data"`
// }

// type Datum struct {
// 	Competition string `json:"competition"`
// 	Year        int64  `json:"year"`
// 	Round       string `json:"round"`
// 	Team1       string `json:"team1"`
// 	Team2       string `json:"team2"`
// 	Team1Goals  string `json:"team1goals"`
// 	Team2Goals  string `json:"team2goals"`
// }

// func doHttpCalls(year int32, page int32) Football {

// 	resp, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?year=%d&page=%d", year, page))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	resBytesx, _ := ioutil.ReadAll(resp.Body)
// 	football, err := UnmarshalFootball(resBytesx)
// 	if err != nil {
// 		fmt.Println(err)
// 		return Football{}
// 	}
// 	return football
// }
// func compareMatchesInfo(data []Datum) []string {
// 	var trueData []string
// 	for _, match := range data {
// 		if match.Team1Goals == match.Team2Goals {
// 			trueData = append(trueData, match.Competition)
// 		}
// 	}
// 	return trueData
// }

// /*
//  * Complete the 'getNumDraws' function below.
//  *
//  * The function is expected to return an INTEGER.
//  * The function accepts INTEGER year as parameter.
//  */

// func getNumDraws(year int32) int32 {
// 	var competitions []string
// 	football := doHttpCalls(year, 1)
// 	var total int32 = int32(football.TotalPages)
// 	listOfTrue := compareMatchesInfo(football.Data)
// 	competitions = append(competitions, listOfTrue...)

// 	var i int32
// 	var wg sync.WaitGroup
// 	wg.Add(int(total))
// 	for i = 0; i < total; i++ {
// 		go func(x int32) {
// 			defer wg.Done()
// 			football := doHttpCalls(year, x)
// 			total = int32(football.TotalPages)
// 			listOfTrue := compareMatchesInfo(football.Data)
// 			competitions = append(competitions, listOfTrue...)
// 			fmt.Println(len(competitions))
// 		}(i)

// 	}
// 	wg.Wait()

// 	return int32(len(competitions))
// }

// func main() {
// 	// reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

// 	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	// checkError(err)

// 	// defer stdout.Close()

// 	// writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

// 	// yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	// checkError(err)
// 	// year := int32(yearTemp)

// 	res := getNumDraws(2011)

// 	fmt.Printf("%d\n", res)

// 	//     writer.Flush()
// 	// }
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }