/* 
Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3. 
*/

const missing = (arr) => {
    let max = Math.max(...arr)
    for(let i = 1; i < max; i++){
        if(!arr.includes(i)){
            return i
        }
    }
    return max + 1
}

console.log(missing([5, 4, 2, -1, 0]))
console.log(missing([3, 4, -1, 1]))
console.log(missing([1, 2, 0]))