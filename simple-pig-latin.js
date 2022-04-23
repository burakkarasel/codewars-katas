function pigIt(str){
    let tempArray = str.split(' ')
    tempArray.forEach((item, index, array)=>{
      for(i of item){
        let cc = i.charCodeAt(0)
        if ((cc > 47 && cc < 58) || (cc > 64 && cc < 91) || (cc > 96 && cc < 123)){
          array[index] = item.slice(1) + item[0] + 'ay' 
        }  
      }
    }
  )
    let newStr = tempArray.join(' ')
    return newStr
  }