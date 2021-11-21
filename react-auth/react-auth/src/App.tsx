import React, { useEffect, useState } from 'react';
import './App.css';

import { BrowserRouter, Route, Routes } from 'react-router-dom';

import Nav from './components/Nav';
import Home from './pages/Home';
import Register from './pages/Register';
import Login from './pages/Login';

function App() {

  const[name,setName] = useState('');
  useEffect(()=>{ 
      (
          async() => {
             const response =  await fetch('http://localhost:8000/api/user',{
                  method: 'GET',
                  headers: {'Content-Type':'application/json'},
                  credentials: 'include'
              });

              const content = await response.json();
              
              console.log("Name in App", content.name)
              setName(content.name)
          }
          
      )();
  });
  
  return (


    <div className="App">
      <BrowserRouter>
       
      <Nav name={name} setName={setName}/>
      
      <main className="form-signin">
      
      <Routes>
        <Route path="/" element={ <Home name={name}/>}/>
        <Route path ="/register" element={<Register/>}/>
        <Route path ="/login/*" element={<Login setName={setName}/>}/>
      </Routes>
      
        
      </main>
      </BrowserRouter>
    </div>
    
  );
}

export default App;
