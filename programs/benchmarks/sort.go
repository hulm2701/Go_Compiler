func main(){
  print("Start")
  var data, sorted_data []int
  var size int = 100000
  // Add values in data with pseudo-random values
  j := 5000000
  for i:=0; i<size; i++ {
    data = append(data, j)
    j -= 2
  }

  var inf = 1000000000
  var min, pointer_min = inf, -1;


  for j:=0; j<size-1; j++ {
    for i:=0; i<size; i++ {
        var temp int
        temp = data[i]
        if temp < min {
            min = temp
            pointer_min = i
        }
    }
    sorted_data = append(sorted_data, min)
    data[pointer_min] = inf
    min = inf
  }
}
