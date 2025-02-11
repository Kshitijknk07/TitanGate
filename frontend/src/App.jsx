import React from 'react';
import { motion } from 'framer-motion';
import Header from './components/Header';
import Sidebar from './components/sidebar';
import Home from './pages/Home';

const App = () => {
  return (
    <div className="bg-black text-white min-h-screen font-sans">
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 0.6 }}
      >
        <Header />
      </motion.div>

      <div className="flex">
        <motion.div
          className="w-72 bg-gray-900 shadow-xl p-4"
          initial={{ x: -300 }}
          animate={{ x: 0 }}
          transition={{ type: 'spring', stiffness: 100 }}
        >
          <Sidebar />
        </motion.div>

        <motion.main
          className="flex-1 p-8 bg-gray-950 rounded-lg shadow-md mx-4 mt-4"
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
