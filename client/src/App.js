import "./App.css";
import { Routes, Route } from "react-router-dom";
import { AnimatePresence } from "framer-motion";
import Signin from "./pages/Signin";
import Signup from "./pages/Signup";
import Mainpage from "./pages/Mainpage";
import Notfound from "./components/Notfound";
import { useStateValue } from "./context/StateProvider";

function App() {
  // eslint-disable-next-line
  const [state, dispatch] = useStateValue();
  return (
    <AnimatePresence exitBeforeEnter>
      <Routes>
        <Route
          path="/"
          element={state.user.token !== "NA" ? <Mainpage /> : <Signin />}
        />

        <Route path="/signup" element={<Signup />} />

        <Route path="*" element={<Notfound />} />
      </Routes>
    </AnimatePresence>
  );
}

export default App;
