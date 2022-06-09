import React from "react";
import "./notfound.css";
import { Link } from "react-router-dom";

const Notfound = () => {
  return (
    <div className="notFound">
      <p>404 Page Not Found</p>
      <Link to="/">Go Back</Link>
    </div>
  );
};

export default Notfound;
