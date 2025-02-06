import React from "react";
import { motion } from "framer-motion";
import APIAnalytics from "../components/APIAnalytics";
import RateLimiting from "../components/RateLimiting";

const Home = () => {
  return (
    <motion.div
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      transition={{ duration: 0.5 }}
      className="flex flex-col space-y-8"
    >
      <h2 className="text-3xl font-bold text-white">Welcome to TitanGate</h2>
      <p className="text-xl text-gray-400">
        A powerful API Gateway with rate limiting, caching, authentication, and more.
      </p>
      <div className="grid grid-cols-2 gap-8">
        <APIAnalytics />
        <RateLimiting />
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 0.2 }}
          className="bg-black p-6 rounded-lg shadow-lg border border-gray-700"
        >
          <h3 className="text-xl font-semibold">Load Balancing</h3>
          <p className="text-gray-400">Distribute traffic evenly across backend services.</p>
        </motion.div>
      </div>
    </motion.div>
  );
};

export default Home;
