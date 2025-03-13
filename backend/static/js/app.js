const { useState, useEffect } = React;
const { Line } = ReactChartjs2;

const App = () => {
    const [metrics, setMetrics] = useState({
        totalRequests: 0,
        cacheHitRate: 0,
        activeBackends: 0,
        rateLimited: 0
    });

    const [backends, setBackends] = useState([]);
    const [chartData, setChartData] = useState({
        labels: [],
        datasets: [{
            label: 'Requests per Second',
            data: [],
            borderColor: '#2563eb',
            tension: 0.4
        }]
    });

    useEffect(() => {
        const fetchMetrics = async () => {
            try {
                const response = await fetch('/metrics');
                const data = await response.json();
                setMetrics(data);
            } catch (error) {
                console.error('Error fetching metrics:', error);
            }
        };

        const fetchBackends = async () => {
            try {
                const response = await fetch('/api/v1/backends');
                const data = await response.json();
                setBackends(data);
            } catch (error) {
                console.error('Error fetching backends:', error);
            }
        };

        const updateChart = () => {
            const now = new Date().toLocaleTimeString();
            setChartData(prev => {
                const newData = {
                    labels: [...prev.labels.slice(-10), now],
                    datasets: [{
                        ...prev.datasets[0],
                        data: [...prev.datasets[0].data.slice(-10), Math.random() * 100]
                    }]
                };
                return newData;
            });
        };

        const interval = setInterval(() => {
            fetchMetrics();
            fetchBackends();
            updateChart();
        }, 1000);

        return () => clearInterval(interval);
    }, []);

    return (
        <div className="dashboard">
            <aside className="sidebar">
                <h1>TitanGate</h1>
                <nav className="nav-menu">
                    <a href="#overview" className="nav-item active">Overview</a>
                    <a href="#load-balancer" className="nav-item">Load Balancer</a>
                    <a href="#rate-limits" className="nav-item">Rate Limits</a>
                    <a href="#cache" className="nav-item">Cache</a>
                    <a href="#metrics" className="nav-item">Metrics</a>
                </nav>
            </aside>

            <main className="main-content">
                <div className="header">
                    <h2>System Overview</h2>
                    <div className="last-update">
                        Last updated: {new Date().toLocaleTimeString()}
                    </div>
                </div>

                <div className="metrics-grid">
                    <div className="metric-card">
                        <h3>Total Requests</h3>
                        <p>{metrics.totalRequests.toLocaleString()}</p>
                    </div>
                    <div className="metric-card">
                        <h3>Cache Hit Rate</h3>
                        <p>{metrics.cacheHitRate.toFixed(1)}%</p>
                    </div>
                    <div className="metric-card">
                        <h3>Active Backends</h3>
                        <p>{metrics.activeBackends}/3</p>
                    </div>
                    <div className="metric-card">
                        <h3>Rate Limited</h3>
                        <p>{metrics.rateLimited.toLocaleString()}</p>
                    </div>
                </div>

                <div className="backend-grid">
                    {backends.map((backend, index) => (
                        <div key={index} className="backend-card">
                            <h3>
                                <span className={`status-indicator ${backend.active ? 'status-active' : 'status-inactive'}`}></span>
                                Backend {index + 1}
                            </h3>
                            <p>URL: {backend.url}</p>
                            <p>Weight: {backend.weight}</p>
                            <p>Active Connections: {backend.activeConns}</p>
                            <p>Response Time: {backend.responseTime}ms</p>
                        </div>
                    ))}
                </div>

                <div className="chart-container">
                    <h3>Request Rate</h3>
                    <Line data={chartData} options={{
                        responsive: true,
                        scales: {
                            y: {
                                beginAtZero: true
                            }
                        }
                    }} />
                </div>
            </main>
        </div>
    );
};

ReactDOM.render(<App />, document.getElementById('root')); 