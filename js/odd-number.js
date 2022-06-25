function findOdd(A) {
    let numberOBJ = {}
    let newNumbers = [... new Set(A)]
    let oddOne
    
    for (i of A){
      if(numberOBJ[i] == undefined ){
          numberOBJ[i] = 1
      }
      else{
          numberOBJ[i] += 1
      }
  }
  
  for(i of newNumbers){
      if(numberOBJ[i] % 2 == 0)
      continue
      else
      oddOne = i
  }
    return oddOne;
  }