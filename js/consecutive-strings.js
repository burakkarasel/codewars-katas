function longestConsec(strarr, k){
    if(k <= strarr.length){
        let newStrings = []
        for(i = 0; i < strarr.length - (k-1); i++){
            let newString = ''
            let counter = 0
            for(ii = i; counter < k; ii++){
                newString += strarr[ii]
                counter++
            }
        newStrings.push(newString)
        }
        let newStringsCounter = newStrings.map(item => item.length)
        let maxLetter = Math.max(...newStringsCounter)
        let indexOfMaxLetter = newStringsCounter.indexOf(maxLetter)
        return newStrings[indexOfMaxLetter]
    }
    else
    return ''
}