import React from "react";
import { useStateValue } from "../context/StateProvider";
import "./confirm.css";

const Confirm = ({ Id, setId, fetchData }) => {
  // eslint-disable-next-line
  const [state, dispatch] = useStateValue();

  const handleClick = async () => {
    setId(null);
    const jsonResponse = await fetch(
      process.env.REACT_APP_API + "/deleteproduct?id=" + Id,
      {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          Authorization: state.user.token,
        },
      }
    );

    const response = await jsonResponse.json();

    if (jsonResponse.status === 200) {
      fetchData();
    } else {
      fetchData();
      window.alert(response.message);
    }

    console.log(response);
  };

  return (
    <div className="confirm">
      <p>You want to delete Product: {"s" + Id.slice(-6)} ?</p>
      <div>
        <button
          style={{ backgroundColor: "green" }}
          onClick={(e) => {
            setId(null);
          }}
        >
          No
        </button>
        <button style={{ backgroundColor: "red" }} onClick={handleClick}>
          Yes
        </button>
      </div>
    </div>
  );
};

export default Confirm;
