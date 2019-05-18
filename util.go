package gomapreduce

import (
	"strings"
)

type Alphabetic []string

func (list Alphabetic) Len() int { return len(list) }

func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list Alphabetic) Less(i, j int) bool {
    si := list[i]
    sj := list[j]
    siLower := strings.ToLower(si)
    sjLower := strings.ToLower(sj)
    if siLower == sjLower {
        return si < sj
    }
    return siLower < sjLower
}