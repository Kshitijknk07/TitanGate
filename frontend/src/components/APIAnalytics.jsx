import React, { useEffect, useState } from 'react';

const APIAnalytics = () => {
  const [traffic, setTraffic] = useState(0);
  const [errors, setErrors] = useState(0);

  useEffect(() => {
    // Fetch real data from your API or simulate it here
    setTimeout(() => {
      setTraffic(5000);
      setErrors(20);
    }, 1000);
  }, []);

  return (
    <div className="bg-gray-900 p-6 rounded-lg shadow-lg">
      <h3 className="text-xl font-semibold">API Analytics</h3>
      <p className="text-gray-400">Track request metrics, performance, and errors.</p>
      <div className="mt-4">
        <p className="text-white">Requests: {traffic}</p>
        <p className="text-white">Errors: {errors}</p>
        <p className="text-white">Error Rate: {(errors / traffic * 100).toFixed(2)}%</p>
      </div>
    </div>
  );
};

export default APIAnalytics;
