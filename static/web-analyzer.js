function analyzeWebPage(event) {
    event.preventDefault();
    const inputVal = document.getElementById('webURL').value;
    const payload = {
        requestId: '12345',
        webUrl: inputVal
    };

    fetch('/analyze', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    })
        .then(res => res.json())
        .then(data => {
            document.getElementById('detailsTable').innerHTML = renderDetailsTable(data);
            document.getElementById('headersCountTable').innerHTML = renderMapAsTable(data.headersCount);
            document.getElementById('hyperlinksTable').innerHTML = renderListAsTable("Hyperlinks", data.hyperlinks);
            document.getElementById('brokenLinksTable').innerHTML = renderListAsTable("Broken Links", data.brokenLinks);
        })
        .catch(err => console.error('Error:', err));
}

function renderDetailsTable(data) {
    return `
            <table class="table table-bordered">
                <tbody>
                    <tr><th>Web URL</th><td>${data.webUrl}</td></tr>
                    <tr><th>HTML Version</th><td>${data.htmlVersion}</td></tr>
                    <tr><th>Page Title</th><td>${data.pageTitle}</td></tr>
                    <tr><th>Is Login Page</th><td>${data.isLoginPage}</td></tr>
                </tbody>
            </table>
        `;
}

function renderListAsTable(title, items) {
    let rows = items.map((item, index) => `<tr><td>${index + 1}</td><td>${item}</td></tr>`)
                    .join('');
    return `
            <h4>${title}</h4>
            <table class="table table-bordered">
                <thead><tr><th>#</th><th>URL</th></tr></thead>
                <tbody>${rows}</tbody>
            </table>
        `;
}

function renderMapAsTable(map) {
    if (!map || Object.keys(map).length === 0) {
        return `<h4>Headers Count</h4><p>No header tags found.</p>`;
    }

    let rows = Object.entries(map)
        .map(([tag, count]) => `<tr><td>${tag}</td><td>${count}</td></tr>`)
        .join('');
    return `
            <h4>Headers Count</h4>
            <table class="table table-bordered">
                <thead><tr><th>Header Tag</th><th>Count</th></tr></thead>
                <tbody>${rows}</tbody>
            </table>
        `;
}