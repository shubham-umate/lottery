import React, { SyntheticEvent, useState } from 'react';
import { Route, Routes } from 'react-router-dom';
import Home from './Home';

const Login = (props: {setName : (name:string)=> void}) => {

  const [email,setEmail] = useState('');
  const [password,setPassword] = useState('');
  const [redirect, setRedirect] = useState(false);
  const [name,setName] = useState('');

  

  const submit = async(e: SyntheticEvent) => {
    
    e.preventDefault()
    
    const response  = await fetch('http://localhost:8000/api/login',{
      method: 'POST',
      headers: {'Content-Type':'application/json'},
      credentials: 'include',
      body: JSON.stringify({
        email,
        password
      })
    });
    
    const content = await response.json();
    setRedirect(true)
    setName(content.name)

    console.log("set name:",name);

    console.log("content name:", content.name);
    
    
    props.setName(content.name);
    console.log("props name:", content.name);
   
    
    
  };

  if(redirect){
    
    return (
      <Routes>
        <Route path="/" element={<Home name={name}/>}/>
      </Routes>
    );
  }; 
  
    return(

      
        <form onSubmit={submit}>
              
        <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
    
       
          <input type="email" className="form-control"  placeholder="Email" required
          onChange={e => setEmail(e.target.value)}
          />
          
       
      
          <input type="password" className="form-control"  placeholder="Password" required
             onChange={e => setPassword(e.target.value)}
          />
          
      
    
       
        <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
       
      </form>
    );
};

export default Login;