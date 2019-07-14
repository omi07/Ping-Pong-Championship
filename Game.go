package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//var refree = map[string]interface{}{}
	var playerarr = []string{"Joey", "Monica", "Chandler", "Ross", "Phoebe", "Rachel", "Sachin", "Rohan"}
	var players = map[string]int{}
	players["Joey"] = 7
	players["Monica"] = 6
	players["Chandler"] = 6
	players["Ross"] = 5
	players["Phoebe"] = 5
	players["Rachel"] = 6
	players["Sachin"] = 4
	players["Rohan"] = 5

	var finalscorecard = map[string]map[string]int{}

	fmt.Printf("Player array : %v \n", players)
	fmt.Printf("Player names : %v \n", playerarr)
	shuffledarr := shuffle(playerarr)
	fmt.Printf("Random array  :%v \n", shuffledarr)
	finalscorecard["Round1"] = gameplay(shuffledarr, players, "Round1")
	shuffledarr = getwinners(finalscorecard["Round1"])
	finalscorecard["Round2"] = gameplay(shuffledarr, players, "Round2")
	shuffledarr = getwinners(finalscorecard["Round2"])
	finalscorecard["Round3"] = gameplay(shuffledarr, players, "Round3")
	shuffledarr = getwinners(finalscorecard["Round3"])
	finalscores, err := json.Marshal(finalscorecard)
	if err != nil {
		fmt.Printf("JSON Encoding Failed")
		fmt.Printf("Final Score Card: %v \n", finalscorecard)
	}
	fmt.Printf("Final Score Card: %v \n", string(finalscores))
	fmt.Printf("CHAMPION ==> %v \n\n", shuffledarr)
}
func getwinners(scorearr map[string]int) []string {
	var winnerarr = []string{}
	for k, v := range scorearr {
		if v == 5 {
			winnerarr = append(winnerarr, k)
		}
	}
	winners := shuffle(winnerarr)
	return winners

}
func gameplay(shuffledarr []string, players map[string]int, round string) map[string]int {
	var scorecard = map[string]int{}
	for k := 0; k < len(shuffledarr); {
		play(shuffledarr[k:k+2], players, scorecard)
		k = k + 2
	}
	fmt.Printf("SCORECARD %v :%v \n", round, scorecard)
	return scorecard

}
func play(playerrr []string, players map[string]int, scorecard map[string]int) {
	offensive := playerrr[0]
	defender := playerrr[1]
	fmt.Printf("Offensive : %v Defender : %v \n ", offensive, defender)
	num := getrandomnumber()
	defensivearr := getrandomarray(players[defender])
	fmt.Printf("Generated number : %v  Defensive array : %v \n", num, defensivearr)
	exist := checknumberexist(num, defensivearr)
	if exist == 0 {
		//fmt.Println("Offensive is Winner!!!")
		scorecard[offensive] = scorecard[offensive] + 1
		if scorecard[offensive] < 5 {
			play(playerrr, players, scorecard)
		} else if scorecard[offensive] == 5 {
			fmt.Println(offensive + " is Winner! of Game !!!")
			return
		}

	} else {
		//fmt.Println("Defender is Winner!!!")
		scorecard[defender] = scorecard[defender] + 1
		if scorecard[defender] < 5 {
			fmt.Printf("Switching roles \n")
			playerswitch := []string{defender, offensive}
			play(playerswitch, players, scorecard)
		} else if scorecard[defender] == 5 {
			fmt.Println(defender + " is Winner! of Game !!!")
			return
		}
	}

}

func shuffle(src []string) []string {
	final := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return final
}

func getrandomnumber() int {
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Intn(10)
	return num + 1

}

func getrandomarray(size int) []int {
	var defenarr = []int{}
	for i := 0; i < size; size-- {
		rand.Seed(time.Now().UTC().UnixNano())
		defenarr = append(defenarr, rand.Intn(10)+1)
	}
	return defenarr

}

func checknumberexist(num int, defensivearr []int) int {
	for _, v := range defensivearr {
		if num == v {
			return 1
		}
	}
	return 0
}
