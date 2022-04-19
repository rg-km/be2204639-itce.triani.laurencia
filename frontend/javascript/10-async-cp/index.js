<<<<<<< HEAD
// lakukan call pada API ini https://jsonplaceholder.typicode.com/todos/1
// dan console.log hasilnya menggunakan async await
// console log error apabila terjadi error

// karena ini adalah environtment node-js, fetch tidak bisa digunakan
// tuliskan program kalian dibawah, dan test kode tersebut di
// https://playcode.io/new/

const callApi = async () => {
  // kerjakan disini
  //beginanswer
  try {
    const res = await (
      await fetch("https://jsonplaceholder.typicode.com/todos/1")
    ).json();

    console.log(res);
  } catch (error) {
    console.log(error);
  } finally {
    console.log("API selesai dipanggil");
  }
  //endanswer
};

callApi();
=======
/**
 * Buatlah fungsi async yang tugasnya mengambil data karakter starwars dan mengembalikannya dalam bentuk string sebagai berikut:
 * 
 * Input: 1 // Karakter Id.
 * Output: "Luke Skywalker, memiliki tinggi 175cm dan lahir pada tahun 19BBY"
 * 
 * Untuk informasi yang lebih lengkap mengenai endpoint yang digunakan, silahkah akses https://swapi.dev/documentation#people
 * 
 * Gunakan keyword async/await untuk menyelesaikan fungsi ini.
 */

/**
 * Gunakan fungsi berikut untuk mempermudah http request.
 */
 const https = require("https");

 function getStarWarsData(url) {
   return new Promise((resolve, reject) => {
     https
       .get(url, (res) => {
         let result = "";
 
         if (res.statusCode !== 200) {
           reject(new Error(res.statusCode));
         }
 
         res.on("data", (d) => {
           result += d;
         });
 
         res.on("end", () => {
           resolve(result);
         });
       })
       .on("error", (e) => {
         reject(e);
       });
   });
 }
 
 
 async function getDataPeopleById(id) {
   // TODO: answer here
 }

 module.exports = {
  getDataPeopleById
 }
 
 
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
