var maxSequence = function(arr){
    let max = 0
    let tempMax = 0
    let newArr = arr.filter(item => 
        item <= 0
    )

    if(arr.length === 0){
        return 0
    }
    else if( newArr.length === arr.length )
        return 0
    else{
        for(item of arr){
        let newList = []
        newList.push(item)
        tempMax = newList.reduce((total, item) =>{
            return total + item
        }, 0)
        if(tempMax > max)
            max = tempMax
        for(i = arr.indexOf(item) + 1; i < arr.length; i++){
            newList.push(arr[i])
            tempMax = newList.reduce((total, item) =>{
                return total + item
            }, 0)
            if(tempMax > max)
            max = tempMax
        }
    }
    return max
}
}