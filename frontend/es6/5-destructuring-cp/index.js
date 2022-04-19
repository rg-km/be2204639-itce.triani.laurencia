// format nama user menjadi "John - john@example.com", dan akses nama user dengan teknik destrucuring assignment.

const format = (user) => {
<<<<<<< HEAD
  //beginanswer
  const { name, email } = user;
  return `${name} - ${email}`;
  //endanswer
=======
  // TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
};

console.log(format({ name: "John", email: "john@example.com" }))

module.exports = format