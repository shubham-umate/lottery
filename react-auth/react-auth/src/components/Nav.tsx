import React from 'react';
import { Link, Route, Routes } from 'react-router-dom';


const Nav = (props : {name: string, setName : (name:string)=> void}) => {
  console.log("Name in nav", props.name)
  const logout = async () =>{
    await fetch('http://localhost:8000/api/logout',{
      method: 'POST',
      headers: {'Content-Type':'application/json'},
      credentials: 'include'
     
    });

    props.setName('');
    console.log("Hitting logout", props.name)
  
    return (
      <Routes>
        <Route path="/login"/>
      </Routes>
    );
  
  }

  let menu;
  if(props.name === '' || props.name === undefined){
    menu = (
      <ul className="navbar-nav me-auto mb-2 mb-md-0">
      <li className="nav-item">
        <Link to="/login" className="nav-link active" >Login</Link>
      </li>
      <li className="nav-item">
        <Link to="/register" className="nav-link active" >Register</Link>
      </li>
      </ul>
    )
  } else {
    menu = (
      
      <ul className="navbar-nav me-auto mb-2 mb-md-0">
        
        <li className="nav-item">
          <Link to="/login" className="nav-link active" onClick={logout}>Logout {props.name}</Link>
        </li>
      </ul>
    )
  }

  
    return(
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
        <div className="container-fluid">
          <Link to="/" className="navbar-brand" >Home</Link>
         
          <div >
           
            {menu}
           
          </div>
        </div>
      </nav>
    );
};

export default Nav;