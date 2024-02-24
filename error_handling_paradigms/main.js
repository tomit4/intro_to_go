function main() {
  try {
    const user = getUser();
    try {
      const profile = getUserProfile(user.ID);
    } catch (err) {
      console.error("something went wrong when getting profile :=>", err);
    }
  } catch (err) {
    console.error("something went wrong when getting the user :=>", err);
  }
}

function getUser() {
  // do some get user logic here
  return user;
}
