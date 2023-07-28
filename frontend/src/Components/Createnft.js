import React from "react";
import defaultPic from "./Images/Defaultpic.png";

export default function Createnft() {
  return (
    <div>
      <div className="w-[50%] mx-auto mt-2">
        <h1 className="text-4xl font-bold my-7">Create new NFT</h1>
        <p className="text-xl font-semibold">Upload Image</p>
        <p className="mt-1 text-sm">
          File types supported: JPG, PNG, GIF, SVG, MP4, WEBM, MP3, WAV, OGG,
          GLB, GLTF. Max size: 100 MB
        </p>
        <div className="w-[26.14vw] h-[28.24vh] bg-[#D0D0D0] mt-3 flex justify-center items-center">
          <img src={defaultPic} alt="" />
        </div>
        <div className="mt-3 w-[60%]">
          <p className="text-xl font-semibold">Name</p>
          <p className="mt-1 text-sm">Name for your digital asset.</p>
          <input
            type="text"
            placeholder="Name"
            className="mt-3 border-[1px] border-[#676767] w-full px-3 py-4 placeholder:text-black"
          />
        </div>
        <div className="mt-3 w-[60%]">
          <p className="text-xl font-semibold">Description</p>
          <p className="mt-1 text-sm">Additional info for your digital asset.</p>
          <textarea
            placeholder="Description"
            className="mt-3 border-[1px] border-[#676767] w-full px-3 py-4 h-40 placeholder:text-black"
          />
        </div>
        <button className="border-[1px] border-[#676767] bg-[#F5F5F5] my-6 w-32 py-3">Create NFT</button>
      </div>
    </div>
  );
}
