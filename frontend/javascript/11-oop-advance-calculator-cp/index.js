/**
 * Kali ini kalian akan mengimplementasikan Method Chaining pada OOP dengan membuat sebuah Advance Calculator.
 * Class Calculator memiliki properti operand
 * Terdapat beberapa method yang perlu kalian buat diantaranya:
 * - `add` : untuk menambahkan operand yang diterima
 * - `subtract` : untuk mengurangi operand yang diterima
 * - `multiply` : untuk mengaklikan operand yang diterima
 * - `divide` : untuk membagi operand yang diterima
 * - `square` : untuk mencari pangkat dari operand yang diterima
 * - `squareRoot` : untuk mencari akar dari operand yang diterima
 * setiap method yang dibuat akan menerima argumen berupa integer
 */

class Calculator {
  constructor (operand) {
    this.operand = operand;
  }
<<<<<<< HEAD
  //beginanswer
  add (angka) {
    this.operand += angka;
    return this
  }
  subtract (angka) {
    this.operand -= angka;
    return this
  }
  multiply (angka) {
    this.operand *= angka;
    return this
  }
  divide (angka) {
    this.operand /= angka;
    return this
  }
  square (angka) {
    this.operand = Math.pow(this.operand, angka)
    return this;
  }
  squareRoot () {
    this.operand = Math.sqrt(this.operand)
    return this
  }
  //endanswer
=======
  // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}

module.exports = Calculator
