import React from "react";
import { motion } from "framer-motion";

const Dashboard = () => {
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
            <p className="text-white">Requests: 5,000</p>
            <p className="text-white">Errors: 20</p>
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
            <p className="text-white">Error Rate: 0.4%</p>
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
            <p className="text-white">Service A: 45%</p>
            <p className="text-white">Service B: 55%</p>
          </div>
        </motion.div>
      </div>
    </motion.div>
  );
};

export default Dashboard;
