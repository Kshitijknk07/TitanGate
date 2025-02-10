import React from "react";
import { motion } from "framer-motion";

const Header = () => {
  return (
    <motion.header
      initial={{ opacity: 0, y: -20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.8, ease: "easeOut", bounce: 0.3 }}
      className="bg-black text-white p-5 shadow-lg border-b border-gray-800"
    >
      <div className="flex justify-between items-center">
        <motion.h1
          whileHover={{ scale: 1.05, transition: { duration: 0.3 } }}
          className="text-3xl font-extrabold tracking-wide"
        >
          TitanGate
        </motion.h1>
        <nav className="space-x-6">
          {["Home", "Dashboard", "Login"].map((item, index) => (
            <motion.a
              key={index}
              href={`/${item.toLowerCase()}`}
              whileHover={{ scale: 1.1, color: "#9ca3af" }}
              transition={{ duration: 0.3 }}
              className="text-white text-lg font-medium"
            >
              {item}
            </motion.a>
          ))}
        </nav>
      </div>
    </motion.header>
  );
};

export default Header;
