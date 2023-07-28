import React, { useState } from "react";
import KeyboardArrowDownIcon from "@mui/icons-material/KeyboardArrowDown";
import Navmenu from "./Navmenu";
import axios from "axios";
import ImageCards from "./ImageCards";
import { Link } from "react-router-dom";

export default function Homepage() {
  const [images, setImages] = useState([]);

  axios
    .get(
      //api here
    )
    .then(function (response) {
      setImages(response.data.results);       //to set images to the url of
      console.log(response);
    });

  return (
    <div>
      <div className="mt-12 px-10">
        <h1 className="font-bold text-3xl">Explore AI integrated art!</h1>
        <p className="py-4 text-sm w-5/12">
          Lorem ipsum, dolor sit amet consectetur adipisicing elit. Eveniet
          quaerat deleniti consequatur numquam enim, delectus reiciendis impedit
          officiis perspiciatis id fugit praesentium nulla doloribus facilis
          laudantium optio itaque, totam consequuntur.
        </p>
      </div>
      <div className="flex px-10 justify-between mt-5">
        <div className="drop-down flex list-none gap-3">
          <li className=" border-[1px] border-black px-3 py-1 transition-all hover:bg-black hover:text-white hover:border-white hover:ease-in-out hover:cursor-pointer">
            <btn>
              Popular
              <KeyboardArrowDownIcon />
            </btn>
          </li>
          <li className="border-[1px] border-black px-3 py-1 transition-all hover:bg-black hover:text-white hover:border-white hover:ease-in-out hover:cursor-pointer">
            <btn>
              Filter
              <KeyboardArrowDownIcon />
            </btn>
          </li>
        </div>
        <Navmenu />
      </div>
      <div className="px-10 py-4 flex justify-center w-screen">
        <div className="container columns-6 gap-0 max-w-8xl">
          {images?.map((image) => {
            return <Link to='/image'><ImageCards key={image.id} url={image.urls.regular} /></Link>; //put some variable you used to differentitate between nft and image
          })}
        </div>
      </div>
    </div>
  );
}
