export const actionType = {
  ADD_USER: "add user",
  REMOVE_USER: "remove contact",
  ADD_POP: "add pop header message",
  ADD_PRODUCT: "add product to product array",
};

const reducer = (state, action) => {
  switch (action.type) {
    case actionType.ADD_USER:
      localStorage.setItem("user", JSON.stringify(action.payload.user));
      return {
        ...state,
        user: action.payload.user,
      };

    case actionType.REMOVE_USER:
      localStorage.clear();
      return {
        ...state,
        user: { username: "NA", token: "NA" },
      };

    case actionType.ADD_POP:
      return {
        ...state,
        pop: action.payload.msg,
      };

    case actionType.ADD_PRODUCT:
      return {
        ...state,
        products: action.payload.products,
      };

    default:
      return state;
  }
};

export default reducer;
