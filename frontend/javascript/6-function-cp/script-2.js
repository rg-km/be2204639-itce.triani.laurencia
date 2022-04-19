// Mengembalikan boolean untuk mengecek apakah teks palindrom
//
// contoh: 
// teks1 = hello
// teks2 = madam
// teks3 = kasur ini rusak
// hasil:
// teks1 dibalik menjadi olleh, maka balikan akan false
// teks2 dibalik sama menjadi madam, maka balikan akan true
// teks3 dibalik sama menjadi kasur ini rusak, maka balikan akan true

function checkPalindrome(string) {
<<<<<<< HEAD
<<<<<<< HEAD
    // TODO: answer here
=======
    // beginanswer
    const len = string.length;

    for (let i = 0; i < len / 2; i++) {
        if (string[i] !== string[len - 1 - i]) {
            return 'Bukan palindrom';
        }
    }
    return 'Palindrom';
    // endanswer
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
    // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}

// masukan teks
const string = prompt('Masukan string: ');

// memanggil fungsi palindrom
const value = checkPalindrome(string);

console.log(value);