import React from "react";
import Header from "../components/Header";
import Controller from "../components/Controller";
import Table from "../components/Table";

const Mainpage = () => {
  return (
    <div className="mainContainer">
      <Header />
      <Controller />
      <Table />
    </div>
  );
};

export default Mainpage;
