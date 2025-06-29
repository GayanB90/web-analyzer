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
            document.getElementById('result-container').innerText = JSON.stringify(data);
        })
        .catch(err => console.error('Error:', err));
}