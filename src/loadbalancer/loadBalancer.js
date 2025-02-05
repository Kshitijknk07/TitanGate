// Configurable list of backends, typically coming from environment variables
const backends = process.env.BACKEND_SERVICES ? process.env.BACKEND_SERVICES.split(',') : [
    "http://backend-service-1:3001",
    "http://backend-service-2:3002",
    "http://backend-service-3:3003"
  ];
  
  let currentIndex = 0;
  let failedBackends = new Set();
  
  // Function to get the next available backend with error handling
  export function getNextBackend() {
    // Filter out the failed backends
    const availableBackends = backends.filter((url) => !failedBackends.has(url));
  
    if (availableBackends.length === 0) {
      // All backends are down, handle it gracefully (maybe return a fallback or throw an error)
      console.error('No available backends. Please check service status.');
      throw new Error('No available backends.');
    }
  
    // Round robin logic
    const target = availableBackends[currentIndex];
    currentIndex = (currentIndex + 1) % availableBackends.length; // Rotate index for Round Robin
  
    console.log(`Using backend: ${target}`);
    return target;
  }
  
  // Function to mark a backend as failed (could be used in response handling or health checks)
  export function markBackendAsFailed(url) {
    if (backends.includes(url)) {
      failedBackends.add(url);
      console.warn(`Backend ${url} marked as failed.`);
    }
  }
  
  // Function to mark a backend as recovered
  export function markBackendAsRecovered(url) {
    if (failedBackends.has(url)) {
      failedBackends.delete(url);
      console.info(`Backend ${url} marked as recovered.`);
    }
  }
      