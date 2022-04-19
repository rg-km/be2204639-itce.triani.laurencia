/** Format nama user menjadi "John - john@example.com - Unemployed",
 *  dan akses data user dengan teknik destrucuring assignment
 *  gunakan default value ketika atribut tidak ditemukan
 */

const format = (user) => {
<<<<<<< HEAD
  //beginanswer
  const { name, email, job = "Unemployed" } = user;
  return `${name} - ${email} - ${job}`;
  //endanswer
=======
  // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
};

console.log(format({ name: "John", email: "john@example.com" }));

module.exports = format