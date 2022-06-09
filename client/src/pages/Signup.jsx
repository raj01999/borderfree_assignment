import React, { useState, useRef } from "react";
import bigCircleL from "../utils/bigCircleL.svg";
import bigCircleR from "../utils/bigCircleR.svg";
import eye from "../utils/eye.svg";
import dot from "../utils/dot.svg";
import Alert from "../components/Alert";
import { motion } from "framer-motion";
import { Link, useNavigate } from "react-router-dom";

const Signup = () => {
  const navigate = useNavigate();
  const passRef = useRef();
  const [msg, setMsg] = useState(null);
  const [wrongPass, setWrongPass] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (
      e.target.elements.reg_password.value !==
      e.target.elements.c_reg_password.value
    ) {
      setTimeout(() => {
        setWrongPass(false);
      }, 2500);
      return setWrongPass(true);
    }

    const data = {
      username: e.target.elements.reg_email.value,
      password: e.target.elements.reg_password.value,
    };

    const jsonResponse = await fetch(process.env.REACT_APP_API + "/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    const response = await jsonResponse.json();
    if (jsonResponse.status === 200) {
      navigate("/");
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
        {msg && <Alert msg={msg} />}
        {wrongPass && <Alert msg={"Password does't match"} />}

        <img src={dot} alt="dotLeft" className="dotLeft" />

        <form action="" className="signForm" onSubmit={handleSubmit}>
          <div className="logo">Logo</div>

          <div className="detail">Create New Account</div>
          <input
            type="email"
            placeholder="User Id"
            id="reg_email"
            className="inputauth"
            required
          />
          <input
            type="password"
            placeholder="Password"
            id="reg_password"
            className="inputauth"
            required
            ref={passRef}
          />
          <input
            type="password"
            id="c_reg_password"
            placeholder="Confirm Password"
            className="inputauth"
            required
          />
          <motion.img
            whileTap={{ scale: 0.95 }}
            src={eye}
            alt="eye"
            onClick={showPassword}
            className="signupeye"
          />
          <motion.button
            whileTap={{ scale: 0.95 }}
            type="submit"
            className="btn"
          >
            Sign Up
          </motion.button>
          <Link to="/" className="linkLogin">
            Already Have Account ?
          </Link>
        </form>

        <img src={dot} alt="dotRight" className="dotRight" />
      </div>
      <img src={bigCircleR} alt="bigCircle" className="bigCircle right" />
    </section>
  );
};

export default Signup;
