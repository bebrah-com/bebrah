import React from "react";
import CloseIcon from "@mui/icons-material/Close";
import FavoriteBorderIcon from "@mui/icons-material/FavoriteBorder";
import ChatOutlinedIcon from '@mui/icons-material/ChatOutlined';
import RemoveRedEyeOutlinedIcon from '@mui/icons-material/RemoveRedEyeOutlined';
import ShareOutlinedIcon from '@mui/icons-material/ShareOutlined';

export default function NFTpage() {
  return (
    <div className="overflow-hidden">
      <nav className="bg-[#9C9C9C] flex justify-end px-3 h-[5.55%] place-items-center">
        <span className="">
          <CloseIcon />
        </span>
      </nav>
      <div className=" w-screen h-[94.45vh] flex justify-around items-center">
        <div className="nft w-[56.4%]">
          <div className="border-[1px] border-[#E6E6E6] w-full h-[37rem] flex justify-center">
            <img
              src="https://images.unsplash.com/photo-1581833971358-2c8b550f87b3?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w0NzEyOTl8MHwxfHNlYXJjaHwzfHxhbmltZXxlbnwwfHx8fDE2OTAzNzY5Nzd8MA&ixlib=rb-4.0.3&q=80&w=1080"
              alt=""
              className=" max-w-lg inline place-self-center"
            />
          </div>
          <button className="w-full font-semibold my-4 border-[1px] border-[#E6E6E6] h-8">
            Buy
          </button>
        </div>
        <div className="profile w-[37.55%] h-[37rem] p-2">
          <div>
            <p className="font-semibold text-3xl">Title</p>
            <p className="text-lg">contents</p>
          </div>
          <div className="flex justify-between w-2/3 mt-10 h-10">
            <button className="border-[1px] border-[#E6E6E6] w-[48%]">
              <FavoriteBorderIcon fontSize="20px" className="mr-2 mb-1" />
              Like
            </button>
            <button className="border-[1px] border-[#E6E6E6] w-[48%]">
              <ChatOutlinedIcon fontSize="20px" className="mr-2 mb-1" />
              Comment
            </button>
          </div>
          <div className="mt-7 flex-col">
            <p>Minted time ago</p>
            <div className="flex justify-between">
              <ul className="list-none flex gap-5 content-center flex-wrap">
                <li><FavoriteBorderIcon fontSize="20px" className="mr-1 mb-1"/>Likes</li>
                <li><RemoveRedEyeOutlinedIcon fontSize="20px" className="mr-1 mb-1"/>Views</li>
                <li><ChatOutlinedIcon fontSize="20px" className="mr-1 mb-1"/>Comments</li>
              </ul>
              <button className="border-[1px] border-[#E6E6E6] w-24 h-10"><ShareOutlinedIcon fontSize="20px" className="mr-1 mb-1"/>Share</button>
            </div>
          </div>
          <div className="mt-4">
            <p className="font-semibold">Creator</p>
            <div className="flex list-none justify-between mt-2">
                <li className="flex gap-2"><img src="" alt="" className="avatar rounded-full w-5 h-5" />Username</li>
                <li>Key</li>
            </div>
          </div>
          <div className="flex-col mt-5 p-4 bg-[#E6E6E6]">
            <div className="flex justify-between">
                <p>Quantity</p>
                <p className="font-semibold">31</p>
            </div>
            <div className="flex justify-between">
                <p>Current Owner</p>
                <p className="font-semibold">2031</p>
            </div>
            <div className="flex justify-between">
                <p>Opening Network</p>
                <p className="font-semibold">Ethereum</p>
            </div>
          </div>
          <div className="mt-4">
            <p><ChatOutlinedIcon fontSize="20px" className="mr-2 mb-1" />
              3 Comment</p>
            <div className="comment-section overflow-y-scroll max-h-32">

            <p className="w-[80%] border-[1px] border-[#E6E6E6] mt-2 p-[10px]">
                <p className="flex gap-2"><img src="" alt="" className="avatar rounded-full w-5 h-5" />Username</p>
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Perspiciatis praesentium quam rem voluptate recusandae hic illo.
            </p>
            <p className="w-[80%] border-[1px] border-[#E6E6E6] mt-2 p-[10px]">
                <p className="flex gap-2"><img src="" alt="" className="avatar rounded-full w-5 h-5" />Username</p>
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Perspiciatis praesentium quam rem voluptate recusandae hic illo.
            </p>
            <p className="w-[80%] border-[1px] border-[#E6E6E6] mt-2 p-[10px]">
                <p className="flex gap-2"><img src="" alt="" className="avatar rounded-full w-5 h-5" />Username</p>
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Perspiciatis praesentium quam rem voluptate recusandae hic illo.
            </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
