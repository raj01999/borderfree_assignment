import React from "react";
import { useStateValue } from "../context/StateProvider";
import { actionType } from "../context/reducer";

const SingleProduct = ({
  _id,
  productname,
  productdetail,
  productprice,
  setId,
  setProd,
}) => {
  const [state, dispatch] = useStateValue();

  const updateClick = (e) => {
    if (state.pop) return;

    setId(null);

    setProd({
      id: _id,
      productname: productname,
      productdetail: productdetail,
      productprice: productprice,
    });
    dispatch({
      type: actionType.ADD_POP,
      payload: { msg: `Product: s${_id.slice(-6)}` },
    });
  };

  const deleteClick = (e) => {
    dispatch({ type: actionType.ADD_POP, payload: { msg: null } });

    setId(_id);

    setProd({
      productname: "",
      productdetail: "",
      productprice: "",
    });
  };

  return (
    <tr>
      <td className="productId">{"s" + _id.slice(-6)}</td>
      <td className="productName">{productname}</td>
      <td className="productDescription">{productdetail}</td>
      <td className="productPrice">{productprice}</td>
      <td className="productUpdate">
        <button style={{ backgroundColor: "green" }} onClick={updateClick}>
          Update
        </button>
      </td>
      <td className="productDelete">
        <button style={{ backgroundColor: "red" }} onClick={deleteClick}>
          Delete
        </button>
      </td>
    </tr>
  );
};

export default SingleProduct;
