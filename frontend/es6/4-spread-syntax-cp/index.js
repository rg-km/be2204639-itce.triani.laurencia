/**
 * Buatlah sebuah fungsi yang menerima berapapun jumlah argumen dalam bentuk angka, kemudian mengembalikan nilai maksimal dari semua angka tersebut.
 *
 * max(1, 2, 3, 4, 12, 3, 4) output: 12
 * max(1, 2, 5) output: 5
 * max(1) output: 1
 *
 * Dilarang menggunakan Math.max
 */
const max = (...args) => {
<<<<<<< HEAD
  //beginanswer
  return args.reduce((previous, current) => {
    if (current > previous) {
      return current;
    } else {
      return previous;
    }
  }, 0);
  //endanswer
=======
  // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
};

module.exports = max