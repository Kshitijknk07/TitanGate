import React from "react";
import { motion } from "framer-motion";

const Sidebar = () => {
  return (
    <motion.div
      initial={{ opacity: 0, x: -20 }}
      animate={{ opacity: 1, x: 0 }}
      transition={{ duration: 0.8, ease: "easeOut", bounce: 0.3 }}
      className="bg-black text-white w-64 h-screen p-6 border-r border-gray-800 shadow-xl"
    >
      <motion.h2
        whileHover={{ scale: 1.05, transition: { duration: 0.3 } }}
        className="text-2xl font-extrabold mb-6 tracking-wide"
      >
        Navigation
      </motion.h2>
      <ul className="space-y-4">
        {["API Analytics", "Load Balancing", "GraphQL Gateway"].map((item, index) => (
          <motion.li
            key={index}
            whileHover={{ scale: 1.05, color: "#9ca3af" }}
            transition={{ duration: 0.3 }}
            className="text-lg font-medium cursor-pointer"
          >
            <a href="#">{item}</a>
          </motion.li>
        ))}
      </ul>
    </motion.div>
  );
};

export default Sidebar;
