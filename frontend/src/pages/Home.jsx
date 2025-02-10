import React from "react";
import { motion } from "framer-motion";
import APIAnalytics from "../components/APIAnalytics";
import RateLimiting from "../components/RateLimiting";

const Home = () => {
  return (
    <motion.div
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      transition={{ duration: 0.8, ease: "easeOut", bounce: 0.3 }}
      className="flex flex-col space-y-8"
    >
      <motion.h2
        whileHover={{ scale: 1.05, transition: { duration: 0.3 } }}
        className="text-3xl font-extrabold text-white tracking-wide"
      >
        Welcome to TitanGate
      </motion.h2>
      <motion.p
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6, ease: "easeOut" }}
        className="text-xl text-gray-400"
      >
        A powerful API Gateway with rate limiting, caching, authentication, and more.
      </motion.p>
      <div className="grid grid-cols-2 gap-8">
        <APIAnalytics />
        <RateLimiting />
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8, delay: 0.2, ease: "easeOut", bounce: 0.3 }}
          whileHover={{ scale: 1.02, transition: { duration: 0.3 } }}
          className="bg-black p-6 rounded-2xl shadow-xl border border-gray-800"
        >
          <h3 className="text-xl font-extrabold">Load Balancing</h3>
          <p className="text-gray-400">Distribute traffic evenly across backend services.</p>
        </motion.div>
      </div>
    </motion.div>
  );
};

export default Home;
