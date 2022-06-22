function moveZeros(arr) {
    return arr.filter(item => item !== 0).concat(arr.filter(item => item === 0))
}

// obsolote
/* function moveZeros(arr) {
    for(i of arr){
        if(i === 0){
            let index = arr.indexOf(i)
            arr.splice(index, 1)
            arr.push(0)
        }
    }
    return arr
  } */