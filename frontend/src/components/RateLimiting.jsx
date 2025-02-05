import React, { useState, useEffect } from 'react';

const RateLimiting = () => {
  const [remaining, setRemaining] = useState(100); // Example remaining requests

  useEffect(() => {
    // Simulate a countdown of remaining requests
    const interval = setInterval(() => {
      if (remaining > 0) setRemaining(remaining - 1);
    }, 1000);

    return () => clearInterval(interval); // Clean up on component unmount
  }, [remaining]);

  return (
    <div className="bg-gray-900 p-6 rounded-lg shadow-lg">
      <h3 className="text-xl font-semibold">Rate Limiting</h3>
      <p className="text-gray-400">Requests left: {remaining}</p>
      <div className="mt-4 bg-gray-800 p-4 rounded-lg">
        <p className="text-white">Current limit: 100 requests/min</p>
        <p className="text-white">Remaining: {remaining} requests</p>
      </div>
    </div>
  );
};

export default RateLimiting;
