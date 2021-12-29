import React from 'react';
import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Dashboard from '../Dashboard/Dashboard';
import Preferences from '../Preferences/Preferences';
import Login from '../Login/Login';
import Register from '../Register/Register';
import useToken from './useToken';

function App() {
    // eslint-disable-next-line no-unused-vars
    const { token, setToken } = useToken();

    // if (!token) {
    //     return <Register setToken={setToken} />
    // }

    return(
    <div className='wrapper'>
        <h1>Application</h1>
        <BrowserRouter>
            <Routes>
                <Route path="/dashboard" element={<Dashboard />} />
                <Route path="/preferences" element={<Preferences />} />
                <Route path="/login" element={<Login setToken={setToken} />} />
                <Route path="/register" element={<Register setToken={setToken} />} />
            </Routes>
        </BrowserRouter>
    </div>
    );
}

export default App;