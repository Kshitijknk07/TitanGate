import React from "react";
import { motion } from "framer-motion";

const Footer = () => {
  return (
    <motion.footer
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.8, ease: "easeOut", bounce: 0.3 }}
      className="bg-black text-white py-6 text-center border-t border-gray-800"
    >
      <motion.p
        whileHover={{ scale: 1.05, transition: { duration: 0.3 } }}
        className="text-gray-400 text-sm tracking-wide"
      >
        &copy; {new Date().getFullYear()} <span className="text-white font-semibold">TitanGate</span>. All rights reserved.
      </motion.p>
    </motion.footer>
  );
};

export default Footer;
