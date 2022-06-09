import React from "react";
import "./form.css";
import { useStateValue } from "../context/StateProvider";
import { actionType } from "../context/reducer";

const Form = ({ pop, prod, setProd, fetchData }) => {
  const [state, dispatch] = useStateValue();
  const handleSubmit = async (e) => {
    e.preventDefault();

    if (prod.id) {
      const jsonResponse = await fetch(
        process.env.REACT_APP_API + "/updateProduct?id=" + prod.id,
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
            Authorization: state.user.token,
          },
          body: JSON.stringify({
            productname: prod.productname,
            productdetail: prod.productdetail,
            productprice: prod.productprice,
          }),
        }
      );

      const response = await jsonResponse.json();

      if (jsonResponse.status === 200) {
        fetchData();
      } else {
        dispatch({ type: actionType.REMOVE_USER });
      }

      setProd({
        productname: "",
        productdetail: "",
        productprice: "",
      });

      dispatch({ type: actionType.ADD_POP, payload: { msg: null } });

      console.log(response);

      return;
    }

    const jsonResponse = await fetch(
      process.env.REACT_APP_API + "/addproduct",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: state.user.token,
        },
        body: JSON.stringify({
          productname: prod.productname,
          productdetail: prod.productdetail,
          productprice: prod.productprice,
        }),
      }
    );

    const response = await jsonResponse.json();

    if (jsonResponse.status === 200) {
      fetchData();
    } else {
      dispatch({ type: actionType.REMOVE_USER });
    }

    setProd({
      productname: "",
      productdetail: "",
      productprice: "",
    });

    dispatch({ type: actionType.ADD_POP, payload: { msg: null } });

    console.log(response);
  };

  const handleClick = (e) => {
    setProd({
      productname: "",
      productdetail: "",
      productprice: "",
    });

    dispatch({ type: actionType.ADD_POP, payload: { msg: null } });
  };

  return (
    <div className="formContainer">
      <form autoComplete="off" onSubmit={handleSubmit}>
        <h3>{pop}</h3>
        <p className="cross" onClick={handleClick}>
          x
        </p>
        <input
          type="text"
          placeholder="Product Name"
          required
          value={prod.productname}
          onChange={(e) => {
            setProd({ ...prod, productname: e.target.value });
          }}
          id="name"
        />
        <input
          type="text"
          placeholder="Product Description"
          required
          value={prod.productdetail}
          onChange={(e) => {
            setProd({ ...prod, productdetail: e.target.value });
          }}
          id="description"
        />
        <input
          type="number"
          placeholder="Product Price"
          required
          value={prod.productprice}
          onChange={(e) => {
            setProd({ ...prod, productprice: e.target.value });
          }}
          id="price"
        />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default Form;
