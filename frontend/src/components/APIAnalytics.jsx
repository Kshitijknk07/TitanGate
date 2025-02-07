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
      className="bg-black text-white p-6 rounded-2xl shadow-xl border border-gray-700 w-80 mx-auto"
    >
      <h3 className="text-2xl font-bold mb-2">API Analytics</h3>
      <p className="text-gray-400 text-sm">Track request metrics, performance, and errors.</p>
      <div className="mt-4 space-y-2">
        <motion.p 
          initial={{ opacity: 0 }} 
          animate={{ opacity: 1 }} 
          transition={{ delay: 0.3 }}
          className="text-lg"
        >
          <span className="font-semibold">Requests:</span> {traffic}
        </motion.p>
        <motion.p 
          initial={{ opacity: 0 }} 
          animate={{ opacity: 1 }} 
          transition={{ delay: 0.5 }}
          className="text-lg"
        >
          <span className="font-semibold">Errors:</span> {errors}
        </motion.p>
        <motion.p 
          initial={{ opacity: 0 }} 
          animate={{ opacity: 1 }} 
          transition={{ delay: 0.7 }}
          className={`text-lg ${errors / traffic > 0.01 ? "text-red-500" : "text-green-500"}`}
        >
          <span className="font-semibold">Error Rate:</span> {(errors / traffic * 100).toFixed(2)}%
        </motion.p>
      </div>
    </motion.div>
  );
};

export default APIAnalytics;