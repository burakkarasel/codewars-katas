function amountOfPages(summary){
    let total = 0
    let pages = 0
    while(total < summary){
      for(i = 1; i < Infinity ; i++){
        total += String(i).length
        pages++
        if(total == summary){
          break
        }
      }
    }
    return pages
}