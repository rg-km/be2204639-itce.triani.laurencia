// Mengembalikan pola * yang menurun di setiap barisnya
//
// contoh: 
// baris = 5
// hasil:
// *****
// ****
// ***
// **
// *

// Masukan jumlah baris
const n = parseInt(prompt("Masukan jumlah baris: "));

<<<<<<< HEAD
// TODO: answer here
=======
// beginanswer
let stars = "";
for (let i = 0; i < n; i++) {
  for (let k = 0; k < n - i; k++) {
    stars += "*";
  }
  stars += "\n";
}
console.log(stars);
// endanswer
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
