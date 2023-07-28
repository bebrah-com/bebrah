import React from "react";
import twitter from "../Components/Icons/twitter.png";
import google from "../Components/Icons/Google.png";
import background from "../Components/Images/BgSignIn.png";
import { Link } from 'react-router-dom'

export default function Signin() {
  return (
    <div>
      <span className="w-52 h-24 inline-block pl-24 pt-8 text-2xl font-semibold">
        bebrah
      </span>
      <div className="w-[31.25%] h-24  pl-24 translate-y-24">
        <span className=" inline-block font-bold text-2xl">
          Sign in and create, improve and work with AI!
        </span>
        <form className="mt-3 font-semibold">
          <label htmlFor="email" className="mt-1">
            Email
          </label>
          <div className="mb-2">
            <input
              type="text"
              placeholder="Enter your email"
              className="border-[1px] w-full  border-black py-4 pl-3"
            />
          </div>
          <label htmlFor="password" className="my-1">
            Password
          </label>
          <div className="mb-2">
            <input
              type="text"
              placeholder="Enter your password"
              className="border-[1px] w-full border-black py-4 pl-3"
            />
          </div>
          <button
            type="submit"
            className="w-full border-[1px] border-[#676767] mt-5 py-4 bg-[#F5F5F5]"
          >
            Sign In
          </button>
        </form>
        <p className="mt-2">
          Don’t have an account? <span className="font-semibold"><Link to="/signup">Sign Up!</Link></span>
        </p>
        <div className="flex justify-between mt-5 font-semibold text-[14px]">
          <button className="w-[49%] border-[1px] border-[#676767] bg-[#F5F5F5] py-3">
            <span className="py-[1px]">Sign In with Twitter</span>
            <img src={twitter} alt="" className="w-[38px] h-[28px] inline-block pl-2" />
          </button>
          <button className="w-[49%] border-[1px] border-[#676767] bg-[#F5F5F5] py-3">
            Sign In with Google
            <img src={google} alt="" className="w-[38px] h-[28px] inline-block pl-2" />
          </button>
        </div>
      </div>
      <img src={background} alt="" className="absolute top-0 right-0 w-[55.5%]"/>
    </div>
  );
}
