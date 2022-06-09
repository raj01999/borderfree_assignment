import React from "react";
import "./header.css";
import user from "../utils/user.svg";
import { useStateValue } from "../context/StateProvider";
import { actionType } from "../context/reducer";

const Header = () => {
  const [state, dispatch] = useStateValue();

  return (
    <header>
      <ul>
        <li style={{ fontSize: "1.2rem" }}>Crud Operation</li>
        <li>About</li>
        <li>Contact</li>
      </ul>

      <div className="righSide">
        <div className="userField">
          <img src={user} alt="user" />
          <div>
            <p>
              {state.user.username.split("@")[0][0].toUpperCase() +
                state.user.username.split("@")[0].slice(1, 15).toLowerCase()}
            </p>
            <p className="userType">Normal User</p>
          </div>
        </div>

        <button
          onClick={() => {
            dispatch({ type: actionType.REMOVE_USER });
            dispatch({ type: actionType.ADD_POP, payload: { msg: null } });
          }}
        >
          Logout
        </button>
      </div>
    </header>
  );
};

export default Header;
