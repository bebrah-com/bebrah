import React from 'react'

export default function ImageCards(props) {

  return (
    <div className='object-cover relative'>
      {props.image === 'nft' && <span className='absolute bg-white bg-opacity-70 left-4 top-4 px-1'>NFT</span>}
      <img className='p-2 object-cover' src={props.url} alt="pins" />
    </div>
  )
}
