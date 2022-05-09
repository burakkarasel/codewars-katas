function dontGiveMeFive(start, end){
    let counter = 0
    for(let i = start; i <= end; i++){
      if(String(i).indexOf('5') === - 1){
        counter += 1
      }
    }
    return counter
}