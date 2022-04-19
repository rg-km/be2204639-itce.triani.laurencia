/** Akses user name, namun dengan nama variabel yang berbeda dari atribut pada objek user
 */
const getUserFirstName = (user) => {
<<<<<<< HEAD
  //beginanswer
  const { name: firstName } = user;
  return firstName;
  //endanswer
=======
  // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
};

console.log(getUserFirstName({ name: "John", email: "john@example.com" }));

module.exports = getUserFirstName