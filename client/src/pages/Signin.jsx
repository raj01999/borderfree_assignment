import "./pages.css";
import React, { useState, useRef } from "react";
import bigCircleL from "../utils/bigCircleL.svg";
import dot from "../utils/dot.svg";
import Alert from "../components/Alert";
import { motion } from "framer-motion";
import eye from "../utils/eye.svg";
import { Link } from "react-router-dom";
import bigCircleR from "../utils/bigCircleR.svg";
import { useStateValue } from "../context/StateProvider";
import { actionType } from "../context/reducer";

const Signin = () => {
  const passRef = useRef();
  const [msg, setMsg] = useState(null);
  // eslint-disable-next-line
  const [state, dispatch] = useStateValue();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const data = {
      username: e.target.elements.log_email.value,
      password: e.target.elements.log_password.value,
    };

    const jsonResponse = await fetch(process.env.REACT_APP_API + "/signin", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });

    const response = await jsonResponse.json();

    if (jsonResponse.status === 200) {
      dispatch({
        type: actionType.ADD_USER,
        payload: {
          user: { username: response.username, token: response.token },
        },
      });
    } else {
      setMsg(response.message);
      setTimeout(() => {
        setMsg(null);
      }, 2500);
    }

    console.log(response);
  };

  const showPassword = (e) => {
    var x = passRef.current;
    if (x.type === "password") {
      x.type = "text";
    } else {
      x.type = "password";
    }
  };

  return (
    <section className="loginContainer">
      <img src={bigCircleL} alt="bigCircle" className="bigCircle left" />
      <div className="mainLogIn">
        <img src={dot} alt="dotLeft" className="dotLeft" />

        {msg ? <Alert msg={msg} /> : ""}

        <form action="" className="signForm" onSubmit={handleSubmit}>
          <div className="logo">Logo</div>

          <div className="detail">
            Enter your credentials to access your account
          </div>
          <input
            type="email"
            placeholder="User Id"
            id="log_email"
            className="inputauth"
            required
          />
          <input
            type="password"
            placeholder="Password"
            id="log_password"
            className="inputauth"
            required
            ref={passRef}
          />
          <motion.img
            whileTap={{ scale: 0.95 }}
            src={eye}
            alt="eye"
            onClick={showPassword}
            className="signineye"
          />
          <motion.button
            whileTap={{ scale: 0.95 }}
            type="submit"
            className="btn"
          >
            Sign In
          </motion.button>
          <Link to="/signup" className="linkLogin">
            Sign Up
          </Link>
        </form>

        <img src={dot} alt="dotRight" className="dotRight" />
      </div>
      <img src={bigCircleR} alt="bigCircle" className="bigCircle right" />
    </section>
  );
};

export default Signin;
