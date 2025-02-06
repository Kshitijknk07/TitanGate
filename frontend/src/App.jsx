import React from 'react';
import { motion } from 'framer-motion';
import Header from './components/Header';
import Sidebar from './components/sidebar';
import Home from './pages/Home';

const App = () => {
  return (
    <div className="bg-black text-white min-h-screen">
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 0.6 }}
      >
        <Header />
      </motion.div>

      <div className="flex">
        <motion.div
          className="w-64 bg-gray-800 shadow-lg"
          initial={{ x: -300 }}
          animate={{ x: 0 }}
          transition={{ type: 'spring', stiffness: 100 }}
        >
          <Sidebar />
        </motion.div>

        <motion.main
          className="flex-1 p-6"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ duration: 0.6 }}
        >
          <Home />
        </motion.main>
      </div>
    </div>
  );
};

export default App;
