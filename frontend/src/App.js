import "./App.css";
import Homepage from "./Components/Homepage";
import Signin from "./Components/Signin";
import Signup from "./Components/Signup";
import NFTpage from "./Components/NFTpage";
import Artpage from "./Components/Artpage";
import Navbar from "./Components/Navbar";
import Createpost from "./Components/Createpost";
import Createnft from "./Components/Createnft";
import { BrowserRouter, Route, Routes, useLocation } from "react-router-dom";

function App() {
  return (
    <>
      <BrowserRouter>
          <Navbar/>
        <Routes>
          <Route path="/" element={<Homepage />} />
          <Route path="/signin" element={<Signin />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/createpost" element={<Createpost/>} />
          <Route path="/createNFT" element={<Createnft/>} />
        </Routes>
      </BrowserRouter>


      {/* <NFTpage/>
     <Artpage/>
    <Createpost/>
    <Createnft/> */}
    </>
  );
}

export default App;
