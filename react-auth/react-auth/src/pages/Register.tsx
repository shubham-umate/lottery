import React, { SyntheticEvent, useState } from 'react';
import { Navigate, Route, Routes } from 'react-router-dom';


const Register = () => {

  const [name,setName] = useState('');
  const [email,setEmail] = useState('');
  const [password,setPassword] = useState('');
  const [redirect, setRedirect] = useState(false);

  const submit = async(e: SyntheticEvent) => {
    e.preventDefault()
    
    await fetch('http://localhost:8000/api/register',{
      method: 'POST',
      headers: {'Content-Type':'application/json'},
      body: JSON.stringify({
        name,
        email,
        password
      })
    });
    setRedirect(true);
    
    
  };

if(redirect){
  return (
    <Routes>

      <Route path="/" element={<Navigate replace to="/login" />} />
     

    </Routes>
  );
}
 

    return(
        <form onSubmit={submit}>
            
           
        <h1 className="h3 mb-3 fw-normal">Please register</h1>
    
       
          <input  type="name" className="form-control"  placeholder="Name" required 
          onChange={e => setName(e.target.value)}
          />
         
          <input type="email" className="form-control"  placeholder="Email" required
             onChange={e => setEmail(e.target.value)}
          />
       
      
          <input type="password" className="form-control"  placeholder="Password" required
             onChange={e => setPassword(e.target.value)}
          />
       
        <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
       
      </form>
        
    );
};

export default Register;