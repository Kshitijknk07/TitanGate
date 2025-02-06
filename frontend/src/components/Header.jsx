import React from "react";
import { motion } from "framer-motion";

const Header = () => {
  return (
    <motion.header 
      initial={{ opacity: 0, y: -20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
      className="bg-black text-white p-4 shadow-md"
    >
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">TitanGate</h1>
        <nav className="space-x-6">
          <a href="/" className="text-white hover:text-gray-400">Home</a>
          <a href="/dashboard" className="text-white hover:text-gray-400">Dashboard</a>
          <a href="/login" className="text-white hover:text-gray-400">Login</a>
        </nav>
      </div>
    </motion.header>
  );
};

export default Header;