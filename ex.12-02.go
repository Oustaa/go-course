// package main

// import (
// 	"fmt"
// 	"time"
// )

// // func elapsedTime(start time.Time, end time.Time) string {
// // 	elapsed := end.Sub(start)
// // 	hours := strconv.Itoa(int(elapsed.Hours()))
// // 	minutes := strconv.Itoa(int(elapsed.Minutes()))
// // 	seconds := strconv.Itoa(int(elapsed.Seconds()))
// // 	return "The total execution time elapsed is: " + hours +
// // 		" hour(s) and " + minutes + " minute(s) and " + seconds + " second(s)!"
// // }

// func main() {
// 	// start := time.Now()
// 	// time.Sleep(1*time.Minute + 12*time.Second)
// 	// end := time.Now()
// 	// fmt.Println(elapsedTime(start, end))

// 	t1, err := time.Parse(time.RFC3339, "2019-09-27T22:18:11+00:00")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	t2, err := time.Parse(time.UnixDate, "2019-09-27T22:18:11+00:00")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	t3, err := time.Parse(time.ANSIC, "2019-09-27T22:18:11+00:00")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("RFC3339:", t1)
// 	fmt.Println("UnixDate", t2)
// 	fmt.Println("ANSIC", t3)
// 	fmt.Println("///////////////////////////")
// 	date := time.Date(2019, 9, 27, 18, 50, 48, 0, time.UTC)
// 	nextDate := date.AddDate(1, 2, 3)
// 	fmt.Println(nextDate)

// 	fmt.Println("///////////////////////////")
// 	current := time.Now()
// 	losAngeles, err := time.LoadLocation("America/Los_Angeles")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("The local current time is:", current.Format(time.
// 		ANSIC))
// 	fmt.Println("The time in Los Angeles is:", current.In(losAngeles).
// 		Format(time.ANSIC))
// }

package main

import (
	"fmt"
	"time"
)

func timeDiff(timezone string) (string, string) {
	current := time.Now()
	remoteZone, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println(err)
	}
	remoteTime := current.In(remoteZone)
	fmt.Println("The current time is:", current.Format(time.
		ANSIC))
	fmt.Println("The timezone:", timezone, "time is:", remoteTime)
	return current.Format(time.ANSIC), remoteTime.Format(time.
		ANSIC)
}

func ma2112in() {
	fmt.Println(timeDiff("America/Los_Angeles"))
}
