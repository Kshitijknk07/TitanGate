import React from "react";
import { motion } from "framer-motion";

const Footer = () => {
  return (
    <motion.footer 
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
      className="bg-black text-white py-4 text-center border-t border-gray-700"
    >
      <p className="text-gray-400 text-sm">&copy; {new Date().getFullYear()} TitanGate. All rights reserved.</p>
    </motion.footer>
  );
};

export default Footer;
