import React from "react";
import { motion } from "framer-motion";

const Sidebar = () => {
  return (
    <motion.div
      initial={{ opacity: 0, x: -20 }}
      animate={{ opacity: 1, x: 0 }}
      transition={{ duration: 0.5 }}
      className="bg-black text-white w-64 h-screen p-4 border-r border-gray-700"
    >
      <h2 className="text-2xl font-bold mb-6">Navigation</h2>
      <ul>
        <li className="mb-4 hover:text-gray-400"><a href="#">API Analytics</a></li>
        <li className="mb-4 hover:text-gray-400"><a href="#">Load Balancing</a></li>
        <li className="mb-4 hover:text-gray-400"><a href="#">GraphQL Gateway</a></li>
      </ul>
    </motion.div>
  );
};

export default Sidebar;
