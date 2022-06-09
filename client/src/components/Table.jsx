import React, { useEffect, useState } from "react";
import "./table.css";
import { useStateValue } from "../context/StateProvider";
import Form from "./Form";
import SingleProduct from "./SingleProduct";
import { actionType } from "../context/reducer";
import Confirm from "./Confirm";

const Table = () => {
  const [state, dispatch] = useStateValue();
  const [Id, setId] = useState(null);
  const [prod, setProd] = useState({
    productname: "",
    productdetail: "",
    productprice: "",
  });

  const fetchData = async () => {
    const jsonResponse = await fetch(
      process.env.REACT_APP_API + "/getproduct",
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: state.user.token,
        },
      }
    );

    const response = await jsonResponse.json();

    if (jsonResponse.status === 200) {
      dispatch({
        type: actionType.ADD_PRODUCT,
        payload: { products: response.data },
      });
    } else {
      dispatch({ type: actionType.REMOVE_USER });
    }

    console.log(response);
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div className="tableContainer">
      {state.pop && (
        <Form
          pop={state.pop}
          prod={prod}
          setProd={setProd}
          fetchData={fetchData}
        />
      )}
      {Id && <Confirm Id={Id} setId={setId} fetchData={fetchData} />}
      <table>
        <thead>
          <tr>
            <th className="productId">Product Id</th>
            <th className="productName">Product Name</th>
            <th className="productDescription">Description</th>
            <th className="productPrice">Price</th>
            <th className="productUpdate">Update</th>
            <th className="productDelete">Delete</th>
          </tr>
        </thead>

        <tbody>
          {state.products.map((product) => (
            <SingleProduct
              key={product._id}
              {...product}
              setId={setId}
              setProd={setProd}
            />
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Table;
