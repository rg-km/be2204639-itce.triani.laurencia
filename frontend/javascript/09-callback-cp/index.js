// Buatlah callback yang memenuhi function call di line 29

// function ini menduplikasi tiap karakter pada sebuah string
// contoh: hehe => hheehhee
function doubleChars(str) {
<<<<<<< HEAD
  // kerjakan disini
  //beginanswer
  return str
    .split("")
    .map((c) => `${c}${c}`)
    .join("");
  //endanswer
=======
  // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}

// function ini mengulang pengaplikasian callback pada str sejumlah num
function repeat(str, num, cb) {
<<<<<<< HEAD
  // kerjakan disini
  //beginanswer
  let result = str;

  for (i = 0; i < num; i++) {
    result = cb(result);
  }

  return result;
  //endanswer
}

console.log(repeat("triple", 2, doubleChars)); // ttttrrrriiiipppplllleeee
=======
  // TODO: answer here
}

console.log(repeat("triple", 2, doubleChars)); //  ttttrrrriiiipppplllleeee

module.exports = {
  doubleChars,
  repeat
}
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
