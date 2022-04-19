<<<<<<< HEAD
// Teman kalian berjanji akan membuatkan kalian kue ulang tahun.
// Kalau semuanya berjalan lancar dan dia tidak sakit, maka kalian akan mendapat kue.

const onMyBirthday = (isMyFriendSick) => {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (!isMyFriendSick) {
        return resolve(true);
      }

      return reject(false);
    }, 2000);
  });
};

onMyBirthday(false)
  .then((_) => console.log("Dapat kue"))
  .catch((_) => console.log("Tidak dapat kue"));

// underscore pada then dan catch menandakan bahwa hasil dari promise itu sendiri tidak dipakai
// contoh dibawah untuk yang dipakai

const onMyNextBirthday = (isMyFriendSick, numberOfCakes) => {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (!isMyFriendSick) {
        return resolve(numberOfCakes);
      }

      return reject(0);
    }, 2000);
  });
};

onMyNextBirthday(true, 3)
  .then((res) =>
    console.log(`Teman saya tidak sakit, dia membuatkan saya ${res} kue`)
  )
  .catch((err) =>
    console.log(`Teman saya sakit, sehingga saya mendapat ${err} kue`)
  );
=======
/**
 * Fungsi getStarWarsData merupakan sebuah fungsi yang mengembalikan promise.
 * Yang mana promise tersebut akan mengembalikan data dari swapi tergantung parameter request yang diberikan.
 *
 * Sebelumnya kita sudah belajar menggunakan cara penggunaan promise. 
 * 
 * Kali ini kita akan melihat bagaimana perbedaan ketika menggunakan async.
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

function getPromisePeopleById(id) {
  return getStarWarsData("https://swapi.dev/api/people/" + id).then(
    (stringResult) => {
      return JSON.parse(stringResult);
    }
  );
}

async function getDataPeopleById(id) {
  const stringResult = await getStarWarsData(
    "https://swapi.dev/api/people/" + id
  );

  return JSON.parse(stringResult);
}

function accessPeoplePromise() {
  getPromisePeopleById(1).then((people) => console.log(people));
}
accessPeoplePromise();

async function accessPeople() {
  const people = await getDataPeopleById(1);
  console.log(people);
}
accessPeople();
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
