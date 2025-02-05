import React from 'react';

const Dashboard = () => {
  return (
    <div className="space-y-8">
      <h2 className="text-3xl font-bold">Dashboard</h2>
      <div className="grid grid-cols-3 gap-8">
        {/* API Analytics */}
        <div className="bg-gray-900 p-6 rounded-lg shadow-lg">
          <h3 className="text-xl font-semibold">API Traffic</h3>
          <p className="text-gray-400">Real-time API traffic data and request count.</p>
          <div className="mt-4 bg-gray-800 p-4 rounded-lg">
            {/* Example chart or metrics */}
            <p className="text-white">Requests: 5,000</p>
            <p className="text-white">Errors: 20</p>
          </div>
        </div>

        {/* Error Rates */}
        <div className="bg-gray-900 p-6 rounded-lg shadow-lg">
          <h3 className="text-xl font-semibold">Error Rates</h3>
          <p className="text-gray-400">Track and analyze error trends in your APIs.</p>
          <div className="mt-4 bg-gray-800 p-4 rounded-lg">
            {/* Example chart or metrics */}
            <p className="text-white">Error Rate: 0.4%</p>
          </div>
        </div>

        {/* Load Balancing */}
        <div className="bg-gray-900 p-6 rounded-lg shadow-lg">
          <h3 className="text-xl font-semibold">Load Balancing</h3>
          <p className="text-gray-400">View load distribution across services.</p>
          <div className="mt-4 bg-gray-800 p-4 rounded-lg">
            {/* Example chart or metrics */}
            <p className="text-white">Service A: 45%</p>
            <p className="text-white">Service B: 55%</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
