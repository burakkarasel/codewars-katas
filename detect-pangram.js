function isPangram(string){
    let alphabet = [...'abcdefghijklmnopqrstuvwxyz']
    string = string.toLowerCase()
    let value
    for(i of alphabet){
        if(!string.includes(i)){
            value = false
            break
        }
        else
            value = true
    }
    return value
}