package bw_87

import (
	"math"
	"strconv"
	"strings"
)

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	hm := make(map[int]int)

	hm[1] = 31
	hm[2] = 28
	hm[3] = 31
	hm[4] = 30
	hm[5] = 31
	hm[6] = 30
	hm[7] = 31
	hm[8] = 31
	hm[9] = 30
	hm[10] = 31
	hm[11] = 30
	hm[12] = 31

	// ALICE

	aliceArriveSlice := strings.Split(arriveAlice, "-")
	aliceLeaveSlice := strings.Split(leaveAlice, "-")

	aliceArriveM, _ := strconv.Atoi(aliceArriveSlice[0])
	aliceArriveDate, _ := strconv.Atoi(aliceArriveSlice[1])
	aliceLeaveM, _ := strconv.Atoi(aliceLeaveSlice[0])
	aliceLeaveDate, _ := strconv.Atoi(aliceLeaveSlice[1])

	// BOB

	bobArriveSlice := strings.Split(arriveBob, "-")
	bobLeaveSlice := strings.Split(leaveBob, "-")

	bobArriveM, _ := strconv.Atoi(bobArriveSlice[0])
	bobArriveDate, _ := strconv.Atoi(bobArriveSlice[1])
	bobLeaveM, _ := strconv.Atoi(bobLeaveSlice[0])
	bobLeaveDate, _ := strconv.Atoi(bobLeaveSlice[1])

	// not crossing
	if aliceArriveM-bobArriveM < 0 {
		return 0
	}

	// same month
	if aliceArriveM-bobLeaveM == 0 {
		return int(math.Min(float64(aliceLeaveDate-aliceArriveDate), float64(bobLeaveDate-bobArriveDate)))
	}

	bobArriveM = aliceArriveM

	if bobLeaveM > aliceLeaveM {
		bobLeaveM = aliceLeaveM
	}



	return 0
}
