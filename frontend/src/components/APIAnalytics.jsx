import React, { useEffect, useState } from "react";
import { motion } from "framer-motion";
import { authenticatedFetch } from "../utils/api";

const APIAnalytics = () => {
  const [traffic, setTraffic] = useState(0);
  const [errors, setErrors] = useState(0);
  const API_URL = import.meta.env.VITE_API_URL;

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await authenticatedFetch(`${API_URL}/metrics`);
        if (!response.ok) throw new Error('Failed to fetch metrics');
        const data = await response.json();
        setTraffic(data.traffic);
        setErrors(data.errors);
      } catch (error) {
        console.error("Error fetching data:", error);
        setTraffic(0);
        setErrors(0);
      }
    };

    fetchData();
    
    const interval = setInterval(fetchData, 30000);
    return () => clearInterval(interval);
  }, []);

  return (
    <motion.div 
      initial={{ opacity: 0, y: 20 }} 
      animate={{ opacity: 1, y: 0 }} 
      transition={{ duration: 0.5 }}
      className="bg-gradient-to-br from-gray-900 to-black text-white p-8 rounded-3xl shadow-2xl border border-gray-800 w-96 mx-auto"
    >
      <motion.h3 
        initial={{ opacity: 0, y: -10 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.2, duration: 0.5 }}
        className="text-3xl font-bold mb-4 bg-gradient-to-r from-purple-500 to-blue-500 bg-clip-text text-transparent"
      >
        API Analytics
      </motion.h3>
      <motion.p 
        initial={{ opacity: 0, y: -10 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.4, duration: 0.5 }}
        className="text-gray-400 text-sm mb-6"
      >
        Track request metrics, performance, and errors in real-time.
      </motion.p>
      <div className="space-y-4">
        <motion.div 
          initial={{ opacity: 0, x: -20 }}
          animate={{ opacity: 1, x: 0 }}
          transition={{ delay: 0.6, duration: 0.5 }}
          className="bg-gray-800 p-4 rounded-xl shadow-lg"
        >
          <p className="text-lg">
            <span className="font-semibold text-purple-400">Requests:</span>{" "}
            <span className="text-white">{traffic}</span>
          </p>
        </motion.div>
        <motion.div 
          initial={{ opacity: 0, x: -20 }}
          animate={{ opacity: 1, x: 0 }}
          transition={{ delay: 0.8, duration: 0.5 }}
          className="bg-gray-800 p-4 rounded-xl shadow-lg"
        >
          <p className="text-lg">
            <span className="font-semibold text-red-400">Errors:</span>{" "}
            <span className="text-white">{errors}</span>
          </p>
        </motion.div>
        <motion.div 
          initial={{ opacity: 0, x: -20 }}
          animate={{ opacity: 1, x: 0 }}
          transition={{ delay: 1, duration: 0.5 }}
          className="bg-gray-800 p-4 rounded-xl shadow-lg"
        >
          <p className="text-lg">
            <span className="font-semibold text-blue-400">Error Rate:</span>{" "}
            <span className={`${errors / traffic > 0.01 ? "text-red-500" : "text-green-500"}`}>
              {(errors / traffic * 100).toFixed(2)}%
            </span>
          </p>
        </motion.div>
      </div>
    </motion.div>
  );
};

export default APIAnalytics;