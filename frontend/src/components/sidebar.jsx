import React from 'react';

const Sidebar = () => {
  return (
    <div className="bg-gray-800 text-white w-64 h-screen p-4">
      <h2 className="text-xl font-semibold mb-6">Navigation</h2>
      <ul>
        <li className="mb-4 hover:text-gray-400"><a href="#">API Analytics</a></li>
        <li className="mb-4 hover:text-gray-400"><a href="#">Load Balancing</a></li>
        <li className="mb-4 hover:text-gray-400"><a href="#">GraphQL Gateway</a></li>
      </ul>
    </div>
  );
};

export default Sidebar;
