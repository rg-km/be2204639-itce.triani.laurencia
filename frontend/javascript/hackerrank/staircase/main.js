// Staircase
// - Problem Solving (Basic)
// - Difficulty: Easy

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

// Full Problem: https://www.hackerrank.com/challenges/staircase/problem

function staircase(n) {
  // Write your code here
<<<<<<< HEAD
<<<<<<< HEAD
  // TODO: answer here
=======
  //beginanswer
  var stair = [];
  var count = 0;
  for (var i = 0; i < n; i++) {
    var horiz = [];

    for (var j = 0; j < n; j++) {
      if (i == count && j >= n - 1 - count) {
        horiz[j] = '#';
      } else {
        horiz[j] = ' ';
      }
    }
    count++;
    stair.push(horiz.join(''));
  }
  // console.log(stair.join('\n'));
  return stair.join('\n')
  //endanswer
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
  // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}

function main() {
  const n = 6

  let result = staircase(n);
  console.log(result)
}

main(); // execute program

module.exports = staircase