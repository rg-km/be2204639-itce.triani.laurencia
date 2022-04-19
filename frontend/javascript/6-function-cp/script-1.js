// Mengembalikan teks yang berkebalikan
//
// contoh: 
// teks = hello world
// hasil:
// dlrow olleh

function reverseString(str) {
<<<<<<< HEAD
<<<<<<< HEAD
    // TODO: answer here
=======
    // beginanswer
    let newString = "";
    for (let i = str.length - 1; i >= 0; i--) {
        newString += str[i];
    }
    return newString;
    // endanswer
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
    // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}

const string = prompt('Masukan teks: ');

const result = reverseString(string);
console.log(result);