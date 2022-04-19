/**
 * Carilah angka palindrome yang LEBIH BESAR dan TERDEKAT dari angka-angka dibawah ini
 * 
 * Misal:
 * - Jika angka pada parameter bernilai 90 maka angka palindrome terdekat adalah 99
 * - Jika angka pada parameter bernilai 102 maka angka palindrome terdekat adalah 111
 */

function angkaPalindrome(num) {
<<<<<<< HEAD
<<<<<<< HEAD
  // TODO: answer here
=======
  //beginanswer
  var cek = true;
  while (cek) {
    num++;
    var numStr = num.toString();
    var hasil = '';
    for (var j = numStr.length - 1; j >= 0; j--) {
      hasil += numStr[j];
    }
    if (hasil === numStr) {
      return num;
    }
  }
  //endanswer
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
  // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}

console.log(angkaPalindrome(10)); //11
console.log(angkaPalindrome(17)); //22
console.log(angkaPalindrome(175)); //181

module.exports = angkaPalindrome