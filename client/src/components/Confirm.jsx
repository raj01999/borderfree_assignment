import React from "react";
import { useStateValue } from "../context/StateProvider";
import "./confirm.css";
import { actionType } from "../context/reducer";

const Confirm = ({ Id, setId, fetchData }) => {
  const [state, dispatch] = useStateValue();

  const handleClick = async () => {
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
      setId(null);
    } else {
      dispatch({ type: actionType.REMOVE_USER });
    }

    console.log(response);
  };

  return (
    <div className="confirm">
      <p>Sure you want to delete this Product?</p>
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
