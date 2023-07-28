import React, { useState } from "react";
import { Link, useLocation } from "react-router-dom";
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';

export default function Navbar() {
  const { pathname } = useLocation();
  const { signedIn, setSignedIn } = useState(false);

  const [anchorEl, setAnchorEl] = React.useState(null);
  const open = Boolean(anchorEl);
  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <div
      className={`${
        pathname === "/signin" || pathname === "/signup" ? "hidden" : "block"
      }`}
    >
      <div className=" navbar flex justify-between pt-11 px-10 text-lg">
        <span className="font-semibold flex items-center">
          <Link to="/">bebrah</Link>
        </span>
        <div className="flex list-none gap-10 items-center">
          <li className="">Explore</li>
          <li className="font-semibold">
            {signedIn ? (
              <Link to="/signin">Sign In</Link>
            ) : (
              <div className="flex items-center gap-10">
                <div>
                <button
                   id="basic-button"
                   aria-controls={open ? 'basic-menu' : undefined}
                   aria-haspopup="true"
                   aria-expanded={open ? 'true' : undefined}
                   onClick={handleClick}
                    className="bg-[#F5F5F5] p-2 w-[6rem]">Create</button>
                <Menu
                  id="basic-menu"
                  anchorEl={anchorEl}
                  open={open}
                  onClose={handleClose}
                  MenuListProps={{
                    'aria-labelledby': 'basic-button',
                  }}
                >
                  <MenuItem onClick={handleClose}><Link to="/createpost">Create Post</Link></MenuItem>
                  <MenuItem onClick={handleClose}><Link to="/createNFT">Create NFT</Link></MenuItem>
                </Menu> 
                </div>
                <p className="flex gap-2">
                  <img src="" alt="" className="my-auto avatar rounded-full w-5 h-5" />
                  Username
                </p>
              </div>
            )}
          </li>
        </div>
      </div>
    </div>
  );
}
