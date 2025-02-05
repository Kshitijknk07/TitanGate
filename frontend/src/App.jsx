import React from 'react';
import Header from './components/Header';
import Sidebar from './components/sidebar';
import Home from './pages/Home'; 

const App = () => {
  return (
    <div className="bg-black text-white min-h-screen">
      <Header />
      <div className="flex">
        <Sidebar />
        <main className="flex-1 p-6">
          <Home />
        </main>
      </div>
    </div>
  );
};

export default App;
