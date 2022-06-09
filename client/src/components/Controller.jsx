import React from "react";
import { useStateValue } from "../context/StateProvider";
import "./controller.css";
import { actionType } from "../context/reducer";

const Controller = () => {
  const [state, dispatch] = useStateValue();

  const handleClick = (e) => {
    if (state.pop) return;

    dispatch({
      type: actionType.ADD_POP,
      payload: { msg: "Add Product" },
    });
  };

  return (
    <div className="controller">
      <p>
        Your Products are here, you can add, delete and update the product
        detail :
      </p>
      <button onClick={handleClick}>Add Product</button>
    </div>
  );
};

export default Controller;
