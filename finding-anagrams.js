// cleaner
function anagrams(word, words) {
    return words.filter(item => item.split('').sort().join('') == word.split('').sort().join(''))
}


//obsolote
/* function anagrams(word, words) {
    const finalArr = []
    word = word.split('').sort().join('')
    const tempArr = words.map(item => item.split('').sort().join('')).filter((item, index) => {
        if(item == word){
            finalArr.push(words[index])
        }
    })
    return finalArr
} */