// Mini-Max Sum
// - Problem Solving (Basic)
// - Difficulty: Easy

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

// Full Problem: https://www.hackerrank.com/challenges/mini-max-sum/problem

function miniMaxSum(arr) {
  // Write your code here
<<<<<<< HEAD
<<<<<<< HEAD
  // TODO: answer here
=======
  //beginanswer
  var sumArr = [];
  var cursor = 0;
  for (var i = 0; i < arr.length; i++) {
    var sum = 0;
    for (var j = 0; j < arr.length; j++) {
      if (cursor != j)
        sum += arr[j];
    }
    cursor++;
    sumArr.push(sum);
  }
  var min = Math.min.apply(null, sumArr);
  var max = Math.max.apply(null, sumArr);
  // console.log(min + ' ' + max);
  return min + ' ' + max 
  //endanswer
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
  // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}

function main() {
  //arr = readLine().split(' ');
  //arr = arr.map(Number);
  var arr = [1, 2, 3, 4, 5]; // override input
  let result = miniMaxSum(arr);
  console.log(result)
}

main(); // execute program

module.exports = miniMaxSum