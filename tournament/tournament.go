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

type team struct {
	name string
	mp   int
	w    int
	d    int
	l    int
	p    int
}

type teams map[string]*team

func (t *team) win() {
	t.mp++
	t.w++
	t.p += 3
}

func (t *team) loss() {
	t.mp++
	t.l++
}

func (t *team) draw() {
	t.mp++
	t.d++
	t.p += 1
}

func (t teams) getOrCreate(n string) *team {
	if t[n] == nil {
		t[n] = &team{name: n}
	}
	return t[n]
}

func (t teams) sort() []*team {
	var list []*team
	for _, v := range t {
		list = append(list, v)
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].p == list[j].p {
			return list[i].name < list[j].name
		}
		return list[i].p > list[j].p
	})
	return list
}

func (t teams) format(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}
	for _, team := range t.sort() {
		_, err := fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n", team.name, team.mp, team.w, team.d, team.l, team.p)
		if err != nil {
			return err
		}
	}
	return nil
}

func Tally(reader io.Reader, writer io.Writer) error {
	t := make(teams)

	scan := bufio.NewScanner(reader)
	for scan.Scan() {
		line := scan.Text()
		if line == "" || line[0] == '#' {
			continue
		}

		strs := strings.Split(line, ";")
		if len(strs) != 3 {
			return errors.New("")
		}

		t1, t2, out := strs[0], strs[1], strs[2]
		switch out {
		case "win":
			t.getOrCreate(t1).win()
			t.getOrCreate(t2).loss()
		case "loss":
			t.getOrCreate(t1).loss()
			t.getOrCreate(t2).win()
		case "draw":
			t.getOrCreate(t1).draw()
			t.getOrCreate(t2).draw()
		default:
			return errors.New("")
		}
	}

	return t.format(writer)
}
