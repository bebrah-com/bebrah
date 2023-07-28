import React, {useState} from 'react'

export default function Navmenu() {
    const [activeMenuItem, setActiveMenuItem] = useState('All')

    const changeToAll=()=>{
        setActiveMenuItem('All')
        console.log(activeMenuItem);
    }
    const changeToArt=()=>{
        setActiveMenuItem('Art')
        console.log(activeMenuItem);
    }
    const changeToGaming=()=>{
        setActiveMenuItem('Gaming')
        console.log(activeMenuItem);
    }
    const changeToGraphics=()=>{
        setActiveMenuItem('Graphics')
        console.log(activeMenuItem);
    }
    const changeToPhoto=()=>{
        setActiveMenuItem('Photo')
        console.log(activeMenuItem);
    }

  return (
    <div className='flex list-none gap-3'>
      <li className={`hover:cursor-pointer ${activeMenuItem==='All'?'active-nav-menu-item':'nav-menu-item'}`} onClick={changeToAll}>All</li>
      <li className={`hover:cursor-pointer ${activeMenuItem==='Art'?'active-nav-menu-item':'nav-menu-item'}`} onClick={changeToArt}>Art</li>
      <li className={`hover:cursor-pointer ${activeMenuItem==='Gaming'?'active-nav-menu-item':'nav-menu-item'}`} onClick={changeToGaming}>Gaming</li>
      <li className={`hover:cursor-pointer ${activeMenuItem==='Graphics'?'active-nav-menu-item':'nav-menu-item'}`} onClick={changeToGraphics}>Graphics</li>
      <li className={`hover:cursor-pointer ${activeMenuItem==='Photo'?'active-nav-menu-item':'nav-menu-item'}`} onClick={changeToPhoto}>Photo</li>
    </div>
  )
}
