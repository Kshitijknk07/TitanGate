import React, { useEffect, useState } from "react";
import { motion } from "framer-motion";

const Dashboard = () => {
  const [traffic, setTraffic] = useState(0);
  const [errors, setErrors] = useState(0);
  const [errorRate, setErrorRate] = useState(0);
  const [loadDistribution, setLoadDistribution] = useState({ serviceA: 0, serviceB: 0 });

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:3000/metrics');
        const data = await response.json();
        setTraffic(data.traffic);
        setErrors(data.errors);
        setErrorRate((data.errors / data.traffic) * 100);
        setLoadDistribution(data.loadDistribution);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, []);

  return (
    <motion.div
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      transition={{ duration: 0.5 }}
      className="space-y-8"
    >
      <h2 className="text-3xl font-bold text-white">Dashboard</h2>
      <div className="grid grid-cols-3 gap-8">
        {/* API Analytics */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="bg-black p-6 rounded-lg shadow-lg border border-gray-700"
        >
          <h3 className="text-xl font-semibold">API Traffic</h3>
          <p className="text-gray-400">Real-time API traffic data and request count.</p>
          <div className="mt-4 bg-gray-800 p-4 rounded-lg">
            <p className="text-white">Requests: {traffic}</p>
            <p className="text-white">Errors: {errors}</p>
          </div>
        </motion.div>

        {/* Error Rates */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 0.2 }}
          className="bg-black p-6 rounded-lg shadow-lg border border-gray-700"
        >
          <h3 className="text-xl font-semibold">Error Rates</h3>
          <p className="text-gray-400">Track and analyze error trends in your APIs.</p>
          <div className="mt-4 bg-gray-800 p-4 rounded-lg">
            <p className="text-white">Error Rate: {errorRate.toFixed(2)}%</p>
          </div>
        </motion.div>

        {/* Load Balancing */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 0.4 }}
          className="bg-black p-6 rounded-lg shadow-lg border border-gray-700"
        >
          <h3 className="text-xl font-semibold">Load Balancing</h3>
          <p className="text-gray-400">View load distribution across services.</p>
          <div className="mt-4 bg-gray-800 p-4 rounded-lg">
            <p className="text-white">Service A: {loadDistribution.serviceA}%</p>
            <p className="text-white">Service B: {loadDistribution.serviceB}%</p>
          </div>
        </motion.div>
      </div>
    </motion.div>
  );
};

export default Dashboard;