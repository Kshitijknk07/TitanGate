import React, { useState } from 'react';
import { login } from '../utils/api';
import { useHistory } from 'react-router-dom';

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const history = useHistory();

  const handleLogin = async (e) => {
    e.preventDefault();
    const success = await login(username, password);
    if (success) {
      history.push('/dashboard'); // Redirect to dashboard
    } else {
      setError('Invalid credentials');
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-900">
      <div className="bg-gray-800 p-8 rounded-lg shadow-lg w-full max-w-md">
        <h2 className="text-3xl font-bold text-white mb-4">Login</h2>
        <form onSubmit={handleLogin} className="space-y-6">
          <div>
            <label htmlFor="username" className="text-white">Username</label>
            <input
              type="text"
              id="username"
              className="w-full p-2 mt-2 rounded bg-gray-700 text-white"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <div>
            <label htmlFor="password" className="text-white">Password</label>
            <input
              type="password"
              id="password"
              className="w-full p-2 mt-2 rounded bg-gray-700 text-white"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          {error && <p className="text-red-500">{error}</p>}
          <button type="submit" className="w-full bg-white text-black p-2 rounded mt-4 hover:bg-gray-200">Login</button>
        </form>
      </div>
    </div>
  );
};

export default Login;
