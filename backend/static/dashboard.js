document.addEventListener('DOMContentLoaded', () => {
    setupNavigation();
    initializeMetrics();
    initializeCharts();
});

function setupNavigation() {
    const links = document.querySelectorAll('.sidebar a');
    links.forEach(link => {
        link.addEventListener('click', (e) => {
            e.preventDefault();
            const target = e.target.getAttribute('href').substring(1);
            showSection(target);
        });
    });
}

function showSection(id) {
    document.querySelectorAll('section').forEach(s => s.classList.remove('active'));
    document.querySelector(`#${id}`).classList.add('active');
    document.querySelectorAll('.sidebar a').forEach(a => a.classList.remove('active'));
    document.querySelector(`[href="#${id}"]`).classList.add('active');
}

function initializeMetrics() {
    const eventSource = new EventSource('/events');
    
    eventSource.onmessage = (event) => {
        const data = JSON.parse(event.data);
        updateMetrics(data);
        updateBackends(data.backends);
        updateCharts(data);
    };

    eventSource.onerror = () => {
        eventSource.close();
        setTimeout(() => initializeMetrics(), 5000);
    };
}

function updateMetrics(data) {
    document.getElementById('totalRequests').textContent = data.totalRequests;
    document.getElementById('cacheHitRate').textContent = `${data.cacheHitRate}%`;
    document.getElementById('activeBackends').textContent = `${data.activeBackends}/3`;
    document.getElementById('rateLimited').textContent = data.rateLimited;
}

function updateBackends(backends) {
    const container = document.getElementById('backendList');
    container.innerHTML = backends.map(backend => `
        <div class="backend-card ${backend.active ? 'active' : 'inactive'}">
            <h3>${backend.url}</h3>
            <p>Requests: ${backend.requests}</p>
            <p>Status: ${backend.active ? 'Active' : 'Inactive'}</p>
        </div>
    `).join('');
}

let metricsChart;

function initializeCharts() {
    const ctx = document.getElementById('metricsChart').getContext('2d');
    metricsChart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: [],
            datasets: [{
                label: 'Requests/s',
                data: [],
                borderColor: '#1a73e8',
                tension: 0.4
            }]
        },
        options: {
            responsive: true,
            animation: false
        }
    });
}

function updateCharts(data) {
    const timestamp = new Date().toLocaleTimeString();
    
    metricsChart.data.labels.push(timestamp);
    metricsChart.data.datasets[0].data.push(data.requestsPerSecond);
    
    if (metricsChart.data.labels.length > 20) {
        metricsChart.data.labels.shift();
        metricsChart.data.datasets[0].data.shift();
    }
    
    metricsChart.update();
}