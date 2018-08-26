package slices

//ContainsString returns true if the container contains the item
func ContainsString(container []string, item string) (contained bool) {
  for _, i := range container{
    if i == item {
      return true
    }
  }
  return false
}

//AddStringIfNotContained adds items to the container if it does not already contain them
func AddStringIfNotContained(container []string, items ...string)  []string{
  for _, i := range items {
    if !ContainsString(container, i){
      container = append(container, i)
    }
  }
  return container
}
