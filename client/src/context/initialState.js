const getUser = () => {
  return localStorage.getItem("user")
    ? JSON.parse(localStorage.getItem("user"))
    : { username: "NA", token: "NA" };
};

const initialState = {
  user: getUser(),
  pop: null,
  products: [],
};

export default initialState;
