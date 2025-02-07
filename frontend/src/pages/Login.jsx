import React, { useState } from 'react';
import { login } from '../utils/api';
import { useHistory } from 'react-router-dom';
import { motion } from 'framer-motion';

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
    <div className="flex justify-center items-center min-h-screen bg-black">
      <motion.div
        className="bg-white p-10 rounded-lg shadow-lg w-full max-w-md"
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 0.6 }}
      >
        <motion.h2
          className="text-3xl font-bold text-black mb-6"
          initial={{ y: -50 }}
          animate={{ y: 0 }}
          transition={{ duration: 0.6 }}
        >
          Login
        </motion.h2>

        <motion.form
          onSubmit={handleLogin}
          className="space-y-6"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ duration: 0.8 }}
        >
          <div>
            <label htmlFor="username" className="text-black font-semibold">
              Username
            </label>
            <motion.input
              type="text"
              id="username"
              className="w-full p-3 mt-2 rounded border border-gray-300 focus:outline-none focus:ring-2 focus:ring-black"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              whileFocus={{ scale: 1.05 }}
              transition={{ type: 'spring', stiffness: 200 }}
            />
          </div>

          <div>
            <label htmlFor="password" className="text-black font-semibold">
              Password
            </label>
            <motion.input
              type="password"
              id="password"
              className="w-full p-3 mt-2 rounded border border-gray-300 focus:outline-none focus:ring-2 focus:ring-black"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              whileFocus={{ scale: 1.05 }}
              transition={{ type: 'spring', stiffness: 200 }}
            />
          </div>

          {error && (
            <motion.p
              className="text-red-500 text-sm"
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              transition={{ duration: 0.6 }}
            >
              {error}
            </motion.p>
          )}

          <motion.button
            type="submit"
            className="w-full bg-black text-white p-3 rounded mt-4 hover:bg-gray-900"
            whileHover={{ scale: 1.05 }}
            transition={{ type: 'spring', stiffness: 200 }}
          >
            Login
          </motion.button>
        </motion.form>
      </motion.div>
    </div>
  );
};

export default Login;