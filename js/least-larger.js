function leastLarger(a,i) {
    return a.indexOf(Math.min(...a.filter(item => item > a[i])))
}


// obsolote
/* function leastLarger(a,i) {
    if(!a.length){
      return undefined 
    } else if (a[i] === Math.max(...a)){
      return -1;
    }else {
      return a.indexOf((a.filter(item => item > a[i]).sort((b,c) => b - c))[0])
    }    
} */