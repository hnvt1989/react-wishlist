import React from 'react';
import { Link } from 'react-router-dom';
import GoogleAuth from './GoogleAuth';

const Header = () => {
  return (
    <div className="ui secondary pointing menu">
      <Link to={`/wishes/mywishes`} className="item">
        My Wish Lists
      </Link>
      <div className="right menu">
        <Link to="/" className="item">
          All Wish Lists
        </Link>
        <GoogleAuth />
      </div>
    </div>
  );
};

export default Header;
