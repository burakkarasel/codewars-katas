function solution(number){
    let total = 0
    let i = 0
    for(; i < number; i++){
      if(i % 3 == 0 && i % 5 == 0){
        total += i
      }
      else if(i % 3 == 0){
        total += i
      }
      else if(i % 5 == 0){
        total += i
      }
    }
    return total
  }