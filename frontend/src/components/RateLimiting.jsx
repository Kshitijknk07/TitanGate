import React, { useState, useEffect } from "react";
import { motion } from "framer-motion";

const RateLimiting = () => {
  const [remaining, setRemaining] = useState(100);

  useEffect(() => {
    const interval = setInterval(() => {
      setRemaining((prev) => (prev > 0 ? prev - 1 : 0));
    }, 1000);

    return () => clearInterval(interval);
  }, []);

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.8, ease: "easeOut", bounce: 0.3 }}
      className="bg-black text-white p-6 rounded-2xl shadow-2xl border border-gray-800 w-80 mx-auto"
    >
      <motion.h3
        whileHover={{ scale: 1.05, transition: { duration: 0.3 } }}
        className="text-2xl font-extrabold mb-2 tracking-wide"
      >
        Rate Limiting
      </motion.h3>
      <p className="text-gray-400 text-lg font-medium">Requests left: {remaining}</p>
      <motion.div
        whileHover={{ scale: 1.02, transition: { duration: 0.3 } }}
        className="mt-4 bg-gray-900 p-5 rounded-xl shadow-md"
      >
        <p className="text-gray-300 text-lg">Current limit: <span className="text-white">100 requests/min</span></p>
        <p className="text-gray-300 text-lg">Remaining: <span className="text-white">{remaining} requests</span></p>
      </motion.div>
    </motion.div>
  );
};

export default RateLimiting;
