// Birthday Cake Candles
// - Problem Solving (Basic)
// - Difficulty: Easy

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

// Full Problem: https://www.hackerrank.com/challenges/birthday-cake-candles/problem

function birthdayCakeCandles(candles) {
  // Write your code here
<<<<<<< HEAD
<<<<<<< HEAD
	// TODO: answer here
=======
	//beginanswer
  var blownOut = 0;
  var tallest = Math.max.apply(null, candles)
  
  for (var i = 0; i < candles.length; i++) {
      if (candles[i] == tallest) blownOut++;
  }

  return blownOut;
	//endanswer
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
	// TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}

function main() {
  //var n = parseInt(readLine());
  //ar = readLine().split(' ');
  //ar = ar.map(Number);
  var ar = [3, 2, 1, 3]; // override input
  var result = birthdayCakeCandles(ar);
  console.log(result);

}

main(); // execute program

module.exports = birthdayCakeCandles
