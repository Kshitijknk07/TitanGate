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
        const response = await fetch("http://localhost:3000/metrics");
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
      transition={{ duration: 0.8, ease: "easeOut", bounce: 0.3 }}
      className="space-y-8"
    >
      <motion.h2
        whileHover={{ scale: 1.05, transition: { duration: 0.3 } }}
        className="text-3xl font-extrabold text-white tracking-wide"
      >
        Dashboard
      </motion.h2>
      <div className="grid grid-cols-3 gap-8">
        {[
          {
            title: "API Traffic",
            description: "Real-time API traffic data and request count.",
            content: (
              <>
                <p className="text-white">Requests: {traffic}</p>
                <p className="text-white">Errors: {errors}</p>
              </>
            ),
          },
          {
            title: "Error Rates",
            description: "Track and analyze error trends in your APIs.",
            content: <p className="text-white">Error Rate: {errorRate.toFixed(2)}%</p>,
          },
          {
            title: "Load Balancing",
            description: "View load distribution across services.",
            content: (
              <>
                <p className="text-white">Service A: {loadDistribution.serviceA}%</p>
                <p className="text-white">Service B: {loadDistribution.serviceB}%</p>
              </>
            ),
          },
        ].map((item, index) => (
          <motion.div
            key={index}
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8, delay: index * 0.2, ease: "easeOut", bounce: 0.3 }}
            whileHover={{ scale: 1.02, transition: { duration: 0.3 } }}
            className="bg-black p-6 rounded-2xl shadow-xl border border-gray-800"
          >
            <h3 className="text-xl font-extrabold">{item.title}</h3>
            <p className="text-gray-400">{item.description}</p>
            <motion.div
              whileHover={{ scale: 1.02, transition: { duration: 0.3 } }}
              className="mt-4 bg-gray-900 p-5 rounded-xl shadow-md"
            >
              {item.content}
            </motion.div>
          </motion.div>
        ))}
      </div>
    </motion.div>
  );
};

export default Dashboard;
