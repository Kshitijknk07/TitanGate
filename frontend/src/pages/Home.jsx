import React from 'react';
import APIAnalytics from '../components/APIAnalytics';
import RateLimiting from '../components/RateLimiting';

const Home = () => {
  return (
    <div className="flex flex-col space-y-8">
      <h2 className="text-3xl font-bold">Welcome to TitanGate</h2>
      <p className="text-xl text-gray-400">
        A powerful API Gateway with rate limiting, caching, authentication, and more.
      </p>
      <div className="grid grid-cols-2 gap-8">
        <APIAnalytics />
        <RateLimiting /> {/* Add Rate Limiting component */}
        <div className="bg-gray-900 p-6 rounded-lg shadow-lg">
          <h3 className="text-xl font-semibold">Load Balancing</h3>
          <p className="text-gray-400">Distribute traffic evenly across backend services.</p>
        </div>
      </div>
    </div>
  );
};

export default Home;
