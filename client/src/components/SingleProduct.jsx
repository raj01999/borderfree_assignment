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

    setProd({
      id: _id,
      productname: productname,
      productdetail: productdetail,
      productprice: productprice,
    });
    dispatch({ type: actionType.ADD_POP, payload: { msg: "Update Product" } });
  };

  const deleteClick = (e) => {
    setId(_id);

    setProd({
      productname: "",
      productdetail: "",
      productprice: "",
    });

    dispatch({ type: actionType.ADD_POP, payload: { msg: null } });
  };

  return (
    <tr>
      <td className="productId">{"s" + _id.slice(-5)}</td>
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
