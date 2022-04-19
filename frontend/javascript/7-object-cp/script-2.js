// Buat kalkulator objek dengan 5 methods:
// - read() meminta (prompt) dua nilai dan menyimpannya sebagai properti objek.
// - sum() mengembalikan jumlah nilai yang disimpan.
// - substract() mengurangi jumlah nilai yang disimpan.
// - multiply() mengalikan nilai yang disimpan.
// - division() membagi nilai yang disimpan.

let calculator = {
  // Tulis kode di sini
<<<<<<< HEAD
<<<<<<< HEAD
  // TODO: answer here
=======
  // beginanswer
  sum() {
    return this.a + this.b;
  },

  substract() {
    return this.a - this.b;
  },

  multiply() {
    return this.a * this.b;
  },

  division() {
    return this.a / this.b;
  },

  read() {
    this.a = +prompt('a =', 0);
    this.b = +prompt('b =', 0);
  }
  // endanswer
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
  // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
};
  
calculator.read();
console.log(calculator.sum());
console.log(calculator.substract());
console.log(calculator.multiply());
console.log(calculator.division());