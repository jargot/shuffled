/*
**  UTILITY FUNCTIONS
*/

package main

import (
  "fmt"
  "sort"
  "math/rand"
)

func print_stack(stack []int) {
  fmt.Print("TOP - ")
  fmt.Print(stack)
  fmt.Println(" - BOTTOM")
}

func is_eq(a []int, b []int) (bool){
  for i := 0; i < len(a); i++ {
    if (a[i] != b[i]) {
      return false;
    }
  }
  return true;
}

// TODO: Allow lower_limit with a default
// TODO: Check error condition
func gen_asc_uniq_numbers(count int, upper_limit int) (res []int) {
  if (count + 1 > upper_limit) {
    // TODO: Throw error in go ?
    fmt.Print("ERROR", count, upper_limit)
    return
  }

  // Init the slice with default value = -1
  res = make([]int, count)
  for i := 0; i < count; i++ {
    res[i] = -1
  }

  for j := 0; j < count; {
    candidate := rand.Intn(upper_limit)
    if (isUniqInSlice(candidate, res)) {
      res[j] = candidate;
      j++;
    }
  }
  sort.Ints(res)
  return
}

func isUniqInSlice(single int, slice []int) (bool) {
  for i := 0; i < len(slice) ; i++ {
    if (slice[i] == single) {
      return false;
    }
  }
  return true;
}

// TODO: is_in_cyclical_order

// func find_out_of_order_pos(shuffled_stack []int, correct_stack []int) (positions []int){
//   for i := 0; shuffled_stack[i]; i++ {
//     if ()
//   }
// }
