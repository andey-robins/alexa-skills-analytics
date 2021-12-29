import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { useNavigate } from 'react-router-dom';
import './Register.css'

async function registerUser(credentials) {
    return fetch('http://localhost:8080/register', {
        method: 'Post',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
    })
        .then(data => data.json())
}

export default function Register({ setToken }) {
    const [username, setUsername] = useState();
    const [password, setPassword] = useState();
    const navigate = useNavigate();

    const handleSubmit = async e => {
        e.preventDefault();
        const token = await registerUser({
            username,
            password
        });
        setToken(token);
    }

    return(
        <div className='login-wrapper'>
            <h1>Please Register Below</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    <p>Username</p>
                    <input type='text' onChange={e => setUsername(e.target.value)}/>
                </label>
                <label>
                    <p>Password</p>
                    <input type='password' onChange={e => setPassword(e.target.value)} />
                </label>
                <div>
                    <button type='submit'>Submit</button>
                </div>
            </form>
            <button onClick={() => { navigate('/login'); }}>Login</button>
        </div>
    )
}

Register.propTypes = {
    setToken: PropTypes.func.isRequired
}