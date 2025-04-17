'use client';

import { useState } from 'react';

export default function CredentialsPage() {
  const [accessKey, setAccessKey] = useState('');
  const [secretKey, setSecretKey] = useState('');
  const [region, setRegion] = useState('');

  // Placeholder functions for demonstration
  const handleLogout = () => {
    alert('Logged out!');
  };

  const handleConnect = () => {
    alert(`Connecting with:
    AccessKey: ${accessKey}
    SecretKey: ${secretKey}
    Region: ${region}`);
  };

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col">
      {/* Top Bar */}
      <header className="bg-white border-b border-gray-200 px-6 py-4 flex justify-between items-center shadow-sm">
        <h1 className="text-xl font-bold text-blue-600">CloudCleanup</h1>
        <button
          onClick={handleLogout}
          className="text-sm font-semibold text-gray-600 hover:text-gray-900"
        >
          LOGOUT
        </button>
      </header>

      {/* Main Content */}
      <main className="flex-1 flex items-center justify-center p-6">
        <div className="max-w-2xl w-full bg-white border border-gray-200 rounded p-6 shadow">
          <h2 className="text-2xl font-bold mb-6">Enter Your AWS Credentials</h2>

          {/* Input Fields */}
          <div className="grid grid-cols-1 gap-4 mb-6">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Access Key
              </label>
              <input
                type="text"
                value={accessKey}
                onChange={(e) => setAccessKey(e.target.value)}
                placeholder="AKIA..."
                className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Secret Key
              </label>
              <input
                type="password"
                value={secretKey}
                onChange={(e) => setSecretKey(e.target.value)}
                placeholder="••••••••"
                className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Region
              </label>
              <input
                type="text"
                value={region}
                onChange={(e) => setRegion(e.target.value)}
                placeholder="ca-central-1"
                className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:border-blue-500"
              />
            </div>
          </div>

          {/* Footer with Illustration and Connect Button */}
          <div className="flex items-center justify-between">
            {/* Placeholder for your cloud/plant illustration */}
            <div className="flex-shrink-0">
              <img
                src="cloudgarden.png"
                alt="Cloud and plants illustration"
                className="w-32 h-24"
              />
            </div>

            <button
              onClick={handleConnect}
              className="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-6 rounded"
            >
              CONNECT
            </button>
          </div>
        </div>
      </main>
    </div>
  );
}

