import React from 'react';

const Header = () => {
  return (
    <header className="bg-white text-black p-4 shadow-md">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">TitanGate</h1>
        <nav className="space-x-6">
          <a href="/" className="text-black hover:text-gray-500">Home</a>
          <a href="/dashboard" className="text-black hover:text-gray-500">Dashboard</a>
          <a href="/login" className="text-black hover:text-gray-500">Login</a>
        </nav>
      </div>
    </header>
  );
};

export default Header;
