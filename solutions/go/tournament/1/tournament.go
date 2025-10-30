package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Team                           | MP |  W |  D |  L |  P

type competition struct {
	team string
	mp   int
	w    int
	d    int
	l    int
	p    int
}

func Tally(reader io.Reader, writer io.Writer) error {
	teams := make(map[string]*competition)

	scan := bufio.NewScanner(reader)
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			continue
		}

		strs := strings.Split(line, ";")
		if len(strs) != 3 {
			return errors.New("")
		}

		if _, ok := teams[strs[0]]; !ok {
			teams[strs[0]] = &competition{
				team: strs[0],
			}
		}
		if _, ok := teams[strs[1]]; !ok {
			teams[strs[1]] = &competition{
				team: strs[1],
			}
		}

		switch strs[2] {
		case "win":
			win(teams[strs[0]])
			loss(teams[strs[1]])
		case "loss":
			win(teams[strs[1]])
			loss(teams[strs[0]])
		case "draw":
			draw(teams[strs[0]])
			draw(teams[strs[1]])
		default:
			return errors.New("")
		}
	}

	res := mpSort(teams)
	format(res, writer)
	return nil
}

func win(com *competition) {
	com.mp++
	com.w++
	com.p += 3
}

func draw(com *competition) {
	com.mp++
	com.d++
	com.p += 1
}

func loss(com *competition) {
	com.mp++
	com.l++
}

func mpSort(mp map[string]*competition) []*competition {
	teams := make([]*competition, 0, len(mp))
	for _, v := range mp {
		teams = append(teams, v)
	}

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].p != teams[j].p {
			return teams[i].p > teams[j].p
		}
		return teams[i].team < teams[j].team
	})

	return teams
}

func format(teams []*competition, writer io.Writer) {
	writer.Write([]byte(fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")))

	for _, t := range teams {
		writer.Write([]byte(fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n",
			t.team,
			t.mp,
			t.w,
			t.d,
			t.l,
			t.p)))
	}
}
