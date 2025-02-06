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
      transition={{ duration: 0.5 }}
      className="bg-black text-white p-6 rounded-2xl shadow-xl border border-gray-700 w-80 mx-auto"
    >
      <h3 className="text-2xl font-bold mb-2">Rate Limiting</h3>
      <p className="text-gray-400">Requests left: {remaining}</p>
      <div className="mt-4 bg-gray-800 p-4 rounded-lg">
        <p className="text-white">Current limit: 100 requests/min</p>
        <p className="text-white">Remaining: {remaining} requests</p>
      </div>
    </motion.div>
  );
};

export default RateLimiting;
